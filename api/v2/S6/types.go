// CODE GENERATED, DO NOT EDIT

// Package S6 contains the implementation for the Sonos Play:5 (Model S6) services.
package S6

// Services struct contains service clients for different functionalities of a Sonos device.
type RightLineInLevel int

type AlarmState string

type GroupMute bool

type DialogLevel string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type TransportErrorHttpHeaders string

type ChannelFreq uint

type GroupVolumeChangeable bool

type SortCriteria string

type Browseable bool

type RoomCalibrationAvailable bool

type Track uint

type AlarmIDRunning uint

type TargetRoomName string

type VolumeDB int

type RoomCalibrationID string

type DailyIndexRefreshTime string

type SavedQueuesUpdateID string

type HdmiCecAvailable bool

type TrackNumbersCSV string

type ServiceId uint

type UpdateExtraOptions string

type SnoozeRunning bool

type SleepTimerState string

type SourceProtocolInfo string

type SerialNumber string

type RedirectURI string

type DiagnosticID uint

type AudioInputName string

type IsZoneBridge bool

type SupportsAudioClip bool

type RoomCalibrationCalibrationMode string

type Count uint

type ServiceDescriptorList string

type Bass int

type AccountMd string

type AlarmProgramURI string

type TimeGeneration uint

type Icon string

type Prefix string

type AuthorizationCode string

type HasConfiguredSSID bool

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type Version string

type TimeZoneAutoAdjustDst bool

type TransportActions string

type Queue string

type SystemUpdateID uint

type AutoplayIncludeLinkedZones bool

type GroupVolume uint

type CachedOnly bool

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type Filter string

type ZoneName string

type Orientation int

type AutoplayUseVolume bool

type SurroundEnabled bool

type TransportErrorURI string

type TrackNumber uint

type Configuration string

type ServiceListVersion string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ConfigMode string

type SurroundLevel string

type ZoneGroupName string

type UpdateItem string

type IsIdle bool

type HardwareVersion string

type Code string

type MuseHouseholdId string

type VariableName string

type ResetVolumeAfter bool

type DID string

type Treble int

type Loudness bool

type SinkProtocolInfo string

type SortCapabilities string

type BootSeq uint

type SpeakerSize uint

type TrackURI string

type URIMetaData string

type Result string

type VirtualLineInGroupID string

type QueueOwnerID string

type AudioDelayRightRear string

type TimeFormat string

type AlbumArtistDisplayOption string

type ContainerUpdateIDs string

type ShareIndexLastError string

type BufferingResultCode int

type MobileDeviceUDN string

type EnqueuedTransportURIMetaData string

type HTBondedZoneCommitState uint

type VoiceUpdateID uint

type ThirdPartyHash string

type PresetNameList string

type AlarmList string

const ()

type RecordStorageMedium string

type SoftwareVersion string

type ButtonState string

type RoomDetectionChirpChannel uint

type AccountNickname string

type IsExpired bool

type ZonePlayerUUIDsInGroup string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type LeftLineInLevel int

type MuseSessions string

type ConnectionManager string

type MobileIPAndPort string

type AlarmListVersion string

type AlarmRunning bool

type MemberList string

type QueueID uint

type SourceAreaIds string

type SurroundMode string

type AlarmRunSequence string

type TransportStatus string

type EnqueueAs bool

type SupportsOutputFixed bool

type NightMode bool

type AlarmRoomUUID string

type RoomCalibrationEnabled bool

type AccountTier uint

type HouseholdID string

type Curated bool

type ProgramURI string

type Username string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type AccountPassword string

type UpdateIDX uint

type TrackMetaData string

type SearchCriteria string

type VoiceConfigState uint

type SessionId string

type Speed string

type StubsCreated string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type TagValueList string

type GroupCoordinatorIsLocal bool

type MusicSurroundLevel string

type RejoinGroup bool

type TrackList string

type FavoritesUpdateID string

type AutoplayVolume uint

type TimeZoneInformation string

type SleepTimerGeneration uint

type URI string

type SavedQueueTitle string

type WifiEnabled bool

type MID string

type AccountUID uint

type AreasUpdateID string

type AbsoluteTimePosition string

type DirectControlIsSuspended bool

