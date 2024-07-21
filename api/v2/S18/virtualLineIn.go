// CODE GENERATED, DO NOT EDIT

package S18

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceVirtualLineInService is the UPnP namespace for the Sonos service.
	namespaceVirtualLineInService = "urn:schemas-upnp-org:service:VirtualLineIn:1"

	// eventVirtualLineInService is the event path for the Sonos event service.
	eventVirtualLineInService = "/MediaRenderer/VirtualLineIn/Event"

	// controlVirtualLineInService is the action path for the Sonos service.
	controlVirtualLineInService = "/MediaRenderer/VirtualLineIn/Control"
)

// VirtualLineInService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type VirtualLineInService struct {
	doer do.Doer
}

// NewVirtualLineInService initializes a new VirtualLineInService with the provided Doer.
func NewVirtualLineInService(doer do.Doer) *VirtualLineInService {
	return &VirtualLineInService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (v VirtualLineInService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return v.doer.Do(ctx, do.WithPath(eventVirtualLineInService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type VirtualLineInStartTransmissionRes struct {
	CurrentTransportSettings TransportSettings
}

func (v VirtualLineInService) StartTransmission(ctx context.Context, coordinatorId PlayerID, startTransmission *VirtualLineInStartTransmissionRes) error {
	type StartTransmissionReq struct {
		InstanceID    InstanceID
		CoordinatorID PlayerID
	}

	type StartTransmissionNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StartTransmissionReq
	}

	type StartTransmissionBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		StartTransmissionNode StartTransmissionNode `xml:"u:StartTransmission,omitempty"`
	}

	body := StartTransmissionBody{StartTransmissionNode: StartTransmissionNode{
		StartTransmissionReq: StartTransmissionReq{
			CoordinatorID: coordinatorId,
			InstanceID:    0,
		},
		Xmlns: namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "StartTransmission", body, startTransmission)...)
}

func (v VirtualLineInService) StopTransmission(ctx context.Context, coordinatorId PlayerID) error {
	type StopTransmissionReq struct {
		InstanceID    InstanceID
		CoordinatorID PlayerID
	}

	type StopTransmissionNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StopTransmissionReq
	}

	type StopTransmissionBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		StopTransmissionNode StopTransmissionNode `xml:"u:StopTransmission,omitempty"`
	}

	body := StopTransmissionBody{StopTransmissionNode: StopTransmissionNode{
		StopTransmissionReq: StopTransmissionReq{
			CoordinatorID: coordinatorId,
			InstanceID:    0,
		},
		Xmlns: namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "StopTransmission", body, nil)...)
}

func (v VirtualLineInService) Play(ctx context.Context, speed Speed) error {
	type PlayReq struct {
		InstanceID InstanceID
		Speed      Speed
	}

	type PlayNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		PlayReq
	}

	type PlayBody struct {
		XMLName  xml.Name `xml:"s:Body"`
		PlayNode PlayNode `xml:"u:Play,omitempty"`
	}

	body := PlayBody{PlayNode: PlayNode{
		PlayReq: PlayReq{
			InstanceID: 0,
			Speed:      speed,
		},
		Xmlns: namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "Play", body, nil)...)
}

func (v VirtualLineInService) Pause(ctx context.Context) error {
	type PauseReq struct {
		InstanceID InstanceID
	}

	type PauseNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		PauseReq
	}

	type PauseBody struct {
		XMLName   xml.Name  `xml:"s:Body"`
		PauseNode PauseNode `xml:"u:Pause,omitempty"`
	}

	body := PauseBody{PauseNode: PauseNode{
		PauseReq: PauseReq{InstanceID: 0},
		Xmlns:    namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "Pause", body, nil)...)
}

func (v VirtualLineInService) Next(ctx context.Context) error {
	type NextReq struct {
		InstanceID InstanceID
	}

	type NextNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		NextReq
	}

	type NextBody struct {
		XMLName  xml.Name `xml:"s:Body"`
		NextNode NextNode `xml:"u:Next,omitempty"`
	}

	body := NextBody{NextNode: NextNode{
		NextReq: NextReq{InstanceID: 0},
		Xmlns:   namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "Next", body, nil)...)
}

func (v VirtualLineInService) Previous(ctx context.Context) error {
	type PreviousReq struct {
		InstanceID InstanceID
	}

	type PreviousNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		PreviousReq
	}

	type PreviousBody struct {
		XMLName      xml.Name     `xml:"s:Body"`
		PreviousNode PreviousNode `xml:"u:Previous,omitempty"`
	}

	body := PreviousBody{PreviousNode: PreviousNode{
		PreviousReq: PreviousReq{InstanceID: 0},
		Xmlns:       namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "Previous", body, nil)...)
}

func (v VirtualLineInService) Stop(ctx context.Context) error {
	type StopReq struct {
		InstanceID InstanceID
	}

	type StopNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StopReq
	}

	type StopBody struct {
		XMLName  xml.Name `xml:"s:Body"`
		StopNode StopNode `xml:"u:Stop,omitempty"`
	}

	body := StopBody{StopNode: StopNode{
		StopReq: StopReq{InstanceID: 0},
		Xmlns:   namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "Stop", body, nil)...)
}

func (v VirtualLineInService) SetVolume(ctx context.Context, desiredVolume Volume) error {
	type SetVolumeReq struct {
		InstanceID    InstanceID
		DesiredVolume Volume
	}

	type SetVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetVolumeReq
	}

	type SetVolumeBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		SetVolumeNode SetVolumeNode `xml:"u:SetVolume,omitempty"`
	}

	body := SetVolumeBody{SetVolumeNode: SetVolumeNode{
		SetVolumeReq: SetVolumeReq{
			DesiredVolume: desiredVolume,
			InstanceID:    0,
		},
		Xmlns: namespaceVirtualLineInService,
	}}

	return v.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlVirtualLineInService, namespaceVirtualLineInService, "SetVolume", body, nil)...)
}
