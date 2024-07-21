// CODE GENERATED, DO NOT EDIT

// Package S38 contains the implementation for the Sonos One SL (Model S38) services.
package S38

// Services struct contains service clients for different functionalities of a Sonos device.
type VoiceConfigState uint

type DailyIndexRefreshTime string

type TimeFormat string

type DisplaySoftwareVersion string

type ConfigModeOptions string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type AccountUID uint

type RDMEnabled bool

type DirectControlIsSuspended bool

type Count uint

type MACAddress string

type Track uint

type FavoritePresetsUpdateID string

type EnqueuedTransportURIMetaData string

type VolumeDB int

type AccountTier uint

type TimeZoneAutoAdjustDst bool

type TransportStatus string

type NumberOfTracks uint

type AlarmIDRunning uint

type URI string

type RejoinGroup bool

type AirPlayEnabled bool

type HTFreq uint

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AlarmRoomUUID string

type AbsoluteTimePosition string

type AccountUDN string

type HardwareVersion string

type HTAudioIn uint

type Mute bool

type LocalGroupUUID string

type GroupMute bool

type SurroundLevel string

const ()

type RecordStorageMedium string

type FavoritesUpdateID string

type ConfigModeState string

type SortCapabilities string

type LastChangedPlayState string

type Seed string

type VoiceUpdateID uint

type AlarmRunSequence string

type PlayerID string

type SleepTimerState string

type ProtocolInfo string

type SourceAreasUpdateID string

type RecordMediumWriteStatus string

type ObjectID string

type SettingsReplicationState string

type SecureRegState uint

type SourceAreaIds string

type PossibleRecordStorageMedia string

type AutoplayRoomUUID string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type DialogLevel string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

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

type SurroundMode string

type TimeZoneInformation string

type TrackMetaData string

type SurroundEnabled bool

type RoomCalibrationState int

type RoomDetectionPlayId uint

type ServiceId uint

type UpdateIDX uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AlarmProgramMetaData string

type VLIState string

type TrackNumber uint

type MobileDeviceName string

type Configuration string

type ZoneGroupState string

type Version string

type SearchCapabilities string

type SoftwareVersion string

type AutoplayIncludeLinkedZones bool

type WirelessLeafOnly bool

type SupportsOutputFixed bool

type CrossfadeMode bool

type InstanceID uint

type SavedQueueTitle string

type AreasUpdateID string

type VariableName string

type UserIdHashCode string

type ThirdPartyMediaServersX string

type AVTransportURI string

type KeepGrouped bool

type VolumeAVTransportURI string

type RecentlyPlayedUpdateID string

type ExtraInfo string

type AutoplayVolume uint

type HTForwardEnabled bool

type DiagnosticID uint

type TrackList string

type Result string

type SortCriteria string

type UpdateID uint

type CustomerID string

type LastIndexChange string

type ShareListUpdateID string

type DID string

type AudioDelay string

type RoomCalibrationEnabled bool

type RestartPending bool

type NumTracksChange int

type Prefix string

type ZoneGroupID string

type Icon string

type QueuePolicy string

type VariableStringValue string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type TransportActions string

type BehindWifiExtender uint

type ButtonState string

type RadioFavoritesUpdateID uint

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type MuseHouseholdId string

type UpdateItem string

type ISO8601Time string

type AlarmProgramURI string

type SleepTimerGeneration uint

type AccountCredential string

type NumTracks uint

type RcsID int

type AudioDelayLeftRear string

type TransportSettings string

type SatRoomUUID string

type MID string

type LIST_URI_AND_METADATA string

type TimeZone string

type SupportsAudioIn bool

type VirtualLineInGroupID string

type MobileDeviceUDN string

type TransportErrorHttpCode string

type MemberID string

type QueueOwnerID string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type IsIdle bool

type OutputFixed bool

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AlarmList string

type Queue string

type SupportsAudioClip bool

type WirelessMode uint

type VolumeAdjustment int

type TimeZoneIndex int

type PossibleRecordQualityModes string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type GroupCoordinatorIsLocal bool

type MuseSessions string

type GroupID string

type Flags uint

