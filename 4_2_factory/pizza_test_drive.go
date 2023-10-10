package main

import (
	"fmt"

	"head-first-design-patterns-in-golang/4_2_factory/api"
)

func main() {
	nyPizzaStore := api.NewNYPizzaStore()
	chicagoPizzaStore := api.NewChicagoPizzaStore()

	pizza := nyPizzaStore.OrderPizza("cheese")

	fmt.Printf("Ethan ordered %s pizza\n\n", pizza.GetName())

	pizza = chicagoPizzaStore.OrderPizza("cheese")

	fmt.Printf("Joel ordered %s pizza\n\n", pizza.GetName())

}
