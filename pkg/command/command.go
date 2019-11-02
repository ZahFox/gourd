package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Action is what the gourd client is requesting to be done
type Action int

const (
	// NOACTION is a request with no action
	NOACTION Action = iota
	// PING is a request to recieve a PONG response
	PING
	// ECHO is a request to have text echoed back
	ECHO
)

// Target is what should be affected the action of a command
type Target int

const (
	// NOTARGET means no target
	NOTARGET Target = iota
	// HOST is a target that represents the local gourdd
	HOST
)

// Request is what a clients use to issue new commands
type Request struct {
	Seq    uint64      `json:"seq"`
	Method string      `json:"method"`
	Params interface{} `json:"body"`
}

// Header contains all of the metadata for a command
type Header struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Action    Action `json:"action"`
	Target    Target `json:"target"`
}

// Command is a data structure that represents a command issued by a gourd client
type Command struct {
	Header
	Body interface{} `json:"body"`
}

// NewRequest creates a new command request
func NewRequest(seq uint64, action Action, target Target, params interface{}) Request {
	return Request{
		Seq:    seq,
		Method: GenServiceMethod(action, target),
		Params: params,
	}
}

// Clear sets all of the request fields to empty values
func (r *Request) Clear() {
	r.Seq = 0
	r.Method = ""
	r.Params = nil
}

// NewCommand creates a new command
func NewCommand(action Action, target Target, body interface{}) Command {
	return Command{
		Header{
			ID:        uuid.New().String(),
			Timestamp: time.Now().UTC().UnixNano(),
			Action:    action,
			Target:    target,
		},
		body,
	}
}

// Clear sets all of the command fields to empty values
func (c *Command) Clear() {
	c.ID = ""
	c.Action = NOACTION
	c.Target = NOTARGET
	c.Body = nil
}

// Set updates the data for a command
func (c *Command) Set(action Action, target Target, body interface{}) {
	c.ID = uuid.New().String()
	c.Timestamp = time.Now().UTC().UnixNano()
	c.Action = action
	c.Target = target
	c.Body = body
}

// SetFromRequest updates the command data using a request
func (c *Command) SetFromRequest(r *Request) {
	c.ID = uuid.New().String()
	c.Timestamp = time.Now().UTC().UnixNano()
	action, target := ParseServiceMethod(r.Method)
	c.Action = action
	c.Target = target
	c.Body = r.Params
}

func (c *Command) String() string {
	return strings.Join([]string{
		ActionString(c.Action),
		TargetString(c.Target),
		strconv.FormatInt(c.Timestamp, 10),
		c.ID,
	}, " ")
}

// NewHostPing creates a new ping command for the local gourdd
func NewHostPing() Command {
	return NewCommand(PING, HOST, nil)
}

// NewHostEcho creates a new echo command for the local gourdd
func NewHostEcho() Command {
	return NewCommand(ECHO, HOST, nil)
}

// ActionCode converts an action string into its code representation
func ActionCode(action string) Action {
	switch action {
	case "PING":
		return PING
	case "ECHO":
		return ECHO
	}
	return NOACTION
}

// ActionString converts an action value into its string representation
func ActionString(action Action) string {
	switch action {
	case NOACTION:
		return "NOACTION"
	case PING:
		return "PING"
	case ECHO:
		return "ECHO"
	}
	return "NOACTION"
}

// TargetCode converts an target string into its code representation
func TargetCode(target string) Target {
	switch target {
	case "NOTARGET":
		return NOTARGET
	case "HOST":
		return HOST
	}
	return NOTARGET
}

// TargetString converts an target value into its string representation
func TargetString(target Target) string {
	switch target {
	case HOST:
		return "HOST"
	}
	return "NOTARGET"
}

// ParseServiceMethod converts a "Target.Action" string into a pair of target and action codes
func ParseServiceMethod(sm string) (Action, Target) {
	chunks := strings.Split(sm, ".")
	if len(chunks) < 2 {
		return NOACTION, NOTARGET
	}
	return ActionCode(strings.ToUpper(chunks[1])), TargetCode(strings.ToUpper(chunks[0]))
}

// GenServiceMethod converts a pair of target and action codes into a "Target.Action" string
func GenServiceMethod(action Action, target Target) string {
	return strings.Join([]string{
		strings.ToTitle(TargetString(target)),
		strings.ToTitle(ActionString(action)),
	}, ".")
}
