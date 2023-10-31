package main

import (
	"ass3/concrete"
	"ass3/concrete/commands"
	"ass3/contracts"
)

func main() {
	node1 := concrete.NewNode("a1")
	node2 := concrete.NewNode("a2")
	node3 := concrete.NewNode("b1")

	commandsSlice := make([]contracts.Command, 0)
	commandsSlice = append(commandsSlice, &commands.OpenServerCommand{})
	commandsSlice = append(commandsSlice, &commands.PingCommand{
		Node: node1,
	})
	commandsSlice = append(commandsSlice, &commands.ConenctCommand{
		Node: node1,
	})
	commandsSlice = append(commandsSlice, &commands.PingCommand{
		Node: node1,
	})
	commandsSlice = append(commandsSlice, &commands.CloseServerCommand{})
	commandsSlice = append(commandsSlice, &commands.ConenctCommand{
		Node: node2,
	})
	commandsSlice = append(commandsSlice, &commands.ConenctCommand{
		Node: node3,
	})
	commandsSlice = append(commandsSlice, &commands.OpenServerCommand{})
	commandsSlice = append(commandsSlice, &commands.ConenctCommand{
		Node: node2,
	})
	commandsSlice = append(commandsSlice, &commands.ConenctCommand{
		Node: node3,
	})
	commandsSlice = append(commandsSlice, &commands.PingCommand{
		Node: node2,
	})
	commandsSlice = append(commandsSlice, &commands.PingCommand{
		Node: node3,
	})

	concrete.GetServer().SetCommands(commandsSlice)
	concrete.GetServer().Start()
	//  := commands.OpenServerCommand{}

}
