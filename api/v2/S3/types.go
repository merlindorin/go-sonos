// CODE GENERATED, DO NOT EDIT

// Package S3 contains the implementation for the Sonos Play:3 (Model S3) services.
package S3

// Services struct contains service clients for different functionalities of a Sonos device.
type AvailableSoftwareUpdate string

type MobileDeviceUDN string

type TrackList string

type Count uint

type SerialNumber string

type HardwareVersion string

type LocalGroupUUID string

type TrackURI string

type RestartPending bool

type ConnectionIDs string

type MoreInfo string

type ServiceDescriptorList string

type DirectControlAccountID string

type AVTransportID int

type RecentlyPlayedUpdateID string

type Flags uint

type SubEnabled bool

type AccountPassword string

type NetsettingsUpdateID string

type AlarmRoomUUID string

type LIST_URI string

type KeepGrouped bool

type SourceAreaIds string

type QueueOwnerContext string

type AccountUDN string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

const ()

type TransportPlaySpeed string

type TrackMetaData string

type VolumeAdjustment int

type SortOrder string

type SearchCapabilities string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type DialogLevel string

type SinkProtocolInfo string

type RadioFavoritesUpdateID uint

type WirelessLeafOnly bool

type RoomDetectionDurationMilliseconds uint

type CustomerID string

type URI string

type RejoinGroup bool

type SourceProtocolInfo string

type SettingsReplicationState string

type RcsID int

type GroupVolumeChangeable bool

type HeadphoneConnected bool

type AlarmEnabled bool

type TimeZoneInformation string

type DailyIndexRefreshTime string

type EnqueueAs bool

type DID string

type CachedOnly bool

type AlarmProgramURI string

type StreamRestartState string

type SortCapabilities string

type ChannelFreq uint

type AlarmList string

type NumberOfTracks uint

type ShareIndexInProgress bool

type TrackNumbersCSV string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type UpdateItem string

type UpdateURL string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TransportErrorURI string

type TransportErrorHttpCode string

type TransportSettings string

type RecordQualityMode string

type TargetRoomName string

type SecureRegState uint

type UpdateExtraOptions string

type AlarmIDRunning uint

type ConfigModeState string

type VariableStringValue string

type ThirdPartyHash string

type ContainerUpdateIDs string

type OAuthDeviceID string

type ZoneGroupID string

type TransportErrorHttpHeaders string

type Section uint

type MemberID string

type SearchCriteria string

type SubCrossover string

type SurroundLevel string

type RoomCalibrationEnabled bool

type AccountUID uint

type ISO8601Time string

type TimeZoneIndex int

type FavoritesUpdateID string

type VoiceConfigState uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AccountID string

type UserIdHashCode string

type AbsoluteTimePosition string

type DirectControlIsSuspended bool

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type SupportsAudioClip bool

type RoomCalibrationCoefficients string

type AccountNickname string

type ThirdPartyMediaServersX string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type MID string

type Loudness bool

type ProgramURI string

type ShareIndexLastError string

type AutoplaySource string

type VirtualLineInGroupID string

type MobileDeviceName string

type EnqueuedTransportURIMetaData string

type HasConfiguredSSID bool

type MuseHouseholdId string

type ZonePlayerUUIDsInGroup string

type ExtraInfo string

type LeftVolume uint

type EQType string

type RDMEnabled bool

type RelativeCounterPosition int

type LastChange string

type MemberList string

type ConnectionManager string

type QueueUpdateID uint

type DisplaySoftwareVersion string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type RoomCalibrationAvailable bool

type WifiEnabled bool

type SupportsOutputFixed bool

type AccountCredential string

type Username string

type ChannelMap string

type SleepTimerGeneration uint

type AvailableRoomCalibration string

type IPAddress string

type VolumeAVTransportURI string

type SourceState string

type Filter string

type ConfigModeOptions string

type ServiceListVersion string

type Prefix string

type ChannelMapSet string

type RoomDetectionPlayId uint

type GroupCoordinatorIsLocal bool

type UpdateIDX uint

const ()

type RecordStorageMedium string

type UserRadioUpdateID string

type Curated bool

type AccountMd string

