// CODE GENERATED, DO NOT EDIT

package S21

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceZoneGroupTopologyService is the UPnP namespace for the Sonos service.
	namespaceZoneGroupTopologyService = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"

	// eventZoneGroupTopologyService is the event path for the Sonos event service.
	eventZoneGroupTopologyService = "/ZoneGroupTopology/Event"

	// controlZoneGroupTopologyService is the action path for the Sonos service.
	controlZoneGroupTopologyService = "/ZoneGroupTopology/Control"
)

// ZoneGroupTopologyService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type ZoneGroupTopologyService struct {
	doer do.Doer
}

// NewZoneGroupTopologyService initializes a new ZoneGroupTopologyService with the provided Doer.
func NewZoneGroupTopologyService(doer do.Doer) *ZoneGroupTopologyService {
	return &ZoneGroupTopologyService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (z ZoneGroupTopologyService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return z.doer.Do(ctx, do.WithPath(eventZoneGroupTopologyService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type ZoneGroupTopologyCheckForUpdateRes struct {
	UpdateItem UpdateItem
}

func (z ZoneGroupTopologyService) CheckForUpdate(ctx context.Context, updateType UpdateType, cachedOnly CachedOnly, version Version, checkForUpdate *ZoneGroupTopologyCheckForUpdateRes) error {
	type CheckForUpdateReq struct {
		UpdateType UpdateType
		CachedOnly CachedOnly
		Version    Version
	}

	type CheckForUpdateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CheckForUpdateReq
	}

	type CheckForUpdateBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		CheckForUpdateNode CheckForUpdateNode `xml:"u:CheckForUpdate,omitempty"`
	}

	body := CheckForUpdateBody{CheckForUpdateNode: CheckForUpdateNode{
		CheckForUpdateReq: CheckForUpdateReq{
			CachedOnly: cachedOnly,
			UpdateType: updateType,
			Version:    version,
		},
		Xmlns: namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "CheckForUpdate", body, checkForUpdate)...)
}

func (z ZoneGroupTopologyService) BeginSoftwareUpdate(ctx context.Context, updateUrl UpdateURL, flags UpdateFlags, extraOptions UpdateExtraOptions) error {
	type BeginSoftwareUpdateReq struct {
		UpdateURL    UpdateURL
		Flags        UpdateFlags
		ExtraOptions UpdateExtraOptions
	}

	type BeginSoftwareUpdateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BeginSoftwareUpdateReq
	}

	type BeginSoftwareUpdateBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		BeginSoftwareUpdateNode BeginSoftwareUpdateNode `xml:"u:BeginSoftwareUpdate,omitempty"`
	}

	body := BeginSoftwareUpdateBody{BeginSoftwareUpdateNode: BeginSoftwareUpdateNode{
		BeginSoftwareUpdateReq: BeginSoftwareUpdateReq{
			ExtraOptions: extraOptions,
			Flags:        flags,
			UpdateURL:    updateUrl,
		},
		Xmlns: namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "BeginSoftwareUpdate", body, nil)...)
}

func (z ZoneGroupTopologyService) ReportUnresponsiveDevice(ctx context.Context, deviceUuid MemberID, desiredAction UnresponsiveDeviceActionType) error {
	type ReportUnresponsiveDeviceReq struct {
		DeviceUUID    MemberID
		DesiredAction UnresponsiveDeviceActionType
	}

	type ReportUnresponsiveDeviceNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReportUnresponsiveDeviceReq
	}

	type ReportUnresponsiveDeviceBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		ReportUnresponsiveDeviceNode ReportUnresponsiveDeviceNode `xml:"u:ReportUnresponsiveDevice,omitempty"`
	}

	body := ReportUnresponsiveDeviceBody{ReportUnresponsiveDeviceNode: ReportUnresponsiveDeviceNode{
		ReportUnresponsiveDeviceReq: ReportUnresponsiveDeviceReq{
			DesiredAction: desiredAction,
			DeviceUUID:    deviceUuid,
		},
		Xmlns: namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "ReportUnresponsiveDevice", body, nil)...)
}

func (z ZoneGroupTopologyService) ReportAlarmStartedRunning(ctx context.Context) error {
	type ReportAlarmStartedRunningReq struct{}

	type ReportAlarmStartedRunningNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReportAlarmStartedRunningReq
	}

	type ReportAlarmStartedRunningBody struct {
		XMLName                       xml.Name                      `xml:"s:Body"`
		ReportAlarmStartedRunningNode ReportAlarmStartedRunningNode `xml:"u:ReportAlarmStartedRunning,omitempty"`
	}

	body := ReportAlarmStartedRunningBody{ReportAlarmStartedRunningNode: ReportAlarmStartedRunningNode{
		ReportAlarmStartedRunningReq: ReportAlarmStartedRunningReq{},
		Xmlns:                        namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "ReportAlarmStartedRunning", body, nil)...)
}

