package client

import (
	"log"
	"net/rpc"
	"os"

	"github.com/zahfox/gourd/pkg/command"
)

type Client struct {
	sl      *log.Logger
	el      *log.Logger
	c       *rpc.Client
	cmdc    chan *command.Request
	sigc    chan int
	running bool
}

func NewClient() Client {
	return Client{
		log.New(os.Stdout, "", 0),
		log.New(os.Stderr, "", 0),
		getConn(),
		make(chan *command.Request),
		make(chan int), false}
}

func (c *Client) Echo(msg string) {
	if c.running {
		cmd := command.NewRequest(command.ECHO, command.HOST, msg)
		c.cmdc <- &cmd
	}
}

func (c *Client) Ping() {
	if c.running {
		cmd := command.NewRequest(command.PING, command.HOST, nil)
		c.cmdc <- &cmd
	}
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
				c.sl.Println("Client stopped")
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
	var res string
	method := command.GenServiceMethod(action, target)

	switch action {
	case command.PING:
		err = c.c.Call(method, (*cmd).GetParams(), &res)
		break
	case command.ECHO:
		err = c.c.Call(method, (*cmd).GetParams(), &res)
		break
	}

	if err != nil {
		c.el.Println("Error from command", err)
		return
	} else {
		c.sl.Println(res)
	}
}