type AlbumArtistDisplayOption string

type NightMode bool

type AlarmProgramMetaData string

type SavedQueuesUpdateID string

type SatRoomUUID string

type OutputFixed bool

type ServiceId uint

type DiagnosticID uint

type QueuePolicy string

type EQValue int

type SurroundMode string

type Speed string

type TimeGeneration uint

type MediaDuration string

type AlarmRunning bool

type HTSatChanMapSet string

type SortCriteria string

type ServiceTypeList string

type Seed string

type PresetNameList string

type PossiblePlaybackStorageMedia string

type GroupID string

type ResetVolumeAfter bool

type AlarmState string

type RedirectURI string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type RoomCalibrationState int

type ButtonState string

type TimeStamp string

type SpeakerSize uint

type ZoneGroupName string

type LIST_URI_AND_METADATA string

type AreasUpdateID string

type DirectControlClientID string

type CopyrightInfo string

type AudioDelayLeftRear string

type AuthorizationCode string

type ZoneGroupState string

type ConfigMode string

type BufferingResultCode int

type QueueOwnerID string

type AudioDelay string

type TransportActions string

type SystemUpdateID uint

type AutoplayIncludeLinkedZones bool

type IncludeControllers bool

type TimeFormat string

type VLIState string

type IsExpired bool

type Version string

type TimeServer string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type TagValueList string

type Orientation int

type PossibleRecordStorageMedia string

type Invisible bool

type MusicSurroundLevel string

type PlayerID string

type NumTracksChange int

type BootSeq uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type Origin string

type AlarmLoggedStartTime string

type LastIndexChange string

type AirPlayEnabled bool

type MicEnabled uint

type AlarmListVersion string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type AutoplayVolume uint

type MobileIPAndPort string

type Mute bool

type Bass int

type Treble int

type RampTimeSeconds uint

type TransportStatus string

type RelativeTimePosition string

type Queue string

type UpdateID uint

type RoomCalibrationID string

type AccountTier uint

type Result string

type ZoneName string

type GroupVolume uint

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type SavedQueueTitle string

type AutoplayRoomUUID string

type TimeZone string

type Track uint

type LIST_URIMetaData string

type TrackNumber uint

type PossibleRecordQualityModes string

type ShareListUpdateID string

type TVConfigurationError bool

type ResumePlayback bool

type AutoplayUseVolume bool

type GroupMute bool

type DateFormat string

type SleepTimerState string

type Icon string

type Configuration string

type URIMetaData string

type HdmiCecAvailable bool

type AccountType uint

type WirelessMode uint

type CrossfadeMode bool

type AVTransportURI string

type ObjectID string

type ConnectionID int

type VariableName string

type ProtocolInfo string

type Index uint

type IsIdle bool

type Code string

type TrackDuration string

type AbsoluteCounterPosition int

type ValidPlayModes string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type RoomCalibrationCalibrationMode string

type SubGain string

type AlarmIncludeLinkedZones bool

type MuseSessions string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type QueueID uint

type InstanceID uint

type IsZoneBridge bool

type AlarmRunSequence string

type SourceAreasUpdateID string

type AlarmID uint

type TransportErrorDescription string

type MACAddress string

type SessionId string

type Volume uint

type VolumeDB int

type SeekTarget string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type SupportsAudioIn bool

type BehindWifiExtender uint

type AlarmVolume uint

type NumTracks uint

type FavoritePresetsUpdateID string

type HTBondedZoneCommitState uint

type RightVolume uint

type SubPolarity string

type SurroundEnabled bool

type RadioLocationUpdateID uint

type LastChangedPlayState string

type RoomDetectionChirpChannel uint

type RoomDetectionChirpIfPlayingSwappableAudio bool

type TimeZoneAutoAdjustDst bool

type StubsCreated string

type RecordMediumWriteStatus string

type SnoozeRunning bool

type HTAudioIn uint

type VoiceUpdateID uint

type HTFreq uint

type SoftwareVersion string

type AudioDelayRightRear string

type UpdateFlags uint

type AVTransportURIMetaData string

type EnqueuedTransportURI string

type Browseable bool

type HouseholdID string
