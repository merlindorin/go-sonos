// CODE GENERATED, DO NOT EDIT

package S19

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceQPlayService is the UPnP namespace for the Sonos service.
	namespaceQPlayService = "urn:schemas-tencent-com:service:QPlay:1"

	// eventQPlayService is the event path for the Sonos event service.
	eventQPlayService = "/QPlay/Event"

	// controlQPlayService is the action path for the Sonos service.
	controlQPlayService = "/QPlay/Control"
)

// QPlayService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type QPlayService struct {
	doer do.Doer
}

// NewQPlayService initializes a new QPlayService with the provided Doer.
func NewQPlayService(doer do.Doer) *QPlayService {
	return &QPlayService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (q QPlayService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return q.doer.Do(ctx, do.WithPath(eventQPlayService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type QPlayQPlayAuthRes struct {
	Code Code
	MID  MID
	DID  DID
}

func (q QPlayService) QPlayAuth(ctx context.Context, seed Seed, qplayAuth *QPlayQPlayAuthRes) error {
	type QPlayAuthReq struct {
		Seed Seed
	}

	type QPlayAuthNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		QPlayAuthReq
	}

	type QPlayAuthBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		QPlayAuthNode QPlayAuthNode `xml:"u:QPlayAuth,omitempty"`
	}

	body := QPlayAuthBody{QPlayAuthNode: QPlayAuthNode{
		QPlayAuthReq: QPlayAuthReq{Seed: seed},
		Xmlns:        namespaceQPlayService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQPlayService, namespaceQPlayService, "QPlayAuth", body, qplayAuth)...)
}
