// CODE GENERATED, DO NOT EDIT

package S3

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceGroupManagementService is the UPnP namespace for the Sonos service.
	namespaceGroupManagementService = "urn:schemas-upnp-org:service:GroupManagement:1"

	// eventGroupManagementService is the event path for the Sonos event service.
	eventGroupManagementService = "/GroupManagement/Event"

	// controlGroupManagementService is the action path for the Sonos service.
	controlGroupManagementService = "/GroupManagement/Control"
)

// GroupManagementService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type GroupManagementService struct {
	doer do.Doer
}

// NewGroupManagementService initializes a new GroupManagementService with the provided Doer.
func NewGroupManagementService(doer do.Doer) *GroupManagementService {
	return &GroupManagementService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (g GroupManagementService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return g.doer.Do(ctx, do.WithPath(eventGroupManagementService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type GroupManagementAddMemberRes struct {
	CurrentTransportSettings TransportSettings
	CurrentURI               AVTransportURI
	GroupUUIDJoined          LocalGroupUUID
	ResetVolumeAfter         ResetVolumeAfter
	VolumeAVTransportURI     VolumeAVTransportURI
}

func (g GroupManagementService) AddMember(ctx context.Context, memberId MemberID, bootSeq BootSeq, addMember *GroupManagementAddMemberRes) error {
	type AddMemberReq struct {
		MemberID MemberID
		BootSeq  BootSeq
	}

	type AddMemberNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddMemberReq
	}

	type AddMemberBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		AddMemberNode AddMemberNode `xml:"u:AddMember,omitempty"`
	}

	body := AddMemberBody{AddMemberNode: AddMemberNode{
		AddMemberReq: AddMemberReq{
			BootSeq:  bootSeq,
			MemberID: memberId,
		},
		Xmlns: namespaceGroupManagementService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupManagementService, namespaceGroupManagementService, "AddMember", body, addMember)...)
}

func (g GroupManagementService) RemoveMember(ctx context.Context, memberId MemberID) error {
	type RemoveMemberReq struct {
		MemberID MemberID
	}

	type RemoveMemberNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveMemberReq
	}

	type RemoveMemberBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		RemoveMemberNode RemoveMemberNode `xml:"u:RemoveMember,omitempty"`
	}

	body := RemoveMemberBody{RemoveMemberNode: RemoveMemberNode{
		RemoveMemberReq: RemoveMemberReq{MemberID: memberId},
		Xmlns:           namespaceGroupManagementService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupManagementService, namespaceGroupManagementService, "RemoveMember", body, nil)...)
}

func (g GroupManagementService) ReportTrackBufferingResult(ctx context.Context, memberId MemberID, resultCode BufferingResultCode) error {
	type ReportTrackBufferingResultReq struct {
		MemberID   MemberID
		ResultCode BufferingResultCode
	}

	type ReportTrackBufferingResultNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReportTrackBufferingResultReq
	}

	type ReportTrackBufferingResultBody struct {
		XMLName                        xml.Name                       `xml:"s:Body"`
		ReportTrackBufferingResultNode ReportTrackBufferingResultNode `xml:"u:ReportTrackBufferingResult,omitempty"`
	}

	body := ReportTrackBufferingResultBody{ReportTrackBufferingResultNode: ReportTrackBufferingResultNode{
		ReportTrackBufferingResultReq: ReportTrackBufferingResultReq{
			MemberID:   memberId,
			ResultCode: resultCode,
		},
		Xmlns: namespaceGroupManagementService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupManagementService, namespaceGroupManagementService, "ReportTrackBufferingResult", body, nil)...)
}

func (g GroupManagementService) SetSourceAreaIds(ctx context.Context, desiredSourceAreaIds SourceAreaIds) error {
	type SetSourceAreaIdsReq struct {
		DesiredSourceAreaIds SourceAreaIds
	}

	type SetSourceAreaIdsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetSourceAreaIdsReq
	}

	type SetSourceAreaIdsBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		SetSourceAreaIdsNode SetSourceAreaIdsNode `xml:"u:SetSourceAreaIds,omitempty"`
	}

	body := SetSourceAreaIdsBody{SetSourceAreaIdsNode: SetSourceAreaIdsNode{
		SetSourceAreaIdsReq: SetSourceAreaIdsReq{DesiredSourceAreaIds: desiredSourceAreaIds},
		Xmlns:               namespaceGroupManagementService,
	}}

	return g.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlGroupManagementService, namespaceGroupManagementService, "SetSourceAreaIds", body, nil)...)
}
