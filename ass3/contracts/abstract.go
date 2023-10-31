package contracts

type Node interface {
	GetAddress() string
	SendMessage() bool
	ConnectToServer() bool
}

type Server interface {
	EstablishConnectionRequest(node *Node) bool
	Ready() bool
	Open()
	Close()
}
