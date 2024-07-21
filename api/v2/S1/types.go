// CODE GENERATED, DO NOT EDIT

// Package S1 contains the implementation for the Sonos Play:1 (Model S1) services.
package S1

// Services struct contains service clients for different functionalities of a Sonos device.
type SettingsReplicationState string

type AccountUID uint

type SearchCriteria string

type SeekTarget string

type IPAddress string

type HasConfiguredSSID bool

type QueuePolicy string

type AlarmRunSequence string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type IsIdle bool

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

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

type MemberList string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type WifiEnabled bool

type MID string

type AlarmRoomUUID string

type SearchCapabilities string

type CopyrightInfo string

type KeepGrouped bool

type VolumeAdjustment int

type AlarmRunning bool

type Flags uint

type RoomDetectionChirpIfPlayingSwappableAudio bool

type SubGain string

type RoomCalibrationEnabled bool

type AccountID string

type VoiceUpdateID uint

type TrackDuration string

type Configuration string

type ZoneGroupState string

type UpdateURL string

type SourceAreasUpdateID string

type AVTransportURI string

type AutoplayVolume uint

type QueueOwnerID string

type SurroundLevel string

type NumTracks uint

type Code string

type Bass int

type NightMode bool

type AreasUpdateID string

type SystemUpdateID uint

type LastChange string

type LIST_URIMetaData string

type PlayerID string

type PresetNameList string

type AccountType uint

type AlarmIDRunning uint

type SinkProtocolInfo string

type Orientation int

type SourceState string

type ConfigModeOptions string

type MuseHouseholdId string

type HardwareVersion string

type RejoinGroup bool

type TVConfigurationError bool

type OAuthDeviceID string

type Section uint

type QueueOwnerContext string

type RightVolume uint

type Treble int

type RoomCalibrationID string

type UpdateExtraOptions string

type TrackURI string

type TimeServer string

type LastIndexChange string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type PossibleRecordQualityModes string

type ConnectionID int

type SortCriteria string

type ServiceId uint

type TrackNumbersCSV string

type AuthorizationCode string

type CrossfadeMode bool

type SleepTimerState string

type SortCapabilities string

type Icon string

type LastChangedPlayState string

type EQType string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type GroupMute bool

type ZoneGroupName string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type DirectControlClientID string

type BufferingResultCode int

type ServiceListVersion string

type FavoritePresetsUpdateID string

type GroupID string

type SatRoomUUID string

type DisplaySoftwareVersion string

type ExtraInfo string

type TimeStamp string

type NumTracksChange int

type Prefix string

type FavoritesUpdateID string

type ServiceDescriptorList string

type SpeakerSize uint

type RoomCalibrationAvailable bool

type SleepTimerGeneration uint

type PossiblePlaybackStorageMedia string

type SoftwareVersion string

type AudioDelayLeftRear string

type TimeZone string

type ConfigMode string

type Volume uint

type ThirdPartyMediaServersX string

type IncludeControllers bool

type WirelessMode uint

type TrackMetaData string

type RelativeCounterPosition int

type AVTransportID int

type ContainerUpdateIDs string

type AirPlayEnabled bool

type RoomCalibrationState int

type GroupCoordinatorIsLocal bool

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type AlarmListVersion string

type NumberOfTracks uint

type SnoozeRunning bool

type DirectControlAccountID string

type AlarmState string

type AlarmProgramURI string

type ValidPlayModes string

type Result string

type EnqueuedTransportURIMetaData string

type HTBondedZoneCommitState uint

type HTAudioIn uint

type SurroundMode string

type RDMEnabled bool

type RecentlyPlayedUpdateID string

type AvailableRoomCalibration string

type AccountNickname string

type UpdateIDX uint

type SupportsAudioIn bool

type TransportErrorHttpHeaders string

type RadioLocationUpdateID uint

type SurroundEnabled bool

type TransportErrorURI string

type URI string

type ResumePlayback bool

type SortOrder string

type ShareListUpdateID string

type QueueID uint

type RestartPending bool

type WirelessLeafOnly bool

type Index uint

type TransportActions string

type MuseSessions string

type Count uint

type IsZoneBridge bool

