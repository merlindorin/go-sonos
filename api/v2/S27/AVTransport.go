// CODE GENERATED, DO NOT EDIT

package S27

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceAVTransportService is the UPnP namespace for the Sonos service.
	namespaceAVTransportService = "urn:schemas-upnp-org:service:AVTransport:1"

	// eventAVTransportService is the event path for the Sonos event service.
	eventAVTransportService = "/MediaRenderer/AVTransport/Event"

	// controlAVTransportService is the action path for the Sonos service.
	controlAVTransportService = "/MediaRenderer/AVTransport/Control"
)

// AVTransportService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type AVTransportService struct {
	doer do.Doer
}

// NewAVTransportService initializes a new AVTransportService with the provided Doer.
func NewAVTransportService(doer do.Doer) *AVTransportService {
	return &AVTransportService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (a AVTransportService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return a.doer.Do(ctx, do.WithPath(eventAVTransportService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

// SetAVTransportURI Set the transport URI to a song, a stream, the queue, another player-rincon and a lot more
func (a AVTransportService) SetAVTransportURI(ctx context.Context, currentUri AVTransportURI, currentUrimetaData AVTransportURIMetaData) error {
	type SetAVTransportURIReq struct {
		InstanceID         InstanceID
		CurrentURI         AVTransportURI
		CurrentURIMetaData AVTransportURIMetaData
	}

	type SetAVTransportURINode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAVTransportURIReq
	}

	type SetAVTransportURIBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		SetAVTransportURINode SetAVTransportURINode `xml:"u:SetAVTransportURI,omitempty"`
	}

	body := SetAVTransportURIBody{SetAVTransportURINode: SetAVTransportURINode{
		SetAVTransportURIReq: SetAVTransportURIReq{
			CurrentURI:         currentUri,
			CurrentURIMetaData: currentUrimetaData,
			InstanceID:         0,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SetAVTransportURI", body, nil)...)
}

func (a AVTransportService) SetNextAVTransportURI(ctx context.Context, nextUri AVTransportURI, nextUrimetaData AVTransportURIMetaData) error {
	type SetNextAVTransportURIReq struct {
		InstanceID      InstanceID
		NextURI         AVTransportURI
		NextURIMetaData AVTransportURIMetaData
	}

	type SetNextAVTransportURINode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetNextAVTransportURIReq
	}

	type SetNextAVTransportURIBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		SetNextAVTransportURINode SetNextAVTransportURINode `xml:"u:SetNextAVTransportURI,omitempty"`
	}

	body := SetNextAVTransportURIBody{SetNextAVTransportURINode: SetNextAVTransportURINode{
		SetNextAVTransportURIReq: SetNextAVTransportURIReq{
			InstanceID:      0,
			NextURI:         nextUri,
			NextURIMetaData: nextUrimetaData,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SetNextAVTransportURI", body, nil)...)
}

type AVTransportAddURIToQueueRes struct {
	FirstTrackNumberEnqueued TrackNumber
	NumTracksAdded           NumTracks
	NewQueueLength           NumTracks
}

// AddURIToQueue Adds songs to the SONOS queue
func (a AVTransportService) AddURIToQueue(ctx context.Context, enqueuedUri URI, enqueuedUrimetaData URIMetaData, desiredFirstTrackNumberEnqueued TrackNumber, enqueueAsNext EnqueueAs, addUritoQueue *AVTransportAddURIToQueueRes) error {
	type AddURIToQueueReq struct {
		InstanceID                      InstanceID
		EnqueuedURI                     URI
		EnqueuedURIMetaData             URIMetaData
		DesiredFirstTrackNumberEnqueued TrackNumber
		EnqueueAsNext                   EnqueueAs
	}

	type AddURIToQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddURIToQueueReq
	}

	type AddURIToQueueBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		AddURIToQueueNode AddURIToQueueNode `xml:"u:AddURIToQueue,omitempty"`
	}

	body := AddURIToQueueBody{AddURIToQueueNode: AddURIToQueueNode{
		AddURIToQueueReq: AddURIToQueueReq{
			DesiredFirstTrackNumberEnqueued: desiredFirstTrackNumberEnqueued,
			EnqueueAsNext:                   enqueueAsNext,
			EnqueuedURI:                     enqueuedUri,
			EnqueuedURIMetaData:             enqueuedUrimetaData,
			InstanceID:                      0,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "AddURIToQueue", body, addUritoQueue)...)
}

type AVTransportAddMultipleURIsToQueueRes struct {
	FirstTrackNumberEnqueued TrackNumber
	NumTracksAdded           NumTracks
	NewQueueLength           NumTracks
	NewUpdateID              QueueUpdateID
}

func (a AVTransportService) AddMultipleURIsToQueue(ctx context.Context, updateId QueueUpdateID, numberOfUris NumTracks, enqueuedUris LIST_URI, enqueuedUrisMetaData LIST_URIMetaData, containerUri URI, containerMetaData URIMetaData, desiredFirstTrackNumberEnqueued TrackNumber, enqueueAsNext EnqueueAs, addMultipleUrisToQueue *AVTransportAddMultipleURIsToQueueRes) error {
	type AddMultipleURIsToQueueReq struct {
		InstanceID                      InstanceID
		UpdateID                        QueueUpdateID
		NumberOfURIs                    NumTracks
		EnqueuedURIs                    LIST_URI
		EnqueuedURIsMetaData            LIST_URIMetaData
		ContainerURI                    URI
		ContainerMetaData               URIMetaData
		DesiredFirstTrackNumberEnqueued TrackNumber
		EnqueueAsNext                   EnqueueAs
	}

	type AddMultipleURIsToQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddMultipleURIsToQueueReq
	}

	type AddMultipleURIsToQueueBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		AddMultipleURIsToQueueNode AddMultipleURIsToQueueNode `xml:"u:AddMultipleURIsToQueue,omitempty"`
	}

	body := AddMultipleURIsToQueueBody{AddMultipleURIsToQueueNode: AddMultipleURIsToQueueNode{
		AddMultipleURIsToQueueReq: AddMultipleURIsToQueueReq{
			ContainerMetaData:               containerMetaData,
			ContainerURI:                    containerUri,
			DesiredFirstTrackNumberEnqueued: desiredFirstTrackNumberEnqueued,
			EnqueueAsNext:                   enqueueAsNext,
			EnqueuedURIs:                    enqueuedUris,
			EnqueuedURIsMetaData:            enqueuedUrisMetaData,
			InstanceID:                      0,
			NumberOfURIs:                    numberOfUris,
			UpdateID:                        updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "AddMultipleURIsToQueue", body, addMultipleUrisToQueue)...)
}

func (a AVTransportService) ReorderTracksInQueue(ctx context.Context, startingIndex TrackNumber, numberOfTracks NumTracks, insertBefore TrackNumber, updateId QueueUpdateID) error {
	type ReorderTracksInQueueReq struct {
		InstanceID     InstanceID
		StartingIndex  TrackNumber
		NumberOfTracks NumTracks
		InsertBefore   TrackNumber
		UpdateID       QueueUpdateID
	}

	type ReorderTracksInQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReorderTracksInQueueReq
	}

	type ReorderTracksInQueueBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		ReorderTracksInQueueNode ReorderTracksInQueueNode `xml:"u:ReorderTracksInQueue,omitempty"`
	}

	body := ReorderTracksInQueueBody{ReorderTracksInQueueNode: ReorderTracksInQueueNode{
		ReorderTracksInQueueReq: ReorderTracksInQueueReq{
			InsertBefore:   insertBefore,
			InstanceID:     0,
			NumberOfTracks: numberOfTracks,
			StartingIndex:  startingIndex,
			UpdateID:       updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "ReorderTracksInQueue", body, nil)...)
}

func (a AVTransportService) RemoveTrackFromQueue(ctx context.Context, objectId ObjectID, updateId QueueUpdateID) error {
	type RemoveTrackFromQueueReq struct {
		InstanceID InstanceID
		ObjectID   ObjectID
		UpdateID   QueueUpdateID
	}

	type RemoveTrackFromQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveTrackFromQueueReq
	}

	type RemoveTrackFromQueueBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		RemoveTrackFromQueueNode RemoveTrackFromQueueNode `xml:"u:RemoveTrackFromQueue,omitempty"`
	}

	body := RemoveTrackFromQueueBody{RemoveTrackFromQueueNode: RemoveTrackFromQueueNode{
		RemoveTrackFromQueueReq: RemoveTrackFromQueueReq{
			InstanceID: 0,
			ObjectID:   objectId,
			UpdateID:   updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "RemoveTrackFromQueue", body, nil)...)
}

type AVTransportRemoveTrackRangeFromQueueRes struct {
	NewUpdateID QueueUpdateID
}

// RemoveTrackRangeFromQueue Removes the specified range of songs from the SONOS queue.
func (a AVTransportService) RemoveTrackRangeFromQueue(ctx context.Context, updateId QueueUpdateID, startingIndex TrackNumber, numberOfTracks NumTracks, removeTrackRangeFromQueue *AVTransportRemoveTrackRangeFromQueueRes) error {
	type RemoveTrackRangeFromQueueReq struct {
		InstanceID     InstanceID
		UpdateID       QueueUpdateID
		StartingIndex  TrackNumber
		NumberOfTracks NumTracks
	}

	type RemoveTrackRangeFromQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveTrackRangeFromQueueReq
	}

	type RemoveTrackRangeFromQueueBody struct {
		XMLName                       xml.Name                      `xml:"s:Body"`
		RemoveTrackRangeFromQueueNode RemoveTrackRangeFromQueueNode `xml:"u:RemoveTrackRangeFromQueue,omitempty"`
	}

	body := RemoveTrackRangeFromQueueBody{RemoveTrackRangeFromQueueNode: RemoveTrackRangeFromQueueNode{
		RemoveTrackRangeFromQueueReq: RemoveTrackRangeFromQueueReq{
			InstanceID:     0,
			NumberOfTracks: numberOfTracks,
			StartingIndex:  startingIndex,
			UpdateID:       updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "RemoveTrackRangeFromQueue", body, removeTrackRangeFromQueue)...)
}

// RemoveAllTracksFromQueue Flushes the SONOS queue.
func (a AVTransportService) RemoveAllTracksFromQueue(ctx context.Context) error {
	type RemoveAllTracksFromQueueReq struct {
		InstanceID InstanceID
	}

	type RemoveAllTracksFromQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveAllTracksFromQueueReq
	}

	type RemoveAllTracksFromQueueBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		RemoveAllTracksFromQueueNode RemoveAllTracksFromQueueNode `xml:"u:RemoveAllTracksFromQueue,omitempty"`
	}

	body := RemoveAllTracksFromQueueBody{RemoveAllTracksFromQueueNode: RemoveAllTracksFromQueueNode{
		RemoveAllTracksFromQueueReq: RemoveAllTracksFromQueueReq{InstanceID: 0},
		Xmlns:                       namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "RemoveAllTracksFromQueue", body, nil)...)
}

