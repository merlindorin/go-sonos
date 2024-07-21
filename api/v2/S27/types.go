// CODE GENERATED, DO NOT EDIT

// Package S27 contains the implementation for the Sonos Roam (Model S27) services.
package S27

// Services struct contains service clients for different functionalities of a Sonos device.
type SortCriteria string

type ShareIndexLastError string

type HdmiCecAvailable bool

type Curated bool

type SubCrossover string

type ISO8601Time string

type TrackMetaData string

type SleepTimerGeneration uint

type IsExpired bool

type ServiceId uint

type SessionId string

type OutputFixed bool

type ZonePlayerUUIDsInGroup string

type EnqueuedTransportURI string

type ShareListUpdateID string

type WirelessLeafOnly bool

type Version string

type SearchCriteria string

type AccountCredential string

type UpdateURL string

type RecordQualityMode string

type Prefix string

type AlarmEnabled bool

type DiagnosticID uint

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type ConfigMode string

type RDMEnabled bool

type SeekTarget string

type RecentlyPlayedUpdateID string

type RightVolume uint

type ButtonState string

type TimeZoneIndex int

type AVTransportID int

type TargetRoomName string

type AudioDelayRightRear string

type MuseHouseholdId string

type AlarmProgramURI string

type SortOrder string

type IPAddress string

type UpdateIDX uint

type MediaDuration string

type TransportActions string

type SinkProtocolInfo string

type HasConfiguredSSID bool

type QueueOwnerContext string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type CachedOnly bool

type URIMetaData string

type TrackList string

type Flags uint

type VolumeAdjustment int

type ProgramURI string

type AudioDelay string

type AccountUID uint

type Speed string

type DailyIndexRefreshTime string

type TransportErrorDescription string

type RoomDetectionChirpChannel uint

type SavedQueuesUpdateID string

type ChannelMapSet string

type QueueID uint

type TrackNumbersCSV string

type TransportErrorHttpCode string

type Section uint

type ConnectionIDs string

type TimeZone string

type StreamRestartState string

type SourceAreaIds string

type SpeakerSize uint

type SurroundLevel string

type Track uint

type AlarmIDRunning uint

type HouseholdID string

type AlarmRunSequence string

type FavoritesUpdateID string

type SupportsOutputFixed bool

type SubEnabled bool

type AlarmID uint

type ServiceDescriptorList string

type AccountTier uint

type EQValue int

type RoomCalibrationID string

type VoiceUpdateID uint

type TransportStatus string

type AlarmLoggedStartTime string

type MemberList string

type Volume uint

type SurroundEnabled bool

type SurroundMode string

type AccountType uint

type OAuthDeviceID string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type SerialNumber string

type ServiceListVersion string

type NetsettingsUpdateID string

type SatRoomUUID string

type LocalGroupUUID string

type GroupVolumeChangeable bool

type RedirectURI string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type Browseable bool

type Configuration string

type DID string

type SubGain string

type ZoneGroupState string

type TimeZoneInformation string

type URI string

type SoftwareVersion string

type AlarmRunning bool

type Seed string

type DialogLevel string

type TimeFormat string

type TransportErrorURI string

type CrossfadeMode bool

type UpdateItem string

type Origin string

type TimeGeneration uint

type AvailableRoomCalibration string

type ConfigModeOptions string

type SystemUpdateID uint

type SecureRegState uint

type MusicSurroundLevel string

type PossibleRecordStorageMedia string

type MuseSessions string

type NumTracksChange int

type TimeStamp string

type AccountID string

type RoomCalibrationEnabled bool

const ()

type RecordStorageMedium string

type NumTracks uint

type EQType string

type UserIdHashCode string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type RestartPending bool

type RoomDetectionChirpIfPlayingSwappableAudio bool

type Queue string

type LIST_URIMetaData string

type UserRadioUpdateID string

type HTAudioIn uint

type ChannelMap string

type AlarmProgramMetaData string

type PossiblePlaybackStorageMedia string

type DirectControlAccountID string

type AlbumArtistDisplayOption string

type VoiceConfigState uint

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type AreasUpdateID string

type TimeZoneAutoAdjustDst bool

