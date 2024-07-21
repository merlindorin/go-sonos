// CODE GENERATED, DO NOT EDIT

package S33

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceDevicePropertiesService is the UPnP namespace for the Sonos service.
	namespaceDevicePropertiesService = "urn:schemas-upnp-org:service:DeviceProperties:1"

	// eventDevicePropertiesService is the event path for the Sonos event service.
	eventDevicePropertiesService = "/DeviceProperties/Event"

	// controlDevicePropertiesService is the action path for the Sonos service.
	controlDevicePropertiesService = "/DeviceProperties/Control"
)

// DevicePropertiesService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type DevicePropertiesService struct {
	doer do.Doer
}

// NewDevicePropertiesService initializes a new DevicePropertiesService with the provided Doer.
func NewDevicePropertiesService(doer do.Doer) *DevicePropertiesService {
	return &DevicePropertiesService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (d DevicePropertiesService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return d.doer.Do(ctx, do.WithPath(eventDevicePropertiesService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

// SetLEDState Set the LED state
func (d DevicePropertiesService) SetLEDState(ctx context.Context, desiredLedstate LEDState) error {
	type SetLEDStateReq struct {
		DesiredLEDState LEDState
	}

	type SetLEDStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetLEDStateReq
	}

	type SetLEDStateBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SetLEDStateNode SetLEDStateNode `xml:"u:SetLEDState,omitempty"`
	}

	body := SetLEDStateBody{SetLEDStateNode: SetLEDStateNode{
		SetLEDStateReq: SetLEDStateReq{DesiredLEDState: desiredLedstate},
		Xmlns:          namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetLEDState", body, nil)...)
}

type DevicePropertiesGetLEDStateRes struct {
	CurrentLEDState LEDState
}

// GetLEDState Get the current LED state
func (d DevicePropertiesService) GetLEDState(ctx context.Context, getLedstate *DevicePropertiesGetLEDStateRes) error {
	type GetLEDStateReq struct{}

	type GetLEDStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetLEDStateReq
	}

	type GetLEDStateBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		GetLEDStateNode GetLEDStateNode `xml:"u:GetLEDState,omitempty"`
	}

	body := GetLEDStateBody{GetLEDStateNode: GetLEDStateNode{
		GetLEDStateReq: GetLEDStateReq{},
		Xmlns:          namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetLEDState", body, getLedstate)...)
}

func (d DevicePropertiesService) AddBondedZones(ctx context.Context, channelMapSet ChannelMapSet) error {
	type AddBondedZonesReq struct {
		ChannelMapSet ChannelMapSet
	}

	type AddBondedZonesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddBondedZonesReq
	}

	type AddBondedZonesBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		AddBondedZonesNode AddBondedZonesNode `xml:"u:AddBondedZones,omitempty"`
	}

	body := AddBondedZonesBody{AddBondedZonesNode: AddBondedZonesNode{
		AddBondedZonesReq: AddBondedZonesReq{ChannelMapSet: channelMapSet},
		Xmlns:             namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "AddBondedZones", body, nil)...)
}

func (d DevicePropertiesService) RemoveBondedZones(ctx context.Context, channelMapSet ChannelMapSet, keepGrouped KeepGrouped) error {
	type RemoveBondedZonesReq struct {
		ChannelMapSet ChannelMapSet
		KeepGrouped   KeepGrouped
	}

	type RemoveBondedZonesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveBondedZonesReq
	}

	type RemoveBondedZonesBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		RemoveBondedZonesNode RemoveBondedZonesNode `xml:"u:RemoveBondedZones,omitempty"`
	}

	body := RemoveBondedZonesBody{RemoveBondedZonesNode: RemoveBondedZonesNode{
		RemoveBondedZonesReq: RemoveBondedZonesReq{
			ChannelMapSet: channelMapSet,
			KeepGrouped:   keepGrouped,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "RemoveBondedZones", body, nil)...)
}

// CreateStereoPair Create a stereo pair (left, right speakers), right one becomes hidden
func (d DevicePropertiesService) CreateStereoPair(ctx context.Context, channelMapSet ChannelMapSet) error {
	type CreateStereoPairReq struct {
		ChannelMapSet ChannelMapSet
	}

	type CreateStereoPairNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CreateStereoPairReq
	}

	type CreateStereoPairBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		CreateStereoPairNode CreateStereoPairNode `xml:"u:CreateStereoPair,omitempty"`
	}

	body := CreateStereoPairBody{CreateStereoPairNode: CreateStereoPairNode{
		CreateStereoPairReq: CreateStereoPairReq{ChannelMapSet: channelMapSet},
		Xmlns:               namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "CreateStereoPair", body, nil)...)
}

