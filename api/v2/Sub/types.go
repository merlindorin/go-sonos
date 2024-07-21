// CODE GENERATED, DO NOT EDIT

// Package Sub contains the implementation for the Sonos Sub (Model Sub) services.
package Sub

// Services struct contains service clients for different functionalities of a Sonos device.
type ResumePlayback bool

type SerialNumber string

type IPAddress string

type UpdateIDX uint

type Speed string

type AVTransportURIMetaData string

type DirectControlAccountID string

type ConnectionID int

type Flags uint

type HardwareVersion string

type MobileDeviceUDN string

type MediaDuration string

type SourceState string

type SavedQueueTitle string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type TrackList string

type Configuration string

type VolumeDB int

type MuseHouseholdId string

type AlarmProgramURI string

type Icon string

type ConfigModeState string

type ZoneGroupName string

type SeekTarget string

type AlarmState string

type ShareListUpdateID string

type SoftwareVersion string

type Orientation int

type QueueID uint

type AccountUDN string

type ValidPlayModes string

type NumTracksChange int

type RejoinGroup bool

type SupportsAudioIn bool

type LocalGroupUUID string

type GroupMute bool

type GroupVolumeChangeable bool

type AccountType uint

type AuthorizationCode string

type MemberID string

type SatRoomUUID string

type SourceAreaIds string

type RoomCalibrationAvailable bool

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type TrackURI string

type RcsID int

type Result string

type Treble int

type RoomCalibrationCoefficients string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type TagValueList string

type GroupVolume uint

type ServiceId uint

type TransportErrorHttpHeaders string

type CopyrightInfo string

type Origin string

type MuseSessions string

type DirectControlIsSuspended bool

type HdmiCecAvailable bool

type ISO8601Time string

type AlarmEnabled bool

type AlarmListVersion string

const ()

type TransportPlaySpeed string

type RoomCalibrationState int

type TimeZone string

type CrossfadeMode bool

type SearchCapabilities string

type HouseholdID string

type SubPolarity string

type OAuthDeviceID string

type TimeZoneAutoAdjustDst bool

type EnqueuedTransportURIMetaData string

type ShareIndexLastError string

type VoiceConfigState uint

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type SubGain string

type UpdateItem string

type PossibleRecordStorageMedia string

type DirectControlClientID string

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type QueueOwnerContext string

type BufferingResultCode int

type RightVolume uint

type PresetNameList string

type RelativeCounterPosition int

type PlayerID string

type RadioLocationUpdateID uint

type HTSatChanMapSet string

type TimeZoneInformation string

type HasConfiguredSSID bool

type RoomDetectionDurationMilliseconds uint

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type NetsettingsUpdateID string

const ()

type RecordStorageMedium string

type WirelessLeafOnly bool

type Mute bool

type DiagnosticID uint

type IsZoneBridge bool

type HTAudioIn uint

type MicEnabled uint

type NightMode bool

type URIMetaData string

type NumTracks uint

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type RoomCalibrationID string

type GroupID string

type MACAddress string

type EQType string

type AlarmRunSequence string

type AlarmVolume uint

type ObjectID string

type SinkProtocolInfo string

type VariableName string

type QueuePolicy string

type AccountUID uint

type MobileIPAndPort string

type TrackMetaData string

type EnqueueAs bool

type SettingsReplicationState string

type ChannelMapSet string

type AccountTier uint

type Index uint

type ChannelFreq uint

type Volume uint

type AudioDelayLeftRear string

type AlarmRoomUUID string

type ResetVolumeAfter bool

type TrackNumbersCSV string

type AreasUpdateID string

type LIST_URI_AND_METADATA string

type Track uint

type Filter string

type UpdateID uint

type AutoplayUseVolume bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type TimeStamp string

type AVTransportID int

type SavedQueuesUpdateID string

type AccountNickname string

type TransportErrorHttpCode string

type UserRadioUpdateID string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type SubCrossover string

type RoomCalibrationCalibrationMode string

type CachedOnly bool

type UpdateExtraOptions string

type AVTransportURI string

type MemberList string

type URI string

type AirPlayEnabled bool

