// CODE GENERATED, DO NOT EDIT

// Package S27 contains the implementation for the Sonos Roam (Model S27) services.
package S27

// Services struct contains service clients for different functionalities of a Sonos device.
type OutputFixed bool

type Queue string

type UpdateFlags uint

type AutoplayVolume uint

type RoomCalibrationEnabled bool

type Speed string

type UpdateItem string

type LIST_URIMetaData string

type Browseable bool

type ZoneName string

type HeadphoneConnected bool

type RedirectURI string

type RecordQualityMode string

type DirectControlAccountID string

type ServiceDescriptorList string

type ResetVolumeAfter bool

type RampTimeSeconds uint

type ZoneGroupState string

type ZoneGroupName string

type SortCapabilities string

type TVConfigurationError bool

type AccountMd string

type TimeZoneIndex int

type TimeZoneAutoAdjustDst bool

type ChannelMapSet string

type MusicSurroundLevel string

type CachedOnly bool

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type SettingsReplicationState string

type HdmiCecAvailable bool

type Orientation int

type RelativeTimePosition string

type AlarmLoggedStartTime string

type UpdateID uint

type Filter string

type SatRoomUUID string

type ServiceListVersion string

type LeftVolume uint

type RoomCalibrationID string

type AlarmList string

type LastChange string

type LIST_URI string

type Origin string

type MobileDeviceName string

type Mute bool

type EQType string

type AccountUDN string

type OAuthDeviceID string

type UpdateIDX uint

type AlarmProgramURI string

type InstanceID uint

type Icon string

type MobileDeviceUDN string

type AirPlayEnabled bool

type SoftwareVersion string

type DID string

type ThirdPartyMediaServersX string

type AlarmRunSequence string

type ConnectionManager string

type SystemUpdateID uint

type RadioLocationUpdateID uint

type ShareListUpdateID string

type VariableName string

type ThirdPartyHash string

type IsExpired bool

type URIMetaData string

type MicEnabled uint

type AudioDelayLeftRear string

type PossiblePlaybackStorageMedia string

type AVTransportURIMetaData string

type ConfigMode string

type TrackList string

type Index uint

type SortOrder string

type GroupMute bool

type DialogLevel string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type PossibleRecordStorageMedia string

type TrackMetaData string

type SurroundLevel string

type Version string

type EnqueueAs bool

type RoomDetectionPlayId uint

type QueueOwnerContext string

type ZoneGroupID string

type TransportStatus string

type ConnectionID int

type Volume uint

type SourceProtocolInfo string

type SecureRegState uint

type AuthorizationCode string

const ()

type TransportPlaySpeed string

type DirectControlIsSuspended bool

type NumTracks uint

type FavoritesUpdateID string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type HardwareVersion string

type VoiceConfigState uint

type DailyIndexRefreshTime string

type SeekTarget string

type SearchCapabilities string

type StreamRestartState string

type AutoplayRoomUUID string

type VolumeAVTransportURI string

type AccountPassword string

type TransportErrorDescription string

const ()

type RecordStorageMedium string

type NumberOfTracks uint

type ServiceId uint

type AlarmRoomUUID string

type GroupID string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ProtocolInfo string

type ShareIndexInProgress bool

type BehindWifiExtender uint

type UserIdHashCode string

type ISO8601Time string

type RecordMediumWriteStatus string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type SerialNumber string

type SubEnabled bool

type RoomCalibrationCalibrationMode string

type HTAudioIn uint

type AvailableSoftwareUpdate string

type DateFormat string

type TrackDuration string

type SupportsAudioIn bool

type BufferingResultCode int

type GroupVolume uint

type Seed string

type RoomCalibrationAvailable bool

type TimeZone string

type AutoplayUseVolume bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type WirelessMode uint

type Treble int

type CustomerID string

type PlayerID string

type SinkProtocolInfo string

type AlbumArtistDisplayOption string

type TimeZoneInformation string

type RoomCalibrationState int

type QueueUpdateID uint

type Invisible bool

type RoomDetectionChirpChannel uint

