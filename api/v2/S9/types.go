// CODE GENERATED, DO NOT EDIT

// Package S9 contains the implementation for the Sonos Playbar (Model S9) services.
package S9

// Services struct contains service clients for different functionalities of a Sonos device.
type TransportErrorURI string

type LIST_URI string

type IsZoneBridge bool

type LocalGroupUUID string

type GroupVolume uint

type SubPolarity string

type TransportErrorHttpHeaders string

type AutoplaySource string

type VolumeDB int

type DialogLevel string

type IsExpired bool

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type DailyIndexRefreshTime string

type TimeFormat string

type VLIState string

type RedirectURI string

type TimeZoneInformation string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type CrossfadeMode bool

type Orientation int

type AccountType uint

type ZoneGroupName string

type DirectControlIsSuspended bool

type HTBondedZoneCommitState uint

type RoomDetectionDurationMilliseconds uint

type VariableName string

type AlarmRoomUUID string

type AlarmState string

type Icon string

type ChannelMap string

type AudioDelay string

type ThirdPartyHash string

type HdmiCecAvailable bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type ServiceListVersion string

type RoomCalibrationCalibrationMode string

type MobileDeviceUDN string

type RoomDetectionPlayId uint

type TimeZoneAutoAdjustDst bool

type ResetVolumeAfter bool

type TagValueList string

type SatRoomUUID string

type AutoplayUseVolume bool

type BehindWifiExtender uint

type TransportErrorDescription string

type MemberID string

type ChannelFreq uint

type UpdateExtraOptions string

type Section uint

type ConnectionIDs string

type HTFreq uint

type GroupCoordinatorIsLocal bool

const (
	OnIRRepeaterState       IRRepeaterState = "On"
	OffIRRepeaterState      IRRepeaterState = "Off"
	DisabledIRRepeaterState IRRepeaterState = "Disabled"
)

type IRRepeaterState string

type SurroundLevel string

type Browseable bool

type AutoplayVolume uint

type AudioDelayRightRear string

type UpdateURL string

type SleepTimerGeneration uint

type FavoritePresetsUpdateID string

type WifiEnabled bool

type Bass int

type TransportStatus string

type VolumeAdjustment int

type AvailableSoftwareUpdate string

type HasConfiguredSSID bool

type SubCrossover string

type SubEnabled bool

type PresetNameList string

type ProtocolInfo string

type VoiceConfigState uint

type Invisible bool

type AccountTier uint

type VoiceUpdateID uint

type Version string

type MuseHouseholdId string

type MobileDeviceName string

type TransportErrorHttpCode string

type TrackDuration string

type EnqueueAs bool

type SavedQueueTitle string

type FavoritesUpdateID string

type NightMode bool

type InstanceID uint

type QueueUpdateID uint

type TrackList string

type SavedQueuesUpdateID string

type ProgramURI string

type Origin string

type RoomCalibrationID string

type ZoneGroupID string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type RadioLocationUpdateID uint

type LIST_URI_AND_METADATA string

type AbsoluteCounterPosition int

type AVTransportID int

type RoomCalibrationAvailable bool

type UpdateFlags uint

type AlarmID uint

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type EnqueuedTransportURIMetaData string

type Flags uint

type AccountUID uint

type AlarmList string

type LastIndexChange string

type TVConfigurationError bool

type GroupMute bool

type TOSLinkConnected bool

type AccountID string

type Queue string

type IsIdle bool

type HTAudioIn uint

type RightVolume uint

type AccountPassword string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type LastChange string

type SortOrder string

type RemoteConfigured bool

type DID string

type MusicSurroundLevel string

type RecordMediumWriteStatus string

type UserRadioUpdateID string

type QueueOwnerID string

type UpdateIDX uint

type Loudness bool

type TrackMetaData string

type ShareIndexLastError string

type MACAddress string

type ConfigMode string

type AudioDelayLeftRear string

type HardwareVersion string

type IPAddress string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type PlayerID string

type RcsID int

type ShareIndexInProgress bool

type SupportsAudioIn bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type AutoplayIncludeLinkedZones bool

type ConfigModeOptions string

type UpdateItem string

type TimeZone string

type MuseSessions string

type MoreInfo string

type StubsCreated string

type ZonePlayerUUIDsInGroup string

type TrackURI string

type DirectControlAccountID string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type AlbumArtistDisplayOption string

const ()

