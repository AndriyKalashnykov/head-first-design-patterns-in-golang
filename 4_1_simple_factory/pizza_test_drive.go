package main

import "head-first-design-patterns-in-golang/4_1_simple_factory/api"

func main() {
	nyPizzaFactory := &api.NyPizzaFactory{}

	pizzaStore := api.NewPizzaStore(nyPizzaFactory)
	pizzaStore.OrderPizza("pepperoni")
	pizzaStore.OrderPizza("new york cheese")
	pizzaStore.OrderPizza("greek")
}
