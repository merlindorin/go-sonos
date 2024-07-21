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
	// namespaceQueueService is the UPnP namespace for the Sonos service.
	namespaceQueueService = "urn:schemas-sonos-com:service:Queue:1"

	// eventQueueService is the event path for the Sonos event service.
	eventQueueService = "/MediaRenderer/Queue/Event"

	// controlQueueService is the action path for the Sonos service.
	controlQueueService = "/MediaRenderer/Queue/Control"
)

// QueueService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type QueueService struct {
	doer do.Doer
}

// NewQueueService initializes a new QueueService with the provided Doer.
func NewQueueService(doer do.Doer) *QueueService {
	return &QueueService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (q QueueService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return q.doer.Do(ctx, do.WithPath(eventQueueService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

type QueueAddURIRes struct {
	FirstTrackNumberEnqueued TrackNumber
	NumTracksAdded           NumTracks
	NewQueueLength           NumTracks
	NewUpdateID              UpdateID
}

func (q QueueService) AddURI(ctx context.Context, queueId QueueID, updateId UpdateID, enqueuedUri URI, enqueuedUrimetaData URIMetaData, desiredFirstTrackNumberEnqueued TrackNumber, enqueueAsNext EnqueueAs, addUri *QueueAddURIRes) error {
	type AddURIReq struct {
		QueueID                         QueueID
		UpdateID                        UpdateID
		EnqueuedURI                     URI
		EnqueuedURIMetaData             URIMetaData
		DesiredFirstTrackNumberEnqueued TrackNumber
		EnqueueAsNext                   EnqueueAs
	}

	type AddURINode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddURIReq
	}

	type AddURIBody struct {
		XMLName    xml.Name   `xml:"s:Body"`
		AddURINode AddURINode `xml:"u:AddURI,omitempty"`
	}

	body := AddURIBody{AddURINode: AddURINode{
		AddURIReq: AddURIReq{
			DesiredFirstTrackNumberEnqueued: desiredFirstTrackNumberEnqueued,
			EnqueueAsNext:                   enqueueAsNext,
			EnqueuedURI:                     enqueuedUri,
			EnqueuedURIMetaData:             enqueuedUrimetaData,
			QueueID:                         queueId,
			UpdateID:                        updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "AddURI", body, addUri)...)
}

type QueueAddMultipleURIsRes struct {
	FirstTrackNumberEnqueued TrackNumber
	NumTracksAdded           NumTracks
	NewQueueLength           NumTracks
	NewUpdateID              UpdateID
}

func (q QueueService) AddMultipleURIs(ctx context.Context, queueId QueueID, updateId UpdateID, containerUri URI, containerMetaData URIMetaData, desiredFirstTrackNumberEnqueued TrackNumber, enqueueAsNext EnqueueAs, numberOfUris NumTracks, enqueuedUrisAndMetaData LIST_URI_AND_METADATA, addMultipleUris *QueueAddMultipleURIsRes) error {
	type AddMultipleURIsReq struct {
		QueueID                         QueueID
		UpdateID                        UpdateID
		ContainerURI                    URI
		ContainerMetaData               URIMetaData
		DesiredFirstTrackNumberEnqueued TrackNumber
		EnqueueAsNext                   EnqueueAs
		NumberOfURIs                    NumTracks
		EnqueuedURIsAndMetaData         LIST_URI_AND_METADATA
	}

	type AddMultipleURIsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddMultipleURIsReq
	}

	type AddMultipleURIsBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		AddMultipleURIsNode AddMultipleURIsNode `xml:"u:AddMultipleURIs,omitempty"`
	}

	body := AddMultipleURIsBody{AddMultipleURIsNode: AddMultipleURIsNode{
		AddMultipleURIsReq: AddMultipleURIsReq{
			ContainerMetaData:               containerMetaData,
			ContainerURI:                    containerUri,
			DesiredFirstTrackNumberEnqueued: desiredFirstTrackNumberEnqueued,
			EnqueueAsNext:                   enqueueAsNext,
			EnqueuedURIsAndMetaData:         enqueuedUrisAndMetaData,
			NumberOfURIs:                    numberOfUris,
			QueueID:                         queueId,
			UpdateID:                        updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "AddMultipleURIs", body, addMultipleUris)...)
}

type QueueAttachQueueRes struct {
	QueueID           QueueID
	QueueOwnerContext QueueOwnerContext
}

func (q QueueService) AttachQueue(ctx context.Context, queueOwnerId QueueOwnerID, attachQueue *QueueAttachQueueRes) error {
	type AttachQueueReq struct {
		QueueOwnerID QueueOwnerID
	}

	type AttachQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AttachQueueReq
	}

	type AttachQueueBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		AttachQueueNode AttachQueueNode `xml:"u:AttachQueue,omitempty"`
	}

	body := AttachQueueBody{AttachQueueNode: AttachQueueNode{
		AttachQueueReq: AttachQueueReq{QueueOwnerID: queueOwnerId},
		Xmlns:          namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "AttachQueue", body, attachQueue)...)
}

func (q QueueService) Backup(ctx context.Context) error {
	type BackupReq struct{}

	type BackupNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		BackupReq
	}

	type BackupBody struct {
		XMLName    xml.Name   `xml:"s:Body"`
		BackupNode BackupNode `xml:"u:Backup,omitempty"`
	}

	body := BackupBody{BackupNode: BackupNode{
		BackupReq: BackupReq{},
		Xmlns:     namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "Backup", body, nil)...)
}

