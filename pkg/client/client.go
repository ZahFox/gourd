package client

import (
	"net/rpc"

	"github.com/zahfox/gourd/pkg/command"
	"github.com/zahfox/gourd/pkg/utils"
)

type Client struct {
	c       *rpc.Client
	cmdc    chan *command.Request
	sigc    chan int
	resc    chan string
	running bool
}

func NewClient() Client {
	return Client{
		getConn(),
		make(chan *command.Request),
		make(chan int),
		make(chan string),
		false}
}

func (c *Client) Echo(msg string) string {
	if c.running {
		cmd := command.NewRequest(command.ECHO, command.HOST, msg)
		c.cmdc <- &cmd
		return <-c.resc
	}
	return ""
}

func (c *Client) Ping() string {
	if c.running {
		cmd := command.NewRequest(command.PING, command.HOST, nil)
		c.cmdc <- &cmd
		return <-c.resc
	}
	return ""
}

func (c *Client) Exit() {
	if c.running {
		c.sigc <- 1
		c.running = false
	}
	c.c.Close()
}

func (c *Client) Stop() {
	c.sigc <- 1
	c.running = false
}

func (c *Client) Wait() {
	for c.running {
	}
}

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
				utils.LogDebug("Client stopped")
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
	}

	if err != nil {
		utils.LogError("Error from command", err)
	} else {
		c.resc <- msg
	}
}
