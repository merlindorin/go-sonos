// CODE GENERATED, DO NOT EDIT

// Package S19 contains the implementation for the Sonos Arc (Model S19) services.
package S19

// Services struct contains service clients for different functionalities of a Sonos device.
type AutoplayUseVolume bool

type Timeout uint

type TrackNumbersCSV string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TimeZoneIndex int

type AlarmLoggedStartTime string

type VLIState string

type RadioLocationUpdateID uint

type AlarmRunning bool

type GroupMute bool

type ZonePlayerUUIDsInGroup string

type AccountTier uint

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type TimeGeneration uint

type MemberList string

type SourceState string

type AccountPassword string

type RoomCalibrationCalibrationMode string

type UpdateFlags uint

type TimeFormat string

type SleepTimerState string

type HTForwardEnabled bool

type IRCode string

type OutputFixed bool

type TransportSettings string

type AutoplayVolume uint

type VariableStringValue string

type CustomerID string

type PossibleRecordStorageMedia string

type VirtualLineInGroupID string

type AccountType uint

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type HardwareVersion string

type HasConfiguredSSID bool

type SourceAreaIds string

type GroupID string

type UpdateID uint

type TVConfigurationError bool

const (
	OnLEDFeedbackState  LEDFeedbackState = "On"
	OffLEDFeedbackState LEDFeedbackState = "Off"
)

type LEDFeedbackState string

type QueueOwnerID string

type PossibleRecordQualityModes string

type SinkProtocolInfo string

type StubsCreated string

type NetsettingsUpdateID string

type PlayerID string

type IsIdle bool

type SerialNumber string

type Bass int

type AudioDelayLeftRear string

type QueueUpdateID uint

type AirPlayEnabled bool

type ChannelMapSet string

type VoiceUpdateID uint

type ZoneGroupState string

type SurroundEnabled bool

type VariableName string

type TrackDuration string

type EnqueuedTransportURI string

type UserRadioUpdateID string

type Username string

type Code string

type WirelessMode uint

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type Loudness bool

type AlarmRoomUUID string

type TimeZoneAutoAdjustDst bool

type RelativeCounterPosition int

type DirectControlIsSuspended bool

type StreamRestartState string

type SourceProtocolInfo string

type ServiceListVersion string

type DirectControlAccountID string

type SupportsAudioClip bool

type GroupCoordinatorIsLocal bool

type RoomCalibrationID string

type AccountUDN string

type ISO8601Time string

type Filter string

type BootSeq uint

type ObjectID string

type AlarmState string

type RoomCalibrationState int

type MicEnabled uint

type AccountUID uint

type Volume uint

type HeadphoneConnected bool

type TimeZone string

type SleepTimerGeneration uint

type SnoozeRunning bool

type ConnectionIDs string

type WifiEnabled bool

type AbsoluteTimePosition string

type ShareIndexLastError string

type EQValue int

type RampTimeSeconds uint

type Speed string

type AlarmRunSequence string

type Track uint

type SavedQueueTitle string

type SortCriteria string

type LastIndexChange string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

const ()

type TransportPlaySpeed string

type AlbumArtistDisplayOption string

type DialogLevel string

type RoomCalibrationCoefficients string

type UpdateIDX uint

type TransportStatus string

type Section uint

type TrackList string

type Index uint

type VoiceConfigState uint

type SurroundLevel string

type UpdateItem string

type AlarmList string

type RadioFavoritesUpdateID uint

type Configuration string

type HTBondedZoneCommitState uint

type AutoplayRoomUUID string

type TrackMetaData string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type Prefix string

type LIST_URI_AND_METADATA string

type AlarmIncludeLinkedZones bool

type RelativeTimePosition string

type NumTracksChange int

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type ValidPlayModes string

type SupportsAudioIn bool

type RDMEnabled bool

type AlarmProgramMetaData string

type NumTracks uint

type Count uint

type RightVolume uint

type IsExpired bool

type WirelessLeafOnly bool

type ChannelMap string

type AudioDelayRightRear string

type MusicSurroundLevel string

type RoomDetectionDurationMilliseconds uint

type SubEnabled bool

type MobileDeviceUDN string

type AlarmID uint

type SortCapabilities string

type ZoneName string

type HTFreq uint

type SatRoomUUID string

type RestartPending bool

type RecentlyPlayedUpdateID string

type IRRemoteName string

type RemoteConfigured bool

