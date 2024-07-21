// CODE GENERATED, DO NOT EDIT

// Package S14 contains the implementation for the Sonos Beam (Model S14) services.
package S14

// Services struct contains service clients for different functionalities of a Sonos device.
type AutoplayIncludeLinkedZones bool

type WifiEnabled bool

type AccountID string

type TimeGeneration uint

type RelativeCounterPosition int

type RampTimeSeconds uint

type AlarmRunSequence string

type TimeZoneAutoAdjustDst bool

type Timeout uint

type SeekTarget string

type ExtraInfo string

type TVConfigurationError bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AudioDelayRightRear string

type AlarmList string

type EnqueuedTransportURI string

type SubCrossover string

type ThirdPartyMediaServersX string

type GroupCoordinatorIsLocal bool

type LeftVolume uint

type VariableName string

type Invisible bool

type SatRoomUUID string

type LastChangedPlayState string

type ServiceId uint

type ZonePlayerUUIDsInGroup string

type SleepTimerGeneration uint

type SourceState string

type AvailableRoomCalibration string

type TrackNumber uint

type SortCriteria string

type Count uint

type TagValueList string

type Seed string

const ()

type TransportPlaySpeed string

type AlarmLoggedStartTime string

type RecentlyPlayedUpdateID string

type TrackNumbersCSV string

type VolumeDB int

type StubsCreated string

type AbsoluteCounterPosition int

type URIMetaData string

type AVTransportURIMetaData string

type LIST_URI string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type SortCapabilities string

type Browseable bool

type FavoritePresetsUpdateID string

type TransportErrorHttpHeaders string

type PossibleRecordStorageMedia string

type AutoplayRoomUUID string

type SourceAreaIds string

type FavoritesUpdateID string

type MACAddress string

type AvailableSoftwareUpdate string

type DirectControlClientID string

type SearchCapabilities string

type ServiceDescriptorList string

type SessionId string

type QueueUpdateID uint

type ShareIndexLastError string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type TransportSettings string

type EnqueueAs bool

type ResetVolumeAfter bool

type SortOrder string

type ButtonState string

type DailyIndexRefreshTime string

type TimeFormat string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type SleepTimerState string

type LastIndexChange string

type ContainerUpdateIDs string

type HTAudioIn uint

type PossibleRecordQualityModes string

type NumTracks uint

type AlbumArtistDisplayOption string

type UserRadioUpdateID string

type Configuration string

type IRCode string

type RoomCalibrationCalibrationMode string

type UpdateIDX uint

type TimeZoneInformation string

type AlarmListVersion string

type Origin string

type RadioLocationUpdateID uint

type WirelessLeafOnly bool

type SnoozeRunning bool

type URI string

type ZoneName string

type ConfigMode string

type DialogLevel string

type MediaDuration string

type RelativeTimePosition string

type RejoinGroup bool

type RightVolume uint

type MobileDeviceUDN string

type ISO8601Time string

type AlarmEnabled bool

type TransportStatus string

type MobileDeviceName string

type VolumeAVTransportURI string

type AccountCredential string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type RoomDetectionDurationMilliseconds uint

type ShareIndexInProgress bool

type RoomDetectionChirpChannel uint

type ConfigModeState string

type Username string

type Bass int

type EQType string

type AccountUID uint

type SourceAreasUpdateID string

type VLIState string

type HasConfiguredSSID bool

type Queue string

type SurroundMode string

type RoomCalibrationID string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type RecordQualityMode string

type SourceProtocolInfo string

type ConnectionManager string

type RadioFavoritesUpdateID uint

type UserIdHashCode string

type Speed string

type AreasUpdateID string

type AlarmProgramURI string

type ResumePlayback bool

type TransportErrorDescription string

type LastChange string

type SupportsAudioClip bool

type HardwareVersion string

type GroupVolume uint

type TimeZone string

type DateFormat string

type DirectControlIsSuspended bool

type GroupID string

type Prefix string

type Icon string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

const ()

type RecordStorageMedium string

type EnqueuedTransportURIMetaData string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type TOSLinkConnected bool

type RDMEnabled bool

type Track uint

type SearchCriteria string

type Index uint

type DisplaySoftwareVersion string

type QueueOwnerContext string

type AccountMd string

type TimeStamp string

type Section uint

type ProtocolInfo string

type ConnectionID int

type SystemUpdateID uint

