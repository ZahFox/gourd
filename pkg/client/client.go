package client

import (
	"net/rpc"

	"github.com/zahfox/gourd/pkg/command"
	"github.com/zahfox/gourd/pkg/utils"
)

// Client is used to send commands to a local or remote gourdd
type Client struct {
	c       *rpc.Client
	cmdc    chan *command.Request
	sigc    chan int
	resc    chan string
	running bool
}

// NewClient creates a new client connected to a local gourdd
func NewClient(socketPath string) Client {
	return Client{
		getConn(socketPath),
		make(chan *command.Request),
		make(chan int),
		make(chan string),
		false}
}

// Echo sends an echo command
func (c *Client) Echo(msg string) string {
	if c.running {
		cmd := command.NewRequest(command.ECHO, command.HOST, msg)
		c.cmdc <- &cmd
		return <-c.resc
	}
	return ""
}

// Ping sends a ping command
func (c *Client) Ping() string {
	if c.running {
		cmd := command.NewRequest(command.PING, command.HOST, nil)
		c.cmdc <- &cmd
		return <-c.resc
	}
	return ""
}

// Install sends an install command
func (c *Client) Install(item string) string {
	if c.running {
		params := command.InstallRequestParams{Item: item, User: utils.Username()}
		cmd := command.NewRequest(command.INSTALL, command.HOST, params)
		c.cmdc <- &cmd
		return <-c.resc
	}
	return ""
}

// Exit closes the connection to the local gourdd
func (c *Client) Exit() {
	if c.running {
		c.sigc <- 1
		c.running = false
	}
	c.c.Close()
}

// Stop prevents the client from sending commands
func (c *Client) Stop() {
	c.sigc <- 1
	c.running = false
}

// Wait will block a thread until the client stops sending commands
func (c *Client) Wait() {
	for c.running {
	}
}

// Run will make it so that the client can send commands
func (c *Client) Run() {
	if c.running {
		return
	}
	c.running = true
	go c.run()
}

func (c *Client) run() {
	for {
		select {
		case cmd := <-c.cmdc:
			c.handleCommand(cmd)
		case sig := <-c.sigc:
			if sig > 0 {
				utils.LogDebug("client stopped")
				return
			}
		}
	}
}

func (c *Client) handleCommand(cmd *command.Request) {
	action, target := (*cmd).GetAction(), (*cmd).GetTarget()
	if target != command.HOST {
		return
	}

	var err error
	var msg string
	method := command.GenServiceMethod(action, target)

	switch action {
	case command.PING:
		var res command.PingResponse
		err = c.c.Call(method, (*cmd).GetParams(), &res)
		msg = res.Message
		break
	case command.ECHO:
		var res command.EchoResponse
		err = c.c.Call(method, (*cmd).GetParams(), &res)
		msg = res.Message
		break
	case command.INSTALL:
		var res command.InstallResponse
		err = c.c.Call(method, (*cmd).GetParams(), &res)
		msg = res.Message
		break
	default:
		return
	}

	if err != nil {
		utils.LogError("error from command", err)
	} else {
		c.resc <- msg
	}
}
