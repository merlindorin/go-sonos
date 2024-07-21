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
	// namespaceAlarmClockService is the UPnP namespace for the Sonos service.
	namespaceAlarmClockService = "urn:schemas-upnp-org:service:AlarmClock:1"

	// eventAlarmClockService is the event path for the Sonos event service.
	eventAlarmClockService = "/AlarmClock/Event"

	// controlAlarmClockService is the action path for the Sonos service.
	controlAlarmClockService = "/AlarmClock/Control"
)

// AlarmClockService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type AlarmClockService struct {
	doer do.Doer
}

// NewAlarmClockService initializes a new AlarmClockService with the provided Doer.
func NewAlarmClockService(doer do.Doer) *AlarmClockService {
	return &AlarmClockService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (a AlarmClockService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return a.doer.Do(ctx, do.WithPath(eventAlarmClockService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

func (a AlarmClockService) SetFormat(ctx context.Context, desiredTimeFormat TimeFormat, desiredDateFormat DateFormat) error {
	type SetFormatReq struct {
		DesiredTimeFormat TimeFormat
		DesiredDateFormat DateFormat
	}

	type SetFormatNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetFormatReq
	}

	type SetFormatBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		SetFormatNode SetFormatNode `xml:"u:SetFormat,omitempty"`
	}

	body := SetFormatBody{SetFormatNode: SetFormatNode{
		SetFormatReq: SetFormatReq{
			DesiredDateFormat: desiredDateFormat,
			DesiredTimeFormat: desiredTimeFormat,
		},
		Xmlns: namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "SetFormat", body, nil)...)
}

type AlarmClockGetFormatRes struct {
	CurrentTimeFormat TimeFormat
	CurrentDateFormat DateFormat
}

func (a AlarmClockService) GetFormat(ctx context.Context, getFormat *AlarmClockGetFormatRes) error {
	type GetFormatReq struct{}

	type GetFormatNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetFormatReq
	}

	type GetFormatBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		GetFormatNode GetFormatNode `xml:"u:GetFormat,omitempty"`
	}

	body := GetFormatBody{GetFormatNode: GetFormatNode{
		GetFormatReq: GetFormatReq{},
		Xmlns:        namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetFormat", body, getFormat)...)
}

func (a AlarmClockService) SetTimeZone(ctx context.Context, index TimeZoneIndex, autoAdjustDst TimeZoneAutoAdjustDst) error {
	type SetTimeZoneReq struct {
		Index         TimeZoneIndex
		AutoAdjustDst TimeZoneAutoAdjustDst
	}

	type SetTimeZoneNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetTimeZoneReq
	}

	type SetTimeZoneBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		SetTimeZoneNode SetTimeZoneNode `xml:"u:SetTimeZone,omitempty"`
	}

	body := SetTimeZoneBody{SetTimeZoneNode: SetTimeZoneNode{
		SetTimeZoneReq: SetTimeZoneReq{
			AutoAdjustDst: autoAdjustDst,
			Index:         index,
		},
		Xmlns: namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "SetTimeZone", body, nil)...)
}

type AlarmClockGetTimeZoneRes struct {
	Index         TimeZoneIndex
	AutoAdjustDst TimeZoneAutoAdjustDst
}

