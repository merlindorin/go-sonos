// CODE GENERATED, DO NOT EDIT

package S5

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceMusicServicesService is the UPnP namespace for the Sonos service.
	namespaceMusicServicesService = "urn:schemas-upnp-org:service:MusicServices:1"

	// eventMusicServicesService is the event path for the Sonos event service.
	eventMusicServicesService = "/MusicServices/Event"

	// controlMusicServicesService is the action path for the Sonos service.
	controlMusicServicesService = "/MusicServices/Control"
)

// MusicServicesService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type MusicServicesService struct {
	doer do.Doer
}

// NewMusicServicesService initializes a new MusicServicesService with the provided Doer.
func NewMusicServicesService(doer do.Doer) *MusicServicesService {
	return &MusicServicesService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (m MusicServicesService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return m.doer.Do(ctx, do.WithPath(eventMusicServicesService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type MusicServicesGetSessionIdRes struct {
	SessionId SessionId
}

func (m MusicServicesService) GetSessionId(ctx context.Context, serviceId ServiceId, username Username, getSessionId *MusicServicesGetSessionIdRes) error {
	type GetSessionIdReq struct {
		ServiceId ServiceId
		Username  Username
	}

	type GetSessionIdNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetSessionIdReq
	}

	type GetSessionIdBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		GetSessionIdNode GetSessionIdNode `xml:"u:GetSessionId,omitempty"`
	}

	body := GetSessionIdBody{GetSessionIdNode: GetSessionIdNode{
		GetSessionIdReq: GetSessionIdReq{
			ServiceId: serviceId,
			Username:  username,
		},
		Xmlns: namespaceMusicServicesService,
	}}

	return m.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlMusicServicesService, namespaceMusicServicesService, "GetSessionId", body, getSessionId)...)
}

type MusicServicesListAvailableServicesRes struct {
	AvailableServiceDescriptorList ServiceDescriptorList
	AvailableServiceTypeList       ServiceTypeList
	AvailableServiceListVersion    ServiceListVersion
}

// ListAvailableServices Load music service list as xml
func (m MusicServicesService) ListAvailableServices(ctx context.Context, listAvailableServices *MusicServicesListAvailableServicesRes) error {
	type ListAvailableServicesReq struct{}

	type ListAvailableServicesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ListAvailableServicesReq
	}

	type ListAvailableServicesBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		ListAvailableServicesNode ListAvailableServicesNode `xml:"u:ListAvailableServices,omitempty"`
	}

	body := ListAvailableServicesBody{ListAvailableServicesNode: ListAvailableServicesNode{
		ListAvailableServicesReq: ListAvailableServicesReq{},
		Xmlns:                    namespaceMusicServicesService,
	}}

	return m.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlMusicServicesService, namespaceMusicServicesService, "ListAvailableServices", body, listAvailableServices)...)
}

func (m MusicServicesService) UpdateAvailableServices(ctx context.Context) error {
	type UpdateAvailableServicesReq struct{}

	type UpdateAvailableServicesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		UpdateAvailableServicesReq
	}

	type UpdateAvailableServicesBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		UpdateAvailableServicesNode UpdateAvailableServicesNode `xml:"u:UpdateAvailableServices,omitempty"`
	}

	body := UpdateAvailableServicesBody{UpdateAvailableServicesNode: UpdateAvailableServicesNode{
		UpdateAvailableServicesReq: UpdateAvailableServicesReq{},
		Xmlns:                      namespaceMusicServicesService,
	}}

	return m.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlMusicServicesService, namespaceMusicServicesService, "UpdateAvailableServices", body, nil)...)
}
