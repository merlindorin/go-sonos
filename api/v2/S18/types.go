// CODE GENERATED, DO NOT EDIT

// Package S18 contains the implementation for the Sonos One (Model S18) services.
package S18

// Services struct contains service clients for different functionalities of a Sonos device.
type AlarmRunning bool

type Treble int

type HeadphoneConnected bool

type AudioDelayLeftRear string

type TrackDuration string

type SavedQueueTitle string

type LocalGroupUUID string

type AvailableSoftwareUpdate string

type UpdateIDX uint

type MuseHouseholdId string

type UpdateExtraOptions string

type ChannelMapSet string

type Orientation int

type ConfigMode string

type AccountNickname string

type IsExpired bool

type Track uint

type TrackURI string

type OutputFixed bool

type CachedOnly bool

type ChannelFreq uint

type Curated bool

type TimeZoneIndex int

type AlarmIDRunning uint

type AVTransportID int

type TagValueList string

type HardwareVersion string

type AlarmIncludeLinkedZones bool

type SearchCapabilities string

type AccountType uint

type DirectControlClientID string

type RejoinGroup bool

type AirPlayEnabled bool

type QueueOwnerContext string

type RoomCalibrationCoefficients string

type AccountTier uint

type ThirdPartyMediaServersX string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type SleepTimerState string

type SearchCriteria string

type HouseholdID string

type AvailableRoomCalibration string

type SerialNumber string

type TVConfigurationError bool

type BootSeq uint

type RecordMediumWriteStatus string

type DirectControlAccountID string

type MemberID string

type ResetVolumeAfter bool

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type MobileDeviceName string

type RoomCalibrationEnabled bool

type SleepTimerGeneration uint

type SnoozeRunning bool

type FavoritesUpdateID string

type Mute bool

type SubCrossover string

type VolumeAVTransportURI string

type TransportErrorHttpHeaders string

type EnqueuedTransportURI string

type ShareIndexLastError string

type AutoplayRoomUUID string

type HasConfiguredSSID bool

type SourceState string

type VLIState string

type RightVolume uint

type SurroundLevel string

type Version string

type RecentlyPlayedUpdateID string

type DID string

type HdmiCecAvailable bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AVTransportURI string

type Browseable bool

type SettingsReplicationState string

type IsZoneBridge bool

type IsIdle bool

type SubEnabled bool

type IncludeControllers bool

type TrackNumber uint

type EthLink bool

type VoiceUpdateID uint

type AlarmRoomUUID string

type SourceProtocolInfo string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type ButtonState string

type ProgramURI string

type VirtualLineInGroupID string

type Username string

type LIST_URI_AND_METADATA string

type UpdateItem string

type ConnectionID int

type UpdateID uint

type SupportsAudioClip bool

type MemberList string

type AlbumArtistDisplayOption string

type Icon string

type QueueOwnerID string

type DailyIndexRefreshTime string

type CopyrightInfo string

type ThirdPartyHash string

type PlayerID string

type RadioLocationUpdateID uint

type SurroundEnabled bool

type DirectControlIsSuspended bool

type NumTracks uint

type SystemUpdateID uint

type AutoplayIncludeLinkedZones bool

type Volume uint

type AudioDelay string

type NetsettingsUpdateID string

type AlarmVolume uint

type TimeGeneration uint

type RelativeCounterPosition int

type SeekTarget string

type AutoplayVolume uint

type MediaDuration string

type ServiceTypeList string

type EQValue int

type RoomCalibrationID string

type HTAudioIn uint

type GroupVolume uint

type ServiceDescriptorList string

type VariableName string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type ObjectID string

type RoomDetectionPlayId uint

type AccountCredential string

type AlarmID uint

type AlarmEnabled bool

type SortCriteria string

type VoiceConfigState uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type AlarmState string

type AccountMd string

type RadioFavoritesUpdateID uint

type WifiEnabled bool

type GroupCoordinatorIsLocal bool

type UpdateURL string

type MobileDeviceUDN string

type TransportErrorDescription string

type ValidPlayModes string

type HTSatChanMapSet string

type MobileIPAndPort string

