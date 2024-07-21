package opts

import (
	"encoding/xml"
	"fmt"
)

// Package opts defines types and functions to interact with SOAP-based web services.
// It provides struct types for constructing SOAP envelopes and for parsing SOAP faults.

// Envelope represents a SOAP envelope, which is the top-level element of a SOAP message.
type Envelope struct {
	XMLName       xml.Name    `xml:"s:Envelope"`
	Xmlns         string      `xml:"xmlns:s,attr"`
	EncodingStyle string      `xml:"s:encodingStyle,attr"`
	Body          interface{} `xml:"s:Body"`
}

// soapEnvelope is a private struct type that defines the XML structure of a SOAP envelope.
type soapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          soapBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

// soapBody is a private struct type that defines the structure of the SOAP body.
// A SOAP body can contain a Fault in case of an error, or the raw XML of the SOAP action.
type soapBody struct {
	Fault     *SOAPFaultError `xml:"Fault"`
	RawAction []byte          `xml:",innerxml"`
}

// SOAPFaultError represents a SOAP fault, which is a way for a web service to indicate an error.
// It implements the error interface, allowing it to be used as an error in Go.
type SOAPFaultError struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      struct {
		UPnPError struct {
			Errorcode        int    `xml:"errorCode"`
			ErrorDescription string `xml:"errorDescription"`
		} `xml:"UPnPError"`
		Raw []byte `xml:",innerxml"`
	} `xml:"detail"`
}

// Error returns a string representation of the SOAPFaultError, compliant with the Go error interface.
func (err *SOAPFaultError) Error() string {
	return fmt.Sprintf("SOAP fault. Code: %s | Explanation: %s | Detail: %s",
		err.FaultCode, err.FaultString, string(err.Detail.Raw))
}

// newSOAPEnvelope is an unexported factory function that creates a new instance of a SOAP envelope with a default
// encoding style. As the soapEnvelope type and newSOAPEnvelope function are private to the package, users of the
// package are encouraged to use the Envelope type directly.
func newSOAPEnvelope() *soapEnvelope {
	const soapEncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"

	return &soapEnvelope{
		EncodingStyle: soapEncodingStyle,
	}
}