type AutoplayRoomUUID string

type SecureRegState uint

type VoiceConfigState uint

type AbsoluteCounterPosition int

type LocalGroupUUID string

type HdmiCecAvailable bool

type BootSeq uint

type LeftVolume uint

type RampTimeSeconds uint

type AudioDelayRightRear string

type Version string

type RecordQualityMode string

type UpdateFlags uint

type TransportSettings string

type ChannelFreq uint

type QueueUpdateID uint

type AlbumArtistDisplayOption string

type BehindWifiExtender uint

type SubPolarity string

type RoomCalibrationCalibrationMode string

type Speed string

type PossibleRecordStorageMedia string

const ()

type TransportPlaySpeed string

type SavedQueueTitle string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ShareIndexInProgress bool

type DialogLevel string

type AvailableSoftwareUpdate string

type MobileIPAndPort string

type TimeGeneration uint

type AlarmProgramMetaData string

type DateFormat string

type HouseholdID string

type Invisible bool

type HTSatChanMapSet string

type ISO8601Time string

type AlarmEnabled bool

type DirectControlIsSuspended bool

type ChannelMapSet string

type HTFreq uint

type RoomDetectionPlayId uint

type GroupVolumeChangeable bool

type ServiceTypeList string

type AlarmID uint

type Origin string

type RadioFavoritesUpdateID uint

type ThirdPartyHash string

type ZonePlayerUUIDsInGroup string

type TimeZoneAutoAdjustDst bool

type EnqueuedTransportURI string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type MACAddress string

type RoomDetectionDurationMilliseconds uint

type Username string

type AVTransportURIMetaData string

type RoomDetectionChirpChannel uint

type DID string

type StubsCreated string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type SessionId string

type AccountPassword string

type MobileDeviceUDN string

type MoreInfo string

type Queue string

type StreamRestartState string

type Browseable bool

type EthLink bool

type AccountUDN string

type NetsettingsUpdateID string

type TimeZoneIndex int

type RoomCalibrationCoefficients string

type VariableName string

type VariableStringValue string

type RedirectURI string

const ()

type RecordStorageMedium string

type AbsoluteTimePosition string

type VLIState string

type TrackNumber uint

type HTForwardEnabled bool

type ZoneGroupID string

type Track uint

type ConnectionIDs string

type ConnectionManager string

type UserRadioUpdateID string

type ButtonState string

type VolumeAVTransportURI string

type ObjectID string

type TimeZoneInformation string

type TransportErrorDescription string

type AutoplayUseVolume bool

type IsExpired bool

type AlarmList string

type MusicSurroundLevel string

type DiagnosticID uint

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type DailyIndexRefreshTime string

type AutoplayIncludeLinkedZones bool

type AlarmVolume uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type TrackList string

type RcsID int

type UpdateID uint

type AccountMd string

type UpdateItem string

type MobileDeviceName string

type AlarmLoggedStartTime string

type ConfigModeState string

type ChannelMap string

type InstanceID uint

type HeadphoneConnected bool

type OutputFixed bool

type TagValueList string

type GroupVolume uint

type Seed string

type Curated bool

type SupportsOutputFixed bool

type AudioDelay string

type EnqueueAs bool

type ShareIndexLastError string

type ProgramURI string

type AccountTier uint

type TransportStatus string

type MediaDuration string

type SourceProtocolInfo string

type ProtocolInfo string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type SerialNumber string

type Mute bool

type TimeFormat string

type SubCrossover string

type Loudness bool

type VirtualLineInGroupID string

type VolumeDB int

type UserIdHashCode string

type URIMetaData string

type Filter string

type ZoneName string

type TargetRoomName string

type EQValue int

type RelativeTimePosition string

type SupportsAudioClip bool

type AutoplaySource string

type TransportErrorHttpCode string

type ResetVolumeAfter bool

type SavedQueuesUpdateID string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type AccountCredential string

type MemberID string

type MicEnabled uint

type LIST_URI_AND_METADATA string

type SubEnabled bool

type RecordMediumWriteStatus string

type LIST_URI string

type SourceAreaIds string

type CustomerID string

type AlarmIncludeLinkedZones bool