func (a AlarmClockService) GetTimeZone(ctx context.Context, getTimeZone *AlarmClockGetTimeZoneRes) error {
	type GetTimeZoneReq struct{}

	type GetTimeZoneNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTimeZoneReq
	}

	type GetTimeZoneBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		GetTimeZoneNode GetTimeZoneNode `xml:"u:GetTimeZone,omitempty"`
	}

	body := GetTimeZoneBody{GetTimeZoneNode: GetTimeZoneNode{
		GetTimeZoneReq: GetTimeZoneReq{},
		Xmlns:          namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetTimeZone", body, getTimeZone)...)
}

type AlarmClockGetTimeZoneAndRuleRes struct {
	Index           TimeZoneIndex
	AutoAdjustDst   TimeZoneAutoAdjustDst
	CurrentTimeZone TimeZone
}

func (a AlarmClockService) GetTimeZoneAndRule(ctx context.Context, getTimeZoneAndRule *AlarmClockGetTimeZoneAndRuleRes) error {
	type GetTimeZoneAndRuleReq struct{}

	type GetTimeZoneAndRuleNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTimeZoneAndRuleReq
	}

	type GetTimeZoneAndRuleBody struct {
		XMLName                xml.Name               `xml:"s:Body"`
		GetTimeZoneAndRuleNode GetTimeZoneAndRuleNode `xml:"u:GetTimeZoneAndRule,omitempty"`
	}

	body := GetTimeZoneAndRuleBody{GetTimeZoneAndRuleNode: GetTimeZoneAndRuleNode{
		GetTimeZoneAndRuleReq: GetTimeZoneAndRuleReq{},
		Xmlns:                 namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetTimeZoneAndRule", body, getTimeZoneAndRule)...)
}

type AlarmClockGetTimeZoneRuleRes struct {
	TimeZone TimeZone
}

func (a AlarmClockService) GetTimeZoneRule(ctx context.Context, index TimeZoneIndex, getTimeZoneRule *AlarmClockGetTimeZoneRuleRes) error {
	type GetTimeZoneRuleReq struct {
		Index TimeZoneIndex
	}

	type GetTimeZoneRuleNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTimeZoneRuleReq
	}

	type GetTimeZoneRuleBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		GetTimeZoneRuleNode GetTimeZoneRuleNode `xml:"u:GetTimeZoneRule,omitempty"`
	}

	body := GetTimeZoneRuleBody{GetTimeZoneRuleNode: GetTimeZoneRuleNode{
		GetTimeZoneRuleReq: GetTimeZoneRuleReq{Index: index},
		Xmlns:              namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetTimeZoneRule", body, getTimeZoneRule)...)
}

func (a AlarmClockService) SetTimeServer(ctx context.Context, desiredTimeServer TimeServer) error {
	type SetTimeServerReq struct {
		DesiredTimeServer TimeServer
	}

	type SetTimeServerNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetTimeServerReq
	}

	type SetTimeServerBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		SetTimeServerNode SetTimeServerNode `xml:"u:SetTimeServer,omitempty"`
	}

	body := SetTimeServerBody{SetTimeServerNode: SetTimeServerNode{
		SetTimeServerReq: SetTimeServerReq{DesiredTimeServer: desiredTimeServer},
		Xmlns:            namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "SetTimeServer", body, nil)...)
}

type AlarmClockGetTimeServerRes struct {
	CurrentTimeServer TimeServer
}

func (a AlarmClockService) GetTimeServer(ctx context.Context, getTimeServer *AlarmClockGetTimeServerRes) error {
	type GetTimeServerReq struct{}

	type GetTimeServerNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTimeServerReq
	}

	type GetTimeServerBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		GetTimeServerNode GetTimeServerNode `xml:"u:GetTimeServer,omitempty"`
	}

	body := GetTimeServerBody{GetTimeServerNode: GetTimeServerNode{
		GetTimeServerReq: GetTimeServerReq{},
		Xmlns:            namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetTimeServer", body, getTimeServer)...)
}

func (a AlarmClockService) SetTimeNow(ctx context.Context, desiredTime ISO8601Time, timeZoneForDesiredTime TimeZoneInformation) error {
	type SetTimeNowReq struct {
		DesiredTime            ISO8601Time
		TimeZoneForDesiredTime TimeZoneInformation
	}

	type SetTimeNowNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetTimeNowReq
	}

	type SetTimeNowBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		SetTimeNowNode SetTimeNowNode `xml:"u:SetTimeNow,omitempty"`
	}

	body := SetTimeNowBody{SetTimeNowNode: SetTimeNowNode{
		SetTimeNowReq: SetTimeNowReq{
			DesiredTime:            desiredTime,
			TimeZoneForDesiredTime: timeZoneForDesiredTime,
		},
		Xmlns: namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "SetTimeNow", body, nil)...)
}

type AlarmClockGetHouseholdTimeAtStampRes struct {
	HouseholdUTCTime ISO8601Time
}

