package opts

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/vanyda-official/go-shared/pkg/net/do"
)

const (
	soapEncodingStyle  = "http://schemas.xmlsoap.org/soap/encoding/"
	soapEnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

func WithSonosUPnPOptions(
	path string,
	actionNamespace string,
	actionName string,
	requestValue interface{}, // value to serialize into xml
	responseValue interface{}, // pointer that will be hydrated with unserialize values
) []do.Option {
	return []do.Option{
		do.WithMethod(http.MethodPost),
		do.WithPath(path),
		do.WithExtraHeader("CONTENT-TYPE", "text/xml; charset=\"utf-8\""),
		do.WithExtraHeaderf("SOAPACTION", "%s#%s", actionNamespace, actionName),
		do.WithPreRequestHandler("sonos_prepare_soap_upnp_request", prepareSOAPRequest(requestValue)),
		do.WithPostRequestHandler("sonos_process_soap_upnp_response", processSOAPResponse(responseValue)), //nolint:lll,bodyclose // unrelated
	}
}

func prepareSOAPRequest(body interface{}) func(ctx context.Context, request *http.Request) error {
	return func(_ context.Context, request *http.Request) error {
		e := &Envelope{
			EncodingStyle: soapEncodingStyle,
			Xmlns:         soapEnvelopeSchema,
			Body:          body,
		}

		marshal, err := xml.Marshal(e)
		if err != nil {
			return err
		}

		request.Body = io.NopCloser(bytes.NewBuffer(marshal))
		request.ContentLength = int64(len(marshal))

		return nil
	}
}

func processSOAPResponse(resBody interface{}) func(ctx context.Context, _ *http.Request, res *http.Response) error {
	return func(_ context.Context, _ *http.Request, res *http.Response) error {
		defer res.Body.Close()
		if res.StatusCode != 200 && res.ContentLength == 0 {
			return fmt.Errorf("goupnp: SOAP request got HTTP %s", res.Status)
		}

		responseEnv := newSOAPEnvelope()
		decoder := xml.NewDecoder(res.Body)
		if err := decoder.Decode(responseEnv); err != nil {
			return fmt.Errorf("goupnp: error decoding response body: %w", err)
		}

		if responseEnv.Body.Fault != nil {
			return responseEnv.Body.Fault
		} else if res.StatusCode != http.StatusOK {
			return fmt.Errorf("goupnp: SOAP request got HTTP %s", res.Status)
		}

		if resBody != nil {
			if err := xml.Unmarshal(responseEnv.Body.RawAction, resBody); err != nil {
				return fmt.Errorf("goupnp: error unmarshalling out action: %w, %v", err, responseEnv.Body.RawAction)
			}
		}

		return nil
	}
}
