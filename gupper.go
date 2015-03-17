package main

import ()

func main() {
	json := ReadConfig("gupper.json")
	systems := ReadSystems(json)

	api := NewAPI(systems)
	api.AddRoutes()
	api.Run(":8080")
}