type TrackNumbersCSV string

type SubPolarity string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AlarmProgramURI string

type AlarmProgramMetaData string

type Section uint

type TrackMetaData string

type Filter string

type SurroundMode string

type ISO8601Time string

type ConnectionManager string

type SortCapabilities string

type ContainerUpdateIDs string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type TimeServer string

type AVTransportURIMetaData string

type SoftwareVersion string

type LastChange string

type EnqueueAs bool

type Prefix string

type Seed string

type ZoneName string

type AccountUDN string

type RedirectURI string

type TimeZoneInformation string

const ()

type RecordStorageMedium string

type NumberOfTracks uint

type LIST_URIMetaData string

type FavoritePresetsUpdateID string

type DiagnosticID uint

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type NumTracksChange int

type ConfigModeOptions string

type RoomCalibrationCalibrationMode string

type AccountUID uint

type AlarmListVersion string

type RoomCalibrationState int

type BufferingResultCode int

type UpdateFlags uint

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type MoreInfo string

type RoomDetectionChirpChannel uint

type RoomDetectionDurationMilliseconds uint

type AccountID string

type TargetRoomName string

type GroupMute bool

type ServiceId uint

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TransportStatus string

type CrossfadeMode bool

type QueueUpdateID uint

type Count uint

type UserIdHashCode string

type RcsID int

type RampTimeSeconds uint

type SpeakerSize uint

type LastIndexChange string

type IPAddress string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type Speed string

type Origin string

type StubsCreated string

type AlarmRunSequence string

type TransportSettings string

type SavedQueuesUpdateID string

type ShareListUpdateID string

type Flags uint

type MID string

type MuseSessions string

type SourceAreaIds string

type ZoneGroupName string

type ZoneGroupID string

type ConfigModeState string

type DialogLevel string

type AreasUpdateID string

type TransportErrorURI string

type AlarmLoggedStartTime string

type SortOrder string

type MACAddress string

type WirelessLeafOnly bool

type SourceAreasUpdateID string

type PossibleRecordQualityModes string

type ResumePlayback bool

type LeftVolume uint

type Code string

type TimeStamp string

type TransportActions string

type Index uint

type UserRadioUpdateID string

type HTBondedZoneCommitState uint

type VolumeDB int

type AccountPassword string

const ()

type TransportPlaySpeed string

type URI string

type SinkProtocolInfo string

type ConnectionIDs string

type Invisible bool

type SecureRegState uint

type AlarmList string

type PossiblePlaybackStorageMedia string

type StreamRestartState string

type ProtocolInfo string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type EQType string

type RoomCalibrationAvailable bool

type ZoneGroupState string

type ZonePlayerUUIDsInGroup string

type AbsoluteCounterPosition int

type Queue string

type ServiceListVersion string

type QueuePolicy string

type TransportErrorHttpCode string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type EnqueuedTransportURIMetaData string

type GroupID string

type ShareIndexInProgress bool

type AuthorizationCode string

type TimeFormat string

type DateFormat string

type RestartPending bool

type MusicSurroundLevel string

type RDMEnabled bool

type RecordQualityMode string

type GroupVolumeChangeable bool

type PresetNameList string

type VolumeAdjustment int

type CustomerID string

type AbsoluteTimePosition string

type TrackList string

type SupportsAudioIn bool

type InstanceID uint

type SatRoomUUID string

type SessionId string

type SubGain string

type NightMode bool

type SupportsOutputFixed bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AutoplaySource string

type AutoplayUseVolume bool

type WirelessMode uint

type KeepGrouped bool

type QueueID uint

type TimeZoneAutoAdjustDst bool

type URIMetaData string

type BehindWifiExtender uint

type AudioDelayRightRear string

type OAuthDeviceID string

type TimeZone string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type Result string

type HTFreq uint

type LastChangedPlayState string

type DisplaySoftwareVersion string

type ExtraInfo string

type Bass int

type Loudness bool

type ChannelMap string

type PossibleRecordStorageMedia string

type LIST_URI string

type Configuration string

type MicEnabled uint

type HTForwardEnabled bool

type RelativeTimePosition string

type VariableStringValue string
