package api

import "fmt"

/**
 * We will start with the abstract implement of Pizza.
 * All concrete Pizzas will derive from this.
 */
type IPizza interface {
	prepare()
	bake()
	cut()
	box()
	GetName() string
}

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

// Implementing the default preparation steps for pizza: prepare, bake, cut, box
func (p *Pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *Pizza) bake() {
	fmt.Println("Baking Pizza for 25 min at 350")
}

func (p *Pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *Pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *Pizza) GetName() string {
	return p.name
}

type NyStyleCheesePizza struct {
	*Pizza
}

func NewNYStyleCheesePizza() IPizza {
	p := &Pizza{
		name:     "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}

	return &NyStyleCheesePizza{
		Pizza: p,
	}
}

type ChicagoStyleCheesePizza struct {
	*Pizza
}

func NewChicagoStyleCheesePizza() IPizza {
	p := &Pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}

	return &ChicagoStyleCheesePizza{
		Pizza: p,
	}
}

/**
 * The Chicago Style Pizza overrides the cut method
 * so that the pieces are cut in squares.
 */
func (c *ChicagoStyleCheesePizza) cut() {
	fmt.Println("Cutting the pizza into square slices")
}
