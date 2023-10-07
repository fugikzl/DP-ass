package factoiresTransport

type transport interface {
	GetName() string
}

type TransportFactory interface {
	ConstructTransport() transport
}
