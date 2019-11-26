package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

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
func (c *Command) SetFromRequest(r Request) {
	c.ID = uuid.New().String()
	c.Timestamp = time.Now().UTC().UnixNano()
	c.Action = r.GetAction()
	c.Target = r.GetTarget()
	c.Body = r.GetParams()
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
func NewHostEcho(msg string) Command {
	return NewCommand(ECHO, HOST, msg)
}

// NewHostInstall creates a new install command for the local gourdd
func NewHostInstall(item string) Command {
	return NewCommand(INSTALL, HOST, item)
}

// ActionCode converts an action string into its code representation
func ActionCode(action string) Action {
	switch action {
	case "PING":
		return PING
	case "ECHO":
		return ECHO
	case "INSTALL":
		return INSTALL
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
	case INSTALL:
		return "INSTALL"
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
	return strings.Join([]string{strings.Title(strings.ToLower(TargetString(target))),
		strings.Title(strings.ToLower(ActionString(action)))}, ".")
}
