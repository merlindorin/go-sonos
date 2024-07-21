// CODE GENERATED, DO NOT EDIT

package S13

import (
	"context"
	"encoding/xml"
	"fmt"
	opts "github.com/hoomy-official/go-sonos/opts"
	do "github.com/vanyda-official/go-shared/pkg/net/do"
	"net/url"
)

const (
	// namespaceContentDirectoryService is the UPnP namespace for the Sonos service.
	namespaceContentDirectoryService = "urn:schemas-upnp-org:service:ContentDirectory:1"

	// eventContentDirectoryService is the event path for the Sonos event service.
	eventContentDirectoryService = "/MediaServer/ContentDirectory/Event"

	// controlContentDirectoryService is the action path for the Sonos service.
	controlContentDirectoryService = "/MediaServer/ContentDirectory/Control"
)

// ContentDirectoryService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type ContentDirectoryService struct {
	doer do.Doer
}

// NewContentDirectoryService initializes a new ContentDirectoryService with the provided Doer.
func NewContentDirectoryService(doer do.Doer) *ContentDirectoryService {
	return &ContentDirectoryService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (c ContentDirectoryService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return c.doer.Do(ctx, do.WithPath(eventContentDirectoryService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type ContentDirectoryGetSearchCapabilitiesRes struct {
	SearchCaps SearchCapabilities
}

func (c ContentDirectoryService) GetSearchCapabilities(ctx context.Context, getSearchCapabilities *ContentDirectoryGetSearchCapabilitiesRes) error {
	type GetSearchCapabilitiesReq struct{}

	type GetSearchCapabilitiesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetSearchCapabilitiesReq
	}

	type GetSearchCapabilitiesBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		GetSearchCapabilitiesNode GetSearchCapabilitiesNode `xml:"u:GetSearchCapabilities,omitempty"`
	}

	body := GetSearchCapabilitiesBody{GetSearchCapabilitiesNode: GetSearchCapabilitiesNode{
		GetSearchCapabilitiesReq: GetSearchCapabilitiesReq{},
		Xmlns:                    namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetSearchCapabilities", body, getSearchCapabilities)...)
}

type ContentDirectoryGetSortCapabilitiesRes struct {
	SortCaps SortCapabilities
}

func (c ContentDirectoryService) GetSortCapabilities(ctx context.Context, getSortCapabilities *ContentDirectoryGetSortCapabilitiesRes) error {
	type GetSortCapabilitiesReq struct{}

	type GetSortCapabilitiesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetSortCapabilitiesReq
	}

	type GetSortCapabilitiesBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		GetSortCapabilitiesNode GetSortCapabilitiesNode `xml:"u:GetSortCapabilities,omitempty"`
	}

	body := GetSortCapabilitiesBody{GetSortCapabilitiesNode: GetSortCapabilitiesNode{
		GetSortCapabilitiesReq: GetSortCapabilitiesReq{},
		Xmlns:                  namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetSortCapabilities", body, getSortCapabilities)...)
}

type ContentDirectoryGetSystemUpdateIDRes struct {
	Id SystemUpdateID
}

func (c ContentDirectoryService) GetSystemUpdateID(ctx context.Context, getSystemUpdateId *ContentDirectoryGetSystemUpdateIDRes) error {
	type GetSystemUpdateIDReq struct{}

	type GetSystemUpdateIDNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetSystemUpdateIDReq
	}

	type GetSystemUpdateIDBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		GetSystemUpdateIDNode GetSystemUpdateIDNode `xml:"u:GetSystemUpdateID,omitempty"`
	}

	body := GetSystemUpdateIDBody{GetSystemUpdateIDNode: GetSystemUpdateIDNode{
		GetSystemUpdateIDReq: GetSystemUpdateIDReq{},
		Xmlns:                namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetSystemUpdateID", body, getSystemUpdateId)...)
}

type ContentDirectoryGetAlbumArtistDisplayOptionRes struct {
	AlbumArtistDisplayOption AlbumArtistDisplayOption
}

// GetAlbumArtistDisplayOption Get the current album art display option such as `WMP`, `ITUNES` or `NONE`
func (c ContentDirectoryService) GetAlbumArtistDisplayOption(ctx context.Context, getAlbumArtistDisplayOption *ContentDirectoryGetAlbumArtistDisplayOptionRes) error {
	type GetAlbumArtistDisplayOptionReq struct{}

	type GetAlbumArtistDisplayOptionNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAlbumArtistDisplayOptionReq
	}

	type GetAlbumArtistDisplayOptionBody struct {
		XMLName                         xml.Name                        `xml:"s:Body"`
		GetAlbumArtistDisplayOptionNode GetAlbumArtistDisplayOptionNode `xml:"u:GetAlbumArtistDisplayOption,omitempty"`
	}

	body := GetAlbumArtistDisplayOptionBody{GetAlbumArtistDisplayOptionNode: GetAlbumArtistDisplayOptionNode{
		GetAlbumArtistDisplayOptionReq: GetAlbumArtistDisplayOptionReq{},
		Xmlns:                          namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetAlbumArtistDisplayOption", body, getAlbumArtistDisplayOption)...)
}

type ContentDirectoryGetLastIndexChangeRes struct {
	LastIndexChange LastIndexChange
}

func (c ContentDirectoryService) GetLastIndexChange(ctx context.Context, getLastIndexChange *ContentDirectoryGetLastIndexChangeRes) error {
	type GetLastIndexChangeReq struct{}

	type GetLastIndexChangeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetLastIndexChangeReq
	}

	type GetLastIndexChangeBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		GetLastIndexChangeNode GetLastIndexChangeNode `xml:"u:GetLastIndexChange,omitempty"`
	}

	body := GetLastIndexChangeBody{GetLastIndexChangeNode: GetLastIndexChangeNode{
		GetLastIndexChangeReq: GetLastIndexChangeReq{},
		Xmlns:                 namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetLastIndexChange", body, getLastIndexChange)...)
}

type ContentDirectoryBrowseRes struct {
	Result         Result
	NumberReturned Count
	TotalMatches   Count
	UpdateID       UpdateID
}

// Browse Browse for content: Music library (A), share(S:), Sonos playlists(SQ:), Sonos favorites(FV:2), radio stations(R:0/0), radio shows(R:0/1), queue(Q:)). Recommendation: Send one request, check the `TotalMatches` and - if necessary - do additional requests with higher `StartingIndex`. In case of duplicates only the first is returned! Example: albums with same title, even if artists are different
func (c ContentDirectoryService) Browse(ctx context.Context, objectId ObjectID, browseFlag BrowseFlag, filter Filter, startingIndex Index, requestedCount Count, sortCriteria SortCriteria, browse *ContentDirectoryBrowseRes) error {
	type BrowseReq struct {
		ObjectID       ObjectID
		BrowseFlag     BrowseFlag
		Filter         Filter
		StartingIndex  Index
		RequestedCount Count
		SortCriteria   SortCriteria
	}

	type BrowseNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BrowseReq
	}

	type BrowseBody struct {
		XMLName    xml.Name   `xml:"s:Body"`
		BrowseNode BrowseNode `xml:"u:Browse,omitempty"`
	}

	body := BrowseBody{BrowseNode: BrowseNode{
		BrowseReq: BrowseReq{
			BrowseFlag:     browseFlag,
			Filter:         filter,
			ObjectID:       objectId,
			RequestedCount: requestedCount,
			SortCriteria:   sortCriteria,
			StartingIndex:  startingIndex,
		},
		Xmlns: namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "Browse", body, browse)...)
}

type ContentDirectoryFindPrefixRes struct {
	StartingIndex Index
	UpdateID      UpdateID
}

func (c ContentDirectoryService) FindPrefix(ctx context.Context, objectId ObjectID, prefix Prefix, findPrefix *ContentDirectoryFindPrefixRes) error {
	type FindPrefixReq struct {
		ObjectID ObjectID
		Prefix   Prefix
	}

	type FindPrefixNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		FindPrefixReq
	}

	type FindPrefixBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		FindPrefixNode FindPrefixNode `xml:"u:FindPrefix,omitempty"`
	}

	body := FindPrefixBody{FindPrefixNode: FindPrefixNode{
		FindPrefixReq: FindPrefixReq{
			ObjectID: objectId,
			Prefix:   prefix,
		},
		Xmlns: namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "FindPrefix", body, findPrefix)...)
}

type ContentDirectoryGetAllPrefixLocationsRes struct {
	TotalPrefixes     Count
	PrefixAndIndexCSV Result
	UpdateID          UpdateID
}

func (c ContentDirectoryService) GetAllPrefixLocations(ctx context.Context, objectId ObjectID, getAllPrefixLocations *ContentDirectoryGetAllPrefixLocationsRes) error {
	type GetAllPrefixLocationsReq struct {
		ObjectID ObjectID
	}

	type GetAllPrefixLocationsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetAllPrefixLocationsReq
	}

	type GetAllPrefixLocationsBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		GetAllPrefixLocationsNode GetAllPrefixLocationsNode `xml:"u:GetAllPrefixLocations,omitempty"`
	}

	body := GetAllPrefixLocationsBody{GetAllPrefixLocationsNode: GetAllPrefixLocationsNode{
		GetAllPrefixLocationsReq: GetAllPrefixLocationsReq{ObjectID: objectId},
		Xmlns:                    namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetAllPrefixLocations", body, getAllPrefixLocations)...)
}

type ContentDirectoryCreateObjectRes struct {
	ObjectID ObjectID
	Result   Result
}

func (c ContentDirectoryService) CreateObject(ctx context.Context, containerId ObjectID, elements Result, createObject *ContentDirectoryCreateObjectRes) error {
	type CreateObjectReq struct {
		ContainerID ObjectID
		Elements    Result
	}

	type CreateObjectNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CreateObjectReq
	}

	type CreateObjectBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		CreateObjectNode CreateObjectNode `xml:"u:CreateObject,omitempty"`
	}

	body := CreateObjectBody{CreateObjectNode: CreateObjectNode{
		CreateObjectReq: CreateObjectReq{
			ContainerID: containerId,
			Elements:    elements,
		},
		Xmlns: namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "CreateObject", body, createObject)...)
}

