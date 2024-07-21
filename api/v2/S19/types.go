// CODE GENERATED, DO NOT EDIT

// Package S19 contains the implementation for the Sonos Arc (Model S19) services.
package S19

// Services struct contains service clients for different functionalities of a Sonos device.
type AlarmLoggedStartTime string

type SourceAreaIds string

type Seed string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AVTransportID int

type RedirectURI string

type GroupVolumeChangeable bool

type Volume uint

type AccountTier uint

type TimeZone string

type InstanceID uint

type ProtocolInfo string

type UpdateID uint

type SupportsAudioIn bool

type RoomCalibrationID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TransportActions string

type WifiEnabled bool

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SurroundMode string

type LeftVolume uint

type RoomCalibrationEnabled bool

type AlarmRunSequence string

type TimeZoneInformation string

type Track uint

type AVTransportURI string

type SleepTimerState string

type HTForwardEnabled bool

type NightMode bool

type DirectControlClientID string

type SatRoomUUID string

type SupportsOutputFixed bool

type AccountID string

type DateFormat string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type AvailableRoomCalibration string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type MobileIPAndPort string

type MACAddress string

type DialogLevel string

type IncludeControllers bool

type ProgramURI string

type AudioDelay string

type AudioDelayRightRear string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type SeekTarget string

type LIST_URI string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type OutputFixed bool

type SpeakerSize uint

type AccountPassword string

type ThirdPartyHash string

type SinkProtocolInfo string

type ConnectionID int

type Speed string

type SourceAreasUpdateID string

type ConfigMode string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type AlarmID uint

type CrossfadeMode bool

type RelativeTimePosition string

type RcsID int

type ContainerUpdateIDs string

type MobileDeviceName string

type NumberOfTracks uint

type AVTransportURIMetaData string

type RoomCalibrationState int

type RoomDetectionPlayId uint

type EQValue int

type AlarmList string

type SoftwareVersion string

type AutoplayRoomUUID string

type SurroundEnabled bool

type AccountCredential string

type HasConfiguredSSID bool

type MID string

type AlarmProgramURI string

type PossiblePlaybackStorageMedia string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type TrackList string

type AutoplayUseVolume bool

type LastChange string

type MemberID string

type SessionId string

type SurroundLevel string

type CustomerID string

type UpdateIDX uint

type ZoneGroupID string

type AlarmEnabled bool

type AlarmRunning bool

type URIMetaData string

type FavoritePresetsUpdateID string

type ServiceId uint

type SubEnabled bool

type TimeZoneIndex int

type TimeZoneAutoAdjustDst bool

type HTSatChanMapSet string

type HTFreq uint

type IRRemoteName string

type MediaDuration string

type URI string

type GroupID string

type QueueUpdateID uint

type ConfigModeOptions string

type VLIState string

type EnqueueAs bool

const (
	OnLEDFeedbackState  LEDFeedbackState = "On"
	OffLEDFeedbackState LEDFeedbackState = "Off"
)

type LEDFeedbackState string

type SavedQueueTitle string

type ResumePlayback bool

type HTBondedZoneCommitState uint

type AccountMd string

type TimeGeneration uint

type TransportErrorHttpHeaders string

type AlarmIDRunning uint

type UpdateFlags uint

type TimeStamp string

type MicEnabled uint

type AreasUpdateID string

type TOSLinkConnected bool

type MusicSurroundLevel string

type RelativeCounterPosition int

type SortCriteria string

type Index uint

type AutoplayVolume uint

type LocalGroupUUID string

type AbsoluteTimePosition string

type ServiceDescriptorList string

type Username string

type TransportSettings string

type SourceProtocolInfo string

type Count uint

type IsZoneBridge bool

type CachedOnly bool

type MuseSessions string

type SupportsAudioClip bool

type HdmiCecAvailable bool

type SecureRegState uint

type HeadphoneConnected bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type StubsCreated string

type AudioDelayLeftRear string

type UpdateItem string

type AlarmProgramMetaData string

type SnoozeRunning bool

type MemberList string

type LIST_URI_AND_METADATA string

type ChannelMap string

type ShareListUpdateID string

type TVConfigurationError bool

type RightVolume uint

type SubGain string

type VariableName string

type HouseholdID string

type Timeout uint

type UserRadioUpdateID string

type VoiceConfigState uint

type QueueOwnerContext string

type MuseHouseholdId string

type DiagnosticID uint

