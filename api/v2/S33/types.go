// CODE GENERATED, DO NOT EDIT

// Package S33 contains the implementation for the SYMFONISK Bookshelf (Model S33) services.
package S33

// Services struct contains service clients for different functionalities of a Sonos device.
const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AlarmRoomUUID string

type MemberList string

type GroupCoordinatorIsLocal bool

type LeftVolume uint

type AudioDelayLeftRear string

type PresetNameList string

type AuthorizationCode string

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

const ()

type TransportPlaySpeed string

type ZoneName string

type CopyrightInfo string

type AutoplaySource string

type UpdateURL string

type CrossfadeMode bool

type AudioDelayRightRear string

type SurroundEnabled bool

type AlarmRunSequence string

type TransportErrorHttpCode string

type Prefix string

type SubPolarity string

type AccountNickname string

type TargetRoomName string

type HTFreq uint

type WirelessLeafOnly bool

type MID string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type IncludeControllers bool

type NumberOfTracks uint

type LIST_URI string

type SearchCriteria string

type Icon string

type KeepGrouped bool

type ISO8601Time string

type TimeStamp string

type AVTransportURIMetaData string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type Mute bool

type TimeServer string

type EnqueuedTransportURIMetaData string

type PlayerID string

type ResumePlayback bool

type ConnectionIDs string

type GroupVolumeChangeable bool

type SurroundLevel string

type AccountTier uint

type VoiceUpdateID uint

type AlarmProgramURI string

type TimeZoneIndex int

type DirectControlAccountID string

type TransportSettings string

type Count uint

type HardwareVersion string

type RightVolume uint

type AccountType uint

type TransportErrorDescription string

type Track uint

type ShareIndexInProgress bool

type SavedQueuesUpdateID string

type Configuration string

type AvailableRoomCalibration string

type TransportStatus string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type SortCapabilities string

type DialogLevel string

type NetsettingsUpdateID string

type TagValueList string

type ContainerUpdateIDs string

type HouseholdID string

type LocalGroupUUID string

type GroupMute bool

type Loudness bool

type VariableName string

type AlarmID uint

type Section uint

type SleepTimerGeneration uint

type ExtraInfo string

type EnqueueAs bool

type TrackList string

type AirPlayEnabled bool

type VirtualLineInGroupID string

type LIST_URI_AND_METADATA string

type EQValue int

type Treble int

type SubCrossover string

type RelativeCounterPosition int

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type AVTransportID int

type FavoritePresetsUpdateID string

type AlarmProgramMetaData string

type BufferingResultCode int

type QueueOwnerID string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type MobileDeviceUDN string

type MoreInfo string

type RoomDetectionDurationMilliseconds uint

type GroupVolume uint

type Curated bool

type TrackURI string

type SeekTarget string

type HTBondedZoneCommitState uint

type QueuePolicy string

type UpdateExtraOptions string

type InstanceID uint

type ObjectID string

type ResetVolumeAfter bool

type FavoritesUpdateID string

type ZoneGroupState string

type SortOrder string

type Orientation int

type Flags uint

type HeadphoneConnected bool

type RoomCalibrationAvailable bool

type AlarmVolume uint

type SourceState string

type SystemUpdateID uint

type ChannelFreq uint

type TrackNumbersCSV string

type StubsCreated string

type UpdateFlags uint

type Queue string

type QueueID uint

type TransportErrorHttpHeaders string

type HasConfiguredSSID bool

type RDMEnabled bool

type AlarmList string

type TransportActions string

type SourceAreaIds string

type RampTimeSeconds uint

type MediaDuration string

type URI string

type RecordQualityMode string

type SettingsReplicationState string

type ProgramURI string

type RelativeTimePosition string

type EnqueuedTransportURI string

type RoomCalibrationState int

type ZonePlayerUUIDsInGroup string

type AlarmIDRunning uint

type ConnectionManager string

type EQType string

type ThirdPartyMediaServersX string

type RecordMediumWriteStatus string

type AbsoluteCounterPosition int

type GroupID string

type Filter string

type ChannelMapSet string

type HTSatChanMapSet string

type DateFormat string

