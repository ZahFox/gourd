package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/ttacon/chalk"
	c "github.com/zahfox/gourd/pkg/client"
)

// Color is used internally by gsh to represent different colors
type Color int

const (
	Default Color = iota
	Red
	DarkRed
	Blue
	DarkBlue
	Green
	Black
	White
)

// ClientConfig is the options used for the gourd client
type ClientConfig struct {
	SocketPath string
}

// State is used to hold any data that the active shell might need
type State struct {
	clientStarted bool
	promptText    string
	promptColor   Color
	err           string
}

type shellConfig struct {
	socketPath string
}

// Shell is used to represent the data and behaviour of gsh
type Shell struct {
	state  State
	client c.Client
	config shellConfig
}

// NewShellOpts are the initial options used for a new shell
type NewShellOpts struct {
	PromptColor Color
	PromptText  string
	Client      ClientConfig
}

var errText = chalk.Red.NewStyle().WithBackground(chalk.ResetColor).WithTextStyle(chalk.Bold).Style
var infoText = chalk.Magenta.NewStyle().WithBackground(chalk.ResetColor).WithTextStyle(chalk.Italic).Style

// NewShell creates and configures a new shell
func NewShell(opts NewShellOpts) Shell {
	return Shell{
		config: shellConfig{
			socketPath: opts.Client.SocketPath,
		},
		state: State{
			clientStarted: false,
			promptColor:   opts.PromptColor,
			promptText:    opts.PromptText,
		},
	}
}

// Die disconnects the client and kills the process
func (s *Shell) Die(code int) {
	s.Disconnect()
	os.Exit(code)
}

// Disconnect disables any features beyond the shell's own process
func (s *Shell) Disconnect() {
	if s.state.clientStarted {
		s.state.clientStarted = false
		s.client.Exit()
	}
}

// Client returns a reference to a gourd client
func (s *Shell) Client() *c.Client {
	s.startClient()
	return &s.client
}

// ClearErr set the shell error to an empty value
func (s *Shell) ClearErr() {
	s.state.err = ""
}

// GetErr returns the current value of the shell error
func (s *Shell) GetErr() string {
	return s.state.err
}

// HasErr indicates whether or not the shell currently has an error
func (s *Shell) HasErr() bool {
	return s.state.err != ""
}

// SetErr changes the current value of the shell error
func (s *Shell) SetErr(err string) {
	s.state.err = err
	fmt.Println(errText(err))
}

// SetErrF changes the current value of the shell error
func (s *Shell) SetErrF(err string, a ...interface{}) {
	s.state.err = fmt.Sprintf(err, a...)
	fmt.Println(errText(s.state.err))
}

// GetPromptColor returns the value of the prompt color that the active shell is using
func (s *Shell) GetPromptColor() Color {
	return s.state.promptColor
}

// SetPromptColor updates the prompt color that the active shell is using
func (s *Shell) SetPromptColor(c Color) {
	s.state.promptColor = c
	fmt.Printf(infoText("set prompt color to '%s'\n"), colorToDisplay(c))
}

// GetPromptText returns the value of the prompt text that the active shell is using
func (s *Shell) GetPromptText() string {
	return s.state.promptText
}

// SetPromptText updates the prompt text that the active shell is using
func (s *Shell) SetPromptText(p string) {
	s.state.promptText = p
	fmt.Printf(infoText("set prompt text to '%s'\n"), strings.Trim(p, " "))
}

func (s *Shell) startClient() {
	if s.state.clientStarted {
		return
	}
	s.state.clientStarted = true
	s.client = c.NewClient(s.config.socketPath)
	s.client.Run()
}