type AVTransportSaveQueueRes struct {
	AssignedObjectID ObjectID
}

// SaveQueue Saves the current SONOS queue as a SONOS playlist and outputs objectID
func (a AVTransportService) SaveQueue(ctx context.Context, title SavedQueueTitle, objectId ObjectID, saveQueue *AVTransportSaveQueueRes) error {
	type SaveQueueReq struct {
		InstanceID InstanceID
		Title      SavedQueueTitle
		ObjectID   ObjectID
	}

	type SaveQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SaveQueueReq
	}

	type SaveQueueBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		SaveQueueNode SaveQueueNode `xml:"u:SaveQueue,omitempty"`
	}

	body := SaveQueueBody{SaveQueueNode: SaveQueueNode{
		SaveQueueReq: SaveQueueReq{
			InstanceID: 0,
			ObjectID:   objectId,
			Title:      title,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SaveQueue", body, saveQueue)...)
}

func (a AVTransportService) BackupQueue(ctx context.Context) error {
	type BackupQueueReq struct {
		InstanceID InstanceID
	}

	type BackupQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BackupQueueReq
	}

	type BackupQueueBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		BackupQueueNode BackupQueueNode `xml:"u:BackupQueue,omitempty"`
	}

	body := BackupQueueBody{BackupQueueNode: BackupQueueNode{
		BackupQueueReq: BackupQueueReq{InstanceID: 0},
		Xmlns:          namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "BackupQueue", body, nil)...)
}

