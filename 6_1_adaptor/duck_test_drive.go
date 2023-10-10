package main

import (
	"fmt"

	"head-first-design-patterns-in-golang/6_1_adaptor/api"
)

func main() {
	// Creating a Duck and a Turkey
	mallardDuck := &api.MallardDuck{}

	wildTurkey := &api.WildTurkey{}

	// Wrapping the Turkey in a Turkey adapter, which makes it look like a Duck
	turkeyAdapter := api.NewTurkeyAdaptor(wildTurkey)

	// Passing the Duck to a method testDuck() which expects a Duck object
	fmt.Println("The Duck says...")
	testDuck(mallardDuck)

	// Passing off Turkey as a Duck to the same method
	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(turkeyAdapter)
}

func testDuck(d api.Duck) {
	d.Quack()
	d.Fly()
}