type ZoneGroupTopologySubmitDiagnosticsRes struct {
	DiagnosticID DiagnosticID
}

func (z ZoneGroupTopologyService) SubmitDiagnostics(ctx context.Context, includeControllers IncludeControllers, t Origin, submitDiagnostics *ZoneGroupTopologySubmitDiagnosticsRes) error {
	type SubmitDiagnosticsReq struct {
		IncludeControllers IncludeControllers
		Type               Origin
	}

	type SubmitDiagnosticsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SubmitDiagnosticsReq
	}

	type SubmitDiagnosticsBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		SubmitDiagnosticsNode SubmitDiagnosticsNode `xml:"u:SubmitDiagnostics,omitempty"`
	}

	body := SubmitDiagnosticsBody{SubmitDiagnosticsNode: SubmitDiagnosticsNode{
		SubmitDiagnosticsReq: SubmitDiagnosticsReq{
			IncludeControllers: includeControllers,
			Type:               t,
		},
		Xmlns: namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "SubmitDiagnostics", body, submitDiagnostics)...)
}

func (z ZoneGroupTopologyService) RegisterMobileDevice(ctx context.Context, mobileDeviceName MobileDeviceName, mobileDeviceUdn MobileDeviceUDN, mobileIpandPort MobileIPAndPort) error {
	type RegisterMobileDeviceReq struct {
		MobileDeviceName MobileDeviceName
		MobileDeviceUDN  MobileDeviceUDN
		MobileIPAndPort  MobileIPAndPort
	}

	type RegisterMobileDeviceNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RegisterMobileDeviceReq
	}

	type RegisterMobileDeviceBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		RegisterMobileDeviceNode RegisterMobileDeviceNode `xml:"u:RegisterMobileDevice,omitempty"`
	}

	body := RegisterMobileDeviceBody{RegisterMobileDeviceNode: RegisterMobileDeviceNode{
		RegisterMobileDeviceReq: RegisterMobileDeviceReq{
			MobileDeviceName: mobileDeviceName,
			MobileDeviceUDN:  mobileDeviceUdn,
			MobileIPAndPort:  mobileIpandPort,
		},
		Xmlns: namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "RegisterMobileDevice", body, nil)...)
}

type ZoneGroupTopologyGetZoneGroupAttributesRes struct {
	CurrentZoneGroupName          ZoneGroupName
	CurrentZoneGroupID            ZoneGroupID
	CurrentZonePlayerUUIDsInGroup ZonePlayerUUIDsInGroup
	CurrentMuseHouseholdId        MuseHouseholdId
}

// GetZoneGroupAttributes Get information about the current Zone
func (z ZoneGroupTopologyService) GetZoneGroupAttributes(ctx context.Context, getZoneGroupAttributes *ZoneGroupTopologyGetZoneGroupAttributesRes) error {
	type GetZoneGroupAttributesReq struct{}

	type GetZoneGroupAttributesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetZoneGroupAttributesReq
	}

	type GetZoneGroupAttributesBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		GetZoneGroupAttributesNode GetZoneGroupAttributesNode `xml:"u:GetZoneGroupAttributes,omitempty"`
	}

	body := GetZoneGroupAttributesBody{GetZoneGroupAttributesNode: GetZoneGroupAttributesNode{
		GetZoneGroupAttributesReq: GetZoneGroupAttributesReq{},
		Xmlns:                     namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "GetZoneGroupAttributes", body, getZoneGroupAttributes)...)
}

type ZoneGroupTopologyGetZoneGroupStateRes struct {
	ZoneGroupState ZoneGroupState
}

// GetZoneGroupState Get all the Sonos groups, (as XML)
func (z ZoneGroupTopologyService) GetZoneGroupState(ctx context.Context, getZoneGroupState *ZoneGroupTopologyGetZoneGroupStateRes) error {
	type GetZoneGroupStateReq struct{}

	type GetZoneGroupStateNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetZoneGroupStateReq
	}

	type GetZoneGroupStateBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		GetZoneGroupStateNode GetZoneGroupStateNode `xml:"u:GetZoneGroupState,omitempty"`
	}

	body := GetZoneGroupStateBody{GetZoneGroupStateNode: GetZoneGroupStateNode{
		GetZoneGroupStateReq: GetZoneGroupStateReq{},
		Xmlns:                namespaceZoneGroupTopologyService,
	}}

	return z.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlZoneGroupTopologyService, namespaceZoneGroupTopologyService, "GetZoneGroupState", body, getZoneGroupState)...)
}