type MemberID string

type AlarmState string

type QueueOwnerID string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type RelativeCounterPosition int

type LIST_URI string

type MACAddress string

type Username string

type MobileIPAndPort string

type DateFormat string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type AutoplayIncludeLinkedZones bool

type SearchCapabilities string

type RadioLocationUpdateID uint

type Loudness bool

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type TimeServer string

type EnqueuedTransportURIMetaData string

type ThirdPartyMediaServersX string

type IsZoneBridge bool

type SupportsAudioClip bool

type MoreInfo string

type Invisible bool

type Orientation int

type LastChangedPlayState string

type WirelessMode uint

type KeepGrouped bool

type RelativeTimePosition string

type LastChange string

type SourceProtocolInfo string

type AccountUDN string

type Bass int

type HeadphoneConnected bool

type AudioDelayLeftRear string

type AlarmIncludeLinkedZones bool

type SupportsAudioIn bool

type Index uint

type HTSatChanMapSet string

type WifiEnabled bool

type AlarmRoomUUID string

type AbsoluteTimePosition string

type ResumePlayback bool

type ConfigModeState string

type LeftVolume uint

type VLIState string

type Result string

type FavoritePresetsUpdateID string

type MobileDeviceName string

type SourceAreasUpdateID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type SortCapabilities string

type ZoneGroupName string

type ProtocolInfo string

type MicEnabled uint

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type QueuePolicy string

type RoomCalibrationCalibrationMode string

type ZoneGroupID string

type PossibleRecordQualityModes string

type HTBondedZoneCommitState uint

type Code string

type ConnectionID int

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type TransportErrorHttpHeaders string

type RecordMediumWriteStatus string

type Filter string

type BufferingResultCode int

type PresetNameList string

type MobileDeviceUDN string

type VariableName string

type TrackDuration string

type UpdateID uint

type AirPlayEnabled bool

type VirtualLineInGroupID string

type GroupVolume uint

type AlarmVolume uint

type TrackURI string

type TagValueList string

type TVConfigurationError bool

type RampTimeSeconds uint

type IncludeControllers bool

type ZoneName string

type CopyrightInfo string

type BehindWifiExtender uint

type RoomCalibrationCoefficients string

type AlarmList string

type AlarmListVersion string

type LastIndexChange string

type HardwareVersion string

type RoomCalibrationState int

const ()

type TransportPlaySpeed string

type NumberOfTracks uint

type RcsID int

type SourceState string

type ServiceTypeList string

type DirectControlIsSuspended bool

type HTFreq uint

type RoomDetectionDurationMilliseconds uint

type VariableStringValue string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type AVTransportURIMetaData string

type AbsoluteCounterPosition int

type AutoplaySource string

type GroupMute bool

type AccountPassword string

type AuthorizationCode string

type CustomerID string

type ValidPlayModes string

type Count uint

type RadioFavoritesUpdateID uint

type UpdateExtraOptions string

type Mute bool

type SleepTimerState string

type MID string

type LIST_URI_AND_METADATA string

type BootSeq uint

type NightMode bool

type RoomCalibrationAvailable bool

type ThirdPartyHash string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type TransportSettings string

type EnqueueAs bool

type QueueUpdateID uint

type ExtraInfo string

type AccountNickname string

type InstanceID uint

type SavedQueueTitle string

type SettingsReplicationState string

type Icon string

type ObjectID string

type PlayerID string

type RejoinGroup bool

type Treble int

type StubsCreated string

type ContainerUpdateIDs string

type ShareIndexInProgress bool

type GroupCoordinatorIsLocal bool

type AccountMd string

type SnoozeRunning bool

type RoomDetectionPlayId uint

type SubPolarity string

type VolumeDB int

type IsIdle bool

type AutoplayRoomUUID string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type TrackNumber uint

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type DisplaySoftwareVersion string

type AutoplayVolume uint

type ChannelFreq uint

type AvailableSoftwareUpdate string

type UpdateFlags uint

type AVTransportURI string

type ResetVolumeAfter bool

type ConnectionManager string

type VolumeAVTransportURI string

type DirectControlClientID string

type GroupID string

type AutoplayUseVolume bool
