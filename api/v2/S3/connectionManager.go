// CODE GENERATED, DO NOT EDIT

package S3

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceConnectionManagerService is the UPnP namespace for the Sonos service.
	namespaceConnectionManagerService = "urn:schemas-upnp-org:service:ConnectionManager:1"

	// eventConnectionManagerService is the event path for the Sonos event service.
	eventConnectionManagerService = "/MediaRenderer/ConnectionManager/Event"

	// controlConnectionManagerService is the action path for the Sonos service.
	controlConnectionManagerService = "/MediaRenderer/ConnectionManager/Control"
)

// ConnectionManagerService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type ConnectionManagerService struct {
	doer do.Doer
}

// NewConnectionManagerService initializes a new ConnectionManagerService with the provided Doer.
func NewConnectionManagerService(doer do.Doer) *ConnectionManagerService {
	return &ConnectionManagerService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (c ConnectionManagerService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return c.doer.Do(ctx, do.WithPath(eventConnectionManagerService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type ConnectionManagerGetProtocolInfoRes struct {
	Source SourceProtocolInfo
	Sink   SinkProtocolInfo
}

func (c ConnectionManagerService) GetProtocolInfo(ctx context.Context, getProtocolInfo *ConnectionManagerGetProtocolInfoRes) error {
	type GetProtocolInfoReq struct{}

	type GetProtocolInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetProtocolInfoReq
	}

	type GetProtocolInfoBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		GetProtocolInfoNode GetProtocolInfoNode `xml:"u:GetProtocolInfo,omitempty"`
	}

	body := GetProtocolInfoBody{GetProtocolInfoNode: GetProtocolInfoNode{
		GetProtocolInfoReq: GetProtocolInfoReq{},
		Xmlns:              namespaceConnectionManagerService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlConnectionManagerService, namespaceConnectionManagerService, "GetProtocolInfo", body, getProtocolInfo)...)
}

type ConnectionManagerGetCurrentConnectionIDsRes struct {
	ConnectionIDs ConnectionIDs
}

func (c ConnectionManagerService) GetCurrentConnectionIDs(ctx context.Context, getCurrentConnectionIds *ConnectionManagerGetCurrentConnectionIDsRes) error {
	type GetCurrentConnectionIDsReq struct{}

	type GetCurrentConnectionIDsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetCurrentConnectionIDsReq
	}

	type GetCurrentConnectionIDsBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		GetCurrentConnectionIDsNode GetCurrentConnectionIDsNode `xml:"u:GetCurrentConnectionIDs,omitempty"`
	}

	body := GetCurrentConnectionIDsBody{GetCurrentConnectionIDsNode: GetCurrentConnectionIDsNode{
		GetCurrentConnectionIDsReq: GetCurrentConnectionIDsReq{},
		Xmlns:                      namespaceConnectionManagerService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlConnectionManagerService, namespaceConnectionManagerService, "GetCurrentConnectionIDs", body, getCurrentConnectionIds)...)
}

type ConnectionManagerGetCurrentConnectionInfoRes struct {
	RcsID                 RcsID
	AVTransportID         AVTransportID
	ProtocolInfo          ProtocolInfo
	PeerConnectionManager ConnectionManager
	PeerConnectionID      ConnectionID
	Direction             Direction
	Status                ConnectionStatus
}

func (c ConnectionManagerService) GetCurrentConnectionInfo(ctx context.Context, connectionId ConnectionID, getCurrentConnectionInfo *ConnectionManagerGetCurrentConnectionInfoRes) error {
	type GetCurrentConnectionInfoReq struct {
		ConnectionID ConnectionID
	}

	type GetCurrentConnectionInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetCurrentConnectionInfoReq
	}

	type GetCurrentConnectionInfoBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		GetCurrentConnectionInfoNode GetCurrentConnectionInfoNode `xml:"u:GetCurrentConnectionInfo,omitempty"`
	}

	body := GetCurrentConnectionInfoBody{GetCurrentConnectionInfoNode: GetCurrentConnectionInfoNode{
		GetCurrentConnectionInfoReq: GetCurrentConnectionInfoReq{ConnectionID: connectionId},
		Xmlns:                       namespaceConnectionManagerService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlConnectionManagerService, namespaceConnectionManagerService, "GetCurrentConnectionInfo", body, getCurrentConnectionInfo)...)
}
