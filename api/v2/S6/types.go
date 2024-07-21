// CODE GENERATED, DO NOT EDIT

// Package S6 contains the implementation for the Sonos Play:5 (Model S6) services.
package S6

// Services struct contains service clients for different functionalities of a Sonos device.
type RoomCalibrationState int

type BehindWifiExtender uint

type MusicSurroundLevel string

type AlbumArtistDisplayOption string

type RadioFavoritesUpdateID uint

type Volume uint

type SystemUpdateID uint

type VolumeDB int

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type ISO8601Time string

type AlarmEnabled bool

type AbsoluteCounterPosition int

type Loudness bool

type SurroundMode string

type VariableName string

type Origin string

type TimeServer string

type MicEnabled uint

type RightVolume uint

type Code string

type Track uint

type MuseSessions string

type URIMetaData string

type AccountUID uint

type ThirdPartyMediaServersX string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type RecordQualityMode string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type HTFreq uint

type HTSatChanMapSet string

type WirelessMode uint

type SpeakerSize uint

type IncludeControllers bool

type MobileIPAndPort string

type CrossfadeMode bool

type Section uint

type Prefix string

type ChannelMapSet string

type SubPolarity string

type AlarmID uint

type TimeZoneIndex int

type ConnectionIDs string

type PossiblePlaybackStorageMedia string

type AutoplayIncludeLinkedZones bool

type SubGain string

type EnqueueAs bool

type ZoneGroupID string

type TimeStamp string

type RightLineInLevel int

type LastChange string

type SleepTimerState string

type MACAddress string

type VoiceUpdateID uint

type AudioInputName string

type SearchCapabilities string

type ZoneName string

type SubEnabled bool

type TrackMetaData string

type SourceState string

type ResetVolumeAfter bool

type Filter string

type Configuration string

type GroupCoordinatorIsLocal bool

type VolumeAdjustment int

type LIST_URI_AND_METADATA string

type RelativeCounterPosition int

type AlarmIDRunning uint

type AVTransportID int

type DialogLevel string

type AutoplayRoomUUID string

type NetsettingsUpdateID string

type SnoozeRunning bool

type ConnectionID int

type Browseable bool

type SinkProtocolInfo string

type TrackNumbersCSV string

type ThirdPartyHash string

type RoomDetectionChirpChannel uint

type RoomCalibrationEnabled bool

type TransportStatus string

type AVTransportURIMetaData string

type RadioLocationUpdateID uint

type TagValueList string

type NightMode bool

type Queue string

type LIST_URIMetaData string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type ShareIndexInProgress bool

type Orientation int

type AlarmRunSequence string

type AccountUDN string

type FavoritesUpdateID string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type GroupMute bool

type AudioDelayRightRear string

type AvailableSoftwareUpdate string

type Version string

type SortCapabilities string

type Seed string

type Bass int

type HardwareVersion string

type KeepGrouped bool

type QueuePolicy string

type ZoneGroupName string

type MobileDeviceUDN string

type TransportErrorHttpHeaders string

type LIST_URI string

type IsZoneBridge bool

type ChannelMap string

type ContainerUpdateIDs string

type RoomCalibrationCoefficients string

type AlarmVolume uint

const ()

type TransportPlaySpeed string

type StreamRestartState string

type ConfigMode string

type BootSeq uint

type ServiceListVersion string

type SubCrossover string

type DiagnosticID uint

type DateFormat string

type MemberList string

type RejoinGroup bool

type SortCriteria string

type InstanceID uint

type NumTracks uint

type SearchCriteria string

type SatRoomUUID string

type WifiEnabled bool

type SecureRegState uint

type ButtonState string

type VoiceConfigState uint

type AlarmProgramMetaData string

type TransportErrorHttpCode string

type SavedQueuesUpdateID string

type GroupVolumeChangeable bool

type AccountID string

type AuthorizationCode string

type LeftLineInLevel int

type QueueUpdateID uint

type ConfigModeOptions string

type EQValue int

type ProgramURI string

type SourceAreasUpdateID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TimeFormat string

type ValidPlayModes string

type AccountType uint

type ProtocolInfo string

type Result string

type Mute bool

