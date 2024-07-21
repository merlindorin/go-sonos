// CODE GENERATED, DO NOT EDIT

package S5

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceGroupRenderingControlService is the UPnP namespace for the Sonos service.
	namespaceGroupRenderingControlService = "urn:schemas-upnp-org:service:GroupRenderingControl:1"

	// eventGroupRenderingControlService is the event path for the Sonos event service.
	eventGroupRenderingControlService = "/MediaRenderer/GroupRenderingControl/Event"

	// controlGroupRenderingControlService is the action path for the Sonos service.
	controlGroupRenderingControlService = "/MediaRenderer/GroupRenderingControl/Control"
)

// GroupRenderingControlService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type GroupRenderingControlService struct {
	doer do.Doer
}

// NewGroupRenderingControlService initializes a new GroupRenderingControlService with the provided Doer.
func NewGroupRenderingControlService(doer do.Doer) *GroupRenderingControlService {
	return &GroupRenderingControlService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (g GroupRenderingControlService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return g.doer.Do(ctx, do.WithPath(eventGroupRenderingControlService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type GroupRenderingControlGetGroupMuteRes struct {
	CurrentMute GroupMute
}

// GetGroupMute Get the group mute state.
func (g GroupRenderingControlService) GetGroupMute(ctx context.Context, getGroupMute *GroupRenderingControlGetGroupMuteRes) error {
	type GetGroupMuteReq struct {
		InstanceID InstanceID
	}

	type GetGroupMuteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetGroupMuteReq
	}

	type GetGroupMuteBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		GetGroupMuteNode GetGroupMuteNode `xml:"u:GetGroupMute,omitempty"`
	}

	body := GetGroupMuteBody{GetGroupMuteNode: GetGroupMuteNode{
		GetGroupMuteReq: GetGroupMuteReq{InstanceID: 0},
		Xmlns:           namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "GetGroupMute", body, getGroupMute)...)
}

// SetGroupMute (Un-/)Mute the entire group
func (g GroupRenderingControlService) SetGroupMute(ctx context.Context, desiredMute GroupMute) error {
	type SetGroupMuteReq struct {
		InstanceID  InstanceID
		DesiredMute GroupMute
	}

	type SetGroupMuteNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetGroupMuteReq
	}

	type SetGroupMuteBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		SetGroupMuteNode SetGroupMuteNode `xml:"u:SetGroupMute,omitempty"`
	}

	body := SetGroupMuteBody{SetGroupMuteNode: SetGroupMuteNode{
		SetGroupMuteReq: SetGroupMuteReq{
			DesiredMute: desiredMute,
			InstanceID:  0,
		},
		Xmlns: namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "SetGroupMute", body, nil)...)
}

type GroupRenderingControlGetGroupVolumeRes struct {
	CurrentVolume GroupVolume
}

// GetGroupVolume Get the group volume.
func (g GroupRenderingControlService) GetGroupVolume(ctx context.Context, getGroupVolume *GroupRenderingControlGetGroupVolumeRes) error {
	type GetGroupVolumeReq struct {
		InstanceID InstanceID
	}

	type GetGroupVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetGroupVolumeReq
	}

	type GetGroupVolumeBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		GetGroupVolumeNode GetGroupVolumeNode `xml:"u:GetGroupVolume,omitempty"`
	}

	body := GetGroupVolumeBody{GetGroupVolumeNode: GetGroupVolumeNode{
		GetGroupVolumeReq: GetGroupVolumeReq{InstanceID: 0},
		Xmlns:             namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "GetGroupVolume", body, getGroupVolume)...)
}

// SetGroupVolume Change group volume. Players volume will be changed proportionally based on last snapshot
func (g GroupRenderingControlService) SetGroupVolume(ctx context.Context, desiredVolume GroupVolume) error {
	type SetGroupVolumeReq struct {
		InstanceID    InstanceID
		DesiredVolume GroupVolume
	}

	type SetGroupVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetGroupVolumeReq
	}

	type SetGroupVolumeBody struct {
		XMLName            xml.Name           `xml:"s:Body"`
		SetGroupVolumeNode SetGroupVolumeNode `xml:"u:SetGroupVolume,omitempty"`
	}

	body := SetGroupVolumeBody{SetGroupVolumeNode: SetGroupVolumeNode{
		SetGroupVolumeReq: SetGroupVolumeReq{
			DesiredVolume: desiredVolume,
			InstanceID:    0,
		},
		Xmlns: namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "SetGroupVolume", body, nil)...)
}

type GroupRenderingControlSetRelativeGroupVolumeRes struct {
	NewVolume GroupVolume
}

// SetRelativeGroupVolume Relatively change group volume - returns final group volume. Players volume will be changed proportionally based on last snapshot
func (g GroupRenderingControlService) SetRelativeGroupVolume(ctx context.Context, adjustment VolumeAdjustment, setRelativeGroupVolume *GroupRenderingControlSetRelativeGroupVolumeRes) error {
	type SetRelativeGroupVolumeReq struct {
		InstanceID InstanceID
		Adjustment VolumeAdjustment
	}

	type SetRelativeGroupVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetRelativeGroupVolumeReq
	}

	type SetRelativeGroupVolumeBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		SetRelativeGroupVolumeNode SetRelativeGroupVolumeNode `xml:"u:SetRelativeGroupVolume,omitempty"`
	}

	body := SetRelativeGroupVolumeBody{SetRelativeGroupVolumeNode: SetRelativeGroupVolumeNode{
		SetRelativeGroupVolumeReq: SetRelativeGroupVolumeReq{
			Adjustment: adjustment,
			InstanceID: 0,
		},
		Xmlns: namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "SetRelativeGroupVolume", body, setRelativeGroupVolume)...)
}

// SnapshotGroupVolume Creates a new group volume snapshot,  the volume ratio between all players. It is used by SetGroupVolume and SetRelativeGroupVolume
func (g GroupRenderingControlService) SnapshotGroupVolume(ctx context.Context) error {
	type SnapshotGroupVolumeReq struct {
		InstanceID InstanceID
	}

	type SnapshotGroupVolumeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SnapshotGroupVolumeReq
	}

	type SnapshotGroupVolumeBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		SnapshotGroupVolumeNode SnapshotGroupVolumeNode `xml:"u:SnapshotGroupVolume,omitempty"`
	}

	body := SnapshotGroupVolumeBody{SnapshotGroupVolumeNode: SnapshotGroupVolumeNode{
		SnapshotGroupVolumeReq: SnapshotGroupVolumeReq{InstanceID: 0},
		Xmlns:                  namespaceGroupRenderingControlService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupRenderingControlService, namespaceGroupRenderingControlService, "SnapshotGroupVolume", body, nil)...)
}