type ZonePlayerUUIDsInGroup string

type Filter string

type HouseholdID string

type RoomDetectionDurationMilliseconds uint

type AlbumArtistDisplayOption string

type SpeakerSize uint

type Origin string

type QueueOwnerContext string

type PresetNameList string

type TimeStamp string

const ()

type TransportPlaySpeed string

type MemberList string

type AlarmIncludeLinkedZones bool

type GroupVolumeChangeable bool

type AccountNickname string

type IncludeControllers bool

type SinkProtocolInfo string

type ConfigMode string

type HeadphoneConnected bool

type ChannelMap string

type TrackDuration string

type ValidPlayModes string

type LIST_URIMetaData string

type ServiceDescriptorList string

type Speed string

type Section uint

type SeekTarget string

type BootSeq uint

type ChannelFreq uint

type Username string

type LeftVolume uint

type ZoneGroupName string

type DateFormat string

type ResetVolumeAfter bool

type AutoplaySource string

type SerialNumber string

type Treble int

type AccountType uint

type LIST_URI string

type ContainerUpdateIDs string

type TransportErrorDescription string

type TransportErrorHttpHeaders string

type RecordQualityMode string

type TimeServer string

type RedirectURI string

type UpdateExtraOptions string

type AlarmVolume uint

type ServiceListVersion string

type QueueID uint

type ResumePlayback bool

type SourceProtocolInfo string

type ShareIndexInProgress bool

type ProgramURI string

type AlarmID uint

type MediaDuration string

type URIMetaData string

type TargetRoomName string

type EthLink bool

type MicEnabled uint

type RightVolume uint

type UpdateURL string

type TrackURI string

type AlarmLoggedStartTime string

type ShareIndexLastError string

type NetsettingsUpdateID string

type SavedQueuesUpdateID string

type Invisible bool

type Code string

type AuthorizationCode string

type TimeGeneration uint

type LastChange string

type EnqueueAs bool

type SortOrder string

type TVConfigurationError bool

type RoomCalibrationCalibrationMode string

type RelativeTimePosition string

type AlarmState string

type SearchCriteria string

type GroupVolume uint

type ServiceTypeList string

type SubGain string

type NightMode bool

type PossiblePlaybackStorageMedia string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type Browseable bool

type SessionId string

type BufferingResultCode int

type TrackNumbersCSV string

type UpdateFlags uint

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type ConnectionID int

type RadioLocationUpdateID uint

type RoomDetectionChirpChannel uint

type ThirdPartyHash string

type TransportErrorURI string

type HTSatChanMapSet string

type CopyrightInfo string

type HdmiCecAvailable bool

type OAuthDeviceID string

type IsZoneBridge bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type AutoplayUseVolume bool

type IsExpired bool

type AvailableSoftwareUpdate string

type AlarmEnabled bool

type DirectControlClientID string

type UserRadioUpdateID string

type AudioDelayRightRear string

type MusicSurroundLevel string

type HTBondedZoneCommitState uint

type Volume uint

type Loudness bool

type AccountPassword string

type Index uint

type ZoneName string

type RampTimeSeconds uint

type MobileIPAndPort string

type AlarmListVersion string

type RoomCalibrationAvailable bool

type CachedOnly bool

type EnqueuedTransportURI string

type MoreInfo string

type HasConfiguredSSID bool

type AccountMd string

type AVTransportURIMetaData string

type SnoozeRunning bool

type AlarmRunning bool

type Bass int

type EQValue int

type SubEnabled bool

type SubPolarity string

type AbsoluteCounterPosition int

type Orientation int

type Curated bool

type EQType string

type RoomCalibrationID string

type SourceState string

type ConnectionIDs string

type AVTransportID int

type SubCrossover string

type AccountID string

type QueueUpdateID uint

type ConnectionManager string

type AvailableRoomCalibration string

type TagValueList string

type SystemUpdateID uint

type ChannelMapSet string

type StubsCreated string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type StreamRestartState string

type RoomCalibrationCoefficients string

type WifiEnabled bool

type RelativeCounterPosition int

type DirectControlAccountID string

type IPAddress string
