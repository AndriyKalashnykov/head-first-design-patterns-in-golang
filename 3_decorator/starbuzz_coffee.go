package main

import (
	"fmt"

	"head-first-design-patterns-in-golang/3_decorator/api"
)

func main() {
	// Order up an espresso, no condiments
	beverage1 := &api.Espresso{}
	fmt.Println("Order1:", beverage1.GetDescription())
	fmt.Println("Cost:", beverage1.Cost())

	// Create a DarkRoast coffee with milk and mocha
	beverage2 := api.NewMocha(api.NewMilk(&api.DarkRoast{}))

	fmt.Println("Order2:", beverage2.GetDescription())
	fmt.Println("Cost:", beverage2.Cost())

	// Add a new decorator (Soy) at runtime
	beverage2 = api.NewSoy(beverage2)

	fmt.Println("Updated Order2:", beverage2.GetDescription())
	fmt.Println("Updated Cost:", beverage2.Cost())
}
