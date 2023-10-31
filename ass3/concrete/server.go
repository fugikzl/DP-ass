package concrete

import (
	"ass3/contracts"
	"fmt"
)

var instance *server

type server struct {
	nodes    []contracts.Node
	commands []contracts.Command
	open     bool
}

func GetServer() *server {
	if instance == nil {
		instance = &server{
			nodes:    make([]contracts.Node, 0),
			commands: make([]contracts.Command, 0),
			open:     false,
		}
	}

	return instance
}

func (s *server) Close() {
	fmt.Println("Server is closed for connections")
	s.open = false
}

func (s *server) Open() {
	fmt.Println("Server is opened for connections")
	s.open = true
}

func (s *server) Ready() bool {
	return s.open
}

func (s *server) SendMessage(targetNode contracts.Node, message string) bool {
	if !s.isConnected(targetNode) {
		fmt.Println("Server is not connected to target " + (targetNode).GetAddress())
		return false
	}

	fmt.Println("server received message from " + (targetNode).GetAddress() + " message: " + message)
	return true
}

func (s *server) EstablishConnectionRequest(targetNode contracts.Node) bool {
	if !s.Ready() {
		fmt.Println("Server is not ready for connection")
		return false
	}

	if s.isConnected(targetNode) {
		fmt.Println("Server is already connected with " + targetNode.GetAddress())
		return false
	}

	fmt.Println("Server made connection with node " + targetNode.GetAddress())
	s.nodes = append(s.nodes, targetNode)
	return true
}

func (s *server) isConnected(targetNode contracts.Node) bool {
	for _, node := range s.nodes {
		if node == targetNode {
			return true
		}
	}

	return false
}

func (s *server) Start() {
	for _, command := range s.commands {
		command.Execute()
	}
}

func (s *server) SetCommands(commands []contracts.Command) {
	s.commands = commands
}
