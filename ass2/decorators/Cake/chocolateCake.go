package decoratorsCake

type ChocolateCake struct {
	Cake Cake
}

func (s *ChocolateCake) GetPrice() uint16 {
	return 15 + s.Cake.GetPrice()
}

func (s *ChocolateCake) GetIngredients() []string {
	return append(s.Cake.GetIngredients(), "chocolate")
}