type RoomDetectionChirpIfPlayingSwappableAudio bool

type PresetNameList string

type AlarmVolume uint

type ValidPlayModes string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type RoomCalibrationState int

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type ConfigModeOptions string

type SurroundLevel string

type AccountType uint

type RecordMediumWriteStatus string

type AlarmIDRunning uint

type QueueID uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type IsZoneBridge bool

type LocalGroupUUID string

type BootSeq uint

type EQValue int

type ThirdPartyHash string

type ConnectionIDs string

type MicEnabled uint

type DirectControlAccountID string

type MemberID string

type TrackList string

type SinkProtocolInfo string

type SoftwareVersion string

const (
	OnIRRepeaterState       IRRepeaterState = "On"
	OffIRRepeaterState      IRRepeaterState = "Off"
	DisabledIRRepeaterState IRRepeaterState = "Disabled"
)

type IRRepeaterState string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type TransportErrorHttpCode string

type Version string

type ChannelFreq uint

type VoiceConfigState uint

type Loudness bool

type ChannelMap string

type RoomCalibrationCoefficients string

type AccountPassword string

type TimeServer string

type AlarmRunning bool

type OAuthDeviceID string

type ZoneGroupName string

type AlarmState string

type HTSatChanMapSet string

type OutputFixed bool

type NightMode bool

type UpdateItem string

type AlarmID uint

type RestartPending bool

type UpdateFlags uint

type NetsettingsUpdateID string

type Orientation int

type UpdateURL string

type IsIdle bool

type AutoplaySource string

type QueueOwnerID string

type TrackDuration string

type LIST_URIMetaData string

type RoomCalibrationEnabled bool

type Result string

type SerialNumber string

type AutoplayVolume uint

type ProgramURI string

type RoomCalibrationAvailable bool

type AccountNickname string

type AlarmIncludeLinkedZones bool

type AlarmRoomUUID string

type RcsID int

type TargetRoomName string

type QueuePolicy string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AVTransportID int

type AbsoluteTimePosition string

type SurroundEnabled bool

type KeepGrouped bool

type MID string

type ZoneGroupState string

type HeadphoneConnected bool

type AudioDelayLeftRear string

type MuseHouseholdId string

type IncludeControllers bool

type LIST_URI_AND_METADATA string

type SpeakerSize uint

type DiagnosticID uint

type PossiblePlaybackStorageMedia string

type NumTracksChange int

type ObjectID string

type ShareListUpdateID string

type HTFreq uint

type Treble int

type CrossfadeMode bool

type AVTransportURI string

type MusicSurroundLevel string

type NumberOfTracks uint

type GroupVolumeChangeable bool

type TrackURI string

type MuseSessions string

type PlayerID string

type HouseholdID string

type MoreInfo string

type VirtualLineInGroupID string

type TimeZoneIndex int

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type VolumeAdjustment int

type RemoteConfigured bool

const (
	OnLEDFeedbackState  LEDFeedbackState = "On"
	OffLEDFeedbackState LEDFeedbackState = "Off"
)

type LEDFeedbackState string

type IsExpired bool

type AlarmProgramMetaData string

type SavedQueuesUpdateID string

type WirelessMode uint

type GroupMute bool

type CachedOnly bool

type InstanceID uint

type AuthorizationCode string

type Filter string

type SupportsAudioIn bool

type HTBondedZoneCommitState uint

type RoomDetectionPlayId uint

type SupportsOutputFixed bool

type SubGain string

type StreamRestartState string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type UpdateExtraOptions string

type BufferingResultCode int

type IRRemoteName string

type VariableStringValue string

type AccountUDN string

type UpdateID uint

type HdmiCecAvailable bool

type VoiceUpdateID uint

type ZoneGroupID string

type MobileIPAndPort string

type TransportErrorURI string

type MemberList string

type Volume uint

type RedirectURI string

type AutoplayUseVolume bool

type Curated bool

type SecureRegState uint

type SubPolarity string

type SettingsReplicationState string

type BehindWifiExtender uint

type CopyrightInfo string

type Flags uint

type DID string

type Mute bool

type AudioDelay string

type AccountTier uint

type TrackMetaData string

type SavedQueueTitle string

type IPAddress string

type SubEnabled bool

type TransportActions string

type ServiceListVersion string

type Code string

type ServiceTypeList string

type CustomerID string

type AirPlayEnabled bool

type ChannelMapSet string