func (a AlarmClockService) GetHouseholdTimeAtStamp(ctx context.Context, timeStamp TimeStamp, getHouseholdTimeAtStamp *AlarmClockGetHouseholdTimeAtStampRes) error {
	type GetHouseholdTimeAtStampReq struct {
		TimeStamp TimeStamp
	}

	type GetHouseholdTimeAtStampNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetHouseholdTimeAtStampReq
	}

	type GetHouseholdTimeAtStampBody struct {
		XMLName                     xml.Name                    `xml:"s:Body"`
		GetHouseholdTimeAtStampNode GetHouseholdTimeAtStampNode `xml:"u:GetHouseholdTimeAtStamp,omitempty"`
	}

	body := GetHouseholdTimeAtStampBody{GetHouseholdTimeAtStampNode: GetHouseholdTimeAtStampNode{
		GetHouseholdTimeAtStampReq: GetHouseholdTimeAtStampReq{TimeStamp: timeStamp},
		Xmlns:                      namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetHouseholdTimeAtStamp", body, getHouseholdTimeAtStamp)...)
}

type AlarmClockGetTimeNowRes struct {
	CurrentUTCTime        ISO8601Time
	CurrentLocalTime      ISO8601Time
	CurrentTimeZone       TimeZone
	CurrentTimeGeneration TimeGeneration
}

func (a AlarmClockService) GetTimeNow(ctx context.Context, getTimeNow *AlarmClockGetTimeNowRes) error {
	type GetTimeNowReq struct{}

	type GetTimeNowNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetTimeNowReq
	}

	type GetTimeNowBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		GetTimeNowNode GetTimeNowNode `xml:"u:GetTimeNow,omitempty"`
	}

	body := GetTimeNowBody{GetTimeNowNode: GetTimeNowNode{
		GetTimeNowReq: GetTimeNowReq{},
		Xmlns:         namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetTimeNow", body, getTimeNow)...)
}

type AlarmClockCreateAlarmRes struct {
	AssignedID AlarmID
}

// CreateAlarm Create a single alarm, all properties are required
func (a AlarmClockService) CreateAlarm(ctx context.Context, startLocalTime ISO8601Time, duration ISO8601Time, recurrence Recurrence, enabled AlarmEnabled, roomUuid AlarmRoomUUID, programUri AlarmProgramURI, programMetaData AlarmProgramMetaData, playMode AlarmPlayMode, volume AlarmVolume, includeLinkedZones AlarmIncludeLinkedZones, createAlarm *AlarmClockCreateAlarmRes) error {
	type CreateAlarmReq struct {
		StartLocalTime     ISO8601Time
		Duration           ISO8601Time
		Recurrence         Recurrence
		Enabled            AlarmEnabled
		RoomUUID           AlarmRoomUUID
		ProgramURI         AlarmProgramURI
		ProgramMetaData    AlarmProgramMetaData
		PlayMode           AlarmPlayMode
		Volume             AlarmVolume
		IncludeLinkedZones AlarmIncludeLinkedZones
	}

	type CreateAlarmNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		CreateAlarmReq
	}

	type CreateAlarmBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		CreateAlarmNode CreateAlarmNode `xml:"u:CreateAlarm,omitempty"`
	}

	body := CreateAlarmBody{CreateAlarmNode: CreateAlarmNode{
		CreateAlarmReq: CreateAlarmReq{
			Duration:           duration,
			Enabled:            enabled,
			IncludeLinkedZones: includeLinkedZones,
			PlayMode:           playMode,
			ProgramMetaData:    programMetaData,
			ProgramURI:         programUri,
			Recurrence:         recurrence,
			RoomUUID:           roomUuid,
			StartLocalTime:     startLocalTime,
			Volume:             volume,
		},
		Xmlns: namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "CreateAlarm", body, createAlarm)...)
}

// UpdateAlarm Update an alarm, all parameters are required.
func (a AlarmClockService) UpdateAlarm(ctx context.Context, id AlarmID, startLocalTime ISO8601Time, duration ISO8601Time, recurrence Recurrence, enabled AlarmEnabled, roomUuid AlarmRoomUUID, programUri AlarmProgramURI, programMetaData AlarmProgramMetaData, playMode AlarmPlayMode, volume AlarmVolume, includeLinkedZones AlarmIncludeLinkedZones) error {
	type UpdateAlarmReq struct {
		ID                 AlarmID
		StartLocalTime     ISO8601Time
		Duration           ISO8601Time
		Recurrence         Recurrence
		Enabled            AlarmEnabled
		RoomUUID           AlarmRoomUUID
		ProgramURI         AlarmProgramURI
		ProgramMetaData    AlarmProgramMetaData
		PlayMode           AlarmPlayMode
		Volume             AlarmVolume
		IncludeLinkedZones AlarmIncludeLinkedZones
	}

	type UpdateAlarmNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		UpdateAlarmReq
	}

	type UpdateAlarmBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		UpdateAlarmNode UpdateAlarmNode `xml:"u:UpdateAlarm,omitempty"`
	}

	body := UpdateAlarmBody{UpdateAlarmNode: UpdateAlarmNode{
		UpdateAlarmReq: UpdateAlarmReq{
			Duration:           duration,
			Enabled:            enabled,
			ID:                 id,
			IncludeLinkedZones: includeLinkedZones,
			PlayMode:           playMode,
			ProgramMetaData:    programMetaData,
			ProgramURI:         programUri,
			Recurrence:         recurrence,
			RoomUUID:           roomUuid,
			StartLocalTime:     startLocalTime,
			Volume:             volume,
		},
		Xmlns: namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "UpdateAlarm", body, nil)...)
}

