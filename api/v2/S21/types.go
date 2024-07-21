// CODE GENERATED, DO NOT EDIT

// Package S21 contains the implementation for the SYMFONISK Bookshelf (Model S21) services.
package S21

// Services struct contains service clients for different functionalities of a Sonos device.
type AlarmRoomUUID string

type ShareIndexLastError string

type RadioLocationUpdateID uint

type LastChangedPlayState string

type SubCrossover string

type AccountMd string

const ()

type RecordStorageMedium string

const (
	BrowseMetadataBrowseFlag       BrowseFlag = "BrowseMetadata"
	BrowseDirectChildrenBrowseFlag BrowseFlag = "BrowseDirectChildren"
)

type BrowseFlag string

type SavedQueuesUpdateID string

type IsZoneBridge bool

type ServiceListVersion string

type HeadphoneConnected bool

type SourceAreasUpdateID string

type URI string

type TrackNumber uint

type SortOrder string

type ChannelMap string

type AccountPassword string

type AlarmEnabled bool

type TransportSettings string

const (
	MasterMuteChannel      MuteChannel = "Master"
	LfMuteChannel          MuteChannel = "LF"
	RfMuteChannel          MuteChannel = "RF"
	SpeakerOnlyMuteChannel MuteChannel = "SpeakerOnly"
)

type MuteChannel string

type AudioDelay string

type MobileDeviceName string

type UpdateID uint

type AccountUID uint

type ZonePlayerUUIDsInGroup string

type DiagnosticID uint

type ProtocolInfo string

type ZoneName string

type AuthorizationCode string

type RecordQualityMode string

type TVConfigurationError bool

type QueueID uint

type MobileIPAndPort string

type VariableName string

type AccountNickname string

const (
	NonePlaybackStorageMedium    PlaybackStorageMedium = "NONE"
	NetworkPlaybackStorageMedium PlaybackStorageMedium = "NETWORK"
)

type PlaybackStorageMedium string

const (
	InputDirection  Direction = "Input"
	OutputDirection Direction = "Output"
)

type Direction string

type AlbumArtistDisplayOption string

type ContainerUpdateIDs string

type RightVolume uint

type SurroundLevel string

type VoiceUpdateID uint

type AlarmList string

type AlarmListVersion string

type RadioFavoritesUpdateID uint

type GroupVolume uint

type Code string

type EnqueuedTransportURIMetaData string

type ButtonState string

type SurroundEnabled bool

type MuseHouseholdId string

const (
	OnButtonLockState  ButtonLockState = "On"
	OffButtonLockState ButtonLockState = "Off"
)

type ButtonLockState string

type AlarmVolume uint

type AlarmRunning bool

type SourceProtocolInfo string

type Index uint

type IPAddress string

type AutoplaySource string

type Browseable bool

type Configuration string

type SupportsAudioIn bool

type EQType string

type Speed string

type AVTransportURIMetaData string

const (
	OkConnectionStatus                    ConnectionStatus = "OK"
	ContentFormatMismatchConnectionStatus ConnectionStatus = "ContentFormatMismatch"
	InsufficientBandwidthConnectionStatus ConnectionStatus = "InsufficientBandwidth"
	UnreliableChannelConnectionStatus     ConnectionStatus = "UnreliableChannel"
	UnknownConnectionStatus               ConnectionStatus = "Unknown"
)

type ConnectionStatus string

type Seed string

type EQValue int

type DialogLevel string

type UserIdHashCode string

type NumberOfTracks uint

type IsIdle bool

type ExtraInfo string

type HasConfiguredSSID bool

type QueueOwnerID string

type Loudness bool

const (
	RemoveUnresponsiveDeviceActionType                     UnresponsiveDeviceActionType = "Remove"
	TopologyMonitorProbeUnresponsiveDeviceActionType       UnresponsiveDeviceActionType = "TopologyMonitorProbe"
	VerifyThenRemoveSystemwideUnresponsiveDeviceActionType UnresponsiveDeviceActionType = "VerifyThenRemoveSystemwide"
)

type UnresponsiveDeviceActionType string

type Origin string

type Prefix string

type HTSatChanMapSet string

type ThirdPartyHash string

type AreasUpdateID string

type UpdateItem string

type TimeZoneInformation string