type TransportPlaySpeed string

type AlarmLoggedStartTime string

type HTSatChanMapSet string

type AccountCredential string

type AreasUpdateID string

type SnoozeRunning bool

type SourceProtocolInfo string

type KeepGrouped bool

type ZoneGroupState string

type ValidPlayModes string

type SleepTimerState string

type IRCode string

type LeftVolume uint

type MemberList string

type RejoinGroup bool

type ThirdPartyMediaServersX string

type AutoplayRoomUUID string

type GroupVolumeChangeable bool

type TimeStamp string

type EnqueuedTransportURI string

type NumTracksChange int

type SortCriteria string

type SettingsReplicationState string

type AvailableRoomCalibration string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AirPlayEnabled bool

type RDMEnabled bool

type SystemUpdateID uint

type ShareListUpdateID string

type TargetRoomName string

type VirtualLineInGroupID string

type SpeakerSize uint

type VariableStringValue string

type QueuePolicy string

type LIST_URIMetaData string

type ConnectionManager string

type Index uint

type CopyrightInfo string

type SourceAreaIds string

type MID string

type TransportSettings string

type EthLink bool

type AlarmIncludeLinkedZones bool

type ResumePlayback bool

type Timeout uint

const (
	OnLEDFeedbackState  LEDFeedbackState = "On"
	OffLEDFeedbackState LEDFeedbackState = "Off"
)

type LEDFeedbackState string

type SeekTarget string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ExtraInfo string

type SupportsOutputFixed bool

type AccountNickname string

type SecureRegState uint

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type AlarmListVersion string

type DateFormat string

type PossiblePlaybackStorageMedia string

type NumberOfTracks uint

type TrackNumber uint

type ConnectionID int

type AlarmRunning bool

type ObjectID string

const ()

type RecordStorageMedium string

type Track uint

type Configuration string

type SerialNumber string

type TimeZoneIndex int

type PossibleRecordStorageMedia string

type RelativeTimePosition string

type SinkProtocolInfo string

type RadioFavoritesUpdateID uint

type ZoneName string

type IRRemoteName string

type Seed string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type BootSeq uint

type MobileIPAndPort string

type AlarmVolume uint

type NumTracks uint

type StreamRestartState string

type RoomDetectionChirpChannel uint

type NetsettingsUpdateID string

type EQValue int

type EQType string

type MediaDuration string

type TransportActions string

type RestartPending bool

type URIMetaData string

type ConfigModeState string

type HTForwardEnabled bool

type SubGain string

type SupportsAudioClip bool

type WirelessLeafOnly bool

type SourceAreasUpdateID string

type UpdateID uint

type SessionId string

type Username string

type Mute bool

type HeadphoneConnected bool

type CustomerID string

type GroupID string

type HouseholdID string

type Treble int

type AccountMd string

type UserIdHashCode string

type DisplaySoftwareVersion string

type QueueID uint

type TrackNumbersCSV string

type SurroundEnabled bool

type DiagnosticID uint

type RecentlyPlayedUpdateID string

type LastChangedPlayState string

type AlarmProgramURI string

type SourceState string

type URI string

type Prefix string

type SearchCapabilities string

type SortCapabilities string

type RoomCalibrationState int

type ServiceDescriptorList string

type OutputFixed bool

type RampTimeSeconds uint

type AccountUDN string

type AlarmEnabled bool

type TimeServer string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type WirelessMode uint

type Code string

type QueueOwnerContext string

type AVTransportURIMetaData string

type Result string

type ServiceId uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type RoomCalibrationCoefficients string

type VolumeAVTransportURI string

type Speed string

type RoomCalibrationEnabled bool

type OAuthDeviceID string

type ISO8601Time string

type AlarmProgramMetaData string

type RecordQualityMode string

type SoftwareVersion string

type ButtonState string

type ServiceTypeList string

type PossibleRecordQualityModes string

type Filter string

type ContainerUpdateIDs string

type Curated bool

type AlarmRunSequence string

type AbsoluteTimePosition string

type RelativeCounterPosition int

type DirectControlClientID string

type Count uint

type IncludeControllers bool

type AuthorizationCode string

type CachedOnly bool

type TimeGeneration uint

type SearchCriteria string

type ChannelMapSet string

type BufferingResultCode int

type Volume uint

type SurroundMode string

type AVTransportURI string

type AlarmIDRunning uint

type MicEnabled uint
