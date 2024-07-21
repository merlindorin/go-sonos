// CODE GENERATED, DO NOT EDIT

// Package S13 contains the implementation for the Sonos One (Model S13) services.
package S13

// Services struct contains service clients for different functionalities of a Sonos device.
type Result string

type QueueID uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type TimeStamp string

type AlarmRunning bool

type AlarmLoggedStartTime string

type MemberID string

type ResumePlayback bool

type MobileDeviceName string

type NightMode bool

type IncludeControllers bool

type Section uint

type ValidPlayModes string

type SortCapabilities string

type RecentlyPlayedUpdateID string

type RadioFavoritesUpdateID uint

type AlarmProgramMetaData string

type VLIState string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type AlarmRunSequence string

type UpdateURL string

type SourceProtocolInfo string

type Index uint

type LIST_URI_AND_METADATA string

type ProgramURI string

type AuthorizationCode string

type ZonePlayerUUIDsInGroup string

type AlarmListVersion string

type TransportActions string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type ContainerUpdateIDs string

type CopyrightInfo string

type Filter string

type Prefix string

type SystemUpdateID uint

type ISO8601Time string

type CrossfadeMode bool

type ResetVolumeAfter bool

type QueueUpdateID uint

type ConnectionManager string

type DisplaySoftwareVersion string

type ThirdPartyHash string

type ProtocolInfo string

type AVTransportID int

type LastIndexChange string

type KeepGrouped bool

type ServiceTypeList string

type AlarmState string

type EQValue int

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type BehindWifiExtender uint

type CustomerID string

type Configuration string

type SurroundMode string

type AvailableRoomCalibration string

type VariableName string

type NumberOfTracks uint

type ChannelFreq uint

type Seed string

type AccountCredential string

type NetsettingsUpdateID string

type EQType string

type TrackList string

type SortCriteria string

type HouseholdID string

type AutoplayRoomUUID string

type AutoplayVolume uint

type NumTracksChange int

type LastChange string

type DirectControlClientID string

type FavoritesUpdateID string

type ZoneGroupState string

type UpdateID uint

type VariableStringValue string

type IsExpired bool

type UserIdHashCode string

type SoftwareVersion string

type Bass int

type EnqueuedTransportURIMetaData string

type AirPlayEnabled bool

type ServiceDescriptorList string

type VoiceUpdateID uint

type AlarmIDRunning uint

type HTSatChanMapSet string

type RoomDetectionChirpChannel uint

type Volume uint

type OAuthDeviceID string

type RoomCalibrationCoefficients string

type Speed string

type RDMEnabled bool

type UpdateFlags uint

type AlarmID uint

type ConnectionID int

type AutoplaySource string

type HasConfiguredSSID bool

type Code string

type TimeGeneration uint

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type MuseSessions string

type NumTracks uint

type IsIdle bool

type ChannelMap string

type ConfigModeState string

type ButtonState string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type AlarmProgramURI string

type LIST_URI string

type FavoritePresetsUpdateID string

type TVConfigurationError bool

type SecureRegState uint

type RightVolume uint

type MuseHouseholdId string

type AVTransportURI string

type RejoinGroup bool

type LastChangedPlayState string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type SourceAreasUpdateID string

type SearchCriteria string

type GroupVolumeChangeable bool

type CachedOnly bool

type AbsoluteTimePosition string

type RestartPending bool

type URI string

type TrackNumbersCSV string

type GroupMute bool

type Mute bool

type Loudness bool

type StreamRestartState string

type Count uint

type MoreInfo string

type RoomCalibrationState int

type SerialNumber string

type Origin string

type AutoplayIncludeLinkedZones bool

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type TransportSettings string

type TrackNumber uint

type IsZoneBridge bool

type SnoozeRunning bool

type IPAddress string

type BootSeq uint

type QueueOwnerContext string

type PresetNameList string

type TransportErrorHttpHeaders string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type RoomCalibrationCalibrationMode string

type AlarmList string

type MemberList string

type RcsID int

type Flags uint

type ConfigMode string

type PlayerID string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type Username string

type LeftVolume uint