func (c ContentDirectoryService) UpdateObject(ctx context.Context, objectId ObjectID, currentTagValue TagValueList, newTagValue TagValueList) error {
	type UpdateObjectReq struct {
		ObjectID        ObjectID
		CurrentTagValue TagValueList
		NewTagValue     TagValueList
	}

	type UpdateObjectNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		UpdateObjectReq
	}

	type UpdateObjectBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		UpdateObjectNode UpdateObjectNode `xml:"u:UpdateObject,omitempty"`
	}

	body := UpdateObjectBody{UpdateObjectNode: UpdateObjectNode{
		UpdateObjectReq: UpdateObjectReq{
			CurrentTagValue: currentTagValue,
			NewTagValue:     newTagValue,
			ObjectID:        objectId,
		},
		Xmlns: namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "UpdateObject", body, nil)...)
}

func (c ContentDirectoryService) DestroyObject(ctx context.Context, objectId ObjectID) error {
	type DestroyObjectReq struct {
		ObjectID ObjectID
	}

	type DestroyObjectNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		DestroyObjectReq
	}

	type DestroyObjectBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		DestroyObjectNode DestroyObjectNode `xml:"u:DestroyObject,omitempty"`
	}

	body := DestroyObjectBody{DestroyObjectNode: DestroyObjectNode{
		DestroyObjectReq: DestroyObjectReq{ObjectID: objectId},
		Xmlns:            namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "DestroyObject", body, nil)...)
}