type OAuthDeviceID string

type ZoneGroupID string

type TimeZoneInformation string

type RcsID int

type ShareListUpdateID string

type SecureRegState uint

type LeftVolume uint

type ServiceDescriptorList string

type EQType string

type RecordQualityMode string

type IsZoneBridge bool

type MACAddress string

type EthLink bool

type ConfigMode string

type TimeStamp string

type HdmiCecAvailable bool

type ConfigModeState string

type AccountCredential string

type MuseSessions string

type SurroundMode string

type ZoneGroupName string

type ContainerUpdateIDs string

type AVTransportURIMetaData string

type Orientation int

type ChannelFreq uint

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AbsoluteCounterPosition int

type InstanceID uint

type SearchCriteria string

type AreasUpdateID string

type Origin string

type Queue string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type MoreInfo string

type GroupVolumeChangeable bool

type SpeakerSize uint

type SearchCapabilities string

type Icon string

type RoomDetectionChirpChannel uint

type QueuePolicy string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type UserIdHashCode string

type MobileDeviceName string

type AlarmEnabled bool

type AlarmProgramURI string

type TagValueList string

type SystemUpdateID uint

type SubGain string

type AlarmIDRunning uint

type TrackNumber uint

type Browseable bool

type SupportsOutputFixed bool

type AudioDelay string

type PossiblePlaybackStorageMedia string

type RecordMediumWriteStatus string

type HTSatChanMapSet string

type ProgramURI string

type DiagnosticID uint

type LIST_URIMetaData string

type LastChangedPlayState string

type AutoplaySource string

type TOSLinkConnected bool

type SubCrossover string

type HTAudioIn uint

type DID string

type NightMode bool

type TransportErrorHttpCode string

type DirectControlClientID string

type URIMetaData string

type SettingsReplicationState string

type DisplaySoftwareVersion string

type ThirdPartyMediaServersX string

type GroupVolume uint

type VolumeDB int

type AVTransportURI string

type EnqueueAs bool

type ExtraInfo string

type Mute bool

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type Result string

type QueueID uint

type Treble int

type UpdateURL string

type DateFormat string

type TransportErrorDescription string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type SortOrder string

type VolumeAdjustment int

type MuseHouseholdId string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type MobileIPAndPort string

type AlarmListVersion string

type RoomDetectionPlayId uint

const (
	OnIRRepeaterState       IRRepeaterState = "On"
	OffIRRepeaterState      IRRepeaterState = "Off"
	DisabledIRRepeaterState IRRepeaterState = "Disabled"
)

type IRRepeaterState string

type AccountID string

type AvailableSoftwareUpdate string

type AutoplayIncludeLinkedZones bool

type Version string

type ConnectionID int

type HouseholdID string

type SubPolarity string

type RoomCalibrationAvailable bool

type AccountNickname string

type NumberOfTracks uint

type LIST_URI string

type LocalGroupUUID string

type MID string

type RoomCalibrationEnabled bool

type IncludeControllers bool

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type TrackURI string

type EnqueuedTransportURIMetaData string

type MemberID string

type ResumePlayback bool

type QueueOwnerContext string

type CrossfadeMode bool

type URI string

type SavedQueuesUpdateID string

type FavoritePresetsUpdateID string

type ButtonState string

type DailyIndexRefreshTime string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type SeekTarget string

type BufferingResultCode int

type CachedOnly bool

const ()

type RecordStorageMedium string

type AVTransportID int

type IPAddress string

type SessionId string

type AccountMd string

type BehindWifiExtender uint

type AuthorizationCode string

type TimeServer string

type TransportErrorHttpHeaders string

type TransportActions string

type LastChange string

type FavoritesUpdateID string

type AlarmVolume uint

type ShareIndexInProgress bool

type TargetRoomName string

type KeepGrouped bool

type Curated bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type ThirdPartyHash string

type ResetVolumeAfter bool

type RejoinGroup bool

type AvailableRoomCalibration string

type CopyrightInfo string

type ConfigModeOptions string

type TransportErrorURI string

type Flags uint

type ServiceTypeList string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type MediaDuration string

type ConnectionManager string

type ProtocolInfo string

type Invisible bool

type SoftwareVersion string

type ServiceId uint

type RedirectURI string

type SourceAreasUpdateID string

type VolumeAVTransportURI string

type Seed string

type PresetNameList string

type UpdateExtraOptions string
