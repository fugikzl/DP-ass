package decoratorsCake

type StrawberryCake struct {
	Cake Cake
}

func (s *StrawberryCake) GetPrice() uint16 {
	return 7 + s.Cake.GetPrice()
}

func (s *StrawberryCake) GetIngredients() []string {
	return append(s.Cake.GetIngredients(), "strawberry")
}