type QueueBrowseRes struct {
	Result         Result
	NumberReturned Count
	TotalMatches   Count
	UpdateID       UpdateID
}

func (q QueueService) Browse(ctx context.Context, queueId QueueID, startingIndex Index, requestedCount Count, browse *QueueBrowseRes) error {
	type BrowseReq struct {
		QueueID        QueueID
		StartingIndex  Index
		RequestedCount Count
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
			QueueID:        queueId,
			RequestedCount: requestedCount,
			StartingIndex:  startingIndex,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "Browse", body, browse)...)
}

type QueueCreateQueueRes struct {
	QueueID QueueID
}

func (q QueueService) CreateQueue(ctx context.Context, queueOwnerId QueueOwnerID, queueOwnerContext QueueOwnerContext, queuePolicy QueuePolicy, createQueue *QueueCreateQueueRes) error {
	type CreateQueueReq struct {
		QueueOwnerID      QueueOwnerID
		QueueOwnerContext QueueOwnerContext
		QueuePolicy       QueuePolicy
	}

	type CreateQueueNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CreateQueueReq
	}

	type CreateQueueBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		CreateQueueNode CreateQueueNode `xml:"u:CreateQueue,omitempty"`
	}

	body := CreateQueueBody{CreateQueueNode: CreateQueueNode{
		CreateQueueReq: CreateQueueReq{
			QueueOwnerContext: queueOwnerContext,
			QueueOwnerID:      queueOwnerId,
			QueuePolicy:       queuePolicy,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "CreateQueue", body, createQueue)...)
}

type QueueRemoveAllTracksRes struct {
	NewUpdateID UpdateID
}

func (q QueueService) RemoveAllTracks(ctx context.Context, queueId QueueID, updateId UpdateID, removeAllTracks *QueueRemoveAllTracksRes) error {
	type RemoveAllTracksReq struct {
		QueueID  QueueID
		UpdateID UpdateID
	}

	type RemoveAllTracksNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveAllTracksReq
	}

	type RemoveAllTracksBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		RemoveAllTracksNode RemoveAllTracksNode `xml:"u:RemoveAllTracks,omitempty"`
	}

	body := RemoveAllTracksBody{RemoveAllTracksNode: RemoveAllTracksNode{
		RemoveAllTracksReq: RemoveAllTracksReq{
			QueueID:  queueId,
			UpdateID: updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "RemoveAllTracks", body, removeAllTracks)...)
}

type QueueRemoveTrackRangeRes struct {
	NewUpdateID UpdateID
}

func (q QueueService) RemoveTrackRange(ctx context.Context, queueId QueueID, updateId UpdateID, startingIndex TrackNumber, numberOfTracks NumTracks, removeTrackRange *QueueRemoveTrackRangeRes) error {
	type RemoveTrackRangeReq struct {
		QueueID        QueueID
		UpdateID       UpdateID
		StartingIndex  TrackNumber
		NumberOfTracks NumTracks
	}

	type RemoveTrackRangeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveTrackRangeReq
	}

	type RemoveTrackRangeBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		RemoveTrackRangeNode RemoveTrackRangeNode `xml:"u:RemoveTrackRange,omitempty"`
	}

	body := RemoveTrackRangeBody{RemoveTrackRangeNode: RemoveTrackRangeNode{
		RemoveTrackRangeReq: RemoveTrackRangeReq{
			NumberOfTracks: numberOfTracks,
			QueueID:        queueId,
			StartingIndex:  startingIndex,
			UpdateID:       updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "RemoveTrackRange", body, removeTrackRange)...)
}