type PossiblePlaybackStorageMedia string

type TrackMetaData string

type AvailableSoftwareUpdate string

type MuseHouseholdId string

type LIST_URIMetaData string

type UserRadioUpdateID string

type SerialNumber string

type IPAddress string

type SupportsOutputFixed bool

type AreasUpdateID string

type VLIState string

type Origin string

type SearchCapabilities string

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

type WirelessMode uint

type AudioDelay string

type ThirdPartyHash string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type RestartPending bool

type ButtonState string

type SessionId string

type AccountUID uint

type TransportErrorURI string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type Volume uint

type RoomCalibrationCalibrationMode string

type MemberID string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type Version string

type TimeZone string

type StreamRestartState string

type RejoinGroup bool

type ZoneGroupID string

type RadioLocationUpdateID uint

type MACAddress string

type TVConfigurationError bool

type EthLink bool

type SurroundMode string

type SourceAreasUpdateID string

const ()

type RecordStorageMedium string

type SupportsAudioIn bool

type VolumeDB int

type AccountMd string

type SleepTimerState string

type LastChangedPlayState string

type SecureRegState uint

type ServiceListVersion string

type TimeZoneInformation string

type AVTransportURI string

type ConnectionID int

type Result string

type SatRoomUUID string

type HTAudioIn uint

type HdmiCecAvailable bool

type WifiEnabled bool

type ConfigModeState string

type DID string

type UserIdHashCode string

type UpdateIDX uint

type MobileDeviceName string

type SnoozeRunning bool

type NumTracksChange int

type RcsID int

type RadioFavoritesUpdateID uint

type Bass int

type AbsoluteTimePosition string

type SavedQueueTitle string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type UpdateID uint

type SupportsAudioClip bool

type VariableStringValue string

type ValidPlayModes string

type Index uint

type ServiceId uint

type Username string

type QueueOwnerContext string

type VolumeAVTransportURI string

type SubEnabled bool

type AlarmLoggedStartTime string

type MuseSessions string

type ProtocolInfo string

type NightMode bool

type ZoneGroupName string

type PossibleRecordQualityModes string

type IsIdle bool

type SubGain string

type HTForwardEnabled bool

type Seed string

type RoomCalibrationCoefficients string

type Speed string

type AlarmRunning bool

type DirectControlIsSuspended bool

type AutoplayIncludeLinkedZones bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type RoomDetectionPlayId uint

type VolumeAdjustment int

type Code string

type TimeZoneAutoAdjustDst bool

type URIMetaData string

type SpeakerSize uint

type OAuthDeviceID string

type LastChange string

type TrackNumber uint

type Invisible bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type VoiceConfigState uint

type AccountPassword string

type NumTracks uint

type ShareListUpdateID string

type DisplaySoftwareVersion string

type BootSeq uint

type ServiceDescriptorList string

type RoomCalibrationEnabled bool

type UpdateItem string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type RecentlyPlayedUpdateID string

type AutoplayVolume uint

type MobileIPAndPort string

type PossibleRecordStorageMedia string

type AlbumArtistDisplayOption string

type LastIndexChange string

type DiagnosticID uint

type AlarmIncludeLinkedZones bool

type ConfigModeOptions string

type ServiceTypeList string

type RoomCalibrationID string

type DailyIndexRefreshTime string

type SoftwareVersion string

type CustomerID string

type TimeGeneration uint

type SinkProtocolInfo string

type Browseable bool

type ConfigMode string

type DirectControlClientID string

type MicEnabled uint

type OutputFixed bool

type MusicSurroundLevel string

type RedirectURI string

type CachedOnly bool

type SortCriteria string

type ShareIndexLastError string

type AutoplayRoomUUID string

type ChannelMap string

type IsExpired bool

type QueueUpdateID uint

type SourceProtocolInfo string

type AutoplayUseVolume bool

type AccountUDN string

type AccountID string

type AlarmEnabled bool

type AlarmListVersion string

type TimeFormat string

type TrackDuration string

type IsZoneBridge bool

type BehindWifiExtender uint

type RoomDetectionChirpChannel uint

type AccountCredential string
