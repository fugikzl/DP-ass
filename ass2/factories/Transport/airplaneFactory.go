package factoiresTransport

type airplaneFactory struct {
}

type airplane struct {
}

func (b *airplane) GetName() string {
	return "airplane"
}

func NewAirplaneFactory() *airplaneFactory {
	return &airplaneFactory{}
}

func (f *airplaneFactory) ConstructTransport() transport {
	return &airplane{}
}