// DestroyAlarm Delete an alarm
func (a AlarmClockService) DestroyAlarm(ctx context.Context, id AlarmID) error {
	type DestroyAlarmReq struct {
		ID AlarmID
	}

	type DestroyAlarmNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		DestroyAlarmReq
	}

	type DestroyAlarmBody struct {
		XMLName          xml.Name         `xml:"s:Body"`
		DestroyAlarmNode DestroyAlarmNode `xml:"u:DestroyAlarm,omitempty"`
	}

	body := DestroyAlarmBody{DestroyAlarmNode: DestroyAlarmNode{
		DestroyAlarmReq: DestroyAlarmReq{ID: id},
		Xmlns:           namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "DestroyAlarm", body, nil)...)
}

type AlarmClockListAlarmsRes struct {
	CurrentAlarmList        AlarmList
	CurrentAlarmListVersion AlarmListVersion
}

// ListAlarms Get the AlarmList as XML
func (a AlarmClockService) ListAlarms(ctx context.Context, listAlarms *AlarmClockListAlarmsRes) error {
	type ListAlarmsReq struct{}

	type ListAlarmsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ListAlarmsReq
	}

	type ListAlarmsBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		ListAlarmsNode ListAlarmsNode `xml:"u:ListAlarms,omitempty"`
	}

	body := ListAlarmsBody{ListAlarmsNode: ListAlarmsNode{
		ListAlarmsReq: ListAlarmsReq{},
		Xmlns:         namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "ListAlarms", body, listAlarms)...)
}

func (a AlarmClockService) SetDailyIndexRefreshTime(ctx context.Context, desiredDailyIndexRefreshTime DailyIndexRefreshTime) error {
	type SetDailyIndexRefreshTimeReq struct {
		DesiredDailyIndexRefreshTime DailyIndexRefreshTime
	}

	type SetDailyIndexRefreshTimeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetDailyIndexRefreshTimeReq
	}

	type SetDailyIndexRefreshTimeBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		SetDailyIndexRefreshTimeNode SetDailyIndexRefreshTimeNode `xml:"u:SetDailyIndexRefreshTime,omitempty"`
	}

	body := SetDailyIndexRefreshTimeBody{SetDailyIndexRefreshTimeNode: SetDailyIndexRefreshTimeNode{
		SetDailyIndexRefreshTimeReq: SetDailyIndexRefreshTimeReq{DesiredDailyIndexRefreshTime: desiredDailyIndexRefreshTime},
		Xmlns:                       namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "SetDailyIndexRefreshTime", body, nil)...)
}

type AlarmClockGetDailyIndexRefreshTimeRes struct {
	CurrentDailyIndexRefreshTime DailyIndexRefreshTime
}

func (a AlarmClockService) GetDailyIndexRefreshTime(ctx context.Context, getDailyIndexRefreshTime *AlarmClockGetDailyIndexRefreshTimeRes) error {
	type GetDailyIndexRefreshTimeReq struct{}

	type GetDailyIndexRefreshTimeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetDailyIndexRefreshTimeReq
	}

	type GetDailyIndexRefreshTimeBody struct {
		XMLName                      xml.Name                     `xml:"s:Body"`
		GetDailyIndexRefreshTimeNode GetDailyIndexRefreshTimeNode `xml:"u:GetDailyIndexRefreshTime,omitempty"`
	}

	body := GetDailyIndexRefreshTimeBody{GetDailyIndexRefreshTimeNode: GetDailyIndexRefreshTimeNode{
		GetDailyIndexRefreshTimeReq: GetDailyIndexRefreshTimeReq{},
		Xmlns:                       namespaceAlarmClockService,
	}}

	return a.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlAlarmClockService, namespaceAlarmClockService, "GetDailyIndexRefreshTime", body, getDailyIndexRefreshTime)...)
}
