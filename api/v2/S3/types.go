// CODE GENERATED, DO NOT EDIT

// Package S3 contains the implementation for the Sonos Play:3 (Model S3) services.
package S3

// Services struct contains service clients for different functionalities of a Sonos device.
const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type Version string

type CrossfadeMode bool

type NumberOfTracks uint

type AccountNickname string

type OAuthDeviceID string

type TransportStatus string

const ()

type RecordStorageMedium string

type AlbumArtistDisplayOption string

type ZonePlayerUUIDsInGroup string

type AlarmVolume uint

type TransportErrorDescription string

type SubEnabled bool

type AccountUDN string

type SortOrder string

type IPAddress string

type HdmiCecAvailable bool

type AutoplayRoomUUID string

type VoiceConfigState uint

type BufferingResultCode int

type RoomCalibrationCoefficients string

type RelativeCounterPosition int

type AlarmRunning bool

type SourceState string

type Icon string

type TransportSettings string

type SoftwareVersion string

type VolumeAdjustment int

type AlarmEnabled bool

type AVTransportURI string

type DirectControlClientID string

type InstanceID uint

const ()

type TransportPlaySpeed string

type Queue string

type QueuePolicy string

type EQType string

type DisplaySoftwareVersion string

type DID string

type Curated bool

type RoomCalibrationEnabled bool

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type Filter string

type SerialNumber string

type AccountUID uint

type RejoinGroup bool

type MoreInfo string

type HTFreq uint

type Treble int

type AlarmProgramURI string

type TimeZoneAutoAdjustDst bool

type MemberID string

type GroupID string

type HeadphoneConnected bool

type SubCrossover string

type SurroundLevel string

type StreamRestartState string

type KeepGrouped bool

type AccountTier uint

type URI string

type TrackNumber uint

type HardwareVersion string

type AuthorizationCode string

type AreasUpdateID string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type RoomDetectionPlayId uint

type PresetNameList string

type CachedOnly bool

type TransportErrorHttpCode string

type AirPlayEnabled bool

type ServiceId uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type MemberList string

type Count uint

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type RoomCalibrationState int

type LeftVolume uint

type TimeZoneIndex int

type AlarmListVersion string

type RestartPending bool

type SupportsAudioIn bool

type LIST_URI string

type MusicSurroundLevel string

type PossiblePlaybackStorageMedia string

type TrackURI string

type SupportsAudioClip bool

type Bass int

type EQValue int

type MediaDuration string

type SnoozeRunning bool

type UpdateID uint

type ChannelMapSet string

type TVConfigurationError bool

type SessionId string

type AudioDelay string

type RDMEnabled bool

type PossibleRecordStorageMedia string

type EnqueuedTransportURIMetaData string

type LIST_URIMetaData string

type Invisible bool

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AVTransportID int

type ThirdPartyMediaServersX string

type UpdateExtraOptions string

type HTSatChanMapSet string

type Orientation int

type Flags uint

type QueueID uint

type DateFormat string

type Track uint

type UserRadioUpdateID string

type SettingsReplicationState string

type RightVolume uint

type AccountPassword string

type UpdateIDX uint

type RelativeTimePosition string

type URIMetaData string

type MicEnabled uint

type DialogLevel string

type SurroundMode string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ExtraInfo string

type AutoplaySource string

type SupportsOutputFixed bool

type WirelessMode uint

type Username string

type TimeGeneration uint

type TimeFormat string

type TransportErrorHttpHeaders string

type SortCapabilities string

type TargetRoomName string

type SourceProtocolInfo string

type ConnectionManager string

type Prefix string

type Index uint

type TimeServer string

type SearchCapabilities string

type VirtualLineInGroupID string

type Volume uint

type AlarmState string

type ButtonState string

type ZoneGroupName string

type TimeStamp string

type MuseSessions string

type WifiEnabled bool

type VolumeDB int

type VolumeAVTransportURI string

type AvailableSoftwareUpdate string

type AlarmProgramMetaData string

type TransportActions string

type SearchCriteria string

type ShareIndexLastError string

type TrackMetaData string

