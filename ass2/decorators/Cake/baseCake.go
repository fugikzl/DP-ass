package decoratorsCake

type BaseCake struct {
	price       uint16
	ingredients []string
}

func (cake *BaseCake) GetPrice() uint16 {
	return cake.price
}

func (cake *BaseCake) GetIngredients() []string {
	return cake.ingredients
}

func NewBaseCake(price uint16, ingredients []string) *BaseCake {
	return &BaseCake{
		price:       price,
		ingredients: ingredients,
	}
}