type AVTransportCreateSavedQueueRes struct {
	NumTracksAdded   NumTracks
	NewQueueLength   NumTracks
	AssignedObjectID ObjectID
	NewUpdateID      QueueUpdateID
}

func (a AVTransportService) CreateSavedQueue(ctx context.Context, title SavedQueueTitle, enqueuedUri URI, enqueuedUrimetaData URIMetaData, createSavedQueue *AVTransportCreateSavedQueueRes) error {
	type CreateSavedQueueReq struct {
		InstanceID          InstanceID
		Title               SavedQueueTitle
		EnqueuedURI         URI
		EnqueuedURIMetaData URIMetaData
	}

	type CreateSavedQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CreateSavedQueueReq
	}

	type CreateSavedQueueBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		CreateSavedQueueNode CreateSavedQueueNode `xml:"u:CreateSavedQueue,omitempty"`
	}

	body := CreateSavedQueueBody{CreateSavedQueueNode: CreateSavedQueueNode{
		CreateSavedQueueReq: CreateSavedQueueReq{
			EnqueuedURI:         enqueuedUri,
			EnqueuedURIMetaData: enqueuedUrimetaData,
			InstanceID:          0,
			Title:               title,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "CreateSavedQueue", body, createSavedQueue)...)
}

type AVTransportAddURIToSavedQueueRes struct {
	NumTracksAdded NumTracks
	NewQueueLength NumTracks
	NewUpdateID    QueueUpdateID
}

func (a AVTransportService) AddURIToSavedQueue(ctx context.Context, objectId ObjectID, updateId QueueUpdateID, enqueuedUri URI, enqueuedUrimetaData URIMetaData, addAtIndex TrackNumber, addUritoSavedQueue *AVTransportAddURIToSavedQueueRes) error {
	type AddURIToSavedQueueReq struct {
		InstanceID          InstanceID
		ObjectID            ObjectID
		UpdateID            QueueUpdateID
		EnqueuedURI         URI
		EnqueuedURIMetaData URIMetaData
		AddAtIndex          TrackNumber
	}

	type AddURIToSavedQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddURIToSavedQueueReq
	}

	type AddURIToSavedQueueBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		AddURIToSavedQueueNode AddURIToSavedQueueNode `xml:"u:AddURIToSavedQueue,omitempty"`
	}

	body := AddURIToSavedQueueBody{AddURIToSavedQueueNode: AddURIToSavedQueueNode{
		AddURIToSavedQueueReq: AddURIToSavedQueueReq{
			AddAtIndex:          addAtIndex,
			EnqueuedURI:         enqueuedUri,
			EnqueuedURIMetaData: enqueuedUrimetaData,
			InstanceID:          0,
			ObjectID:            objectId,
			UpdateID:            updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "AddURIToSavedQueue", body, addUritoSavedQueue)...)
}

type AVTransportReorderTracksInSavedQueueRes struct {
	QueueLengthChange NumTracksChange
	NewQueueLength    NumTracks
	NewUpdateID       QueueUpdateID
}

func (a AVTransportService) ReorderTracksInSavedQueue(ctx context.Context, objectId ObjectID, updateId QueueUpdateID, trackList TrackList, newPositionList TrackList, reorderTracksInSavedQueue *AVTransportReorderTracksInSavedQueueRes) error {
	type ReorderTracksInSavedQueueReq struct {
		InstanceID      InstanceID
		ObjectID        ObjectID
		UpdateID        QueueUpdateID
		TrackList       TrackList
		NewPositionList TrackList
	}

	type ReorderTracksInSavedQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReorderTracksInSavedQueueReq
	}

	type ReorderTracksInSavedQueueBody struct {
		XMLName                       xml.Name                      `xml:"s:Body"`
		ReorderTracksInSavedQueueNode ReorderTracksInSavedQueueNode `xml:"u:ReorderTracksInSavedQueue,omitempty"`
	}

	body := ReorderTracksInSavedQueueBody{ReorderTracksInSavedQueueNode: ReorderTracksInSavedQueueNode{
		ReorderTracksInSavedQueueReq: ReorderTracksInSavedQueueReq{
			InstanceID:      0,
			NewPositionList: newPositionList,
			ObjectID:        objectId,
			TrackList:       trackList,
			UpdateID:        updateId,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "ReorderTracksInSavedQueue", body, reorderTracksInSavedQueue)...)
}