// RefreshShareIndex Updates the music library (share) index
func (c ContentDirectoryService) RefreshShareIndex(ctx context.Context, albumArtistDisplayOption AlbumArtistDisplayOption) error {
	type RefreshShareIndexReq struct {
		AlbumArtistDisplayOption AlbumArtistDisplayOption
	}

	type RefreshShareIndexNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RefreshShareIndexReq
	}

	type RefreshShareIndexBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		RefreshShareIndexNode RefreshShareIndexNode `xml:"u:RefreshShareIndex,omitempty"`
	}

	body := RefreshShareIndexBody{RefreshShareIndexNode: RefreshShareIndexNode{
		RefreshShareIndexReq: RefreshShareIndexReq{AlbumArtistDisplayOption: albumArtistDisplayOption},
		Xmlns:                namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "RefreshShareIndex", body, nil)...)
}

func (c ContentDirectoryService) RequestResort(ctx context.Context, sortOrder SortOrder) error {
	type RequestResortReq struct {
		SortOrder SortOrder
	}

	type RequestResortNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RequestResortReq
	}

	type RequestResortBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		RequestResortNode RequestResortNode `xml:"u:RequestResort,omitempty"`
	}

	body := RequestResortBody{RequestResortNode: RequestResortNode{
		RequestResortReq: RequestResortReq{SortOrder: sortOrder},
		Xmlns:            namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "RequestResort", body, nil)...)
}

