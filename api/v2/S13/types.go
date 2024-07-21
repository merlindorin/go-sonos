// CODE GENERATED, DO NOT EDIT

// Package S13 contains the implementation for the Sonos One (Model S13) services.
package S13

// Services struct contains service clients for different functionalities of a Sonos device.
type RecordQualityMode string

type AbsoluteCounterPosition int

type SavedQueueTitle string

type Prefix string

type AutoplayRoomUUID string

type QueuePolicy string

type ZoneGroupState string

type TimeZoneIndex int

type MobileIPAndPort string

type AlarmRunning bool

type URIMetaData string

type ConnectionIDs string

type TrackNumbersCSV string

type TimeFormat string

type SubPolarity string

type RDMEnabled bool

type RoomCalibrationState int

type Result string

type MicEnabled uint

type BufferingResultCode int

type Bass int

type DialogLevel string

type RejoinGroup bool

type ConnectionID int

type ShareIndexLastError string

type TrackDuration string

type AVTransportURI string

type InstanceID uint

type VariableStringValue string

type TimeServer string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type ServiceId uint

type VLIState string

type MemberID string

type Count uint

type HTAudioIn uint

type AccountCredential string

type UpdateIDX uint

type TransportSettings string

type AccountUDN string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type ChannelFreq uint

type SearchCriteria string

type HTFreq uint

type AlarmProgramMetaData string

type HardwareVersion string

type GroupVolume uint

type QueueOwnerID string

type Volume uint

type TargetRoomName string

type ZoneName string

type VolumeDB int

type NetsettingsUpdateID string

type EnqueuedTransportURI string

type TVConfigurationError bool

type RoomDetectionDurationMilliseconds uint

type OutputFixed bool

type PlayerID string

type TrackList string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type CachedOnly bool

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type SystemUpdateID uint

type SerialNumber string

type BootSeq uint

type AudioDelay string

type NightMode bool

type UpdateItem string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type ThirdPartyMediaServersX string

type SourceState string

type Orientation int

type ConfigModeState string

type RoomDetectionChirpChannel uint

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type SubCrossover string

type Track uint

type TrackURI string

type IsIdle bool

type TransportErrorHttpCode string

type TrackMetaData string

type LIST_URIMetaData string

type SourceProtocolInfo string

type SoftwareVersion string

type RedirectURI string

type TimeZoneInformation string

type DirectControlIsSuspended bool

type HouseholdID string

type RoomCalibrationEnabled bool

type MuseHouseholdId string

type ISO8601Time string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type UpdateID uint

type DisplaySoftwareVersion string

type IPAddress string

type ConfigMode string

type AuthorizationCode string

type AlarmListVersion string

type PossiblePlaybackStorageMedia string

type PossibleRecordQualityModes string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type RoomDetectionPlayId uint

type TimeGeneration uint

type EnqueueAs bool

type RcsID int

type Configuration string

type SupportsAudioIn bool

type ButtonState string

type Treble int

type TrackNumber uint

type ValidPlayModes string

type NumTracks uint

type RadioFavoritesUpdateID uint

type AlarmLoggedStartTime string

type QueueUpdateID uint

type HTBondedZoneCommitState uint

type VolumeAVTransportURI string

type AlarmRoomUUID string

type CopyrightInfo string

type WirelessLeafOnly bool

type GroupMute bool

type EQValue int

type MusicSurroundLevel string

type RoomCalibrationAvailable bool

type ZoneGroupID string

type AlarmID uint

type Index uint

type LastIndexChange string

type SessionId string

type DirectControlAccountID string

type AVTransportID int

type Loudness bool

type ProgramURI string

type RelativeCounterPosition int

type SortCriteria string

type SortCapabilities string

type AutoplayVolume uint

type VirtualLineInGroupID string

type RoomCalibrationID string

type TimeStamp string

type RightVolume uint

type DateFormat string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type SurroundLevel string

type ResetVolumeAfter bool

type Section uint

type WirelessMode uint

const ()

type RecordStorageMedium string

type EnqueuedTransportURIMetaData string

type DirectControlClientID string

type AlarmState string

