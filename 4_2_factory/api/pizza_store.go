package api

import "fmt"

// Abstract Interface
type IPizzaStore interface {
	OrderPizza(pizzaType string) IPizza
	CreatePizza(pizzaType string) (IPizza, error)
}

// Abstract Concrete Type
type APizzaStore struct {
	CreatePizza func(pizzaType string) (IPizza, error)
}

func (a *APizzaStore) OrderPizza(pizzaType string) IPizza {
	if pizza, err := a.CreatePizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
		return pizza
	}
}

/**
 * Each subclass implements the createPizza() method
 * and make use of the orderPizza() method in the parent class
 */
type NyPizzaStore struct {
	*APizzaStore
}

func NewNYPizzaStore() IPizzaStore {
	basePizzaStore := &APizzaStore{}

	nyPizzaStore := &NyPizzaStore{basePizzaStore}

	nyPizzaStore.APizzaStore.CreatePizza = nyPizzaStore.CreatePizza

	return nyPizzaStore
}

func (n *NyPizzaStore) CreatePizza(pizzaType string) (IPizza, error) {
	switch pizzaType {
	case "cheese":
		return NewNYStyleCheesePizza(), nil
		//case "greek":
		//	return newNYStyleGreekPizza(), nil
		//case "pepperoni":
		//	return newNYStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}

type ChicagoPizzaStore struct {
	*APizzaStore
}

func NewChicagoPizzaStore() IPizzaStore {
	basePizzaStore := &APizzaStore{}

	chicagoPizzaStore := &ChicagoPizzaStore{basePizzaStore}

	chicagoPizzaStore.APizzaStore.CreatePizza = chicagoPizzaStore.CreatePizza

	return chicagoPizzaStore
}

func (c *ChicagoPizzaStore) CreatePizza(pizzaType string) (IPizza, error) {
	switch pizzaType {
	case "cheese":
		return NewChicagoStyleCheesePizza(), nil
		//case "greek":
		//	return newChicagoStyleGreekPizza(), nil
		//case "pepperoni":
		//	return newChicagoStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}