type RcsID int

type AccountCredential string

type PossibleRecordQualityModes string

type UserIdHashCode string

type AudioDelayLeftRear string

type ShareIndexInProgress bool

type MACAddress string

type ServiceTypeList string

type LIST_URI_AND_METADATA string

type StreamRestartState string

type QueueUpdateID uint

type ConnectionIDs string

type DisplaySoftwareVersion string

type AlarmEnabled bool

type RecordMediumWriteStatus string

type RelativeCounterPosition int

type LastChange string

type SecureRegState uint

type HTFreq uint

type LastChangedPlayState string

type EQType string

type UpdateFlags uint

type AvailableSoftwareUpdate string

type NumTracks uint

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type ChannelMap string

type MobileDeviceName string

type HTAudioIn uint

type KeepGrouped bool

type Mute bool

type ZoneGroupState string

type DateFormat string

type OAuthDeviceID string

type AVTransportID int

type Invisible bool

type AccountUDN string

type RDMEnabled bool

type InstanceID uint

type SupportsAudioIn bool

type BehindWifiExtender uint

type SubGain string

type ValidPlayModes string

type DirectControlAccountID string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type CrossfadeMode bool

type RecordQualityMode string

type RestartPending bool

type IncludeControllers bool

type AVTransportURIMetaData string

type VLIState string

type Seed string

type HeadphoneConnected bool

type TimeServer string

type OutputFixed bool

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type RoomCalibrationCoefficients string

type AlarmVolume uint

type Section uint

type AVTransportURI string

type HTSatChanMapSet string

type TransportErrorDescription string

type LIST_URI string

type GroupID string

type CopyrightInfo string

type RightVolume uint

type AudioDelay string

type PossibleRecordStorageMedia string

type DirectControlClientID string

type LIST_URIMetaData string

type ConfigModeState string

type ThirdPartyMediaServersX string

type ChannelMapSet string

type RoomCalibrationState int

type AutoplaySource string

type SubEnabled bool

type QueuePolicy string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AlarmProgramMetaData string

type ResumePlayback bool

type AvailableRoomCalibration string

type RoomDetectionPlayId uint

type ZoneGroupID string

type LineInConnected bool

type TVConfigurationError bool

type Volume uint

type EQValue int

type UpdateID uint

type ExtraInfo string

type VariableStringValue string

type Playing bool

type AlarmLoggedStartTime string

type EnqueuedTransportURI string

type SeekTarget string

type ISO8601Time string

type TransportSettings string

type MediaDuration string

type RecentlyPlayedUpdateID string

const ()

type TransportPlaySpeed string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type ConnectionID int

type LastIndexChange string

type WirelessLeafOnly bool

type SubPolarity string

type WirelessMode uint

type RoomDetectionChirpIfPlayingSwappableAudio bool

type LocalGroupUUID string

type VolumeAVTransportURI string

type PlayerID string

type VolumeAdjustment int

type QueueOwnerContext string

type AccountID string

type TimeStamp string

type TransportErrorHttpCode string

type RelativeTimePosition string

type AbsoluteCounterPosition int

type SatRoomUUID string

type Flags uint

type LeftVolume uint

type RampTimeSeconds uint

type AlarmID uint

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type PossiblePlaybackStorageMedia string

type Index uint

type MemberID string

type RadioFavoritesUpdateID uint

type FavoritePresetsUpdateID string

type SourceAreasUpdateID string

type ObjectID string

type NumTracksChange int

type ProtocolInfo string

type IPAddress string

type NetsettingsUpdateID string

type TimeZone string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type AutoplayRoomUUID string

type Origin string

type SortOrder string

type ShareListUpdateID string

type SettingsReplicationState string

type CustomerID string

type SubCrossover string

type AccountType uint

type NumberOfTracks uint

type TrackDuration string

type SearchCapabilities string

type MoreInfo string

type AlarmIncludeLinkedZones bool

type TimeZoneIndex int

type RadioLocationUpdateID uint

type UpdateURL string

type MicEnabled uint

type RoomDetectionDurationMilliseconds uint

type SourceState string

type UserRadioUpdateID string

type AirPlayEnabled bool

type ConfigModeOptions string
