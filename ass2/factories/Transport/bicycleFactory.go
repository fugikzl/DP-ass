package factoiresTransport

type bicycleFactory struct {
}

type bicycle struct {
}

func (b *bicycle) GetName() string {
	return "bicycle"
}

func NewBicycleFactory() *bicycleFactory {
	return &bicycleFactory{}
}

func (f *bicycleFactory) ConstructTransport() transport {
	return &bicycle{}
}
