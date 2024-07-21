// CODE GENERATED, DO NOT EDIT

// Package S3 contains the implementation for the Sonos Play:3 (Model S3) services.
package S3

import do "github.com/vanyda-official/go-shared/pkg/net/do"

// Services struct contains service clients for different functionalities of a Sonos device.
type Services struct {
	// Control the sonos alarms and times
	AlarmClock *AlarmClockService

	// Service that controls stuff related to transport (play/pause/next/special urls)
	AVTransport *AVTransportService

	// Services related to connections and protocols
	ConnectionManager *ConnectionManagerService

	// Browse for local content
	ContentDirectory *ContentDirectoryService

	// Modify device properties, like LED status and stereo pairs
	DeviceProperties *DevicePropertiesService

	// Services related to groups
	GroupManagement *GroupManagementService

	// Volume related controls for groups
	GroupRenderingControl *GroupRenderingControlService

	// Access to external music services, like Spotify or Youtube Music
	MusicServices *MusicServicesService

	// Services related to Chinese Tencent Qplay service
	QPlay *QPlayService

	// Modify and browse queues
	Queue *QueueService

	// Volume related controls
	RenderingControl *RenderingControlService

	// Manage system-wide settings, mainly account stuff
	SystemProperties *SystemPropertiesService

	VirtualLineIn *VirtualLineInService
	// Zone config stuff, eg getting all the configured sonos zones
	ZoneGroupTopology *ZoneGroupTopologyService
}

// NewService creates and returns a new Services struct initialized with base URLs for Sonos services available.
func NewService(doer do.Doer) *Services {
	return &Services{
		AVTransport:           NewAVTransportService(doer),
		AlarmClock:            NewAlarmClockService(doer),
		ConnectionManager:     NewConnectionManagerService(doer),
		ContentDirectory:      NewContentDirectoryService(doer),
		DeviceProperties:      NewDevicePropertiesService(doer),
		GroupManagement:       NewGroupManagementService(doer),
		GroupRenderingControl: NewGroupRenderingControlService(doer),
		MusicServices:         NewMusicServicesService(doer),
		QPlay:                 NewQPlayService(doer),
		Queue:                 NewQueueService(doer),
		RenderingControl:      NewRenderingControlService(doer),
		SystemProperties:      NewSystemPropertiesService(doer),
		VirtualLineIn:         NewVirtualLineInService(doer),
		ZoneGroupTopology:     NewZoneGroupTopologyService(doer),
	}
}

// Model returns the model identifier of the Sonos device.
func (s Services) Model() string {
	return "S3"
}

// ModelDescription returns a human-readable description of the Sonos model.
func (s Services) ModelDescription() string {
	return "Sonos Play:3"
}

// SoftwareGeneration returns the generation of the software the services are running on.
func (s Services) SoftwareGeneration() int {
	return 2
}

// SoftwareVersion returns the current software version installed on the Sonos device.
func (s Services) SoftwareVersion() string {
	return "64.3-19080"
}

// DiscoveryDate returns a timestamp (in Unix epoch time) representing when the services were discovered.
func (s Services) DiscoveryDate() int {
	return 1629266976
}
