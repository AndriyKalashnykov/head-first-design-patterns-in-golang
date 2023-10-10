package main

import "head-first-design-patterns-in-golang/1_strategy/api"

func main() {
	mallard := api.NewMallardDuck()
	mallard.PerformFly()
	mallard.PerformQuack()

	rubber := api.NewRubberDuck()
	rubber.PerformFly()
	rubber.PerformQuack()

	rubber.SetFlyBehaviour(api.FlyRocketPowered{})
	rubber.PerformFly()
	rubber.SetQuackBehaviour(api.MuteQuack{})
	rubber.PerformQuack()
}