type StreamRestartState string

type MoreInfo string

type SessionId string

type Curated bool

type UpdateIDX uint

type HardwareVersion string

type BehindWifiExtender uint

type TimeServer string

type DateFormat string

const (
	StoppedTransportState        TransportState = "STOPPED"
	PlayingTransportState        TransportState = "PLAYING"
	PausedPlaybackTransportState TransportState = "PAUSED_PLAYBACK"
	TransitioningTransportState  TransportState = "TRANSITIONING"
)

type TransportState string

type LIST_URI string

type RoomCalibrationState int

type SerialNumber string

type GroupCoordinatorIsLocal bool

type GroupVolumeChangeable bool

type Username string

type SubEnabled bool

type Invisible bool

type CopyrightInfo string

type AlarmID uint

type TransportErrorHttpCode string

type TrackMetaData string

type SavedQueueTitle string

type ShareIndexInProgress bool

type RecentlyPlayedUpdateID string

type KeepGrouped bool

type ThirdPartyMediaServersX string

type Version string

type UpdateURL string

type URIMetaData string

type RejoinGroup bool

type Filter string

type AutoplayIncludeLinkedZones bool

type ChannelFreq uint

type RDMEnabled bool

const (
	OnLEDState  LEDState = "On"
	OffLEDState LEDState = "Off"
)

type LEDState string

type Flags uint

type MID string

type PresetNameList string

type TransportErrorHttpHeaders string

type RoomDetectionChirpIfPlayingSwappableAudio bool

type AutoplayVolume uint

type LIST_URI_AND_METADATA string

const (
	AllUpdateType      UpdateType = "All"
	SoftwareUpdateType UpdateType = "Software"
)

type UpdateType string

type AbsoluteTimePosition string

type MemberList string

type SpeakerSize uint

type AutoplayRoomUUID string

type WirelessLeafOnly bool

type HouseholdID string

type SecureRegState uint

type Volume uint

type EnqueuedTransportURI string

type Treble int

const (
	SleepTimerRampTypeRampType RampType = "SLEEP_TIMER_RAMP_TYPE"
	AlarmRampTypeRampType      RampType = "ALARM_RAMP_TYPE"
	AutoplayRampTypeRampType   RampType = "AUTOPLAY_RAMP_TYPE"
)

type RampType string

type TrackDuration string

const (
	TrackNrSeekMode   SeekMode = "TRACK_NR"
	RelTimeSeekMode   SeekMode = "REL_TIME"
	TimeDeltaSeekMode SeekMode = "TIME_DELTA"
)

type SeekMode string

type SatRoomUUID string

type RoomDetectionDurationMilliseconds uint

type VariableStringValue string

const (
	NormalAlarmPlayMode          AlarmPlayMode = "NORMAL"
	RepeatAllAlarmPlayMode       AlarmPlayMode = "REPEAT_ALL"
	ShuffleNorepeatAlarmPlayMode AlarmPlayMode = "SHUFFLE_NOREPEAT"
	ShuffleAlarmPlayMode         AlarmPlayMode = "SHUFFLE"
)

type AlarmPlayMode string

type AVTransportURI string

type TrackList string

type SourceState string

type SleepTimerState string

type AccountCredential string

type AlarmIncludeLinkedZones bool

type ObjectID string

type OutputFixed bool

type IsExpired bool

type CachedOnly bool

type MobileDeviceUDN string

const ()

type TransportPlaySpeed string

type AbsoluteCounterPosition int

type PlayerID string

type AudioDelayRightRear string

const (
	NormalPlayMode           PlayMode = "NORMAL"
	RepeatAllPlayMode        PlayMode = "REPEAT_ALL"
	RepeatOnePlayMode        PlayMode = "REPEAT_ONE"
	ShuffleNorepeatPlayMode  PlayMode = "SHUFFLE_NOREPEAT"
	ShufflePlayMode          PlayMode = "SHUFFLE"
	ShuffleRepeatOnePlayMode PlayMode = "SHUFFLE_REPEAT_ONE"
)

type PlayMode string

type CrossfadeMode bool

type ValidPlayModes string

type MemberID string

type ResumePlayback bool

type TagValueList string

type StubsCreated string

type UpdateFlags uint