// SeparateStereoPair Separate a stereo pair
func (d DevicePropertiesService) SeparateStereoPair(ctx context.Context, channelMapSet ChannelMapSet) error {
	type SeparateStereoPairReq struct {
		ChannelMapSet ChannelMapSet
	}

	type SeparateStereoPairNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SeparateStereoPairReq
	}

	type SeparateStereoPairBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		SeparateStereoPairNode SeparateStereoPairNode `xml:"u:SeparateStereoPair,omitempty"`
	}

	body := SeparateStereoPairBody{SeparateStereoPairNode: SeparateStereoPairNode{
		SeparateStereoPairReq: SeparateStereoPairReq{ChannelMapSet: channelMapSet},
		Xmlns:                 namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SeparateStereoPair", body, nil)...)
}

func (d DevicePropertiesService) SetZoneAttributes(ctx context.Context, desiredZoneName ZoneName, desiredIcon Icon, desiredConfiguration Configuration, desiredTargetRoomName TargetRoomName) error {
	type SetZoneAttributesReq struct {
		DesiredZoneName       ZoneName
		DesiredIcon           Icon
		DesiredConfiguration  Configuration
		DesiredTargetRoomName TargetRoomName
	}

	type SetZoneAttributesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetZoneAttributesReq
	}

	type SetZoneAttributesBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		SetZoneAttributesNode SetZoneAttributesNode `xml:"u:SetZoneAttributes,omitempty"`
	}

	body := SetZoneAttributesBody{SetZoneAttributesNode: SetZoneAttributesNode{
		SetZoneAttributesReq: SetZoneAttributesReq{
			DesiredConfiguration:  desiredConfiguration,
			DesiredIcon:           desiredIcon,
			DesiredTargetRoomName: desiredTargetRoomName,
			DesiredZoneName:       desiredZoneName,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetZoneAttributes", body, nil)...)
}

type DevicePropertiesGetZoneAttributesRes struct {
	CurrentZoneName       ZoneName
	CurrentIcon           Icon
	CurrentConfiguration  Configuration
	CurrentTargetRoomName TargetRoomName
}

func (d DevicePropertiesService) GetZoneAttributes(ctx context.Context, getZoneAttributes *DevicePropertiesGetZoneAttributesRes) error {
	type GetZoneAttributesReq struct{}

	type GetZoneAttributesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetZoneAttributesReq
	}

	type GetZoneAttributesBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		GetZoneAttributesNode GetZoneAttributesNode `xml:"u:GetZoneAttributes,omitempty"`
	}

	body := GetZoneAttributesBody{GetZoneAttributesNode: GetZoneAttributesNode{
		GetZoneAttributesReq: GetZoneAttributesReq{},
		Xmlns:                namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetZoneAttributes", body, getZoneAttributes)...)
}

type DevicePropertiesGetHouseholdIDRes struct {
	CurrentHouseholdID HouseholdID
}

func (d DevicePropertiesService) GetHouseholdID(ctx context.Context, getHouseholdId *DevicePropertiesGetHouseholdIDRes) error {
	type GetHouseholdIDReq struct{}

	type GetHouseholdIDNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetHouseholdIDReq
	}

	type GetHouseholdIDBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		GetHouseholdIDNode GetHouseholdIDNode `xml:"u:GetHouseholdID,omitempty"`
	}

	body := GetHouseholdIDBody{GetHouseholdIDNode: GetHouseholdIDNode{
		GetHouseholdIDReq: GetHouseholdIDReq{},
		Xmlns:             namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetHouseholdID", body, getHouseholdId)...)
}

type DevicePropertiesGetZoneInfoRes struct {
	SerialNumber           SerialNumber
	SoftwareVersion        SoftwareVersion
	DisplaySoftwareVersion DisplaySoftwareVersion
	HardwareVersion        HardwareVersion
	IPAddress              IPAddress
	MACAddress             MACAddress
	CopyrightInfo          CopyrightInfo
	ExtraInfo              ExtraInfo
	HTAudioIn              HTAudioIn
	Flags                  Flags
}

// GetZoneInfo Get information about this specific speaker
func (d DevicePropertiesService) GetZoneInfo(ctx context.Context, getZoneInfo *DevicePropertiesGetZoneInfoRes) error {
	type GetZoneInfoReq struct{}

	type GetZoneInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetZoneInfoReq
	}

	type GetZoneInfoBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		GetZoneInfoNode GetZoneInfoNode `xml:"u:GetZoneInfo,omitempty"`
	}

	body := GetZoneInfoBody{GetZoneInfoNode: GetZoneInfoNode{
		GetZoneInfoReq: GetZoneInfoReq{},
		Xmlns:          namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetZoneInfo", body, getZoneInfo)...)
}

func (d DevicePropertiesService) SetAutoplayLinkedZones(ctx context.Context, includeLinkedZones AutoplayIncludeLinkedZones, source AutoplaySource) error {
	type SetAutoplayLinkedZonesReq struct {
		IncludeLinkedZones AutoplayIncludeLinkedZones
		Source             AutoplaySource
	}

	type SetAutoplayLinkedZonesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAutoplayLinkedZonesReq
	}

	type SetAutoplayLinkedZonesBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		SetAutoplayLinkedZonesNode SetAutoplayLinkedZonesNode `xml:"u:SetAutoplayLinkedZones,omitempty"`
	}

	body := SetAutoplayLinkedZonesBody{SetAutoplayLinkedZonesNode: SetAutoplayLinkedZonesNode{
		SetAutoplayLinkedZonesReq: SetAutoplayLinkedZonesReq{
			IncludeLinkedZones: includeLinkedZones,
			Source:             source,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetAutoplayLinkedZones", body, nil)...)
}

type DevicePropertiesGetAutoplayLinkedZonesRes struct {
	IncludeLinkedZones AutoplayIncludeLinkedZones
}

func (d DevicePropertiesService) GetAutoplayLinkedZones(ctx context.Context, source AutoplaySource, getAutoplayLinkedZones *DevicePropertiesGetAutoplayLinkedZonesRes) error {
	type GetAutoplayLinkedZonesReq struct {
		Source AutoplaySource
	}

	type GetAutoplayLinkedZonesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAutoplayLinkedZonesReq
	}

	type GetAutoplayLinkedZonesBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		GetAutoplayLinkedZonesNode GetAutoplayLinkedZonesNode `xml:"u:GetAutoplayLinkedZones,omitempty"`
	}

	body := GetAutoplayLinkedZonesBody{GetAutoplayLinkedZonesNode: GetAutoplayLinkedZonesNode{
		GetAutoplayLinkedZonesReq: GetAutoplayLinkedZonesReq{Source: source},
		Xmlns:                     namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetAutoplayLinkedZones", body, getAutoplayLinkedZones)...)
}

func (d DevicePropertiesService) SetAutoplayRoomUUID(ctx context.Context, roomUuid AutoplayRoomUUID, source AutoplaySource) error {
	type SetAutoplayRoomUUIDReq struct {
		RoomUUID AutoplayRoomUUID
		Source   AutoplaySource
	}

	type SetAutoplayRoomUUIDNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAutoplayRoomUUIDReq
	}

	type SetAutoplayRoomUUIDBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		SetAutoplayRoomUUIDNode SetAutoplayRoomUUIDNode `xml:"u:SetAutoplayRoomUUID,omitempty"`
	}

	body := SetAutoplayRoomUUIDBody{SetAutoplayRoomUUIDNode: SetAutoplayRoomUUIDNode{
		SetAutoplayRoomUUIDReq: SetAutoplayRoomUUIDReq{
			RoomUUID: roomUuid,
			Source:   source,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetAutoplayRoomUUID", body, nil)...)
}

type DevicePropertiesGetAutoplayRoomUUIDRes struct {
	RoomUUID AutoplayRoomUUID
}

func (d DevicePropertiesService) GetAutoplayRoomUUID(ctx context.Context, source AutoplaySource, getAutoplayRoomUuid *DevicePropertiesGetAutoplayRoomUUIDRes) error {
	type GetAutoplayRoomUUIDReq struct {
		Source AutoplaySource
	}

	type GetAutoplayRoomUUIDNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAutoplayRoomUUIDReq
	}

	type GetAutoplayRoomUUIDBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		GetAutoplayRoomUUIDNode GetAutoplayRoomUUIDNode `xml:"u:GetAutoplayRoomUUID,omitempty"`
	}

	body := GetAutoplayRoomUUIDBody{GetAutoplayRoomUUIDNode: GetAutoplayRoomUUIDNode{
		GetAutoplayRoomUUIDReq: GetAutoplayRoomUUIDReq{Source: source},
		Xmlns:                  namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetAutoplayRoomUUID", body, getAutoplayRoomUuid)...)
}

func (d DevicePropertiesService) SetAutoplayVolume(ctx context.Context, volume AutoplayVolume, source AutoplaySource) error {
	type SetAutoplayVolumeReq struct {
		Volume AutoplayVolume
		Source AutoplaySource
	}

	type SetAutoplayVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAutoplayVolumeReq
	}

	type SetAutoplayVolumeBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		SetAutoplayVolumeNode SetAutoplayVolumeNode `xml:"u:SetAutoplayVolume,omitempty"`
	}

	body := SetAutoplayVolumeBody{SetAutoplayVolumeNode: SetAutoplayVolumeNode{
		SetAutoplayVolumeReq: SetAutoplayVolumeReq{
			Source: source,
			Volume: volume,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetAutoplayVolume", body, nil)...)
}

type DevicePropertiesGetAutoplayVolumeRes struct {
	CurrentVolume AutoplayVolume
}

func (d DevicePropertiesService) GetAutoplayVolume(ctx context.Context, source AutoplaySource, getAutoplayVolume *DevicePropertiesGetAutoplayVolumeRes) error {
	type GetAutoplayVolumeReq struct {
		Source AutoplaySource
	}

	type GetAutoplayVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAutoplayVolumeReq
	}

	type GetAutoplayVolumeBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		GetAutoplayVolumeNode GetAutoplayVolumeNode `xml:"u:GetAutoplayVolume,omitempty"`
	}

	body := GetAutoplayVolumeBody{GetAutoplayVolumeNode: GetAutoplayVolumeNode{
		GetAutoplayVolumeReq: GetAutoplayVolumeReq{Source: source},
		Xmlns:                namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetAutoplayVolume", body, getAutoplayVolume)...)
}

func (d DevicePropertiesService) SetUseAutoplayVolume(ctx context.Context, useVolume AutoplayUseVolume, source AutoplaySource) error {
	type SetUseAutoplayVolumeReq struct {
		UseVolume AutoplayUseVolume
		Source    AutoplaySource
	}

	type SetUseAutoplayVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetUseAutoplayVolumeReq
	}

	type SetUseAutoplayVolumeBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		SetUseAutoplayVolumeNode SetUseAutoplayVolumeNode `xml:"u:SetUseAutoplayVolume,omitempty"`
	}

	body := SetUseAutoplayVolumeBody{SetUseAutoplayVolumeNode: SetUseAutoplayVolumeNode{
		SetUseAutoplayVolumeReq: SetUseAutoplayVolumeReq{
			Source:    source,
			UseVolume: useVolume,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetUseAutoplayVolume", body, nil)...)
}

type DevicePropertiesGetUseAutoplayVolumeRes struct {
	UseVolume AutoplayUseVolume
}

func (d DevicePropertiesService) GetUseAutoplayVolume(ctx context.Context, source AutoplaySource, getUseAutoplayVolume *DevicePropertiesGetUseAutoplayVolumeRes) error {
	type GetUseAutoplayVolumeReq struct {
		Source AutoplaySource
	}

	type GetUseAutoplayVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetUseAutoplayVolumeReq
	}

	type GetUseAutoplayVolumeBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		GetUseAutoplayVolumeNode GetUseAutoplayVolumeNode `xml:"u:GetUseAutoplayVolume,omitempty"`
	}

	body := GetUseAutoplayVolumeBody{GetUseAutoplayVolumeNode: GetUseAutoplayVolumeNode{
		GetUseAutoplayVolumeReq: GetUseAutoplayVolumeReq{Source: source},
		Xmlns:                   namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetUseAutoplayVolume", body, getUseAutoplayVolume)...)
}

// AddHTSatellite Adds satellites and/or a sub woofer to a (main) player. The satellites become hidden. The main player RINCON_* is mandatory. RR: right - rear, LF: left - front, SW: subwoofer
func (d DevicePropertiesService) AddHTSatellite(ctx context.Context, htsatChanMapSet HTSatChanMapSet) error {
	type AddHTSatelliteReq struct {
		HTSatChanMapSet HTSatChanMapSet
	}

	type AddHTSatelliteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddHTSatelliteReq
	}

	type AddHTSatelliteBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		AddHTSatelliteNode AddHTSatelliteNode `xml:"u:AddHTSatellite,omitempty"`
	}

	body := AddHTSatelliteBody{AddHTSatelliteNode: AddHTSatelliteNode{
		AddHTSatelliteReq: AddHTSatelliteReq{HTSatChanMapSet: htsatChanMapSet},
		Xmlns:             namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "AddHTSatellite", body, nil)...)
}

// RemoveHTSatellite Removes a satellite or a sub woofer from (main) player. The satellite becomes visible.
func (d DevicePropertiesService) RemoveHTSatellite(ctx context.Context, satRoomUuid SatRoomUUID) error {
	type RemoveHTSatelliteReq struct {
		SatRoomUUID SatRoomUUID
	}

	type RemoveHTSatelliteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveHTSatelliteReq
	}

	type RemoveHTSatelliteBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		RemoveHTSatelliteNode RemoveHTSatelliteNode `xml:"u:RemoveHTSatellite,omitempty"`
	}

	body := RemoveHTSatelliteBody{RemoveHTSatelliteNode: RemoveHTSatelliteNode{
		RemoveHTSatelliteReq: RemoveHTSatelliteReq{SatRoomUUID: satRoomUuid},
		Xmlns:                namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "RemoveHTSatellite", body, nil)...)
}

type DevicePropertiesEnterConfigModeRes struct {
	State ConfigModeState
}

func (d DevicePropertiesService) EnterConfigMode(ctx context.Context, mode ConfigMode, options ConfigModeOptions, enterConfigMode *DevicePropertiesEnterConfigModeRes) error {
	type EnterConfigModeReq struct {
		Mode    ConfigMode
		Options ConfigModeOptions
	}

	type EnterConfigModeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		EnterConfigModeReq
	}

	type EnterConfigModeBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		EnterConfigModeNode EnterConfigModeNode `xml:"u:EnterConfigMode,omitempty"`
	}

	body := EnterConfigModeBody{EnterConfigModeNode: EnterConfigModeNode{
		EnterConfigModeReq: EnterConfigModeReq{
			Mode:    mode,
			Options: options,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "EnterConfigMode", body, enterConfigMode)...)
}

func (d DevicePropertiesService) ExitConfigMode(ctx context.Context, options ConfigModeOptions) error {
	type ExitConfigModeReq struct {
		Options ConfigModeOptions
	}

	type ExitConfigModeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ExitConfigModeReq
	}

	type ExitConfigModeBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		ExitConfigModeNode ExitConfigModeNode `xml:"u:ExitConfigMode,omitempty"`
	}

	body := ExitConfigModeBody{ExitConfigModeNode: ExitConfigModeNode{
		ExitConfigModeReq: ExitConfigModeReq{Options: options},
		Xmlns:             namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "ExitConfigMode", body, nil)...)
}

