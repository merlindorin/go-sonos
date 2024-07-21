package specs

import (
	"time"
)

type StateVariable struct {
	Name          string      `json:"name"`
	DataType      string      `json:"dataType"`
	SendEvents    bool        `json:"sendEvents"`
	AllowedValues interface{} `json:"allowedValues,omitempty"`
}

type Input struct {
	Name                     string `json:"name"`
	Direction                string `json:"direction"`
	RelatedStateVariableName string `json:"relatedStateVariableName"`
}

type Output struct {
	Name                     string `json:"name"`
	Direction                string `json:"direction"`
	RelatedStateVariableName string `json:"relatedStateVariableName"`
}

type Actions struct {
	Name    string   `json:"name"`
	Inputs  []Input  `json:"inputs,omitempty"`
	Outputs []Output `json:"outputs,omitempty"`
}

type Service struct {
	Name           string         `json:"name"`
	ServiceName    string         `json:"serviceName"`
	DiscoveryURI   string         `json:"discoveryUri"`
	ServiceID      string         `json:"serviceId"`
	ServiceType    string         `json:"serviceType"`
	ControlURL     string         `json:"controlURL"`
	EventSubURL    string         `json:"eventSubURL"`
	StateVariables StateVariables `json:"stateVariables"`
	Actions        []Actions      `json:"actions"`
}

type StateVariables []StateVariable

func (s StateVariables) Get(name string, stateVariable *StateVariable) {
	for _, variable := range s {
		if variable.Name == name {
			*stateVariable = variable
		}
	}
}

type APISpec struct {
	Model              string    `json:"model"`
	ModelDescription   string    `json:"modelDescription"`
	SoftwareGeneration int       `json:"softwareGeneration"`
	SoftwareVersion    string    `json:"softwareVersion"`
	DiscoveryDate      time.Time `json:"discoveryDate"`
	Services           []Service `json:"services"`
}