type GroupCoordinatorIsLocal bool

type QueuePolicy string

type OAuthDeviceID string

type ZoneGroupState string

type MobileDeviceUDN string

type BootSeq uint

type ServiceTypeList string

type PossibleRecordStorageMedia string

type TrackMetaData string

type TrackURI string

type RecentlyPlayedUpdateID string

type ChannelMapSet string

type EnqueuedTransportURI string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type Flags uint

type VolumeAVTransportURI string

type SubCrossover string

type NumTracksChange int

type StreamRestartState string

type HTAudioIn uint

type ZonePlayerUUIDsInGroup string

type TimeFormat string

type DirectControlAccountID string

type SearchCriteria string

type Filter string

type AlarmRoomUUID string

type PlayerID string

type Prefix string

type AirPlayEnabled bool

type MoreInfo string

type Result string

type Configuration string

type QueueID uint

type TransportStatus string

type ShareIndexLastError string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type DisplaySoftwareVersion string

type NumTracks uint

type LastIndexChange string

type FavoritesUpdateID string

type Invisible bool

type QueueOwnerID string

type Loudness bool

type ThirdPartyMediaServersX string

type TimeServer string

type TransportErrorDescription string

type SortOrder string

type AutoplaySource string

type GroupMute bool

type RecordQualityMode string

type Browseable bool

type TargetRoomName string

type SerialNumber string

type ExtraInfo string

type SleepTimerGeneration uint

type HardwareVersion string

type RoomDetectionChirpChannel uint

type AccountType uint

type IsExpired bool

type AlarmIncludeLinkedZones bool

type LIST_URIMetaData string

type SavedQueuesUpdateID string

type SettingsReplicationState string

type RoomCalibrationCalibrationMode string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type AlarmState string

type WirelessLeafOnly bool

type TransportErrorURI string

type Section uint

type TrackNumber uint

type VolumeDB int

type Version string

type ISO8601Time string

type EnqueuedTransportURIMetaData string

type CopyrightInfo string

type IRCode string

type ServiceListVersion string

type UpdateURL string

type Origin string

type ConnectionManager string

type ZoneName string

type ButtonState string

type RemoteConfigured bool

type Treble int

type ConnectionIDs string

type AbsoluteCounterPosition int

type TagValueList string

type PresetNameList string

type AuthorizationCode string

type RoomDetectionDurationMilliseconds uint

type VirtualLineInGroupID string

type VolumeAdjustment int

type DID string

type AvailableSoftwareUpdate string

type AlarmListVersion string

type RadioFavoritesUpdateID uint

type EQType string

type RoomCalibrationCoefficients string

type VoiceUpdateID uint

type ResetVolumeAfter bool

type SystemUpdateID uint

type ShareIndexInProgress bool

type IsIdle bool

type EthLink bool

type ZoneGroupName string

type UpdateExtraOptions string

type DailyIndexRefreshTime string

type TransportErrorHttpCode string

type PossibleRecordQualityModes string

type IPAddress string

type UserIdHashCode string

type AutoplayIncludeLinkedZones bool

type AccountUDN string

type RDMEnabled bool

type AlarmVolume uint

type RecordMediumWriteStatus string

type TrackDuration string

type AlbumArtistDisplayOption string

type SortCapabilities string

type RejoinGroup bool

type Mute bool

type RampTimeSeconds uint

type NetsettingsUpdateID string

const ()

type RecordStorageMedium string

const ()

type TransportPlaySpeed string

type Code string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type SourceState string

type RadioLocationUpdateID uint

type GroupVolume uint

type TrackNumbersCSV string

type Bass int

type BehindWifiExtender uint

type RoomDetectionChirpIfPlayingSwappableAudio bool

type VariableStringValue string

type AccountNickname string

type SubPolarity string

type RoomCalibrationAvailable bool

type DirectControlIsSuspended bool

type ObjectID string

type SearchCapabilities string

type LastChangedPlayState string

const (
	OnIRRepeaterState       IRRepeaterState = "On"
	OffIRRepeaterState      IRRepeaterState = "Off"
	DisabledIRRepeaterState IRRepeaterState = "Disabled"
)

type IRRepeaterState string

type BufferingResultCode int

type Queue string

type Icon string

type WirelessMode uint

type ChannelFreq uint

type ConfigModeState string

type Curated bool

type AccountUID uint

type RestartPending bool

type ValidPlayModes string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type Orientation int

type KeepGrouped bool
