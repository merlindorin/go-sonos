// CODE GENERATED, DO NOT EDIT

// Package S18 contains the implementation for the Sonos One (Model S18) services.
package S18

// Services struct contains service clients for different functionalities of a Sonos device.
type BufferingResultCode int

type ThirdPartyMediaServersX string

type NumberOfTracks uint

type Queue string

type AVTransportID int

type URI string

type ShareIndexInProgress bool

type ConfigMode string

type AvailableSoftwareUpdate string

type StreamRestartState string

type SortOrder string

type SupportsAudioClip bool

type SourceAreasUpdateID string

type AlarmProgramURI string

type TransportErrorHttpCode string

type CrossfadeMode bool

type EnqueuedTransportURIMetaData string

type HTAudioIn uint

type SourceAreaIds string

type Username string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type TransportSettings string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type UpdateID uint

type MACAddress string

type VirtualLineInGroupID string

type AutoplayUseVolume bool

type SecureRegState uint

type EQType string

type DialogLevel string

type SpeakerSize uint

type IncludeControllers bool

type RadioFavoritesUpdateID uint

type AlarmRoomUUID string

type DateFormat string

type AbsoluteTimePosition string

type ObjectID string

type GroupID string

type QueueUpdateID uint

type UserRadioUpdateID string

type AlarmEnabled bool

type TrackURI string

type ResumePlayback bool

type TargetRoomName string

type IsZoneBridge bool

type IPAddress string

type MID string

type TransportActions string

type RejoinGroup bool

type HTForwardEnabled bool

type BootSeq uint

type QueueID uint

type LeftVolume uint

type AccountTier uint

type AccountPassword string

type TimeZoneIndex int

type AbsoluteCounterPosition int

type LIST_URIMetaData string

type HdmiCecAvailable bool

type Treble int

type ProgramURI string

type NightMode bool

type MobileDeviceUDN string

type FavoritePresetsUpdateID string

type VoiceConfigState uint

type Loudness bool

type PresetNameList string

type VariableStringValue string

type AccountNickname string

type AccountCredential string

type DirectControlIsSuspended bool

type SavedQueuesUpdateID string

type ConfigModeOptions string

type RoomDetectionPlayId uint

type Code string

type ThirdPartyHash string

type MemberList string

type QueuePolicy string

type RightVolume uint

type UserIdHashCode string

type EthLink bool

type LIST_URI_AND_METADATA string

type RoomCalibrationID string

type AuthorizationCode string

type ZonePlayerUUIDsInGroup string

type ConfigModeState string

type ISO8601Time string

type RecordQualityMode string

type ValidPlayModes string

type SleepTimerState string

type RecentlyPlayedUpdateID string

type SerialNumber string

type WirelessLeafOnly bool

type ChannelMap string

type VoiceUpdateID uint

type TimeServer string

type TimeFormat string

type HasConfiguredSSID bool

type BehindWifiExtender uint

type CachedOnly bool

type UpdateFlags uint

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type Result string

const ()

type TransportPlaySpeed string

type DirectControlAccountID string

type ConnectionIDs string

type SoftwareVersion string

type TrackNumber uint

type NumTracks uint

type Index uint

type Count uint

type GroupVolumeChangeable bool

type OAuthDeviceID string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type DirectControlClientID string

type LastIndexChange string

type ChannelMapSet string

type GroupCoordinatorIsLocal bool

type QueueOwnerID string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AlarmList string

type TransportStatus string

type LastChange string

type LIST_URI string

type SavedQueueTitle string

type RoomDetectionChirpChannel uint

type SubGain string

type SearchCapabilities string

type Icon string

type SurroundMode string

type ZoneGroupState string

type TrackMetaData string

type ResetVolumeAfter bool

type MicEnabled uint

type EQValue int

type MuseHouseholdId string

type SubPolarity string

type AccountID string

type IsExpired bool

type UpdateItem string

type AlarmIncludeLinkedZones bool

type TimeZoneAutoAdjustDst bool

type PossiblePlaybackStorageMedia string

type AVTransportURIMetaData string

type Prefix string

type SupportsAudioIn bool

