package main

import "head-first-design-patterns-in-golang/6_2_facade/api"

func main() {
	homeTheater := api.NewHomeTheaterFacade()

	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}
