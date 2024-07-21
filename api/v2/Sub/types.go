// CODE GENERATED, DO NOT EDIT

// Package Sub contains the implementation for the Sonos Sub (Model Sub) services.
package Sub

// Services struct contains service clients for different functionalities of a Sonos device.
type PresetNameList string

type UserRadioUpdateID string

type IPAddress string

type HasConfiguredSSID bool

type SubEnabled bool

type ZonePlayerUUIDsInGroup string

type MediaDuration string

type IsIdle bool

type SourceAreaIds string

type Code string

type AccountID string

type RecordQualityMode string

type AbsoluteTimePosition string

type TagValueList string

type Curated bool

type ThirdPartyMediaServersX string

type PossibleRecordStorageMedia string

type InstanceID uint

type MicEnabled uint

type VariableStringValue string

type AlarmID uint

type AlarmIDRunning uint

type VLIState string

type RadioLocationUpdateID uint

type SupportsOutputFixed bool

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SourceState string

type MemberID string

type IsZoneBridge bool

type DID string

type MoreInfo string

type RoomDetectionPlayId uint

type GroupCoordinatorIsLocal bool

type QueueID uint

type VirtualLineInGroupID string

type GroupVolumeChangeable bool

type OutputFixed bool

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type URIMetaData string

type ResetVolumeAfter bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type RelativeTimePosition string

type LastChange string

type AlarmList string

type Index uint

type AutoplayVolume uint

type AlarmState string

type LastChangedPlayState string

type AccountUDN string

type TransportStatus string

type ObjectID string

type Count uint

type LeftVolume uint

type AlbumArtistDisplayOption string

type HTBondedZoneCommitState uint

type MACAddress string

type Flags uint

type TrackURI string

type PlayerID string

type ConnectionManager string

type UpdateID uint

type VariableName string

type ISO8601Time string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type ServiceTypeList string

type AudioDelayLeftRear string

const ()

type TransportPlaySpeed string

type SourceProtocolInfo string

type SubGain string

type ShareIndexInProgress bool

type AccountPassword string

type MobileDeviceUDN string

type LIST_URI string

type QueueOwnerID string

type SourceAreasUpdateID string

type BootSeq uint

type TimeZone string

type CrossfadeMode bool

type SavedQueuesUpdateID string

type Browseable bool

type TrackNumbersCSV string

type Mute bool

type SubCrossover string

type SurroundLevel string

type AbsoluteCounterPosition int

type HTAudioIn uint

type HdmiCecAvailable bool

type RoomDetectionDurationMilliseconds uint

type RoomCalibrationID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type ConfigModeOptions string

type ConfigModeState string

type NetsettingsUpdateID string

type TrackDuration string

type AlarmLoggedStartTime string

type SearchCapabilities string

type WifiEnabled bool

type AVTransportID int

type AccountCredential string

type TransportErrorDescription string

type ValidPlayModes string

type MemberList string

type QueueUpdateID uint

type EnqueuedTransportURIMetaData string

type TrackList string

type LIST_URI_AND_METADATA string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type ResumePlayback bool

type ShareListUpdateID string

type VolumeDB int

type EQValue int

type MobileIPAndPort string

type Section uint

type URI string

type RejoinGroup bool

type ProgramURI string

const ()

type RecordStorageMedium string

type SortOrder string

type AirPlayEnabled bool

type AccountMd string

type TimeGeneration uint

type DailyIndexRefreshTime string

type TimeZoneAutoAdjustDst bool

type SavedQueueTitle string

type DiagnosticID uint

type ProtocolInfo string

type ShareIndexLastError string

type FavoritePresetsUpdateID string

type HouseholdID string

type Track uint

type AVTransportURI string

type TransportActions string

type DirectControlClientID string

type ChannelMapSet string

type Orientation int

type AccountNickname string

type UpdateFlags uint

type RelativeCounterPosition int

type Queue string

type LastIndexChange string

type UserIdHashCode string

type ServiceListVersion string

type RightVolume uint

type SurroundMode string

type AlarmRoomUUID string

type EnqueueAs bool

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type LocalGroupUUID string

