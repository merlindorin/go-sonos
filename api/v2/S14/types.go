// CODE GENERATED, DO NOT EDIT

// Package S14 contains the implementation for the Sonos Beam (Model S14) services.
package S14

// Services struct contains service clients for different functionalities of a Sonos device.
type MicEnabled uint

type AccountMd string

type MobileIPAndPort string

type SeekTarget string

type SavedQueuesUpdateID string

type LeftVolume uint

type AudioDelay string

type SurroundMode string

type ThirdPartyMediaServersX string

type AlarmVolume uint

type TransportErrorURI string

type Prefix string

type UpdateURL string

type UpdateExtraOptions string

type TimeFormat string

type LIST_URI string

type AutoplayIncludeLinkedZones bool

type SubPolarity string

type RoomCalibrationCoefficients string

type AccountNickname string

type DirectControlAccountID string

type SourceState string

type TrackNumber uint

type AutoplayVolume uint

type DID string

type IsExpired bool

type TimeZoneInformation string

type TransportErrorHttpHeaders string

const ()

type RecordStorageMedium string

type NumberOfTracks uint

type NumTracks uint

type SleepTimerState string

type Count uint

type TargetRoomName string

const (
	OnLEDFeedbackState  LEDFeedbackState = "On"
	OffLEDFeedbackState LEDFeedbackState = "Off"
)

type LEDFeedbackState string

type CustomerID string

type AlbumArtistDisplayOption string

type Treble int

type TimeGeneration uint

type InstanceID uint

type OutputFixed bool

type RoomCalibrationAvailable bool

type SatRoomUUID string

type MACAddress string

type ExtraInfo string

type SourceAreaIds string

type MemberID string

type ResumePlayback bool

type ConnectionIDs string

type SnoozeRunning bool

type Timeout uint

type Version string

type AlarmEnabled bool

type AutoplaySource string

type WirelessMode uint

type RoomDetectionPlayId uint

type LIST_URIMetaData string

type AVTransportID int

type RadioLocationUpdateID uint

type HardwareVersion string

type WifiEnabled bool

type RightVolume uint

type TrackDuration string

type AVTransportURI string

type DirectControlIsSuspended bool

type TrackList string

type FavoritePresetsUpdateID string

type SettingsReplicationState string

type ServiceListVersion string

type ZoneGroupState string

type RecordQualityMode string

type Track uint

type NumTracksChange int

type SearchCriteria string

type FavoritesUpdateID string

type ChannelFreq uint

type ServiceDescriptorList string

type OAuthDeviceID string

type UpdateFlags uint

type SourceAreasUpdateID string

type AlarmIDRunning uint

type LastIndexChange string

type ShareIndexLastError string

type AccountPassword string

type TimeZone string

type VoiceUpdateID uint

type URIMetaData string

type LocalGroupUUID string

type QueuePolicy string

type VariableName string

type AlarmIncludeLinkedZones bool

type HTSatChanMapSet string

type Orientation int

type RoomCalibrationEnabled bool

type MediaDuration string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type ShareIndexInProgress bool

type ConfigModeState string

type RemoteConfigured bool

type ServiceTypeList string

type NightMode bool

type ZoneName string

type AudioDelayRightRear string

type SessionId string

type QueueID uint

type VolumeDB int

type Loudness bool

type ProgramURI string

type SubCrossover string

type AccountCredential string

type Icon string

type GroupCoordinatorIsLocal bool

type ZonePlayerUUIDsInGroup string

type GroupVolume uint

type Mute bool

type AlarmState string

type ChannelMapSet string

type NetsettingsUpdateID string

type PossiblePlaybackStorageMedia string

type VLIState string

type SystemUpdateID uint

type VolumeAVTransportURI string

type ThirdPartyHash string

type AlarmProgramURI string

type AlarmRoomUUID string

type TimeZoneIndex int

type RoomCalibrationState int

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type PresetNameList string

type AccountUID uint

type TransportStatus string

type SleepTimerGeneration uint

const (
	OnIRRepeaterState       IRRepeaterState = "On"
	OffIRRepeaterState      IRRepeaterState = "Off"
	DisabledIRRepeaterState IRRepeaterState = "Disabled"
)

type IRRepeaterState string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AccountID string

type ZoneGroupID string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type TransportActions string

type ValidPlayModes string

type ConnectionManager string

type CopyrightInfo string