type QueueReorderTracksRes struct {
	NewUpdateID UpdateID
}

func (q QueueService) ReorderTracks(ctx context.Context, queueId QueueID, startingIndex TrackNumber, numberOfTracks NumTracks, insertBefore TrackNumber, updateId UpdateID, reorderTracks *QueueReorderTracksRes) error {
	type ReorderTracksReq struct {
		QueueID        QueueID
		StartingIndex  TrackNumber
		NumberOfTracks NumTracks
		InsertBefore   TrackNumber
		UpdateID       UpdateID
	}

	type ReorderTracksNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReorderTracksReq
	}

	type ReorderTracksBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		ReorderTracksNode ReorderTracksNode `xml:"u:ReorderTracks,omitempty"`
	}

	body := ReorderTracksBody{ReorderTracksNode: ReorderTracksNode{
		ReorderTracksReq: ReorderTracksReq{
			InsertBefore:   insertBefore,
			NumberOfTracks: numberOfTracks,
			QueueID:        queueId,
			StartingIndex:  startingIndex,
			UpdateID:       updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "ReorderTracks", body, reorderTracks)...)
}

type QueueReplaceAllTracksRes struct {
	NewQueueLength NumTracks
	NewUpdateID    UpdateID
}

func (q QueueService) ReplaceAllTracks(ctx context.Context, queueId QueueID, updateId UpdateID, containerUri URI, containerMetaData URIMetaData, currentTrackIndex TrackNumber, newCurrentTrackIndices TrackNumbersCSV, numberOfUris NumTracks, enqueuedUrisAndMetaData LIST_URI_AND_METADATA, replaceAllTracks *QueueReplaceAllTracksRes) error {
	type ReplaceAllTracksReq struct {
		QueueID                 QueueID
		UpdateID                UpdateID
		ContainerURI            URI
		ContainerMetaData       URIMetaData
		CurrentTrackIndex       TrackNumber
		NewCurrentTrackIndices  TrackNumbersCSV
		NumberOfURIs            NumTracks
		EnqueuedURIsAndMetaData LIST_URI_AND_METADATA
	}

	type ReplaceAllTracksNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReplaceAllTracksReq
	}

	type ReplaceAllTracksBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		ReplaceAllTracksNode ReplaceAllTracksNode `xml:"u:ReplaceAllTracks,omitempty"`
	}

	body := ReplaceAllTracksBody{ReplaceAllTracksNode: ReplaceAllTracksNode{
		ReplaceAllTracksReq: ReplaceAllTracksReq{
			ContainerMetaData:       containerMetaData,
			ContainerURI:            containerUri,
			CurrentTrackIndex:       currentTrackIndex,
			EnqueuedURIsAndMetaData: enqueuedUrisAndMetaData,
			NewCurrentTrackIndices:  newCurrentTrackIndices,
			NumberOfURIs:            numberOfUris,
			QueueID:                 queueId,
			UpdateID:                updateId,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "ReplaceAllTracks", body, replaceAllTracks)...)
}

type QueueSaveAsSonosPlaylistRes struct {
	AssignedObjectID ObjectID
}

func (q QueueService) SaveAsSonosPlaylist(ctx context.Context, queueId QueueID, title SavedQueueTitle, objectId ObjectID, saveAsSonosPlaylist *QueueSaveAsSonosPlaylistRes) error {
	type SaveAsSonosPlaylistReq struct {
		QueueID  QueueID
		Title    SavedQueueTitle
		ObjectID ObjectID
	}

	type SaveAsSonosPlaylistNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SaveAsSonosPlaylistReq
	}

	type SaveAsSonosPlaylistBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		SaveAsSonosPlaylistNode SaveAsSonosPlaylistNode `xml:"u:SaveAsSonosPlaylist,omitempty"`
	}

	body := SaveAsSonosPlaylistBody{SaveAsSonosPlaylistNode: SaveAsSonosPlaylistNode{
		SaveAsSonosPlaylistReq: SaveAsSonosPlaylistReq{
			ObjectID: objectId,
			QueueID:  queueId,
			Title:    title,
		},
		Xmlns: namespaceQueueService,
	}}

	return q.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlQueueService, namespaceQueueService, "SaveAsSonosPlaylist", body, saveAsSonosPlaylist)...)
}
