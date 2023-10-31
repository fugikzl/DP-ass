package concrete

type Node struct {
	address string
}

func NewNode(address string) *Node {
	return &Node{
		address: address,
	}
}

func (n *Node) GetAddress() string {
	return n.address
}

func (n *Node) ConnectToServer() bool {
	server := GetServer()
	return server.EstablishConnectionRequest(n)
}

func (n *Node) SendMessage() bool {
	return true
}
