package decoratorsCake

type VanilCake struct {
	Cake Cake
}

func (s *VanilCake) GetPrice() uint16 {
	return 3 + s.Cake.GetPrice()
}

func (s *VanilCake) GetIngredients() []string {
	return append(s.Cake.GetIngredients(), "vanilla")
}