type AlbumArtistDisplayOption string

type ExtraInfo string

type CustomerID string

type Origin string

type AlarmIDRunning uint

type SettingsReplicationState string

type ConfigModeOptions string

type Code string

type IsExpired bool

type UserIdHashCode string

type RestartPending bool

type AbsoluteTimePosition string

type SnoozeRunning bool

type ConnectionManager string

type SavedQueuesUpdateID string

type MoreInfo string

type MACAddress string

type ServiceDescriptorList string

type TimeZoneAutoAdjustDst bool

type AutoplayIncludeLinkedZones bool

type SurroundEnabled bool

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type DiagnosticID uint

type AutoplaySource string

type SleepTimerState string

type Browseable bool

type WifiEnabled bool

type AlarmProgramURI string

const ()

type TransportPlaySpeed string

type MuseSessions string

type Queue string

type StreamRestartState string

type AccountUID uint

type UpdateFlags uint

type AlarmIncludeLinkedZones bool

type HasConfiguredSSID bool

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type RoomCalibrationCalibrationMode string

type AccountType uint

type ChannelMapSet string

type MediaDuration string

type SleepTimerGeneration uint

type GroupID string

type SinkProtocolInfo string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type HdmiCecAvailable bool

type BehindWifiExtender uint

type TimeZone string

type AccountID string

type PresetNameList string

type NumTracksChange int

type KeepGrouped bool

type GroupVolumeChangeable bool

type ServiceTypeList string

type LastChange string

type ShareIndexInProgress bool

type Mute bool

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SearchCapabilities string

type AvailableRoomCalibration string

type LeftVolume uint

type RampTimeSeconds uint

type StubsCreated string

type ThirdPartyHash string

type SortOrder string

type FavoritePresetsUpdateID string

type MemberList string

type TransportErrorURI string

type TransportErrorHttpHeaders string

type Filter string

type VolumeAdjustment int

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type LocalGroupUUID string

type AccountMd string

type TagValueList string

type NumberOfTracks uint

type ProtocolInfo string

type LastChangedPlayState string

type PossibleRecordStorageMedia string

type Flags uint

type AutoplayUseVolume bool

type QueueOwnerContext string

type AccountPassword string

type SourceAreasUpdateID string

type CrossfadeMode bool

type SecureRegState uint

type OAuthDeviceID string

type LIST_URI string

type SourceAreaIds string

type Seed string

type MID string

type UpdateExtraOptions string

type SeekTarget string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type Username string

type Speed string

type MobileDeviceName string

type URI string

type DailyIndexRefreshTime string

type SubEnabled bool

type AlarmList string

type AlarmVolume uint

type ShareListUpdateID string

type SupportsAudioClip bool

type DID string

type ZonePlayerUUIDsInGroup string

type AlarmEnabled bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type VoiceUpdateID uint

type RecentlyPlayedUpdateID string

type AccountTier uint

type UserRadioUpdateID string

type GroupCoordinatorIsLocal bool

type EQType string

type AudioDelayLeftRear string

type AccountNickname string

type SatRoomUUID string

type Icon string

type Invisible bool

type VoiceConfigState uint

type Curated bool

type HeadphoneConnected bool

type ChannelMap string

type SpeakerSize uint

type RelativeTimePosition string

type Version string

type SurroundMode string

type UpdateURL string

type ResumePlayback bool

type RadioLocationUpdateID uint

type LIST_URI_AND_METADATA string

type SubGain string

type RoomCalibrationCoefficients string

type VariableName string

type AlarmRunSequence string

type ZoneGroupName string

type AVTransportURIMetaData string

type AreasUpdateID string

type TransportErrorDescription string

type ContainerUpdateIDs string

type QueueID uint

type IncludeControllers bool

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type ObjectID string

type IsZoneBridge bool

type HTSatChanMapSet string

type MobileDeviceUDN string

type RecordMediumWriteStatus string

type TransportActions string

type FavoritesUpdateID string

type AirPlayEnabled bool

type ServiceListVersion string

type SupportsOutputFixed bool

type AudioDelayRightRear string

type AvailableSoftwareUpdate string

type TransportStatus string
