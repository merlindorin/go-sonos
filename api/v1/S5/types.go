// CODE GENERATED, DO NOT EDIT

// Package S5 contains the implementation for the Sonos Play:5 (Model S5) services.
package S5

// Services struct contains service clients for different functionalities of a Sonos device.
type AlarmListVersion string

type PossiblePlaybackStorageMedia string

type EnqueuedTransportURIMetaData string

type GroupVolumeChangeable bool

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AccountPassword string

type LIST_URI_AND_METADATA string

type TrackDuration string

type QueueUpdateID uint

type RadioLocationUpdateID uint

type HTSatChanMapSet string

type RoomCalibrationAvailable bool

type IsExpired bool

type TimeGeneration uint

type AbsoluteTimePosition string

type ConnectionID int

type Username string

type TimeServer string

type AccountTier uint

type IncludeControllers bool

type AlbumArtistDisplayOption string

type SourceState string

type QueueOwnerID string

type TransportErrorURI string

type ContainerUpdateIDs string

type Section uint

type AlarmIDRunning uint

type DirectControlAccountID string

type Prefix string

type Browseable bool

type AudioDelayLeftRear string

type Origin string

type Track uint

type NetsettingsUpdateID string

type GroupID string

type UpdateID uint

type AccountCredential string

type VLIState string

type ThirdPartyHash string

type TimeFormat string

type RejoinGroup bool

type TransportSettings string

type ShareIndexInProgress bool

type IsZoneBridge bool

type AlarmProgramMetaData string

type SavedQueueTitle string

type BootSeq uint

type SessionId string

type ChannelMap string

type VoiceUpdateID uint

type AVTransportURI string

type ResumePlayback bool

type AlarmState string

type Bass int

type Queue string

type Speed string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type QueueOwnerContext string

type UpdateItem string

type TrackNumber uint

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type EnqueuedTransportURI string

type InstanceID uint

type FavoritesUpdateID string

type SupportsAudioClip bool

type SerialNumber string

type TimeZoneAutoAdjustDst bool

type ConfigModeState string

type RightVolume uint

type HTAudioIn uint

type AuthorizationCode string

type MuseHouseholdId string

type SurroundMode string

type SupportsOutputFixed bool

type CustomerID string

type BehindWifiExtender uint

type Playing bool

type SeekTarget string

type GroupCoordinatorIsLocal bool

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type RedirectURI string

type LineInConnected bool

type GroupMute bool

type AvailableSoftwareUpdate string

type TransportErrorHttpCode string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type CachedOnly bool

type ProtocolInfo string

type ZoneName string

type ExtraInfo string

type MicEnabled uint

type ServiceTypeList string

type NightMode bool

type TransportStatus string

type AutoplaySource string

type RoomCalibrationEnabled bool

type AccountUID uint

type LeftLineInLevel int

type FavoritePresetsUpdateID string

type Orientation int

type ProgramURI string

type AudioDelay string

type UpdateFlags uint

type DirectControlIsSuspended bool

type VoiceConfigState uint

type SurroundEnabled bool

type SecureRegState uint

type TransportErrorDescription string

type IPAddress string

type SourceAreaIds string

type QueueID uint

type Volume uint

type SubCrossover string

type AccountMd string

type AlarmID uint

type TrackURI string

type Result string

type OutputFixed bool

type DailyIndexRefreshTime string

type Filter string

type SavedQueuesUpdateID string

type RecentlyPlayedUpdateID string

type BufferingResultCode int

type MID string

type HeadphoneConnected bool

type TransportErrorHttpHeaders string

type NumTracks uint

type SinkProtocolInfo string

type Curated bool

type RoomCalibrationCalibrationMode string

type AccountNickname string

type UpdateIDX uint

type Icon string

type DirectControlClientID string

type SupportsAudioIn bool

type LastChangedPlayState string

type SleepTimerGeneration uint

type ValidPlayModes string

type SearchCriteria string

type HardwareVersion string

type SubPolarity string

type RoomCalibrationID string

type VariableStringValue string

type AudioInputName string

type PlayerID string

type SourceProtocolInfo string