type PossiblePlaybackStorageMedia string

type SleepTimerGeneration uint

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type ProtocolInfo string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AlarmList string

type TimeFormat string

type TransportStatus string

type FavoritesUpdateID string

type ExtraInfo string

type ConfigModeOptions string

type DialogLevel string

type Version string

type TransportErrorURI string

type SupportsAudioClip bool

type ServiceDescriptorList string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type TimeServer string

type Browseable bool

type VirtualLineInGroupID string

type TransportErrorDescription string

type LastChange string

type AccountCredential string

type OutputFixed bool

type CustomerID string

type TimeZoneIndex int

type EnqueuedTransportURI string

type ButtonState string

type BootSeq uint

type ConfigMode string

type SourceAreasUpdateID string

type UserIdHashCode string

type ThirdPartyHash string

type IncludeControllers bool

type DailyIndexRefreshTime string

type SortCapabilities string

type Seed string

type AccountMd string

type SleepTimerState string

type StreamRestartState string

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type SnoozeRunning bool

type AlarmLoggedStartTime string

type TargetRoomName string

type AutoplayRoomUUID string

type RestartPending bool

type SearchCriteria string

type LastChangedPlayState string

type AutoplaySource string

type TrackDuration string

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type RadioFavoritesUpdateID uint

type Invisible bool

type QueueOwnerID string

type ConnectionManager string

type SurroundEnabled bool

type AvailableSoftwareUpdate string

type SubEnabled bool

type SurroundLevel string

type SurroundMode string

type UpdateFlags uint

type PossibleRecordQualityModes string

type LIST_URI string

type AutoplayVolume uint

type Curated bool

type StubsCreated string

type InstanceID uint

type AutoplayIncludeLinkedZones bool

type ChannelMap string

type MusicSurroundLevel string

type Loudness bool

type AccountID string

type ThirdPartyMediaServersX string

type TransportActions string

type AlarmRunning bool

type BehindWifiExtender uint

type MID string

type Section uint

type MoreInfo string

type AvailableRoomCalibration string

type RDMEnabled bool

type VoiceUpdateID uint

type MobileDeviceName string

type HTFreq uint

type DID string

type SpeakerSize uint

type RoomCalibrationEnabled bool

type RelativeTimePosition string

type QueueUpdateID uint

type RoomDetectionChirpChannel uint

type ServiceListVersion string

type VolumeAVTransportURI string

type AccountPassword string

type AbsoluteTimePosition string

type LIST_URIMetaData string

type ShareIndexInProgress bool

type IsIdle bool

type AlarmIDRunning uint

type Queue string

type SourceProtocolInfo string

type HeadphoneConnected bool

type AbsoluteCounterPosition int

type ZoneGroupState string

type ContainerUpdateIDs string

type AlarmProgramMetaData string

type AlarmIncludeLinkedZones bool

type TransportSettings string

type ConnectionIDs string

type EQValue int

type ZonePlayerUUIDsInGroup string

type DateFormat string

type SortCriteria string

type SecureRegState uint

type GroupCoordinatorIsLocal bool

type AudioDelay string

type RedirectURI string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

type Count uint

type LastIndexChange string

type RecentlyPlayedUpdateID string

type UpdateURL string

type HTBondedZoneCommitState uint

type WirelessMode uint

type RampTimeSeconds uint

type AudioDelayRightRear string

type DisplaySoftwareVersion string

type SessionId string

type LeftVolume uint

type VariableStringValue string

type RecordQualityMode string

type VLIState string

type AlbumArtistDisplayOption string

type SystemUpdateID uint

type SupportsOutputFixed bool

type ProgramURI string

type ZoneGroupID string

type RecordMediumWriteStatus string

type Username string

type Code string

type IsExpired bool

type VolumeAdjustment int

type Bass int

type AlarmID uint

type TrackNumber uint

type FavoritePresetsUpdateID string

type TVConfigurationError bool

type ZoneName string

type WifiEnabled bool

type KeepGrouped bool

type RoomDetectionPlayId uint

type ServiceTypeList string

type TimeGeneration uint

type NumberOfTracks uint

type Prefix string

type SortOrder string
