package factoiresTransport

type carFactory struct {
}

type car struct {
}

func (b *car) GetName() string {
	return "car"
}

func NewCarFactory() *carFactory {
	return &carFactory{}
}

func (f *carFactory) ConstructTransport() transport {
	return &car{}
}