type SupportsOutputFixed bool

type RestartPending bool

type ProtocolInfo string

type SurroundEnabled bool

type Speed string

type HTFreq uint

type LocalGroupUUID string

type HeadphoneConnected bool

type ConnectionManager string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type IsIdle bool

type SessionId string

type VariableName string

type RDMEnabled bool

type Origin string

type MobileDeviceName string

type RecordMediumWriteStatus string

type EnqueueAs bool

type SortCapabilities string

type ContainerUpdateIDs string

type AirPlayEnabled bool

type TVConfigurationError bool

type VolumeAdjustment int

type AreasUpdateID string

type TransportErrorDescription string

type ShareListUpdateID string

type RoomCalibrationState int

type CopyrightInfo string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type TimeZoneInformation string

type HTSatChanMapSet string

type ChannelFreq uint

type AlarmVolume uint

type SettingsReplicationState string

type ServiceListVersion string

type Mute bool

type SubCrossover string

type MediaDuration string

type TrackList string

type ServiceDescriptorList string

type OutputFixed bool

type SortCriteria string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type RoomCalibrationEnabled bool

type Version string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type Curated bool

type DiagnosticID uint

type StubsCreated string

type AlarmLoggedStartTime string

type AutoplayRoomUUID string

type Bass int

type RampTimeSeconds uint

type RoomCalibrationCoefficients string

type SnoozeRunning bool

type SinkProtocolInfo string

type GroupVolume uint

type AccountMd string

type AlarmState string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type TagValueList string

type ShareIndexLastError string

type ServiceId uint

type TrackNumbersCSV string

type AlarmID uint

type TimeStamp string

type MoreInfo string

type AutoplaySource string

type WirelessMode uint

type EnqueuedTransportURI string

type SystemUpdateID uint

type AvailableRoomCalibration string

type HardwareVersion string

type ZoneGroupID string

type TrackDuration string

type SeekTarget string

type MemberID string

type Configuration string

type VolumeAVTransportURI string

type RoomCalibrationCalibrationMode string

type TransportErrorURI string

type RelativeTimePosition string

type Invisible bool

type AutoplayIncludeLinkedZones bool

type ButtonState string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SurroundLevel string

type TransportErrorHttpHeaders string

type Track uint

type ExtraInfo string

type ServiceTypeList string

type Volume uint

type RedirectURI string

type PossibleRecordStorageMedia string

type AlarmRunning bool

type SourceProtocolInfo string

type SatRoomUUID string

type WifiEnabled bool

type VolumeDB int

type AudioDelayRightRear string

type UpdateIDX uint

type AlarmProgramMetaData string

type AlarmListVersion string

type RcsID int

type LastChangedPlayState string

type AutoplayVolume uint

type QueueOwnerContext string

type PossibleRecordQualityModes string

type InstanceID uint

type RoomDetectionDurationMilliseconds uint

type Seed string

type NetsettingsUpdateID string

type AVTransportURI string

type MuseSessions string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type SourceState string

type NumTracksChange int

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type UpdateURL string

type TimeGeneration uint

type Section uint

type Filter string

type Browseable bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type SubEnabled bool

type TimeZone string

type AlarmIDRunning uint

type ConnectionID int

type FavoritesUpdateID string

type ZoneName string

type Orientation int

type DisplaySoftwareVersion string

type MusicSurroundLevel string

type AccountUDN string

type SleepTimerGeneration uint

type PlayerID string

type HouseholdID string

type CustomerID string

type AlarmRunSequence string

type ZoneGroupName string

type RelativeCounterPosition int

type URIMetaData string

type KeepGrouped bool

type GroupMute bool

type AccountUID uint

type UpdateExtraOptions string

const ()

type RecordStorageMedium string

type SearchCriteria string

type AudioDelay string

type RoomCalibrationAvailable bool

type AccountType uint

type AlbumArtistDisplayOption string

type HTBondedZoneCommitState uint

type Flags uint

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type MobileIPAndPort string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type DailyIndexRefreshTime string

type DID string

type AudioDelayLeftRear string

type VLIState string

type RadioLocationUpdateID uint
