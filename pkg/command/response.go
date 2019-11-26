package command

// Response is what a server will send back to a client after a command
type Response interface {
	GetId() string
	GetError() string
	GetData() interface{}
}

// EchoResponse is what a client will recieve after issuing an echo command
type EchoResponse struct {
	ID      string `json:"id"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (r *EchoResponse) GetId() string {
	return r.ID
}

func (r *EchoResponse) GetError() string {
	return r.Error
}

func (r *EchoResponse) GetData() interface{} {
	return r.Message
}

// PingResponse is what a client will recieve after issuing a ping command
type PingResponse struct {
	ID      string `json:"id"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (r *PingResponse) GetId() string {
	return r.ID
}

func (r *PingResponse) GetError() string {
	return r.Error
}

func (r *PingResponse) GetData() interface{} {
	return r.Message
}

// InstallResponse is what a client will recieve after issuing a install command
type InstallResponse struct {
	ID      string `json:"id"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (r *InstallResponse) GetId() string {
	return r.ID
}

func (r *InstallResponse) GetError() string {
	return r.Error
}

func (r *InstallResponse) GetData() interface{} {
	return r.Message
}

// NewResponse creates a new command response
func NewResponse(cmd *Command, err string) Response {
	switch cmd.Action {
	case ECHO:
		return &EchoResponse{
			ID:      cmd.ID,
			Error:   err,
			Message: cmd.Body.(string),
		}
	case PING:
		return &PingResponse{
			ID:      cmd.ID,
			Error:   err,
			Message: cmd.Body.(string),
		}
	case INSTALL:
		return &InstallResponse{
			ID:      cmd.ID,
			Error:   err,
			Message: cmd.Body.(string),
		}
	}
	return nil
}