type ContentDirectoryGetShareIndexInProgressRes struct {
	IsIndexing ShareIndexInProgress
}

func (c ContentDirectoryService) GetShareIndexInProgress(ctx context.Context, getShareIndexInProgress *ContentDirectoryGetShareIndexInProgressRes) error {
	type GetShareIndexInProgressReq struct{}

	type GetShareIndexInProgressNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetShareIndexInProgressReq
	}

	type GetShareIndexInProgressBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		GetShareIndexInProgressNode GetShareIndexInProgressNode `xml:"u:GetShareIndexInProgress,omitempty"`
	}

	body := GetShareIndexInProgressBody{GetShareIndexInProgressNode: GetShareIndexInProgressNode{
		GetShareIndexInProgressReq: GetShareIndexInProgressReq{},
		Xmlns:                      namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetShareIndexInProgress", body, getShareIndexInProgress)...)
}

type ContentDirectoryGetBrowseableRes struct {
	IsBrowseable Browseable
}

func (c ContentDirectoryService) GetBrowseable(ctx context.Context, getBrowseable *ContentDirectoryGetBrowseableRes) error {
	type GetBrowseableReq struct{}

	type GetBrowseableNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetBrowseableReq
	}

	type GetBrowseableBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		GetBrowseableNode GetBrowseableNode `xml:"u:GetBrowseable,omitempty"`
	}

	body := GetBrowseableBody{GetBrowseableNode: GetBrowseableNode{
		GetBrowseableReq: GetBrowseableReq{},
		Xmlns:            namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "GetBrowseable", body, getBrowseable)...)
}

func (c ContentDirectoryService) SetBrowseable(ctx context.Context, browseable Browseable) error {
	type SetBrowseableReq struct {
		Browseable Browseable
	}

	type SetBrowseableNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetBrowseableReq
	}

	type SetBrowseableBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		SetBrowseableNode SetBrowseableNode `xml:"u:SetBrowseable,omitempty"`
	}

	body := SetBrowseableBody{SetBrowseableNode: SetBrowseableNode{
		SetBrowseableReq: SetBrowseableReq{Browseable: browseable},
		Xmlns:            namespaceContentDirectoryService,
	}}

	return c.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlContentDirectoryService, namespaceContentDirectoryService, "SetBrowseable", body, nil)...)
}
