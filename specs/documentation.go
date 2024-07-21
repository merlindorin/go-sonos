package specs

type Documentation struct {
	Schema   string                 `json:"$schema"`
	Language string                 `json:"language"`
	License  string                 `json:"license"`
	Services map[string]*DocService `json:"services"`
	Errors   []struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

func (d Documentation) Service(name string) *DocService {
	for serviceName, service := range d.Services {
		if serviceName == name {
			return service
		}
	}

	return nil
}

type DocService struct {
	Description string                `json:"description"`
	Actions     map[string]*DocAction `json:"actions"`
	Errors      []struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

func (s DocService) Action(name string) *DocAction {
	for actionName, action := range s.Actions {
		if actionName == name {
			return action
		}
	}

	return nil
}

type DocAction struct {
	Description string `json:"description"`
	Remarks     string `json:"remarks"`
	Params      struct {
		NewCoordinator string `json:"NewCoordinator"`
		RejoinGroup    string `json:"RejoinGroup"`
	} `json:"params"`
}
