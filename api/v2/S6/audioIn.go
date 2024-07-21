// CODE GENERATED, DO NOT EDIT

package S6

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceAudioInService is the UPnP namespace for the Sonos service.
	namespaceAudioInService = "urn:schemas-upnp-org:service:AudioIn:1"

	// eventAudioInService is the event path for the Sonos event service.
	eventAudioInService = "/AudioIn/Event"

	// controlAudioInService is the action path for the Sonos service.
	controlAudioInService = "/AudioIn/Control"
)

// AudioInService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type AudioInService struct {
	doer do.Doer
}

// NewAudioInService initializes a new AudioInService with the provided Doer.
func NewAudioInService(doer do.Doer) *AudioInService {
	return &AudioInService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (a AudioInService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return a.doer.Do(ctx, do.WithPath(eventAudioInService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type AudioInStartTransmissionToGroupRes struct {
	CurrentTransportSettings TransportSettings
}

func (a AudioInService) StartTransmissionToGroup(ctx context.Context, coordinatorId MemberID, startTransmissionToGroup *AudioInStartTransmissionToGroupRes) error {
	type StartTransmissionToGroupReq struct {
		CoordinatorID MemberID
	}

	type StartTransmissionToGroupNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StartTransmissionToGroupReq
	}

	type StartTransmissionToGroupBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		StartTransmissionToGroupNode StartTransmissionToGroupNode `xml:"u:StartTransmissionToGroup,omitempty"`
	}

	body := StartTransmissionToGroupBody{StartTransmissionToGroupNode: StartTransmissionToGroupNode{
		StartTransmissionToGroupReq: StartTransmissionToGroupReq{CoordinatorID: coordinatorId},
		Xmlns:                       namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "StartTransmissionToGroup", body, startTransmissionToGroup)...)
}

func (a AudioInService) StopTransmissionToGroup(ctx context.Context, coordinatorId MemberID) error {
	type StopTransmissionToGroupReq struct {
		CoordinatorID MemberID
	}

	type StopTransmissionToGroupNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StopTransmissionToGroupReq
	}

	type StopTransmissionToGroupBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		StopTransmissionToGroupNode StopTransmissionToGroupNode `xml:"u:StopTransmissionToGroup,omitempty"`
	}

	body := StopTransmissionToGroupBody{StopTransmissionToGroupNode: StopTransmissionToGroupNode{
		StopTransmissionToGroupReq: StopTransmissionToGroupReq{CoordinatorID: coordinatorId},
		Xmlns:                      namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "StopTransmissionToGroup", body, nil)...)
}

func (a AudioInService) SetAudioInputAttributes(ctx context.Context, desiredName AudioInputName, desiredIcon Icon) error {
	type SetAudioInputAttributesReq struct {
		DesiredName AudioInputName
		DesiredIcon Icon
	}

	type SetAudioInputAttributesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAudioInputAttributesReq
	}

	type SetAudioInputAttributesBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		SetAudioInputAttributesNode SetAudioInputAttributesNode `xml:"u:SetAudioInputAttributes,omitempty"`
	}

	body := SetAudioInputAttributesBody{SetAudioInputAttributesNode: SetAudioInputAttributesNode{
		SetAudioInputAttributesReq: SetAudioInputAttributesReq{
			DesiredIcon: desiredIcon,
			DesiredName: desiredName,
		},
		Xmlns: namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "SetAudioInputAttributes", body, nil)...)
}

type AudioInGetAudioInputAttributesRes struct {
	CurrentName AudioInputName
	CurrentIcon Icon
}

func (a AudioInService) GetAudioInputAttributes(ctx context.Context, getAudioInputAttributes *AudioInGetAudioInputAttributesRes) error {
	type GetAudioInputAttributesReq struct{}

	type GetAudioInputAttributesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAudioInputAttributesReq
	}

	type GetAudioInputAttributesBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		GetAudioInputAttributesNode GetAudioInputAttributesNode `xml:"u:GetAudioInputAttributes,omitempty"`
	}

	body := GetAudioInputAttributesBody{GetAudioInputAttributesNode: GetAudioInputAttributesNode{
		GetAudioInputAttributesReq: GetAudioInputAttributesReq{},
		Xmlns:                      namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "GetAudioInputAttributes", body, getAudioInputAttributes)...)
}

func (a AudioInService) SetLineInLevel(ctx context.Context, desiredLeftLineInLevel LeftLineInLevel, desiredRightLineInLevel RightLineInLevel) error {
	type SetLineInLevelReq struct {
		DesiredLeftLineInLevel  LeftLineInLevel
		DesiredRightLineInLevel RightLineInLevel
	}

	type SetLineInLevelNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetLineInLevelReq
	}

	type SetLineInLevelBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		SetLineInLevelNode SetLineInLevelNode `xml:"u:SetLineInLevel,omitempty"`
	}

	body := SetLineInLevelBody{SetLineInLevelNode: SetLineInLevelNode{
		SetLineInLevelReq: SetLineInLevelReq{
			DesiredLeftLineInLevel:  desiredLeftLineInLevel,
			DesiredRightLineInLevel: desiredRightLineInLevel,
		},
		Xmlns: namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "SetLineInLevel", body, nil)...)
}

type AudioInGetLineInLevelRes struct {
	CurrentLeftLineInLevel  LeftLineInLevel
	CurrentRightLineInLevel RightLineInLevel
}

func (a AudioInService) GetLineInLevel(ctx context.Context, getLineInLevel *AudioInGetLineInLevelRes) error {
	type GetLineInLevelReq struct{}

	type GetLineInLevelNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetLineInLevelReq
	}

	type GetLineInLevelBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		GetLineInLevelNode GetLineInLevelNode `xml:"u:GetLineInLevel,omitempty"`
	}

	body := GetLineInLevelBody{GetLineInLevelNode: GetLineInLevelNode{
		GetLineInLevelReq: GetLineInLevelReq{},
		Xmlns:             namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "GetLineInLevel", body, getLineInLevel)...)
}

func (a AudioInService) SelectAudio(ctx context.Context, objectId ObjectID) error {
	type SelectAudioReq struct {
		ObjectID ObjectID
	}

	type SelectAudioNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SelectAudioReq
	}

	type SelectAudioBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SelectAudioNode SelectAudioNode `xml:"u:SelectAudio,omitempty"`
	}

	body := SelectAudioBody{SelectAudioNode: SelectAudioNode{
		SelectAudioReq: SelectAudioReq{ObjectID: objectId},
		Xmlns:          namespaceAudioInService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAudioInService, namespaceAudioInService, "SelectAudio", body, nil)...)
}