type DevicePropertiesGetButtonStateRes struct {
	State ButtonState
}

func (d DevicePropertiesService) GetButtonState(ctx context.Context, getButtonState *DevicePropertiesGetButtonStateRes) error {
	type GetButtonStateReq struct{}

	type GetButtonStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetButtonStateReq
	}

	type GetButtonStateBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		GetButtonStateNode GetButtonStateNode `xml:"u:GetButtonState,omitempty"`
	}

	body := GetButtonStateBody{GetButtonStateNode: GetButtonStateNode{
		GetButtonStateReq: GetButtonStateReq{},
		Xmlns:             namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetButtonState", body, getButtonState)...)
}

type DevicePropertiesGetHTForwardStateRes struct {
	IsHTForwardEnabled HTForwardEnabled
}

func (d DevicePropertiesService) GetHTForwardState(ctx context.Context, getHtforwardState *DevicePropertiesGetHTForwardStateRes) error {
	type GetHTForwardStateReq struct{}

	type GetHTForwardStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetHTForwardStateReq
	}

	type GetHTForwardStateBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		GetHTForwardStateNode GetHTForwardStateNode `xml:"u:GetHTForwardState,omitempty"`
	}

	body := GetHTForwardStateBody{GetHTForwardStateNode: GetHTForwardStateNode{
		GetHTForwardStateReq: GetHTForwardStateReq{},
		Xmlns:                namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetHTForwardState", body, getHtforwardState)...)
}

// SetButtonLockState Set the button lock state
func (d DevicePropertiesService) SetButtonLockState(ctx context.Context, desiredButtonLockState ButtonLockState) error {
	type SetButtonLockStateReq struct {
		DesiredButtonLockState ButtonLockState
	}

	type SetButtonLockStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetButtonLockStateReq
	}

	type SetButtonLockStateBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		SetButtonLockStateNode SetButtonLockStateNode `xml:"u:SetButtonLockState,omitempty"`
	}

	body := SetButtonLockStateBody{SetButtonLockStateNode: SetButtonLockStateNode{
		SetButtonLockStateReq: SetButtonLockStateReq{DesiredButtonLockState: desiredButtonLockState},
		Xmlns:                 namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "SetButtonLockState", body, nil)...)
}