type AVTransportID int

type Index uint

type AvailableRoomCalibration string

type DID string

type ZoneGroupID string

type AlarmVolume uint

type MuseSessions string

type VolumeAVTransportURI string

type Loudness bool

type StubsCreated string

type AlarmEnabled bool

type AbsoluteCounterPosition int

type SnoozeRunning bool

type MobileIPAndPort string

type PossibleRecordStorageMedia string

type MediaDuration string

type WirelessLeafOnly bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type QueuePolicy string

type DialogLevel string

type UserIdHashCode string

type NumberOfTracks uint

type LIST_URI string

type SearchCapabilities string

type HouseholdID string

type Seed string

type MobileDeviceName string

type MemberID string

type AlarmRunning bool

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type RcsID int

type SortCapabilities string

type VirtualLineInGroupID string

type Treble int

type ObjectID string

type TransportActions string

type TrackList string

type Flags uint

type TimeZoneIndex int

type CopyrightInfo string

type TVConfigurationError bool

type MusicSurroundLevel string

type ZoneGroupName string

type ConnectionIDs string

type Configuration string

type UpdateURL string

type CrossfadeMode bool

type ChannelMapSet string

type WifiEnabled bool

type ZoneGroupState string

type DiagnosticID uint

type AreasUpdateID string

type TrackMetaData string

type TimeZoneInformation string

const ()

type RecordStorageMedium string

type RelativeCounterPosition int

type SleepTimerState string

type SortCriteria string

type RoomCalibrationState int

type AutoplayVolume uint

type AlarmIncludeLinkedZones bool

type ServiceListVersion string

type AudioDelayRightRear string

type ButtonState string

type TimeStamp string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type SatRoomUUID string

type AlarmList string

type RecordQualityMode string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type AlarmRoomUUID string

type EQValue int

type RampTimeSeconds uint

type VariableName string

type AccountType uint

type LastIndexChange string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type Invisible bool

type MACAddress string

type AutoplayIncludeLinkedZones bool

type SubGain string

type PresetNameList string

type AccountUDN string

type ISO8601Time string

type Version string

type SubEnabled bool

type HTFreq uint

type TagValueList string

type DisplaySoftwareVersion string

type ConnectionManager string

type NumTracksChange int

type ResetVolumeAfter bool

type VolumeAdjustment int

type SurroundLevel string

type AlarmRunSequence string

type MemberList string

type SoftwareVersion string

type AutoplayRoomUUID string

type Code string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type SourceAreasUpdateID string

type SystemUpdateID uint

type AirPlayEnabled bool

type LocalGroupUUID string

type UpdateExtraOptions string

type DateFormat string

const ()

type TransportPlaySpeed string

type RecordMediumWriteStatus string

type RestartPending bool

type LastChange string

type RadioFavoritesUpdateID uint

type TrackNumbersCSV string

type SpeakerSize uint

type AlarmProgramURI string

type StreamRestartState string

type HTBondedZoneCommitState uint

type HasConfiguredSSID bool

type LIST_URIMetaData string

type EnqueueAs bool

type AutoplayUseVolume bool

type PossibleRecordQualityModes string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type UserRadioUpdateID string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type RelativeTimePosition string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type URI string

type ShareListUpdateID string

type MoreInfo string

type OAuthDeviceID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type EQType string

type MobileDeviceUDN string

type LeftVolume uint

type ConfigMode string

type RDMEnabled bool

type Count uint

type WirelessMode uint

type GroupVolume uint

type RoomCalibrationCoefficients string

type ThirdPartyMediaServersX string

type AlarmLoggedStartTime string

type AVTransportURIMetaData string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type ChannelFreq uint

type KeepGrouped bool

type ServiceDescriptorList string

type ServiceId uint

type Mute bool

type RightLineInLevel int

type AccountID string

type URIMetaData string

type ShareIndexLastError string

type IsIdle bool

type HdmiCecAvailable bool

type ConfigModeOptions string

type VolumeDB int

type TimeZone string

type SettingsReplicationState string

type ZonePlayerUUIDsInGroup string

type SortOrder string
