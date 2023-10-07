package decoratorsCake

type BananaCake struct {
	Cake Cake
}

func (s *BananaCake) GetPrice() uint16 {
	return 5 + s.Cake.GetPrice()
}

func (s *BananaCake) GetIngredients() []string {
	return append(s.Cake.GetIngredients(), "banana")
}