type DevicePropertiesGetButtonLockStateRes struct {
	CurrentButtonLockState ButtonLockState
}

// GetButtonLockState Get the current button lock state
func (d DevicePropertiesService) GetButtonLockState(ctx context.Context, getButtonLockState *DevicePropertiesGetButtonLockStateRes) error {
	type GetButtonLockStateReq struct{}

	type GetButtonLockStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetButtonLockStateReq
	}

	type GetButtonLockStateBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		GetButtonLockStateNode GetButtonLockStateNode `xml:"u:GetButtonLockState,omitempty"`
	}

	body := GetButtonLockStateBody{GetButtonLockStateNode: GetButtonLockStateNode{
		GetButtonLockStateReq: GetButtonLockStateReq{},
		Xmlns:                 namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "GetButtonLockState", body, getButtonLockState)...)
}

type DevicePropertiesRoomDetectionStartChirpingRes struct {
	PlayId RoomDetectionPlayId
}

func (d DevicePropertiesService) RoomDetectionStartChirping(ctx context.Context, channel RoomDetectionChirpChannel, durationMilliseconds RoomDetectionDurationMilliseconds, chirpIfPlayingSwappableAudio RoomDetectionChirpIfPlayingSwappableAudio, roomDetectionStartChirping *DevicePropertiesRoomDetectionStartChirpingRes) error {
	type RoomDetectionStartChirpingReq struct {
		Channel                      RoomDetectionChirpChannel
		DurationMilliseconds         RoomDetectionDurationMilliseconds
		ChirpIfPlayingSwappableAudio RoomDetectionChirpIfPlayingSwappableAudio
	}

	type RoomDetectionStartChirpingNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RoomDetectionStartChirpingReq
	}

	type RoomDetectionStartChirpingBody struct {
		XMLName                        xml.Name                       `xml:"s:Body"`
		RoomDetectionStartChirpingNode RoomDetectionStartChirpingNode `xml:"u:RoomDetectionStartChirping,omitempty"`
	}

	body := RoomDetectionStartChirpingBody{RoomDetectionStartChirpingNode: RoomDetectionStartChirpingNode{
		RoomDetectionStartChirpingReq: RoomDetectionStartChirpingReq{
			Channel:                      channel,
			ChirpIfPlayingSwappableAudio: chirpIfPlayingSwappableAudio,
			DurationMilliseconds:         durationMilliseconds,
		},
		Xmlns: namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "RoomDetectionStartChirping", body, roomDetectionStartChirping)...)
}

func (d DevicePropertiesService) RoomDetectionStopChirping(ctx context.Context, playId RoomDetectionPlayId) error {
	type RoomDetectionStopChirpingReq struct {
		PlayId RoomDetectionPlayId
	}

	type RoomDetectionStopChirpingNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RoomDetectionStopChirpingReq
	}

	type RoomDetectionStopChirpingBody struct {
		XMLName                       xml.Name                      `xml:"s:Body"`
		RoomDetectionStopChirpingNode RoomDetectionStopChirpingNode `xml:"u:RoomDetectionStopChirping,omitempty"`
	}

	body := RoomDetectionStopChirpingBody{RoomDetectionStopChirpingNode: RoomDetectionStopChirpingNode{
		RoomDetectionStopChirpingReq: RoomDetectionStopChirpingReq{PlayId: playId},
		Xmlns:                        namespaceDevicePropertiesService,
	}}

	return d.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlDevicePropertiesService, namespaceDevicePropertiesService, "RoomDetectionStopChirping", body, nil)...)
}