type SubGain string

type Origin string

type VolumeAdjustment int

type MemberList string

type Filter string

type SortCriteria string

type UserRadioUpdateID string

type IRRemoteName string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type RampTimeSeconds uint

type VariableStringValue string

type StubsCreated string

type MuseHouseholdId string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type RecordMediumWriteStatus string

type RelativeCounterPosition int

type SecureRegState uint

type BufferingResultCode int

type Volume uint

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type LastChange string

type SourceProtocolInfo string

type Invisible bool

type RoomDetectionDurationMilliseconds uint

type SupportsOutputFixed bool

type DirectControlClientID string

type ConnectionID int

type LastChangedPlayState string

type AlarmListVersion string

type ObjectID string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ConfigMode string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type BootSeq uint

type SubEnabled bool

type IncludeControllers bool

type AbsoluteTimePosition string

type MID string

type AuthorizationCode string

type UpdateIDX uint

type AlarmRunSequence string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type TimeStamp string

type RcsID int

type HTBondedZoneCommitState uint

type ConfigModeOptions string

type RoomCalibrationID string

type AlarmList string

type MuseSessions string

type URI string

type RedirectURI string

type AvailableSoftwareUpdate string

type TimeZoneAutoAdjustDst bool

type TimeServer string

type PlayerID string

type QueueUpdateID uint

type ShareListUpdateID string

type AccountType uint

type IsZoneBridge bool

type DisplaySoftwareVersion string

type GroupMute bool

type Seed string

type MusicSurroundLevel string

type CachedOnly bool

type AlarmRunning bool

type TransportSettings string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type Browseable bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type UserIdHashCode string

type Section uint

type SinkProtocolInfo string

type SerialNumber string

type Username string

type SavedQueueTitle string

type ContainerUpdateIDs string

type HasConfiguredSSID bool

type EQType string

type HeadphoneConnected bool

type AudioDelayLeftRear string

type PossibleRecordStorageMedia string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type EQValue int

type ChannelMap string

type SurroundEnabled bool

type MobileDeviceName string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TransportErrorDescription string

type CrossfadeMode bool

type IPAddress string

type DiagnosticID uint

type GroupID string

type StreamRestartState string

type HouseholdID string

type KeepGrouped bool

type Bass int

type RDMEnabled bool

type EnqueueAs bool

type SearchCapabilities string

type ButtonState string

type UpdateItem string

type DateFormat string

type IRCode string

type ZoneGroupName string

type AlarmID uint

type AutoplayRoomUUID string

type QueueOwnerContext string

type LIST_URI_AND_METADATA string

type AlarmProgramMetaData string

type TrackMetaData string

type Configuration string

type AreasUpdateID string

type EnqueuedTransportURIMetaData string

type Queue string

type ResetVolumeAfter bool

type ProtocolInfo string

type RecentlyPlayedUpdateID string

type HTAudioIn uint

type ServiceId uint

type RestartPending bool

type AirPlayEnabled bool

type AutoplayUseVolume bool

type TransportErrorHttpCode string

type TagValueList string

type SortCapabilities string

type SupportsAudioIn bool

type MoreInfo string

type Flags uint

type TOSLinkConnected bool

type SurroundLevel string

type AlarmLoggedStartTime string

type RejoinGroup bool

type Speed string

type VirtualLineInGroupID string

type MobileDeviceUDN string

type EnqueuedTransportURI string

type RadioFavoritesUpdateID uint

type TVConfigurationError bool

type SpeakerSize uint

type AccountTier uint

type ISO8601Time string

type TrackURI string

type AbsoluteCounterPosition int

type Result string

type Index uint

type VoiceConfigState uint

type RoomDetectionChirpChannel uint

type QueueOwnerID string

type DailyIndexRefreshTime string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type SortOrder string

type HTFreq uint

type WirelessLeafOnly bool

type BehindWifiExtender uint

type RoomCalibrationCalibrationMode string

const ()

type TransportPlaySpeed string

type PossibleRecordQualityModes string

type AVTransportURIMetaData string

type RelativeTimePosition string

type SupportsAudioClip bool

type AvailableRoomCalibration string

type SoftwareVersion string

type Code string

type Curated bool

type UpdateID uint

type IsIdle bool

type HdmiCecAvailable bool

type GroupVolumeChangeable bool

type TrackNumbersCSV string

type DialogLevel string

type AccountUDN string