type NumTracks uint

type QueueOwnerContext string

type OutputFixed bool

type TimeZone string

type IsIdle bool

type MACAddress string

type MobileIPAndPort string

type GroupVolumeChangeable bool

type Loudness bool

type AccountMd string

type Section uint

type AlarmLoggedStartTime string

type DirectControlAccountID string

type RadioFavoritesUpdateID uint

type AccountCredential string

type AlarmIncludeLinkedZones bool

type NumTracksChange int

type SavedQueueTitle string

type Configuration string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type CustomerID string

type MobileDeviceName string

type PossibleRecordQualityModes string

type FavoritePresetsUpdateID string

type RoomCalibrationCalibrationMode string

type IsExpired bool

type ProgramURI string

type HouseholdID string

type AutoplayIncludeLinkedZones bool

type ServiceListVersion string

type QueueOwnerID string

type ChannelMap string

type EnqueuedTransportURI string

type DirectControlIsSuspended bool

type BootSeq uint

type LIST_URI_AND_METADATA string

type SubPolarity string

type Origin string

type SleepTimerGeneration uint

type RadioLocationUpdateID uint

type RoomDetectionDurationMilliseconds uint

type SubGain string

type NightMode bool

type RoomCalibrationAvailable bool

type LastChange string

type EnqueueAs bool

type ResumePlayback bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type SeekTarget string

type ServiceTypeList string

type AudioDelayRightRear string

type SpeakerSize uint

type StubsCreated string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type ChannelFreq uint

type MID string

type RecordQualityMode string

type ProtocolInfo string

type LocalGroupUUID string

type UpdateItem string

type RecentlyPlayedUpdateID string

type ConfigModeOptions string

type SurroundEnabled bool

type AlarmRunSequence string

type FavoritesUpdateID string

type GroupVolume uint

type ThirdPartyHash string

type AccountType uint

type Speed string

type AlarmID uint

type RecordMediumWriteStatus string

type ShareListUpdateID string

type BehindWifiExtender uint

type ConnectionID int

type ContainerUpdateIDs string

type TrackNumbersCSV string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type AbsoluteTimePosition string

type VLIState string

type ZoneName string

type DailyIndexRefreshTime string

type SavedQueuesUpdateID string

type HTAudioIn uint

type SecureRegState uint

type ConfigMode string

type UpdateFlags uint

type TrackDuration string

type ResetVolumeAfter bool

type AvailableRoomCalibration string

type AutoplayVolume uint

type VariableStringValue string

type AccountID string

type SortCriteria string

type LastIndexChange string

type CopyrightInfo string

type GroupMute bool

type ZoneGroupState string

type PlayerID string

type Result string

type WirelessLeafOnly bool

type RoomDetectionChirpChannel uint

type AVTransportURIMetaData string

type HasConfiguredSSID bool

type RoomCalibrationID string

type SourceAreaIds string

type ISO8601Time string

type AlarmList string

type TransportErrorURI string

type RcsID int

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type TrackList string

type LastChangedPlayState string

type Code string

type IsZoneBridge bool

type ConfigModeState string

type Seed string

type UpdateURL string

type AlarmIDRunning uint

type ObjectID string

type Mute bool

type RampTimeSeconds uint

type Browseable bool

type AutoplayUseVolume bool

type ZoneGroupID string

type TimeZoneInformation string

type ConnectionIDs string

type NetsettingsUpdateID string

type VoiceUpdateID uint

type MuseHouseholdId string

type DiagnosticID uint

type MobileDeviceUDN string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AbsoluteCounterPosition int

type ShareIndexInProgress bool

type VariableName string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type IncludeControllers bool

type AlarmRoomUUID string

type ValidPlayModes string

type SinkProtocolInfo string

type TagValueList string

type AudioDelayLeftRear string

type UserIdHashCode string

type SourceAreasUpdateID string

type SleepTimerState string

type SystemUpdateID uint

type ServiceDescriptorList string

type RedirectURI string

type QueueUpdateID uint

type HTBondedZoneCommitState uint

type SatRoomUUID string

type GroupCoordinatorIsLocal bool