type UpdateExtraOptions string

type TrackDuration string

type LIST_URIMetaData string

type HTBondedZoneCommitState uint

type SourceAreaIds string

type GroupVolume uint

type AlarmEnabled bool

type SeekTarget string

type MACAddress string

type TransportErrorURI string

type UserRadioUpdateID string

type HdmiCecAvailable bool

type SubPolarity string

type AccountType uint

type StubsCreated string

type ZoneGroupName string

type TransportErrorHttpCode string

type SinkProtocolInfo string

type VolumeAVTransportURI string

type SupportsOutputFixed bool

type AudioDelayLeftRear string

type SortOrder string

type HTFreq uint

type HardwareVersion string

type SubEnabled bool

type DateFormat string

type ShareIndexInProgress bool

type ShareIndexLastError string

type SupportsAudioClip bool

type SubCrossover string

type MID string

type ThirdPartyMediaServersX string

type Version string

type MobileIPAndPort string

type SurroundLevel string

type AccountUDN string

type AccountID string

type AlarmVolume uint

type TrackURI string

type ConfigModeOptions string

type ServiceId uint

type SessionId string

type SupportsAudioIn bool

type QueuePolicy string

type TimeFormat string

type AVTransportURIMetaData string

type RelativeCounterPosition int

type DirectControlAccountID string

type SettingsReplicationState string

type PossiblePlaybackStorageMedia string

type EnqueueAs bool

type LocalGroupUUID string

type AudioDelayRightRear string

type AccountMd string

type OutputFixed bool

type GroupCoordinatorIsLocal bool

type DialogLevel string

type MusicSurroundLevel string

type TimeZoneInformation string

type TimeServer string

type DirectControlIsSuspended bool

type AutoplayUseVolume bool

type MicEnabled uint

type RedirectURI string

type MobileDeviceUDN string

type RecordQualityMode string

type RoomDetectionDurationMilliseconds uint

type SubGain string

type Invisible bool

type SatRoomUUID string

type DiagnosticID uint

type TransportErrorDescription string

type Track uint

type TagValueList string

type Browseable bool

type TargetRoomName string

type UpdateItem string

type AlbumArtistDisplayOption string

type SavedQueuesUpdateID string

type Orientation int

type HTAudioIn uint

type SpeakerSize uint

type TimeZoneIndex int

type AccountNickname string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type VoiceConfigState uint

type TimeZoneAutoAdjustDst bool

type MediaDuration string

type AvailableSoftwareUpdate string

type AlarmIncludeLinkedZones bool

type VirtualLineInGroupID string

type WifiEnabled bool

type ServiceListVersion string

type Treble int

type SleepTimerGeneration uint

type EnqueuedTransportURI string

type SleepTimerState string

type ChannelMapSet string

type ExtraInfo string

type SurroundEnabled bool

type ConnectionIDs string

type RadioLocationUpdateID uint

type RoomCalibrationID string

type RampTimeSeconds uint

type AccountUID uint

type UpdateIDX uint

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type GroupID string

type WirelessMode uint

type WirelessLeafOnly bool

type QueueOwnerID string

type DailyIndexRefreshTime string

type TrackMetaData string

type RoomCalibrationAvailable bool

type DID string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AlarmRoomUUID string

type AbsoluteCounterPosition int

type SourceState string

type Queue string

type ZoneName string

type VolumeAdjustment int

type RecordMediumWriteStatus string

type ObjectID string

type SearchCapabilities string

type ShareListUpdateID string

type Icon string

type RoomDetectionPlayId uint

type BufferingResultCode int

type AudioDelay string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TimeZone string

type TransportStatus string

type RelativeTimePosition string

type InstanceID uint

type AreasUpdateID string

type PossibleRecordStorageMedia string

type Curated bool

type AccountTier uint

type URIMetaData string

type SavedQueueTitle string

type VolumeDB int

type HeadphoneConnected bool

type AccountPassword string

const ()

type TransportPlaySpeed string

type PossibleRecordQualityModes string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

const ()

type RecordStorageMedium string

type RoomCalibrationEnabled bool

type ZoneGroupID string
