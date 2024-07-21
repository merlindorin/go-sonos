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
	// namespaceRenderingControlService is the UPnP namespace for the Sonos service.
	namespaceRenderingControlService = "urn:schemas-upnp-org:service:RenderingControl:1"

	// eventRenderingControlService is the event path for the Sonos event service.
	eventRenderingControlService = "/MediaRenderer/RenderingControl/Event"

	// controlRenderingControlService is the action path for the Sonos service.
	controlRenderingControlService = "/MediaRenderer/RenderingControl/Control"
)

// RenderingControlService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type RenderingControlService struct {
	doer do.Doer
}

// NewRenderingControlService initializes a new RenderingControlService with the provided Doer.
func NewRenderingControlService(doer do.Doer) *RenderingControlService {
	return &RenderingControlService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (r RenderingControlService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return r.doer.Do(ctx, do.WithPath(eventRenderingControlService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type RenderingControlGetMuteRes struct {
	CurrentMute Mute
}

func (r RenderingControlService) GetMute(ctx context.Context, channel MuteChannel, getMute *RenderingControlGetMuteRes) error {
	type GetMuteReq struct {
		InstanceID InstanceID
		Channel    MuteChannel
	}

	type GetMuteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetMuteReq
	}

	type GetMuteBody struct {
		XMLName     xml.Name    `xml:"s:Body"`
		GetMuteNode GetMuteNode `xml:"u:GetMute,omitempty"`
	}

	body := GetMuteBody{GetMuteNode: GetMuteNode{
		GetMuteReq: GetMuteReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetMute", body, getMute)...)
}

func (r RenderingControlService) SetMute(ctx context.Context, channel MuteChannel, desiredMute Mute) error {
	type SetMuteReq struct {
		InstanceID  InstanceID
		Channel     MuteChannel
		DesiredMute Mute
	}

	type SetMuteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetMuteReq
	}

	type SetMuteBody struct {
		XMLName     xml.Name    `xml:"s:Body"`
		SetMuteNode SetMuteNode `xml:"u:SetMute,omitempty"`
	}

	body := SetMuteBody{SetMuteNode: SetMuteNode{
		SetMuteReq: SetMuteReq{
			Channel:     channel,
			DesiredMute: desiredMute,
			InstanceID:  0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetMute", body, nil)...)
}

type RenderingControlResetBasicEQRes struct {
	Bass        Bass
	Treble      Treble
	Loudness    Loudness
	LeftVolume  LeftVolume
	RightVolume RightVolume
}

func (r RenderingControlService) ResetBasicEQ(ctx context.Context, resetBasicEq *RenderingControlResetBasicEQRes) error {
	type ResetBasicEQReq struct {
		InstanceID InstanceID
	}

	type ResetBasicEQNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ResetBasicEQReq
	}

	type ResetBasicEQBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		ResetBasicEQNode ResetBasicEQNode `xml:"u:ResetBasicEQ,omitempty"`
	}

	body := ResetBasicEQBody{ResetBasicEQNode: ResetBasicEQNode{
		ResetBasicEQReq: ResetBasicEQReq{InstanceID: 0},
		Xmlns:           namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "ResetBasicEQ", body, resetBasicEq)...)
}

func (r RenderingControlService) ResetExtEQ(ctx context.Context, eqtype EQType) error {
	type ResetExtEQReq struct {
		InstanceID InstanceID
		EQType     EQType
	}

	type ResetExtEQNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ResetExtEQReq
	}

	type ResetExtEQBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		ResetExtEQNode ResetExtEQNode `xml:"u:ResetExtEQ,omitempty"`
	}

	body := ResetExtEQBody{ResetExtEQNode: ResetExtEQNode{
		ResetExtEQReq: ResetExtEQReq{
			EQType:     eqtype,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "ResetExtEQ", body, nil)...)
}

type RenderingControlGetVolumeRes struct {
	CurrentVolume Volume
}

// GetVolume Get volume
func (r RenderingControlService) GetVolume(ctx context.Context, channel Channel, getVolume *RenderingControlGetVolumeRes) error {
	type GetVolumeReq struct {
		InstanceID InstanceID
		Channel    Channel
	}

	type GetVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetVolumeReq
	}

	type GetVolumeBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		GetVolumeNode GetVolumeNode `xml:"u:GetVolume,omitempty"`
	}

	body := GetVolumeBody{GetVolumeNode: GetVolumeNode{
		GetVolumeReq: GetVolumeReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetVolume", body, getVolume)...)
}

func (r RenderingControlService) SetVolume(ctx context.Context, channel Channel, desiredVolume Volume) error {
	type SetVolumeReq struct {
		InstanceID    InstanceID
		Channel       Channel
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
			Channel:       channel,
			DesiredVolume: desiredVolume,
			InstanceID:    0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetVolume", body, nil)...)
}

type RenderingControlSetRelativeVolumeRes struct {
	NewVolume Volume
}

func (r RenderingControlService) SetRelativeVolume(ctx context.Context, channel Channel, adjustment VolumeAdjustment, setRelativeVolume *RenderingControlSetRelativeVolumeRes) error {
	type SetRelativeVolumeReq struct {
		InstanceID InstanceID
		Channel    Channel
		Adjustment VolumeAdjustment
	}

	type SetRelativeVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetRelativeVolumeReq
	}

	type SetRelativeVolumeBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		SetRelativeVolumeNode SetRelativeVolumeNode `xml:"u:SetRelativeVolume,omitempty"`
	}

	body := SetRelativeVolumeBody{SetRelativeVolumeNode: SetRelativeVolumeNode{
		SetRelativeVolumeReq: SetRelativeVolumeReq{
			Adjustment: adjustment,
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetRelativeVolume", body, setRelativeVolume)...)
}

type RenderingControlGetVolumeDBRes struct {
	CurrentVolume VolumeDB
}

func (r RenderingControlService) GetVolumeDB(ctx context.Context, channel Channel, getVolumeDb *RenderingControlGetVolumeDBRes) error {
	type GetVolumeDBReq struct {
		InstanceID InstanceID
		Channel    Channel
	}

	type GetVolumeDBNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetVolumeDBReq
	}

	type GetVolumeDBBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		GetVolumeDBNode GetVolumeDBNode `xml:"u:GetVolumeDB,omitempty"`
	}

	body := GetVolumeDBBody{GetVolumeDBNode: GetVolumeDBNode{
		GetVolumeDBReq: GetVolumeDBReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetVolumeDB", body, getVolumeDb)...)
}

func (r RenderingControlService) SetVolumeDB(ctx context.Context, channel Channel, desiredVolume VolumeDB) error {
	type SetVolumeDBReq struct {
		InstanceID    InstanceID
		Channel       Channel
		DesiredVolume VolumeDB
	}

	type SetVolumeDBNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetVolumeDBReq
	}

	type SetVolumeDBBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SetVolumeDBNode SetVolumeDBNode `xml:"u:SetVolumeDB,omitempty"`
	}

	body := SetVolumeDBBody{SetVolumeDBNode: SetVolumeDBNode{
		SetVolumeDBReq: SetVolumeDBReq{
			Channel:       channel,
			DesiredVolume: desiredVolume,
			InstanceID:    0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetVolumeDB", body, nil)...)
}

type RenderingControlGetVolumeDBRangeRes struct {
	MinValue VolumeDB
	MaxValue VolumeDB
}

func (r RenderingControlService) GetVolumeDBRange(ctx context.Context, channel Channel, getVolumeDbrange *RenderingControlGetVolumeDBRangeRes) error {
	type GetVolumeDBRangeReq struct {
		InstanceID InstanceID
		Channel    Channel
	}

	type GetVolumeDBRangeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetVolumeDBRangeReq
	}

	type GetVolumeDBRangeBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		GetVolumeDBRangeNode GetVolumeDBRangeNode `xml:"u:GetVolumeDBRange,omitempty"`
	}

	body := GetVolumeDBRangeBody{GetVolumeDBRangeNode: GetVolumeDBRangeNode{
		GetVolumeDBRangeReq: GetVolumeDBRangeReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetVolumeDBRange", body, getVolumeDbrange)...)
}

type RenderingControlGetBassRes struct {
	CurrentBass Bass
}

// GetBass Get bass level between -10 and 10
func (r RenderingControlService) GetBass(ctx context.Context, getBass *RenderingControlGetBassRes) error {
	type GetBassReq struct {
		InstanceID InstanceID
	}

	type GetBassNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetBassReq
	}

	type GetBassBody struct {
		XMLName     xml.Name    `xml:"s:Body"`
		GetBassNode GetBassNode `xml:"u:GetBass,omitempty"`
	}

	body := GetBassBody{GetBassNode: GetBassNode{
		GetBassReq: GetBassReq{InstanceID: 0},
		Xmlns:      namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetBass", body, getBass)...)
}

// SetBass Set bass level, between -10 and 10
func (r RenderingControlService) SetBass(ctx context.Context, desiredBass Bass) error {
	type SetBassReq struct {
		InstanceID  InstanceID
		DesiredBass Bass
	}

	type SetBassNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetBassReq
	}

	type SetBassBody struct {
		XMLName     xml.Name    `xml:"s:Body"`
		SetBassNode SetBassNode `xml:"u:SetBass,omitempty"`
	}

	body := SetBassBody{SetBassNode: SetBassNode{
		SetBassReq: SetBassReq{
			DesiredBass: desiredBass,
			InstanceID:  0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetBass", body, nil)...)
}

type RenderingControlGetTrebleRes struct {
	CurrentTreble Treble
}

// GetTreble Get treble
func (r RenderingControlService) GetTreble(ctx context.Context, getTreble *RenderingControlGetTrebleRes) error {
	type GetTrebleReq struct {
		InstanceID InstanceID
	}

	type GetTrebleNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTrebleReq
	}

	type GetTrebleBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		GetTrebleNode GetTrebleNode `xml:"u:GetTreble,omitempty"`
	}

	body := GetTrebleBody{GetTrebleNode: GetTrebleNode{
		GetTrebleReq: GetTrebleReq{InstanceID: 0},
		Xmlns:        namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetTreble", body, getTreble)...)
}

// SetTreble Set treble level
func (r RenderingControlService) SetTreble(ctx context.Context, desiredTreble Treble) error {
	type SetTrebleReq struct {
		InstanceID    InstanceID
		DesiredTreble Treble
	}

	type SetTrebleNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetTrebleReq
	}

	type SetTrebleBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		SetTrebleNode SetTrebleNode `xml:"u:SetTreble,omitempty"`
	}

	body := SetTrebleBody{SetTrebleNode: SetTrebleNode{
		SetTrebleReq: SetTrebleReq{
			DesiredTreble: desiredTreble,
			InstanceID:    0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetTreble", body, nil)...)
}

type RenderingControlGetEQRes struct {
	CurrentValue EQValue
}

// GetEQ Get equalizer value
func (r RenderingControlService) GetEQ(ctx context.Context, eqtype EQType, getEq *RenderingControlGetEQRes) error {
	type GetEQReq struct {
		InstanceID InstanceID
		EQType     EQType
	}

	type GetEQNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetEQReq
	}

	type GetEQBody struct {
		XMLName   xml.Name  `xml:"s:Body"`
		GetEQNode GetEQNode `xml:"u:GetEQ,omitempty"`
	}

	body := GetEQBody{GetEQNode: GetEQNode{
		GetEQReq: GetEQReq{
			EQType:     eqtype,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetEQ", body, getEq)...)
}

// SetEQ Set equalizer value for different types
func (r RenderingControlService) SetEQ(ctx context.Context, eqtype EQType, desiredValue EQValue) error {
	type SetEQReq struct {
		InstanceID   InstanceID
		EQType       EQType
		DesiredValue EQValue
	}

	type SetEQNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetEQReq
	}

	type SetEQBody struct {
		XMLName   xml.Name  `xml:"s:Body"`
		SetEQNode SetEQNode `xml:"u:SetEQ,omitempty"`
	}

	body := SetEQBody{SetEQNode: SetEQNode{
		SetEQReq: SetEQReq{
			DesiredValue: desiredValue,
			EQType:       eqtype,
			InstanceID:   0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetEQ", body, nil)...)
}

type RenderingControlGetLoudnessRes struct {
	CurrentLoudness Loudness
}

// GetLoudness Whether or not Loudness is on
func (r RenderingControlService) GetLoudness(ctx context.Context, channel Channel, getLoudness *RenderingControlGetLoudnessRes) error {
	type GetLoudnessReq struct {
		InstanceID InstanceID
		Channel    Channel
	}

	type GetLoudnessNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetLoudnessReq
	}

	type GetLoudnessBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		GetLoudnessNode GetLoudnessNode `xml:"u:GetLoudness,omitempty"`
	}

	body := GetLoudnessBody{GetLoudnessNode: GetLoudnessNode{
		GetLoudnessReq: GetLoudnessReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetLoudness", body, getLoudness)...)
}

// SetLoudness Set loudness on / off
func (r RenderingControlService) SetLoudness(ctx context.Context, channel Channel, desiredLoudness Loudness) error {
	type SetLoudnessReq struct {
		InstanceID      InstanceID
		Channel         Channel
		DesiredLoudness Loudness
	}

	type SetLoudnessNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetLoudnessReq
	}

	type SetLoudnessBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SetLoudnessNode SetLoudnessNode `xml:"u:SetLoudness,omitempty"`
	}

	body := SetLoudnessBody{SetLoudnessNode: SetLoudnessNode{
		SetLoudnessReq: SetLoudnessReq{
			Channel:         channel,
			DesiredLoudness: desiredLoudness,
			InstanceID:      0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetLoudness", body, nil)...)
}

type RenderingControlGetSupportsOutputFixedRes struct {
	CurrentSupportsFixed SupportsOutputFixed
}

func (r RenderingControlService) GetSupportsOutputFixed(ctx context.Context, getSupportsOutputFixed *RenderingControlGetSupportsOutputFixedRes) error {
	type GetSupportsOutputFixedReq struct {
		InstanceID InstanceID
	}

	type GetSupportsOutputFixedNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetSupportsOutputFixedReq
	}

	type GetSupportsOutputFixedBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		GetSupportsOutputFixedNode GetSupportsOutputFixedNode `xml:"u:GetSupportsOutputFixed,omitempty"`
	}

	body := GetSupportsOutputFixedBody{GetSupportsOutputFixedNode: GetSupportsOutputFixedNode{
		GetSupportsOutputFixedReq: GetSupportsOutputFixedReq{InstanceID: 0},
		Xmlns:                     namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetSupportsOutputFixed", body, getSupportsOutputFixed)...)
}

type RenderingControlGetOutputFixedRes struct {
	CurrentFixed OutputFixed
}

func (r RenderingControlService) GetOutputFixed(ctx context.Context, getOutputFixed *RenderingControlGetOutputFixedRes) error {
	type GetOutputFixedReq struct {
		InstanceID InstanceID
	}

	type GetOutputFixedNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetOutputFixedReq
	}

	type GetOutputFixedBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		GetOutputFixedNode GetOutputFixedNode `xml:"u:GetOutputFixed,omitempty"`
	}

	body := GetOutputFixedBody{GetOutputFixedNode: GetOutputFixedNode{
		GetOutputFixedReq: GetOutputFixedReq{InstanceID: 0},
		Xmlns:             namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetOutputFixed", body, getOutputFixed)...)
}

func (r RenderingControlService) SetOutputFixed(ctx context.Context, desiredFixed OutputFixed) error {
	type SetOutputFixedReq struct {
		InstanceID   InstanceID
		DesiredFixed OutputFixed
	}

	type SetOutputFixedNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetOutputFixedReq
	}

	type SetOutputFixedBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		SetOutputFixedNode SetOutputFixedNode `xml:"u:SetOutputFixed,omitempty"`
	}

	body := SetOutputFixedBody{SetOutputFixedNode: SetOutputFixedNode{
		SetOutputFixedReq: SetOutputFixedReq{
			DesiredFixed: desiredFixed,
			InstanceID:   0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetOutputFixed", body, nil)...)
}

type RenderingControlGetHeadphoneConnectedRes struct {
	CurrentHeadphoneConnected HeadphoneConnected
}

func (r RenderingControlService) GetHeadphoneConnected(ctx context.Context, getHeadphoneConnected *RenderingControlGetHeadphoneConnectedRes) error {
	type GetHeadphoneConnectedReq struct {
		InstanceID InstanceID
	}

	type GetHeadphoneConnectedNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetHeadphoneConnectedReq
	}

	type GetHeadphoneConnectedBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		GetHeadphoneConnectedNode GetHeadphoneConnectedNode `xml:"u:GetHeadphoneConnected,omitempty"`
	}

	body := GetHeadphoneConnectedBody{GetHeadphoneConnectedNode: GetHeadphoneConnectedNode{
		GetHeadphoneConnectedReq: GetHeadphoneConnectedReq{InstanceID: 0},
		Xmlns:                    namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetHeadphoneConnected", body, getHeadphoneConnected)...)
}

type RenderingControlRampToVolumeRes struct {
	RampTime RampTimeSeconds
}

func (r RenderingControlService) RampToVolume(ctx context.Context, channel Channel, rampType RampType, desiredVolume Volume, resetVolumeAfter ResetVolumeAfter, programUri ProgramURI, rampToVolume *RenderingControlRampToVolumeRes) error {
	type RampToVolumeReq struct {
		InstanceID       InstanceID
		Channel          Channel
		RampType         RampType
		DesiredVolume    Volume
		ResetVolumeAfter ResetVolumeAfter
		ProgramURI       ProgramURI
	}

	type RampToVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RampToVolumeReq
	}

	type RampToVolumeBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		RampToVolumeNode RampToVolumeNode `xml:"u:RampToVolume,omitempty"`
	}

	body := RampToVolumeBody{RampToVolumeNode: RampToVolumeNode{
		RampToVolumeReq: RampToVolumeReq{
			Channel:          channel,
			DesiredVolume:    desiredVolume,
			InstanceID:       0,
			ProgramURI:       programUri,
			RampType:         rampType,
			ResetVolumeAfter: resetVolumeAfter,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "RampToVolume", body, rampToVolume)...)
}

func (r RenderingControlService) RestoreVolumePriorToRamp(ctx context.Context, channel Channel) error {
	type RestoreVolumePriorToRampReq struct {
		InstanceID InstanceID
		Channel    Channel
	}

	type RestoreVolumePriorToRampNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RestoreVolumePriorToRampReq
	}

	type RestoreVolumePriorToRampBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		RestoreVolumePriorToRampNode RestoreVolumePriorToRampNode `xml:"u:RestoreVolumePriorToRamp,omitempty"`
	}

	body := RestoreVolumePriorToRampBody{RestoreVolumePriorToRampNode: RestoreVolumePriorToRampNode{
		RestoreVolumePriorToRampReq: RestoreVolumePriorToRampReq{
			Channel:    channel,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "RestoreVolumePriorToRamp", body, nil)...)
}

func (r RenderingControlService) SetChannelMap(ctx context.Context, channelMap ChannelMap) error {
	type SetChannelMapReq struct {
		InstanceID InstanceID
		ChannelMap ChannelMap
	}

	type SetChannelMapNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetChannelMapReq
	}

	type SetChannelMapBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		SetChannelMapNode SetChannelMapNode `xml:"u:SetChannelMap,omitempty"`
	}

	body := SetChannelMapBody{SetChannelMapNode: SetChannelMapNode{
		SetChannelMapReq: SetChannelMapReq{
			ChannelMap: channelMap,
			InstanceID: 0,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetChannelMap", body, nil)...)
}

type RenderingControlGetRoomCalibrationStatusRes struct {
	RoomCalibrationEnabled   RoomCalibrationEnabled
	RoomCalibrationAvailable RoomCalibrationAvailable
}

func (r RenderingControlService) GetRoomCalibrationStatus(ctx context.Context, getRoomCalibrationStatus *RenderingControlGetRoomCalibrationStatusRes) error {
	type GetRoomCalibrationStatusReq struct {
		InstanceID InstanceID
	}

	type GetRoomCalibrationStatusNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetRoomCalibrationStatusReq
	}

	type GetRoomCalibrationStatusBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		GetRoomCalibrationStatusNode GetRoomCalibrationStatusNode `xml:"u:GetRoomCalibrationStatus,omitempty"`
	}

	body := GetRoomCalibrationStatusBody{GetRoomCalibrationStatusNode: GetRoomCalibrationStatusNode{
		GetRoomCalibrationStatusReq: GetRoomCalibrationStatusReq{InstanceID: 0},
		Xmlns:                       namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "GetRoomCalibrationStatus", body, getRoomCalibrationStatus)...)
}

func (r RenderingControlService) SetRoomCalibrationStatus(ctx context.Context, roomCalibrationEnabled RoomCalibrationEnabled) error {
	type SetRoomCalibrationStatusReq struct {
		InstanceID             InstanceID
		RoomCalibrationEnabled RoomCalibrationEnabled
	}

	type SetRoomCalibrationStatusNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetRoomCalibrationStatusReq
	}

	type SetRoomCalibrationStatusBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		SetRoomCalibrationStatusNode SetRoomCalibrationStatusNode `xml:"u:SetRoomCalibrationStatus,omitempty"`
	}

	body := SetRoomCalibrationStatusBody{SetRoomCalibrationStatusNode: SetRoomCalibrationStatusNode{
		SetRoomCalibrationStatusReq: SetRoomCalibrationStatusReq{
			InstanceID:             0,
			RoomCalibrationEnabled: roomCalibrationEnabled,
		},
		Xmlns: namespaceRenderingControlService,
	}}

	return r.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlRenderingControlService, namespaceRenderingControlService, "SetRoomCalibrationStatus", body, nil)...)
}
