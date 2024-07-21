// CODE GENERATED, DO NOT EDIT

// Package S33 contains the implementation for the SYMFONISK Bookshelf (Model S33) services.
package S33

// Services struct contains service clients for different functionalities of a Sonos device.
type DirectControlClientID string

type VLIState string

type SettingsReplicationState string

type MID string

type MobileDeviceName string

type AbsoluteCounterPosition int

type AVTransportID int

type AirPlayEnabled bool

type VariableName string

type VariableStringValue string

type EnqueuedTransportURIMetaData string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type HTSatChanMapSet string

type ThirdPartyMediaServersX string

type TimeZoneAutoAdjustDst bool

type TimeZone string

type EnqueueAs bool

type FavoritesUpdateID string

type SupportsAudioClip bool

type LIST_URIMetaData string

type SupportsOutputFixed bool

type AlarmID uint

type AlarmList string

type AlarmRunning bool

type SleepTimerState string

type SortOrder string

type RadioFavoritesUpdateID uint

type Orientation int

type RoomDetectionChirpChannel uint

type Queue string

type AlbumArtistDisplayOption string

type IsIdle bool

type ChannelFreq uint

type SecureRegState uint

type HTForwardEnabled bool

type ServiceListVersion string

type TrackNumbersCSV string

type UpdateFlags uint

type AlarmProgramURI string

type NumTracksChange int

type HdmiCecAvailable bool

type VoiceConfigState uint

type SurroundEnabled bool

type AlarmProgramMetaData string

type RightVolume uint

type SpeakerSize uint

type DiagnosticID uint

type Origin string

type NumberOfTracks uint

type HouseholdID string

type CopyrightInfo string

type ExtraInfo string

type Username string

type SubCrossover string

type TransportStatus string

type TransportErrorHttpHeaders string

type RecordQualityMode string

type PlayerID string

type Volume uint

type DialogLevel string

type SubPolarity string

type Section uint

type EnqueuedTransportURI string

type FavoritePresetsUpdateID string

type LastChangedPlayState string

type ConfigMode string

type AccountNickname string

type UpdateExtraOptions string

type SourceProtocolInfo string

type SoftwareVersion string

type RoomCalibrationAvailable bool

type Browseable bool

type DisplaySoftwareVersion string

type GroupMute bool

type EQValue int

type SleepTimerGeneration uint

type ServiceTypeList string

type RampTimeSeconds uint

type PresetNameList string

type RoomCalibrationID string

type Speed string

type AlarmVolume uint

type ResetVolumeAfter bool

type ShareIndexInProgress bool

type TVConfigurationError bool

type GroupVolumeChangeable bool

type Curated bool

type UserIdHashCode string

type SystemUpdateID uint

type Invisible bool

type VoiceUpdateID uint

type IncludeControllers bool

type TimeGeneration uint

type TransportErrorURI string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type AVTransportURI string

type SavedQueueTitle string

type ConnectionIDs string

type UpdateID uint

type AutoplaySource string

type Loudness bool

type CustomerID string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

const (
	MasterMuteChannel MuteChannel = "Master"
	LfMuteChannel     MuteChannel = "LF"
	RfMuteChannel     MuteChannel = "RF"
)

type MuteChannel string

type AvailableSoftwareUpdate string

type MobileIPAndPort string

type InstanceID uint

type GroupID string

type ConnectionID int

type ChannelMapSet string

type VirtualLineInGroupID string

type ServiceDescriptorList string

type RDMEnabled bool

type AccountTier uint

type RelativeCounterPosition int

type AlarmLoggedStartTime string

type SupportsAudioIn bool

type SatRoomUUID string

type QueueOwnerContext string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type AutoplayIncludeLinkedZones bool

type Treble int

type AccountUDN string

type TransportErrorHttpCode string

type SnoozeRunning bool

type SinkProtocolInfo string

type LastIndexChange string

type ServiceId uint

type Bass int

type UpdateIDX uint

const ()

type TransportPlaySpeed string

type RelativeTimePosition string

type RejoinGroup bool

type RecentlyPlayedUpdateID string

type QueuePolicy string

type RedirectURI string

type ZonePlayerUUIDsInGroup string

type RestartPending bool

type LIST_URI string