type PossiblePlaybackStorageMedia string

type VLIState string

type LIST_URIMetaData string

type HTAudioIn uint

type ConfigMode string

type Bass int

type RelativeCounterPosition int

type ShareListUpdateID string

type AutoplayUseVolume bool

type ConfigModeState string

type AirPlayEnabled bool

type ServiceTypeList string

type DID string

type AccountTier uint

type RecordMediumWriteStatus string

type DirectControlAccountID string

type ResetVolumeAfter bool

type MACAddress string

type RoomCalibrationID string

type TimeZoneAutoAdjustDst bool

type PossibleRecordStorageMedia string

type RestartPending bool

type EnqueueAs bool

type SupportsAudioClip bool

type AvailableRoomCalibration string

type AlarmState string

type ConnectionID int

type HdmiCecAvailable bool

type AudioDelayLeftRear string

type AlarmRunSequence string

type TimeStamp string

type MediaDuration string

type FavoritesUpdateID string

type AccountType uint

type ZoneGroupID string

type TimeZoneIndex int

type PossibleRecordQualityModes string

type ConnectionIDs string

type AVTransportID int

type SortCapabilities string

type MicEnabled uint

type AlarmIDRunning uint

type UserRadioUpdateID string

type ServiceId uint

const (
	MasterChannel Channel = "Master"
	LfChannel     Channel = "LF"
	RfChannel     Channel = "RF"
)

type Channel string

type MusicSurroundLevel string

type IncludeControllers bool

type AlarmLoggedStartTime string

type Orientation int

type RoomDetectionChirpChannel uint

type VirtualLineInGroupID string

type LeftVolume uint

type TimeGeneration uint

type Count uint

type Mute bool

type DisplaySoftwareVersion string

type BufferingResultCode int

type SubPolarity string

type ISO8601Time string

type SnoozeRunning bool

type RoomCalibrationCalibrationMode string

type QueueUpdateID uint

type GroupMute bool

type VolumeDB int

type SubGain string

type DailyIndexRefreshTime string

type SystemUpdateID uint

type HTBondedZoneCommitState uint

type SupportsOutputFixed bool

type NetsettingsUpdateID string

type TrackURI string

type SleepTimerGeneration uint

type RoomCalibrationAvailable bool

type ZoneGroupState string

type TransportErrorDescription string

type Queue string

type SettingsReplicationState string

type HTFreq uint

type WirelessMode uint

type RoomCalibrationCoefficients string

type DirectControlIsSuspended bool

type ConnectionManager string

type Result string

type TargetRoomName string

type OAuthDeviceID string

const (
	OnceRecurrence     Recurrence = "ONCE"
	WeekdaysRecurrence Recurrence = "WEEKDAYS"
	WeekendsRecurrence Recurrence = "WEEKENDS"
	DailyRecurrence    Recurrence = "DAILY"
)

type Recurrence string

type AlarmProgramURI string

type TransportStatus string

type GroupID string

type SinkProtocolInfo string

type VolumeAVTransportURI string

type MuseSessions string

type SearchCapabilities string

type ProgramURI string

type ChannelMapSet string

type SurroundMode string

type AccountID string

type AccountUDN string

type TransportErrorURI string

type SeekTarget string

type FavoritePresetsUpdateID string

type VoiceConfigState uint

type VolumeAdjustment int

type RampTimeSeconds uint

type AlarmProgramMetaData string

type Track uint

type SortCriteria string

type Icon string

type RoomCalibrationEnabled bool

type AvailableSoftwareUpdate string

type TimeFormat string

type RcsID int

type LastIndexChange string

type QueueOwnerContext string

type TrackNumbersCSV string

type NightMode bool

type Section uint

type TransportActions string

type NumTracks uint

type RoomDetectionPlayId uint

type ServiceDescriptorList string

type RelativeTimePosition string

type SearchCriteria string

type SourceAreaIds string

type LastChange string

type SoftwareVersion string

type LocalGroupUUID string

type QueuePolicy string

type RedirectURI string

type ZoneGroupName string

type TimeZone string

type InstanceID uint

type WifiEnabled bool

type CustomerID string

type NumTracksChange int

type BootSeq uint

type DirectControlClientID string

type ConfigModeOptions string

type UpdateExtraOptions string
