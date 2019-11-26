package command

// Request is what a clients use to issue new commands
type Request interface {
	GetAction() Action
	GetTarget() Target
	GetParams() interface{}
}

// EchoRequest is what a clients use to issue new echo commands
type EchoRequest struct {
	Action  Action `json:"action"`
	Target  Target `json:"target"`
	Message string `json:"message"`
}

func (r *EchoRequest) GetAction() Action {
	return ECHO
}

func (r *EchoRequest) GetTarget() Target {
	return r.Target
}

func (r *EchoRequest) GetParams() interface{} {
	return r.Message
}

// PingRequest is what a clients use to issue new ping commands
type PingRequest struct {
	Action Action `json:"action"`
	Target Target `json:"target"`
}

func (r *PingRequest) GetAction() Action {
	return PING
}

func (r *PingRequest) GetTarget() Target {
	return r.Target
}

func (r *PingRequest) GetParams() interface{} {
	return nil
}

// InstallRequest is what a clients use to issue new install commands
type InstallRequest struct {
	Action Action `json:"action"`
	Target Target `json:"target"`
	Item   string `json:"item"`
}

func (r *InstallRequest) GetAction() Action {
	return INSTALL
}

func (r *InstallRequest) GetTarget() Target {
	return r.Target
}

func (r *InstallRequest) GetParams() interface{} {
	return r.Item
}

// NewRequest creates a new command request
func NewRequest(action Action, target Target, params interface{}) Request {
	switch action {
	case ECHO:
		return &EchoRequest{
			Action:  action,
			Target:  target,
			Message: params.(string),
		}
	case PING:
		return &PingRequest{
			Action: action,
			Target: target,
		}
	case INSTALL:
		return &InstallRequest{
			Action: action,
			Target: target,
			Item:   params.(string),
		}
	}
	return nil
}