type SecureRegState uint

type RDMEnabled bool

type AlarmProgramMetaData string

type Filter string

type SerialNumber string

type WirelessMode uint

type RoomCalibrationCalibrationMode string

type AlarmProgramURI string

type SleepTimerState string

type ExtraInfo string

type NightMode bool

type ButtonState string

type AccountType uint

type RoomDetectionChirpChannel uint

type DialogLevel string

type AlarmEnabled bool

type TimeZoneInformation string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type FavoritesUpdateID string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type DisplaySoftwareVersion string

type Username string

type AlarmRunSequence string

type ServiceDescriptorList string

type AuthorizationCode string

type RedirectURI string

type Origin string

type SettingsReplicationState string

type RoomCalibrationState int

type VoiceConfigState uint

type VolumeAdjustment int

type QueuePolicy string

type StubsCreated string

type MuseHouseholdId string

type CachedOnly bool

type ZoneGroupID string

type UpdateItem string

type StreamRestartState string

type SortCapabilities string

type HardwareVersion string

type KeepGrouped bool

type TransportErrorURI string

type AlarmRunning bool

type RoomCalibrationCoefficients string

type UpdateIDX uint

type CustomerID string

type AudioDelay string

type AudioDelayRightRear string

type UpdateURL string

type TimeStamp string

type PossiblePlaybackStorageMedia string

type EnqueuedTransportURI string

type ContainerUpdateIDs string

type TimeZoneIndex int

type Icon string

type SatRoomUUID string

type ChannelFreq uint

type MobileDeviceName string

type EQType string

type AccountTier uint

type ThirdPartyHash string

type ZoneGroupName string

type VoiceUpdateID uint

type UpdateExtraOptions string

type IncludeControllers bool

type SnoozeRunning bool

type SupportsAudioIn bool

type ServiceId uint

type ChannelMap string

type HeadphoneConnected bool

type AlarmVolume uint

type GroupID string

type HTSatChanMapSet string

type Seed string

type Treble int

type OAuthDeviceID string

type DirectControlAccountID string

type RadioFavoritesUpdateID uint

type ZoneName string

type Volume uint

type RecordMediumWriteStatus string

type NumTracksChange int

type Prefix string

type HTFreq uint

type DateFormat string

type PossibleRecordQualityModes string

type NumberOfTracks uint

type Configuration string

type VolumeAVTransportURI string

type MusicSurroundLevel string

type RoomCalibrationAvailable bool

type TrackMetaData string

type RcsID int

type SortCriteria string

type RecentlyPlayedUpdateID string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type Invisible bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type AreasUpdateID string

type Loudness bool

type SurroundEnabled bool

type NumTracks uint

type ConnectionIDs string

type BehindWifiExtender uint

type GroupVolume uint

type LIST_URIMetaData string

type SupportsAudioClip bool

type AutoplayUseVolume bool

type GroupMute bool

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type SubPolarity string

type AccountUID uint

type Speed string

type TransportSettings string

type SoftwareVersion string

type TVConfigurationError bool

type BufferingResultCode int

type ZoneGroupState string

type TransportErrorHttpCode string

type QueueOwnerContext string

type RoomCalibrationEnabled bool

type TimeFormat string

type SinkProtocolInfo string

type MID string

type RampTimeSeconds uint

type MuseSessions string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type AvailableRoomCalibration string

type AutoplayIncludeLinkedZones bool

type TrackNumber uint

type WirelessLeafOnly bool

type SessionId string

type SpeakerSize uint

type TimeServer string

type AlarmListVersion string

type SleepTimerGeneration uint

type SeekTarget string

type Version string

type ConfigMode string

type Bass int

type AvailableSoftwareUpdate string

type AlarmIncludeLinkedZones bool

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type Result string

type TargetRoomName string

type AutoplayRoomUUID string

type AutoplaySource string

type RestartPending bool

type DirectControlIsSuspended bool

type SearchCriteria string

type CopyrightInfo string

type IsExpired bool

type TransportErrorHttpHeaders string

type AVTransportURIMetaData string

type ConnectionID int

type SystemUpdateID uint