type AudioDelayRightRear string

type BufferingResultCode int

type MobileDeviceUDN string

type AlarmEnabled bool

type TransportErrorDescription string

type HTBondedZoneCommitState uint

type ISO8601Time string

type ValidPlayModes string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type AccountType uint

type AlarmRunSequence string

type TrackURI string

type TransportSettings string

type Filter string

type TimeZoneInformation string

type AlarmListVersion string

type AlarmState string

type AutoplayUseVolume bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type KeepGrouped bool

type RoomCalibrationCoefficients string

type AccountID string

type MemberID string

type GroupCoordinatorIsLocal bool

type ZoneGroupID string

type UpdateItem string

type SourceState string

type StreamRestartState string

type SessionId string

type Seed string

type SubGain string

type AccountUID uint

type IsExpired bool

type StubsCreated string

type RadioLocationUpdateID uint

type HasConfiguredSSID bool

type AccountCredential string

type OutputFixed bool

type AccountPassword string

type LastChange string

type SeekTarget string

type RoomDetectionPlayId uint

type GroupVolume uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type MuseHouseholdId string

type SourceAreasUpdateID string

type CrossfadeMode bool

type ProtocolInfo string

type HardwareVersion string

type WifiEnabled bool

type Mute bool

type PossiblePlaybackStorageMedia string

type TrackMetaData string

type URI string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type TimeServer string

type TrackDuration string

type TrackNumber uint

type Index uint

type LeftVolume uint

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type MuseSessions string

type ConnectionManager string

type BehindWifiExtender uint

type RoomDetectionDurationMilliseconds uint

type AvailableRoomCalibration string

type ZoneGroupName string

type AreasUpdateID string

type AlarmIDRunning uint

type WirelessMode uint

const ()

type RecordStorageMedium string

type TrackList string

type TagValueList string

type SearchCapabilities string

type SortCapabilities string

type AudioDelayLeftRear string

type ZoneGroupState string

type CachedOnly bool

type MemberList string

type SearchCriteria string

type UserRadioUpdateID string

type Configuration string

type QueueOwnerID string

type NightMode bool

type Version string

type AlarmIncludeLinkedZones bool

type RcsID int

type Result string

type Count uint

type TargetRoomName string

type IsZoneBridge bool

type AutoplayRoomUUID string

type OAuthDeviceID string

type Track uint

type DirectControlIsSuspended bool

type HeadphoneConnected bool

type UpdateURL string

type TimeZoneIndex int

type RecordMediumWriteStatus string

type WirelessLeafOnly bool

type AVTransportURIMetaData string

type ProgramURI string

type PossibleRecordStorageMedia string

type ConfigModeOptions string

type ConfigModeState string

type AudioDelay string

type AlarmRoomUUID string

type ObjectID string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type LIST_URI_AND_METADATA string

type RoomCalibrationEnabled bool

type AuthorizationCode string

type Prefix string

type RoomCalibrationState int

type LocalGroupUUID string

type MusicSurroundLevel string

type NumTracks uint

type ShareListUpdateID string

type HTAudioIn uint

type ChannelMap string

type ResumePlayback bool

type Icon string

type MoreInfo string

type SerialNumber string

type ButtonState string

type RoomCalibrationCalibrationMode string

type TransportActions string

type DirectControlAccountID string

type IPAddress string

type EthLink bool

type EQType string

type DailyIndexRefreshTime string

type URIMetaData string

type QueueUpdateID uint

type AutoplayVolume uint

type MediaDuration string

type ContainerUpdateIDs string

type SubEnabled bool

type SurroundLevel string

type SortCriteria string

type ZoneName string

type Flags uint

type DateFormat string

type AbsoluteTimePosition string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type VolumeAVTransportURI string

type Code string

type DID string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type TimeFormat string

type VolumeAdjustment int

type VolumeDB int

type ThirdPartyHash string

type TimeStamp string

type PossibleRecordQualityModes string

type SavedQueuesUpdateID string

type HTFreq uint

type SourceAreaIds string

type AccountMd string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type MACAddress string

type MicEnabled uint

type SurroundMode string

type QueueID uint

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type ShareIndexLastError string

type BootSeq uint

type NetsettingsUpdateID string
