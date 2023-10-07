package main

import (
	decoratorsCake "ass2/decorators/Cake"
	factoiresTransport "ass2/factories/Transport"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	cake := decoratorsCake.NewBaseCake(10, []string{"sugar", "bisquit"})
	fmt.Printf("Base cake price is: %d, and ingredients are: %s\n", cake.GetPrice(), strings.Join(cake.GetIngredients(), ", "))

	finalCake := decoratorsCake.VanilCake{Cake: &decoratorsCake.ChocolateCake{Cake: &decoratorsCake.BananaCake{Cake: &decoratorsCake.StrawberryCake{Cake: cake}}}}
	fmt.Printf("Final cake price is: %d, and ingredients are: %s\n", finalCake.GetPrice(), strings.Join(finalCake.GetIngredients(), ", "))

	transportFactories := []factoiresTransport.TransportFactory{factoiresTransport.NewAirplaneFactory(), factoiresTransport.NewBicycleFactory(), factoiresTransport.NewCarFactory()}
	ourChoice := rand.Intn(3)
	ourTransportFactory := transportFactories[ourChoice]
	ourTransport := ourTransportFactory.ConstructTransport()

	fmt.Printf("We have travelled with %s\n", ourTransport.GetName())
}
