package main

import ()

func main() {
	json := ReadConfig("gupper.json")
	systems := ReadSystems(json)
	ReportMetrics(systems)

	api := NewAPI(systems)
	api.AddRoutes()
	api.Run(":81")
}