type DirectControlClientID string

type ConfigModeState string

type Username string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type AccountMd string

type AlarmIncludeLinkedZones bool

type RecordMediumWriteStatus string

type SleepTimerGeneration uint

type MuseHouseholdId string

type Curated bool

type TransportErrorDescription string

type EnqueuedTransportURIMetaData string

type SupportsAudioClip bool

type LastIndexChange string

type LocalGroupUUID string

type OutputFixed bool

type AudioDelayLeftRear string

type PlayerID string

type TrackNumber uint

type SortOrder string

type ResumePlayback bool

type ShareListUpdateID string

type HdmiCecAvailable bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type VolumeAVTransportURI string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type AbsoluteTimePosition string

type AlarmLoggedStartTime string

type GroupVolume uint

type AccountPassword string

type UpdateFlags uint

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AvailableRoomCalibration string

type QueueOwnerContext string

type PossibleRecordQualityModes string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type URI string

type TrackList string

type UpdateID uint

type TimeZone string

type Icon string

type ObjectID string

type MID string

type ZoneGroupState string

type NumTracksChange int

type SupportsAudioIn bool

type HTAudioIn uint

type DID string

type UserIdHashCode string

type TimeGeneration uint

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type IsIdle bool

type AutoplayVolume uint

type QueueOwnerID string

type SurroundEnabled bool

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type HTBondedZoneCommitState uint

type ServiceId uint

type UpdateExtraOptions string

const ()

type RecordStorageMedium string

type VariableStringValue string

type StubsCreated string

type VLIState string

type RedirectURI string

type MemberID string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type PossibleRecordStorageMedia string

type SourceProtocolInfo string

type AutoplayUseVolume bool

type Speed string

type LineInConnected bool

type AlarmRunning bool

type RestartPending bool

type QueueID uint

type Treble int

type EQType string

type TimeZoneAutoAdjustDst bool

type Playing bool

type SerialNumber string

type AlarmRoomUUID string

type MediaDuration string

type RelativeTimePosition string

type MobileDeviceName string

type HouseholdID string

type SettingsReplicationState string

type CustomerID string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type NumberOfTracks uint

type TargetRoomName string

type WirelessLeafOnly bool

type TVConfigurationError bool

type BufferingResultCode int

type HeadphoneConnected bool

type DisplaySoftwareVersion string

type CopyrightInfo string

type ExtraInfo string

type RoomCalibrationAvailable bool

type UpdateItem string

type AlarmListVersion string

type Count uint

type SupportsOutputFixed bool

type TrackDuration string

type IsExpired bool

type OAuthDeviceID string

type IPAddress string

type PresetNameList string

type TransportErrorURI string

type TrackURI string

type FavoritePresetsUpdateID string

type Flags uint

type UpdateIDX uint

type ConnectionManager string

type ShareIndexLastError string

type AreasUpdateID string

type UserRadioUpdateID string

type Invisible bool

type RoomCalibrationID string

type TimeZoneInformation string

type TransportActions string

type DirectControlIsSuspended bool

type SavedQueueTitle string

type MoreInfo string

type RampTimeSeconds uint

type CachedOnly bool

type DailyIndexRefreshTime string

type RcsID int

type RDMEnabled bool

type ServiceTypeList string

type SessionId string

type ZonePlayerUUIDsInGroup string

type RoomDetectionPlayId uint

type SourceAreaIds string

type LeftVolume uint

type RoomCalibrationCalibrationMode string

type LastChangedPlayState string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type RoomDetectionDurationMilliseconds uint

type DirectControlAccountID string

type ServiceDescriptorList string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type SoftwareVersion string

type AudioDelay string

type AccountCredential string

type UpdateURL string

type TransportSettings string

type GroupID string

type AlarmState string

type VirtualLineInGroupID string

type SurroundLevel string

type AccountNickname string

type RecentlyPlayedUpdateID string

type AirPlayEnabled bool

type ChannelFreq uint

type HasConfiguredSSID bool

type AccountTier uint

type AVTransportURI string

type EnqueuedTransportURI string

type AutoplaySource string

type Index uint

type AlarmList string

type AlarmProgramURI string

type SeekTarget string