type PossibleRecordQualityModes string

type Track uint

type AbsoluteCounterPosition int

type IPAddress string

type BootSeq uint

type MuseSessions string

type SearchCriteria string

type DisplaySoftwareVersion string

type AlarmListVersion string

type SurroundEnabled bool

type MobileIPAndPort string

type TimeFormat string

type ServiceTypeList string

type QueueID uint

type EnqueuedTransportURI string

type SupportsAudioClip bool

type HTBondedZoneCommitState uint

type ConfigModeOptions string

type AlarmVolume uint

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

type SleepTimerGeneration uint

type LastIndexChange string

type ExtraInfo string

type Username string

type SupportsOutputFixed bool

type UpdateExtraOptions string

type RestartPending bool

type NumTracksChange int

type AlarmState string

type IncludeControllers bool

type SourceAreasUpdateID string

type WifiEnabled bool

type LocalGroupUUID string

type MID string

type SavedQueueTitle string

type SortCriteria string

type RadioFavoritesUpdateID uint

type VolumeAdjustment int

type Curated bool

type ButtonState string

type Code string

type TransportActions string

type DirectControlClientID string

type IsZoneBridge bool

type MACAddress string

type SurroundMode string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type RelativeCounterPosition int

type HouseholdID string

type IsIdle bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type PresetNameList string

type TransportSettings string

type TrackNumber uint

type Count uint

type ChannelFreq uint

type RightVolume uint

type VariableStringValue string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type Result string

type Flags uint

type SourceAreaIds string

type AccountType uint

type MuseHouseholdId string

type WirelessLeafOnly bool

type GroupVolumeChangeable bool

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type TrackURI string

type URI string

type ConnectionIDs string

type AVTransportID int

type ConfigModeState string

type TimeServer string

type MemberList string

type ResumePlayback bool

type AccountCredential string

type DiagnosticID uint

type ObjectID string

type RcsID int

type RoomCalibrationCoefficients string

type AudioDelay string

type SnoozeRunning bool

type SessionId string

type ChannelMap string

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

type SubPolarity string

type NightMode bool

type SourceState string

type SleepTimerState string

type RejoinGroup bool

type GroupCoordinatorIsLocal bool

type ZonePlayerUUIDsInGroup string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type MemberID string

type RoomDetectionDurationMilliseconds uint

type HTSatChanMapSet string

type Bass int

type AlarmEnabled bool

type AlarmRunning bool

type VLIState string

type TrackNumbersCSV string

type AccountNickname string

type NetsettingsUpdateID string

type AutoplayIncludeLinkedZones bool

type KeepGrouped bool

type VirtualLineInGroupID string

type MediaDuration string

type AVTransportURI string

type SubGain string

type ProgramURI string

type EnqueuedTransportURIMetaData string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type MoreInfo string

type Loudness bool

type TransportErrorURI string

type TransportErrorHttpHeaders string

type Configuration string

type TargetRoomName string

type AutoplaySource string

type LIST_URI_AND_METADATA string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AlarmProgramMetaData string

type CrossfadeMode bool

type Prefix string

type VolumeDB int

type StubsCreated string

type RDMEnabled bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AreasUpdateID string

type TimeGeneration uint

type TagValueList string

type AvailableRoomCalibration string

type EQValue int

type SubCrossover string

type AccountUID uint

type FavoritePresetsUpdateID string

type HasConfiguredSSID bool

type QueuePolicy string

type AbsoluteTimePosition string

type UserRadioUpdateID string

type VoiceUpdateID uint

type TimeStamp string

type ShareIndexLastError string

type AudioDelayRightRear string

type AlarmID uint

type ValidPlayModes string

type AccountTier uint

type LastChangedPlayState string

type CopyrightInfo string

type SpeakerSize uint

type UpdateURL string

type ContainerUpdateIDs string

type SavedQueuesUpdateID string

type HTFreq uint

type RecentlyPlayedUpdateID string

type QueueOwnerID string

type AccountID string

type AlarmIncludeLinkedZones bool

type Section uint

type AlarmIDRunning uint
