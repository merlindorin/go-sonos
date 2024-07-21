// CODE GENERATED, DO NOT EDIT

// Package S5 contains the implementation for the Sonos Play:5 (Model S5) services.
package S5

// Services struct contains service clients for different functionalities of a Sonos device.
type Prefix string

type RDMEnabled bool

type CachedOnly bool

type AlarmRunning bool

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AlarmVolume uint

type ResumePlayback bool

type SavedQueuesUpdateID string

type Configuration string

type Invisible bool

type IsIdle bool

type AlarmProgramURI string

type LeftVolume uint

type RedirectURI string

type MobileDeviceName string

type DisplaySoftwareVersion string

type ObjectID string

type RelativeCounterPosition int

type GroupVolume uint

type DiagnosticID uint

type DateFormat string

const ()

type RecordStorageMedium string

type RestartPending bool

type UserRadioUpdateID string

type MicEnabled uint

type LIST_URI_AND_METADATA string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SpeakerSize uint

type LineInConnected bool

type AccountUDN string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type ServiceTypeList string

type ChannelMap string

type SubPolarity string

type MobileIPAndPort string

type SavedQueueTitle string

type SerialNumber string

type TransportStatus string

type SortCapabilities string

type SurroundMode string

type VariableStringValue string

type ThirdPartyHash string

type PlayerID string

type AudioInputName string

type AccountUID uint

type AccountCredential string

type AlarmRunSequence string

type TimeStamp string

type SortOrder string

type SearchCapabilities string

type PresetNameList string

type InstanceID uint

type DialogLevel string

type AccountMd string

type AccountTier uint

type SoftwareVersion string

type VolumeAVTransportURI string

type TimeZone string

type DID string

type ThirdPartyMediaServersX string

type SleepTimerGeneration uint

type TransportErrorDescription string

type SnoozeRunning bool

type LIST_URI string

type GroupID string

type TrackNumber uint

type TimeZoneInformation string

type StreamRestartState string

type KeepGrouped bool

type Track uint

type SeekTarget string

type AutoplayIncludeLinkedZones bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AlarmRoomUUID string

type AvailableSoftwareUpdate string

type Curated bool

type RoomCalibrationCoefficients string

type MuseHouseholdId string

type ServiceListVersion string

type ProtocolInfo string

type Result string

type AlbumArtistDisplayOption string

type ContainerUpdateIDs string

type VirtualLineInGroupID string

type ServiceDescriptorList string

type Loudness bool

type TransportActions string

type ZoneGroupID string

type LocalGroupUUID string

type Treble int

type CustomerID string

type MuseSessions string

type TransportErrorHttpCode string

const ()

type TransportPlaySpeed string

type RejoinGroup bool

type Flags uint

type ProgramURI string

type IncludeControllers bool

type Playing bool

type Index uint

type RadioFavoritesUpdateID uint

type AutoplayRoomUUID string

type GroupCoordinatorIsLocal bool

type IsExpired bool

type ConnectionID int

type Filter string

type FavoritePresetsUpdateID string

type ConfigModeOptions string

type Origin string

type AVTransportURIMetaData string

type TimeZoneIndex int

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type OutputFixed bool

type VariableName string

type AlarmProgramMetaData string

type TransportErrorHttpHeaders string

type LastIndexChange string

type RightVolume uint

type MemberID string

type AutoplayUseVolume bool

type SubGain string

type RoomCalibrationEnabled bool

type Version string

type NetsettingsUpdateID string

type TrackMetaData string

type LastChange string

type SecureRegState uint

type GroupMute bool

type ServiceId uint

type AuthorizationCode string

type MediaDuration string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type SatRoomUUID string

type VoiceUpdateID uint

type TimeServer string

type LeftLineInLevel int

type CrossfadeMode bool

type DirectControlClientID string

type HTBondedZoneCommitState uint

type TVConfigurationError bool

type Bass int

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type GroupVolumeChangeable bool

type Seed string

type UpdateID uint

type DirectControlIsSuspended bool

type VLIState string

type ConfigModeState string

type TimeZoneAutoAdjustDst bool

