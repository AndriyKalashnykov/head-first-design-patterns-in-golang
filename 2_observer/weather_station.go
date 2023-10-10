package main

import "head-first-design-patterns-in-golang/2_observer/api"

func main() {
	weatherData := api.NewWeatherData()

	currentConditionsDisplay := api.NewCurrentConditionsDisplay()
	statisticsDisplay := api.NewStatisticsDisplay()

	weatherData.RegisterObserver(currentConditionsDisplay)
	weatherData.RegisterObserver(statisticsDisplay)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
