package commands

import (
	"ass3/concrete"
	"ass3/contracts"
)

type OpenServerCommand struct {
}

func (c *OpenServerCommand) Execute() {
	server := concrete.GetServer()
	server.Open()
}

type CloseServerCommand struct {
}

func (c *CloseServerCommand) Execute() {
	server := concrete.GetServer()
	server.Close()
}

type ConenctCommand struct {
	Node contracts.Node
}

func (c *ConenctCommand) Execute() {
	c.Node.ConnectToServer()
}

type PingCommand struct {
	Node contracts.Node
}

func (c *PingCommand) Execute() {
	server := concrete.GetServer()
	server.SendMessage(c.Node, "Ping")
}