type AVTransportGetMediaInfoRes struct {
	NrTracks           NumberOfTracks
	MediaDuration      MediaDuration
	CurrentURI         AVTransportURI
	CurrentURIMetaData AVTransportURIMetaData
	NextURI            AVTransportURI
	NextURIMetaData    AVTransportURIMetaData
	PlayMedium         PlaybackStorageMedium
	RecordMedium       RecordStorageMedium
	WriteStatus        RecordMediumWriteStatus
}

// GetMediaInfo Get information about the current playing media (queue)
func (a AVTransportService) GetMediaInfo(ctx context.Context, getMediaInfo *AVTransportGetMediaInfoRes) error {
	type GetMediaInfoReq struct {
		InstanceID InstanceID
	}

	type GetMediaInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetMediaInfoReq
	}

	type GetMediaInfoBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		GetMediaInfoNode GetMediaInfoNode `xml:"u:GetMediaInfo,omitempty"`
	}

	body := GetMediaInfoBody{GetMediaInfoNode: GetMediaInfoNode{
		GetMediaInfoReq: GetMediaInfoReq{InstanceID: 0},
		Xmlns:           namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetMediaInfo", body, getMediaInfo)...)
}

type AVTransportGetTransportInfoRes struct {
	CurrentTransportState  TransportState
	CurrentTransportStatus TransportStatus
	CurrentSpeed           TransportPlaySpeed
}

// GetTransportInfo Get current transport status, speed and state such as PLAYING, STOPPED, PLAYING, PAUSED_PLAYBACK, TRANSITIONING, NO_MEDIA_PRESENT
func (a AVTransportService) GetTransportInfo(ctx context.Context, getTransportInfo *AVTransportGetTransportInfoRes) error {
	type GetTransportInfoReq struct {
		InstanceID InstanceID
	}

	type GetTransportInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTransportInfoReq
	}

	type GetTransportInfoBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		GetTransportInfoNode GetTransportInfoNode `xml:"u:GetTransportInfo,omitempty"`
	}

	body := GetTransportInfoBody{GetTransportInfoNode: GetTransportInfoNode{
		GetTransportInfoReq: GetTransportInfoReq{InstanceID: 0},
		Xmlns:               namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetTransportInfo", body, getTransportInfo)...)
}

type AVTransportGetPositionInfoRes struct {
	Track         Track
	TrackDuration TrackDuration
	TrackMetaData TrackMetaData
	TrackURI      TrackURI
	RelTime       RelativeTimePosition
	AbsTime       AbsoluteTimePosition
	RelCount      RelativeCounterPosition
	AbsCount      AbsoluteCounterPosition
}

// GetPositionInfo Get information about current position (position in queue and time in current song)
func (a AVTransportService) GetPositionInfo(ctx context.Context, getPositionInfo *AVTransportGetPositionInfoRes) error {
	type GetPositionInfoReq struct {
		InstanceID InstanceID
	}

	type GetPositionInfoNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetPositionInfoReq
	}

	type GetPositionInfoBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		GetPositionInfoNode GetPositionInfoNode `xml:"u:GetPositionInfo,omitempty"`
	}

	body := GetPositionInfoBody{GetPositionInfoNode: GetPositionInfoNode{
		GetPositionInfoReq: GetPositionInfoReq{InstanceID: 0},
		Xmlns:              namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetPositionInfo", body, getPositionInfo)...)
}

type AVTransportGetDeviceCapabilitiesRes struct {
	PlayMedia       PossiblePlaybackStorageMedia
	RecMedia        PossibleRecordStorageMedia
	RecQualityModes PossibleRecordQualityModes
}

func (a AVTransportService) GetDeviceCapabilities(ctx context.Context, getDeviceCapabilities *AVTransportGetDeviceCapabilitiesRes) error {
	type GetDeviceCapabilitiesReq struct {
		InstanceID InstanceID
	}

	type GetDeviceCapabilitiesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetDeviceCapabilitiesReq
	}

	type GetDeviceCapabilitiesBody struct {
		XMLName                   xml.Name                  `xml:"s:Body"`
		GetDeviceCapabilitiesNode GetDeviceCapabilitiesNode `xml:"u:GetDeviceCapabilities,omitempty"`
	}

	body := GetDeviceCapabilitiesBody{GetDeviceCapabilitiesNode: GetDeviceCapabilitiesNode{
		GetDeviceCapabilitiesReq: GetDeviceCapabilitiesReq{InstanceID: 0},
		Xmlns:                    namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetDeviceCapabilities", body, getDeviceCapabilities)...)
}

type AVTransportGetTransportSettingsRes struct {
	PlayMode       PlayMode
	RecQualityMode RecordQualityMode
}

// GetTransportSettings Get transport settings
func (a AVTransportService) GetTransportSettings(ctx context.Context, getTransportSettings *AVTransportGetTransportSettingsRes) error {
	type GetTransportSettingsReq struct {
		InstanceID InstanceID
	}

	type GetTransportSettingsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTransportSettingsReq
	}

	type GetTransportSettingsBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		GetTransportSettingsNode GetTransportSettingsNode `xml:"u:GetTransportSettings,omitempty"`
	}

	body := GetTransportSettingsBody{GetTransportSettingsNode: GetTransportSettingsNode{
		GetTransportSettingsReq: GetTransportSettingsReq{InstanceID: 0},
		Xmlns:                   namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetTransportSettings", body, getTransportSettings)...)
}

type AVTransportGetCrossfadeModeRes struct {
	CrossfadeMode CrossfadeMode
}

// GetCrossfadeMode Get crossfade mode
func (a AVTransportService) GetCrossfadeMode(ctx context.Context, getCrossfadeMode *AVTransportGetCrossfadeModeRes) error {
	type GetCrossfadeModeReq struct {
		InstanceID InstanceID
	}

	type GetCrossfadeModeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetCrossfadeModeReq
	}

	type GetCrossfadeModeBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		GetCrossfadeModeNode GetCrossfadeModeNode `xml:"u:GetCrossfadeMode,omitempty"`
	}

	body := GetCrossfadeModeBody{GetCrossfadeModeNode: GetCrossfadeModeNode{
		GetCrossfadeModeReq: GetCrossfadeModeReq{InstanceID: 0},
		Xmlns:               namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetCrossfadeMode", body, getCrossfadeMode)...)
}

// Stop Stop playback
func (a AVTransportService) Stop(ctx context.Context) error {
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
		Xmlns:   namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Stop", body, nil)...)
}

// Play Start playing the set TransportURI
func (a AVTransportService) Play(ctx context.Context, speed TransportPlaySpeed) error {
	type PlayReq struct {
		InstanceID InstanceID
		Speed      TransportPlaySpeed
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
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Play", body, nil)...)
}

// Pause Pause playback
func (a AVTransportService) Pause(ctx context.Context) error {
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
		Xmlns:    namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Pause", body, nil)...)
}

// Seek Seek track in queue, time delta or absolute time in song
func (a AVTransportService) Seek(ctx context.Context, unit SeekMode, target SeekTarget) error {
	type SeekReq struct {
		InstanceID InstanceID
		Unit       SeekMode
		Target     SeekTarget
	}

	type SeekNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SeekReq
	}

	type SeekBody struct {
		XMLName  xml.Name `xml:"s:Body"`
		SeekNode SeekNode `xml:"u:Seek,omitempty"`
	}

	body := SeekBody{SeekNode: SeekNode{
		SeekReq: SeekReq{
			InstanceID: 0,
			Target:     target,
			Unit:       unit,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Seek", body, nil)...)
}

// Next Go to next song
func (a AVTransportService) Next(ctx context.Context) error {
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
		Xmlns:   namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Next", body, nil)...)
}

// Previous Go to previous song
func (a AVTransportService) Previous(ctx context.Context) error {
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
		Xmlns:       namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "Previous", body, nil)...)
}

// SetPlayMode Set the PlayMode
func (a AVTransportService) SetPlayMode(ctx context.Context, newPlayMode PlayMode) error {
	type SetPlayModeReq struct {
		InstanceID  InstanceID
		NewPlayMode PlayMode
	}

	type SetPlayModeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetPlayModeReq
	}

	type SetPlayModeBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SetPlayModeNode SetPlayModeNode `xml:"u:SetPlayMode,omitempty"`
	}

	body := SetPlayModeBody{SetPlayModeNode: SetPlayModeNode{
		SetPlayModeReq: SetPlayModeReq{
			InstanceID:  0,
			NewPlayMode: newPlayMode,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SetPlayMode", body, nil)...)
}

// SetCrossfadeMode Set crossfade mode
func (a AVTransportService) SetCrossfadeMode(ctx context.Context, crossfadeMode CrossfadeMode) error {
	type SetCrossfadeModeReq struct {
		InstanceID    InstanceID
		CrossfadeMode CrossfadeMode
	}

	type SetCrossfadeModeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetCrossfadeModeReq
	}

	type SetCrossfadeModeBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		SetCrossfadeModeNode SetCrossfadeModeNode `xml:"u:SetCrossfadeMode,omitempty"`
	}

	body := SetCrossfadeModeBody{SetCrossfadeModeNode: SetCrossfadeModeNode{
		SetCrossfadeModeReq: SetCrossfadeModeReq{
			CrossfadeMode: crossfadeMode,
			InstanceID:    0,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SetCrossfadeMode", body, nil)...)
}

func (a AVTransportService) NotifyDeletedURI(ctx context.Context, deletedUri AVTransportURI) error {
	type NotifyDeletedURIReq struct {
		InstanceID InstanceID
		DeletedURI AVTransportURI
	}

	type NotifyDeletedURINode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		NotifyDeletedURIReq
	}

	type NotifyDeletedURIBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		NotifyDeletedURINode NotifyDeletedURINode `xml:"u:NotifyDeletedURI,omitempty"`
	}

	body := NotifyDeletedURIBody{NotifyDeletedURINode: NotifyDeletedURINode{
		NotifyDeletedURIReq: NotifyDeletedURIReq{
			DeletedURI: deletedUri,
			InstanceID: 0,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "NotifyDeletedURI", body, nil)...)
}

type AVTransportGetCurrentTransportActionsRes struct {
	Actions TransportActions
}

// GetCurrentTransportActions Get current transport actions such as Set, Stop, Pause, Play, X_DLNA_SeekTime, Next, X_DLNA_SeekTrackNr
func (a AVTransportService) GetCurrentTransportActions(ctx context.Context, getCurrentTransportActions *AVTransportGetCurrentTransportActionsRes) error {
	type GetCurrentTransportActionsReq struct {
		InstanceID InstanceID
	}

	type GetCurrentTransportActionsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetCurrentTransportActionsReq
	}

	type GetCurrentTransportActionsBody struct {
		XMLName                        xml.Name                       `xml:"s:Body"`
		GetCurrentTransportActionsNode GetCurrentTransportActionsNode `xml:"u:GetCurrentTransportActions,omitempty"`
	}

	body := GetCurrentTransportActionsBody{GetCurrentTransportActionsNode: GetCurrentTransportActionsNode{
		GetCurrentTransportActionsReq: GetCurrentTransportActionsReq{InstanceID: 0},
		Xmlns:                         namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetCurrentTransportActions", body, getCurrentTransportActions)...)
}

type AVTransportBecomeCoordinatorOfStandaloneGroupRes struct {
	DelegatedGroupCoordinatorID PlayerID
	NewGroupID                  GroupID
}

// BecomeCoordinatorOfStandaloneGroup Leave the current group and revert to a single player.
func (a AVTransportService) BecomeCoordinatorOfStandaloneGroup(ctx context.Context, becomeCoordinatorOfStandaloneGroup *AVTransportBecomeCoordinatorOfStandaloneGroupRes) error {
	type BecomeCoordinatorOfStandaloneGroupReq struct {
		InstanceID InstanceID
	}

	type BecomeCoordinatorOfStandaloneGroupNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BecomeCoordinatorOfStandaloneGroupReq
	}

	type BecomeCoordinatorOfStandaloneGroupBody struct {
		XMLName                                xml.Name                               `xml:"s:Body"`
		BecomeCoordinatorOfStandaloneGroupNode BecomeCoordinatorOfStandaloneGroupNode `xml:"u:BecomeCoordinatorOfStandaloneGroup,omitempty"`
	}

	body := BecomeCoordinatorOfStandaloneGroupBody{BecomeCoordinatorOfStandaloneGroupNode: BecomeCoordinatorOfStandaloneGroupNode{
		BecomeCoordinatorOfStandaloneGroupReq: BecomeCoordinatorOfStandaloneGroupReq{InstanceID: 0},
		Xmlns:                                 namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "BecomeCoordinatorOfStandaloneGroup", body, becomeCoordinatorOfStandaloneGroup)...)
}

// DelegateGroupCoordinationTo Delegates the coordinator role to another player in the same group
func (a AVTransportService) DelegateGroupCoordinationTo(ctx context.Context, newCoordinator MemberID, rejoinGroup RejoinGroup) error {
	type DelegateGroupCoordinationToReq struct {
		InstanceID     InstanceID
		NewCoordinator MemberID
		RejoinGroup    RejoinGroup
	}

	type DelegateGroupCoordinationToNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		DelegateGroupCoordinationToReq
	}

	type DelegateGroupCoordinationToBody struct {
		XMLName                         xml.Name                        `xml:"s:Body"`
		DelegateGroupCoordinationToNode DelegateGroupCoordinationToNode `xml:"u:DelegateGroupCoordinationTo,omitempty"`
	}

	body := DelegateGroupCoordinationToBody{DelegateGroupCoordinationToNode: DelegateGroupCoordinationToNode{
		DelegateGroupCoordinationToReq: DelegateGroupCoordinationToReq{
			InstanceID:     0,
			NewCoordinator: newCoordinator,
			RejoinGroup:    rejoinGroup,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "DelegateGroupCoordinationTo", body, nil)...)
}

func (a AVTransportService) BecomeGroupCoordinator(ctx context.Context, currentCoordinator MemberID, currentGroupId GroupID, otherMembers MemberList, transportSettings TransportSettings, currentUri AVTransportURI, currentUrimetaData AVTransportURIMetaData, sleepTimerState SleepTimerState, alarmState AlarmState, streamRestartState StreamRestartState, currentQueueTrackList Queue, currentVlistate VLIState) error {
	type BecomeGroupCoordinatorReq struct {
		InstanceID            InstanceID
		CurrentCoordinator    MemberID
		CurrentGroupID        GroupID
		OtherMembers          MemberList
		TransportSettings     TransportSettings
		CurrentURI            AVTransportURI
		CurrentURIMetaData    AVTransportURIMetaData
		SleepTimerState       SleepTimerState
		AlarmState            AlarmState
		StreamRestartState    StreamRestartState
		CurrentQueueTrackList Queue
		CurrentVLIState       VLIState
	}

	type BecomeGroupCoordinatorNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BecomeGroupCoordinatorReq
	}

	type BecomeGroupCoordinatorBody struct {
		XMLName                    xml.Name                   `xml:"s:Body"`
		BecomeGroupCoordinatorNode BecomeGroupCoordinatorNode `xml:"u:BecomeGroupCoordinator,omitempty"`
	}

	body := BecomeGroupCoordinatorBody{BecomeGroupCoordinatorNode: BecomeGroupCoordinatorNode{
		BecomeGroupCoordinatorReq: BecomeGroupCoordinatorReq{
			AlarmState:            alarmState,
			CurrentCoordinator:    currentCoordinator,
			CurrentGroupID:        currentGroupId,
			CurrentQueueTrackList: currentQueueTrackList,
			CurrentURI:            currentUri,
			CurrentURIMetaData:    currentUrimetaData,
			CurrentVLIState:       currentVlistate,
			InstanceID:            0,
			OtherMembers:          otherMembers,
			SleepTimerState:       sleepTimerState,
			StreamRestartState:    streamRestartState,
			TransportSettings:     transportSettings,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "BecomeGroupCoordinator", body, nil)...)
}

func (a AVTransportService) BecomeGroupCoordinatorAndSource(ctx context.Context, currentCoordinator MemberID, currentGroupId GroupID, otherMembers MemberList, currentUri AVTransportURI, currentUrimetaData AVTransportURIMetaData, sleepTimerState SleepTimerState, alarmState AlarmState, streamRestartState StreamRestartState, currentAvttrackList Queue, currentQueueTrackList Queue, currentSourceState SourceState, resumePlayback ResumePlayback) error {
	type BecomeGroupCoordinatorAndSourceReq struct {
		InstanceID            InstanceID
		CurrentCoordinator    MemberID
		CurrentGroupID        GroupID
		OtherMembers          MemberList
		CurrentURI            AVTransportURI
		CurrentURIMetaData    AVTransportURIMetaData
		SleepTimerState       SleepTimerState
		AlarmState            AlarmState
		StreamRestartState    StreamRestartState
		CurrentAVTTrackList   Queue
		CurrentQueueTrackList Queue
		CurrentSourceState    SourceState
		ResumePlayback        ResumePlayback
	}

	type BecomeGroupCoordinatorAndSourceNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BecomeGroupCoordinatorAndSourceReq
	}

	type BecomeGroupCoordinatorAndSourceBody struct {
		XMLName                             xml.Name                            `xml:"s:Body"`
		BecomeGroupCoordinatorAndSourceNode BecomeGroupCoordinatorAndSourceNode `xml:"u:BecomeGroupCoordinatorAndSource,omitempty"`
	}

	body := BecomeGroupCoordinatorAndSourceBody{BecomeGroupCoordinatorAndSourceNode: BecomeGroupCoordinatorAndSourceNode{
		BecomeGroupCoordinatorAndSourceReq: BecomeGroupCoordinatorAndSourceReq{
			AlarmState:            alarmState,
			CurrentAVTTrackList:   currentAvttrackList,
			CurrentCoordinator:    currentCoordinator,
			CurrentGroupID:        currentGroupId,
			CurrentQueueTrackList: currentQueueTrackList,
			CurrentSourceState:    currentSourceState,
			CurrentURI:            currentUri,
			CurrentURIMetaData:    currentUrimetaData,
			InstanceID:            0,
			OtherMembers:          otherMembers,
			ResumePlayback:        resumePlayback,
			SleepTimerState:       sleepTimerState,
			StreamRestartState:    streamRestartState,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "BecomeGroupCoordinatorAndSource", body, nil)...)
}

func (a AVTransportService) ChangeCoordinator(ctx context.Context, currentCoordinator MemberID, newCoordinator MemberID, newTransportSettings TransportSettings, currentAvtransportUri AVTransportURI) error {
	type ChangeCoordinatorReq struct {
		InstanceID            InstanceID
		CurrentCoordinator    MemberID
		NewCoordinator        MemberID
		NewTransportSettings  TransportSettings
		CurrentAVTransportURI AVTransportURI
	}

	type ChangeCoordinatorNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ChangeCoordinatorReq
	}

	type ChangeCoordinatorBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		ChangeCoordinatorNode ChangeCoordinatorNode `xml:"u:ChangeCoordinator,omitempty"`
	}

	body := ChangeCoordinatorBody{ChangeCoordinatorNode: ChangeCoordinatorNode{
		ChangeCoordinatorReq: ChangeCoordinatorReq{
			CurrentAVTransportURI: currentAvtransportUri,
			CurrentCoordinator:    currentCoordinator,
			InstanceID:            0,
			NewCoordinator:        newCoordinator,
			NewTransportSettings:  newTransportSettings,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "ChangeCoordinator", body, nil)...)
}

func (a AVTransportService) ChangeTransportSettings(ctx context.Context, newTransportSettings TransportSettings, currentAvtransportUri AVTransportURI) error {
	type ChangeTransportSettingsReq struct {
		InstanceID            InstanceID
		NewTransportSettings  TransportSettings
		CurrentAVTransportURI AVTransportURI
	}

	type ChangeTransportSettingsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ChangeTransportSettingsReq
	}

	type ChangeTransportSettingsBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		ChangeTransportSettingsNode ChangeTransportSettingsNode `xml:"u:ChangeTransportSettings,omitempty"`
	}

	body := ChangeTransportSettingsBody{ChangeTransportSettingsNode: ChangeTransportSettingsNode{
		ChangeTransportSettingsReq: ChangeTransportSettingsReq{
			CurrentAVTransportURI: currentAvtransportUri,
			InstanceID:            0,
			NewTransportSettings:  newTransportSettings,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "ChangeTransportSettings", body, nil)...)
}

// ConfigureSleepTimer Stop playing after set sleep timer or cancel
func (a AVTransportService) ConfigureSleepTimer(ctx context.Context, newSleepTimerDuration ISO8601Time) error {
	type ConfigureSleepTimerReq struct {
		InstanceID            InstanceID
		NewSleepTimerDuration ISO8601Time
	}

	type ConfigureSleepTimerNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ConfigureSleepTimerReq
	}

	type ConfigureSleepTimerBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		ConfigureSleepTimerNode ConfigureSleepTimerNode `xml:"u:ConfigureSleepTimer,omitempty"`
	}

	body := ConfigureSleepTimerBody{ConfigureSleepTimerNode: ConfigureSleepTimerNode{
		ConfigureSleepTimerReq: ConfigureSleepTimerReq{
			InstanceID:            0,
			NewSleepTimerDuration: newSleepTimerDuration,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "ConfigureSleepTimer", body, nil)...)
}

type AVTransportGetRemainingSleepTimerDurationRes struct {
	RemainingSleepTimerDuration ISO8601Time
	CurrentSleepTimerGeneration SleepTimerGeneration
}

// GetRemainingSleepTimerDuration Get time left on sleeptimer.
func (a AVTransportService) GetRemainingSleepTimerDuration(ctx context.Context, getRemainingSleepTimerDuration *AVTransportGetRemainingSleepTimerDurationRes) error {
	type GetRemainingSleepTimerDurationReq struct {
		InstanceID InstanceID
	}

	type GetRemainingSleepTimerDurationNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetRemainingSleepTimerDurationReq
	}

	type GetRemainingSleepTimerDurationBody struct {
		XMLName                            xml.Name                           `xml:"s:Body"`
		GetRemainingSleepTimerDurationNode GetRemainingSleepTimerDurationNode `xml:"u:GetRemainingSleepTimerDuration,omitempty"`
	}

	body := GetRemainingSleepTimerDurationBody{GetRemainingSleepTimerDurationNode: GetRemainingSleepTimerDurationNode{
		GetRemainingSleepTimerDurationReq: GetRemainingSleepTimerDurationReq{InstanceID: 0},
		Xmlns:                             namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetRemainingSleepTimerDuration", body, getRemainingSleepTimerDuration)...)
}

func (a AVTransportService) RunAlarm(ctx context.Context, alarmId AlarmIDRunning, loggedStartTime AlarmLoggedStartTime, duration ISO8601Time, programUri AVTransportURI, programMetaData AVTransportURIMetaData, playMode PlayMode, volume AlarmVolume, includeLinkedZones AlarmIncludeLinkedZones) error {
	type RunAlarmReq struct {
		InstanceID         InstanceID
		AlarmID            AlarmIDRunning
		LoggedStartTime    AlarmLoggedStartTime
		Duration           ISO8601Time
		ProgramURI         AVTransportURI
		ProgramMetaData    AVTransportURIMetaData
		PlayMode           PlayMode
		Volume             AlarmVolume
		IncludeLinkedZones AlarmIncludeLinkedZones
	}

	type RunAlarmNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RunAlarmReq
	}

	type RunAlarmBody struct {
		XMLName      xml.Name     `xml:"s:Body"`
		RunAlarmNode RunAlarmNode `xml:"u:RunAlarm,omitempty"`
	}

	body := RunAlarmBody{RunAlarmNode: RunAlarmNode{
		RunAlarmReq: RunAlarmReq{
			AlarmID:            alarmId,
			Duration:           duration,
			IncludeLinkedZones: includeLinkedZones,
			InstanceID:         0,
			LoggedStartTime:    loggedStartTime,
			PlayMode:           playMode,
			ProgramMetaData:    programMetaData,
			ProgramURI:         programUri,
			Volume:             volume,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "RunAlarm", body, nil)...)
}

func (a AVTransportService) StartAutoplay(ctx context.Context, programUri AVTransportURI, programMetaData AVTransportURIMetaData, volume AlarmVolume, includeLinkedZones AlarmIncludeLinkedZones, resetVolumeAfter ResetVolumeAfter) error {
	type StartAutoplayReq struct {
		InstanceID         InstanceID
		ProgramURI         AVTransportURI
		ProgramMetaData    AVTransportURIMetaData
		Volume             AlarmVolume
		IncludeLinkedZones AlarmIncludeLinkedZones
		ResetVolumeAfter   ResetVolumeAfter
	}

	type StartAutoplayNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		StartAutoplayReq
	}

	type StartAutoplayBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		StartAutoplayNode StartAutoplayNode `xml:"u:StartAutoplay,omitempty"`
	}

	body := StartAutoplayBody{StartAutoplayNode: StartAutoplayNode{
		StartAutoplayReq: StartAutoplayReq{
			IncludeLinkedZones: includeLinkedZones,
			InstanceID:         0,
			ProgramMetaData:    programMetaData,
			ProgramURI:         programUri,
			ResetVolumeAfter:   resetVolumeAfter,
			Volume:             volume,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "StartAutoplay", body, nil)...)
}

type AVTransportGetRunningAlarmPropertiesRes struct {
	AlarmID         AlarmIDRunning
	GroupID         GroupID
	LoggedStartTime AlarmLoggedStartTime
}

func (a AVTransportService) GetRunningAlarmProperties(ctx context.Context, getRunningAlarmProperties *AVTransportGetRunningAlarmPropertiesRes) error {
	type GetRunningAlarmPropertiesReq struct {
		InstanceID InstanceID
	}

	type GetRunningAlarmPropertiesNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetRunningAlarmPropertiesReq
	}

	type GetRunningAlarmPropertiesBody struct {
		XMLName                       xml.Name                      `xml:"s:Body"`
		GetRunningAlarmPropertiesNode GetRunningAlarmPropertiesNode `xml:"u:GetRunningAlarmProperties,omitempty"`
	}

	body := GetRunningAlarmPropertiesBody{GetRunningAlarmPropertiesNode: GetRunningAlarmPropertiesNode{
		GetRunningAlarmPropertiesReq: GetRunningAlarmPropertiesReq{InstanceID: 0},
		Xmlns:                        namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "GetRunningAlarmProperties", body, getRunningAlarmProperties)...)
}