type ResetVolumeAfter bool

type SupportsAudioClip bool

type HTSatChanMapSet string

type AutoplayVolume uint

type MemberList string

type RampTimeSeconds uint

type ChannelMapSet string

type HTAudioIn uint

type SupportsOutputFixed bool

type AccountType uint

type UpdateExtraOptions string

type AreasUpdateID string

type URI string

type PossibleRecordQualityModes string

type Section uint

type AlarmIDRunning uint

type EnqueuedTransportURIMetaData string

type EnqueuedTransportURI string

type RcsID int

type FavoritesUpdateID string

type AlarmEnabled bool

type MoreInfo string

type QueueOwnerID string

type VolumeDB int

type AlarmState string

type AVTransportID int

type AirPlayEnabled bool

type HdmiCecAvailable bool

type ZonePlayerUUIDsInGroup string

type SleepTimerState string

type WirelessLeafOnly bool

type ButtonState string

type MusicSurroundLevel string

type ValidPlayModes string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type VoiceConfigState uint

type UpdateFlags uint

type TransportSettings string

type Queue string

type LIST_URIMetaData string

type SinkProtocolInfo string

type ChannelFreq uint

type BehindWifiExtender uint

type QueueID uint

type AlarmIncludeLinkedZones bool

type AlarmList string

type DailyIndexRefreshTime string

type SupportsAudioIn bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type OAuthDeviceID string

type AlarmID uint

type ZoneName string

type SortCriteria string

type NightMode bool

type StubsCreated string

type PossibleRecordStorageMedia string

type NumTracks uint

type AutoplaySource string

type Volume uint

type AccountID string

type SourceState string

type RecordQualityMode string

type AbsoluteCounterPosition int

type DirectControlAccountID string

type EnqueueAs bool

type WirelessMode uint

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type Browseable bool

type IsZoneBridge bool

type MID string

type EQValue int

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type AccountNickname string

type AccountPassword string

type HouseholdID string

type VolumeAdjustment int

type UpdateIDX uint

type ConnectionManager string

type RecordMediumWriteStatus string

type NumberOfTracks uint

type AvailableRoomCalibration string

type QueuePolicy string

type AudioDelayLeftRear string

type Speed string

type TimeGeneration uint

type Icon string

type AbsoluteTimePosition string

type URIMetaData string

type ConnectionIDs string

type WifiEnabled bool

type Mute bool

type RoomCalibrationAvailable bool

type AlarmListVersion string

type UpdateItem string

type ZoneGroupState string

type TransportErrorURI string

type AudioDelayRightRear string

type RecentlyPlayedUpdateID string

type TagValueList string

type ExtraInfo string

type QueueOwnerContext string

type HeadphoneConnected bool

type RightLineInLevel int

type SessionId string

type Username string

type HardwareVersion string

type AVTransportURI string

type SettingsReplicationState string

type SurroundLevel string

type SurroundEnabled bool

type TrackURI string

type NumTracksChange int

type QueueUpdateID uint

type RoomCalibrationID string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type SystemUpdateID uint

type ShareListUpdateID string

type LastChangedPlayState string

type RoomCalibrationState int

type IPAddress string

type TrackNumbersCSV string

type SubEnabled bool

type TimeFormat string

type EQType string

type AudioDelay string

type SubCrossover string

type RoomCalibrationCalibrationMode string

type SearchCriteria string

type BootSeq uint

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type UpdateURL string

type MobileDeviceUDN string

type HTFreq uint

type HasConfiguredSSID bool

type ConfigMode string

type AlarmLoggedStartTime string

type Orientation int

type UserIdHashCode string

type RelativeTimePosition string

type ShareIndexLastError string

type ZoneGroupName string

type PossiblePlaybackStorageMedia string

type TrackList string

type Count uint

type RadioLocationUpdateID uint

type MACAddress string

type CopyrightInfo string

type BufferingResultCode int

type Code string

type TrackDuration string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type SourceProtocolInfo string

type ShareIndexInProgress bool

type SourceAreaIds string

type SourceAreasUpdateID string

type ISO8601Time string
