// CODE GENERATED, DO NOT EDIT

package S14

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceHTControlService is the UPnP namespace for the Sonos service.
	namespaceHTControlService = "urn:schemas-upnp-org:service:HTControl:1"

	// eventHTControlService is the event path for the Sonos event service.
	eventHTControlService = "/HTControl/Event"

	// controlHTControlService is the action path for the Sonos service.
	controlHTControlService = "/HTControl/Control"
)

// HTControlService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type HTControlService struct {
	doer do.Doer
}

// NewHTControlService initializes a new HTControlService with the provided Doer.
func NewHTControlService(doer do.Doer) *HTControlService {
	return &HTControlService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (h HTControlService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return h.doer.Do(ctx, do.WithPath(eventHTControlService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

func (h HTControlService) SetIRRepeaterState(ctx context.Context, desiredIrrepeaterState IRRepeaterState) error {
	type SetIRRepeaterStateReq struct {
		DesiredIRRepeaterState IRRepeaterState
	}

	type SetIRRepeaterStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetIRRepeaterStateReq
	}

	type SetIRRepeaterStateBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		SetIRRepeaterStateNode SetIRRepeaterStateNode `xml:"u:SetIRRepeaterState,omitempty"`
	}

	body := SetIRRepeaterStateBody{SetIRRepeaterStateNode: SetIRRepeaterStateNode{
		SetIRRepeaterStateReq: SetIRRepeaterStateReq{DesiredIRRepeaterState: desiredIrrepeaterState},
		Xmlns:                 namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "SetIRRepeaterState", body, nil)...)
}

type HTControlGetIRRepeaterStateRes struct {
	CurrentIRRepeaterState IRRepeaterState
}

func (h HTControlService) GetIRRepeaterState(ctx context.Context, getIrrepeaterState *HTControlGetIRRepeaterStateRes) error {
	type GetIRRepeaterStateReq struct{}

	type GetIRRepeaterStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetIRRepeaterStateReq
	}

	type GetIRRepeaterStateBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		GetIRRepeaterStateNode GetIRRepeaterStateNode `xml:"u:GetIRRepeaterState,omitempty"`
	}

	body := GetIRRepeaterStateBody{GetIRRepeaterStateNode: GetIRRepeaterStateNode{
		GetIRRepeaterStateReq: GetIRRepeaterStateReq{},
		Xmlns:                 namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "GetIRRepeaterState", body, getIrrepeaterState)...)
}

func (h HTControlService) IdentifyIRRemote(ctx context.Context, timeout Timeout) error {
	type IdentifyIRRemoteReq struct {
		Timeout Timeout
	}

	type IdentifyIRRemoteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		IdentifyIRRemoteReq
	}

	type IdentifyIRRemoteBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		IdentifyIRRemoteNode IdentifyIRRemoteNode `xml:"u:IdentifyIRRemote,omitempty"`
	}

	body := IdentifyIRRemoteBody{IdentifyIRRemoteNode: IdentifyIRRemoteNode{
		IdentifyIRRemoteReq: IdentifyIRRemoteReq{Timeout: timeout},
		Xmlns:               namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "IdentifyIRRemote", body, nil)...)
}

func (h HTControlService) LearnIRCode(ctx context.Context, ircode IRCode, timeout Timeout) error {
	type LearnIRCodeReq struct {
		IRCode  IRCode
		Timeout Timeout
	}

	type LearnIRCodeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		LearnIRCodeReq
	}

	type LearnIRCodeBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		LearnIRCodeNode LearnIRCodeNode `xml:"u:LearnIRCode,omitempty"`
	}

	body := LearnIRCodeBody{LearnIRCodeNode: LearnIRCodeNode{
		LearnIRCodeReq: LearnIRCodeReq{
			IRCode:  ircode,
			Timeout: timeout,
		},
		Xmlns: namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "LearnIRCode", body, nil)...)
}

func (h HTControlService) CommitLearnedIRCodes(ctx context.Context, name IRRemoteName) error {
	type CommitLearnedIRCodesReq struct {
		Name IRRemoteName
	}

	type CommitLearnedIRCodesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CommitLearnedIRCodesReq
	}

	type CommitLearnedIRCodesBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		CommitLearnedIRCodesNode CommitLearnedIRCodesNode `xml:"u:CommitLearnedIRCodes,omitempty"`
	}

	body := CommitLearnedIRCodesBody{CommitLearnedIRCodesNode: CommitLearnedIRCodesNode{
		CommitLearnedIRCodesReq: CommitLearnedIRCodesReq{Name: name},
		Xmlns:                   namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "CommitLearnedIRCodes", body, nil)...)
}

type HTControlIsRemoteConfiguredRes struct {
	RemoteConfigured RemoteConfigured
}

func (h HTControlService) IsRemoteConfigured(ctx context.Context, isRemoteConfigured *HTControlIsRemoteConfiguredRes) error {
	type IsRemoteConfiguredReq struct{}

	type IsRemoteConfiguredNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		IsRemoteConfiguredReq
	}

	type IsRemoteConfiguredBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		IsRemoteConfiguredNode IsRemoteConfiguredNode `xml:"u:IsRemoteConfigured,omitempty"`
	}

	body := IsRemoteConfiguredBody{IsRemoteConfiguredNode: IsRemoteConfiguredNode{
		IsRemoteConfiguredReq: IsRemoteConfiguredReq{},
		Xmlns:                 namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "IsRemoteConfigured", body, isRemoteConfigured)...)
}

func (h HTControlService) SetLEDFeedbackState(ctx context.Context, ledfeedbackState LEDFeedbackState) error {
	type SetLEDFeedbackStateReq struct {
		LEDFeedbackState LEDFeedbackState
	}

	type SetLEDFeedbackStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetLEDFeedbackStateReq
	}

	type SetLEDFeedbackStateBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		SetLEDFeedbackStateNode SetLEDFeedbackStateNode `xml:"u:SetLEDFeedbackState,omitempty"`
	}

	body := SetLEDFeedbackStateBody{SetLEDFeedbackStateNode: SetLEDFeedbackStateNode{
		SetLEDFeedbackStateReq: SetLEDFeedbackStateReq{LEDFeedbackState: ledfeedbackState},
		Xmlns:                  namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "SetLEDFeedbackState", body, nil)...)
}

type HTControlGetLEDFeedbackStateRes struct {
	LEDFeedbackState LEDFeedbackState
}

func (h HTControlService) GetLEDFeedbackState(ctx context.Context, getLedfeedbackState *HTControlGetLEDFeedbackStateRes) error {
	type GetLEDFeedbackStateReq struct{}

	type GetLEDFeedbackStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetLEDFeedbackStateReq
	}

	type GetLEDFeedbackStateBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		GetLEDFeedbackStateNode GetLEDFeedbackStateNode `xml:"u:GetLEDFeedbackState,omitempty"`
	}

	body := GetLEDFeedbackStateBody{GetLEDFeedbackStateNode: GetLEDFeedbackStateNode{
		GetLEDFeedbackStateReq: GetLEDFeedbackStateReq{},
		Xmlns:                  namespaceHTControlService,
	}}

	return h.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlHTControlService, namespaceHTControlService, "GetLEDFeedbackState", body, getLedfeedbackState)...)
}
