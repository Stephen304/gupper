package main

import (
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := ioutil.ReadFile("gupper.json")
	if err != nil {
		log.Println("[!!!] Coudln't open config file")
		os.Exit(1)
	}
	json, err := jason.NewObjectFromBytes(f)
	if err != nil {
		log.Println("[!!!] Couldn't parse config file")
		os.Exit(1)
	}
	systemsCfg, err := json.GetObjectArray("systems")
	if err != nil {
		log.Println("[!!!] Couldn't parse systems object")
		os.Exit(1)
	}
	status := ParseSystems(systemsCfg)
	ReportMetrics(status)

	api := NewAPI()
	api.AddRoutes()
	api.Run(":81")
}