// SnoozeAlarm Snooze the current alarm for some time.
func (a AVTransportService) SnoozeAlarm(ctx context.Context, duration ISO8601Time) error {
	type SnoozeAlarmReq struct {
		InstanceID InstanceID
		Duration   ISO8601Time
	}

	type SnoozeAlarmNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SnoozeAlarmReq
	}

	type SnoozeAlarmBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SnoozeAlarmNode SnoozeAlarmNode `xml:"u:SnoozeAlarm,omitempty"`
	}

	body := SnoozeAlarmBody{SnoozeAlarmNode: SnoozeAlarmNode{
		SnoozeAlarmReq: SnoozeAlarmReq{
			Duration:   duration,
			InstanceID: 0,
		},
		Xmlns: namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "SnoozeAlarm", body, nil)...)
}

func (a AVTransportService) EndDirectControlSession(ctx context.Context) error {
	type EndDirectControlSessionReq struct {
		InstanceID InstanceID
	}

	type EndDirectControlSessionNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		EndDirectControlSessionReq
	}

	type EndDirectControlSessionBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		EndDirectControlSessionNode EndDirectControlSessionNode `xml:"u:EndDirectControlSession,omitempty"`
	}

	body := EndDirectControlSessionBody{EndDirectControlSessionNode: EndDirectControlSessionNode{
		EndDirectControlSessionReq: EndDirectControlSessionReq{InstanceID: 0},
		Xmlns:                      namespaceAVTransportService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAVTransportService, namespaceAVTransportService, "EndDirectControlSession", body, nil)...)
}
