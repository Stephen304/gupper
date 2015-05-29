package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/antonholmquist/jason"
)

func ReadConfig(filename string) *jason.Object {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("[!!!] Coudln't open config file")
		os.Exit(1)
	}
	json, err := jason.NewObjectFromBytes(f)
	if err != nil {
		log.Println("[!!!] Couldn't parse config file")
		os.Exit(1)
	}
	return json
}

func ReadSystems(json *jason.Object) *[]System {
	syscfg, err := json.GetObjectArray("systems")
	if err != nil {
		log.Println("[!!!] Couldn't parse systems object")
		os.Exit(1)
	}
	var systems []System
	for _, system := range syscfg {
		name, err1 := system.GetString("name")
		commands, err2 := system.GetStringArray("commands")
		if err1 == nil && err2 == nil {
			// Add extra values here later - all commands should give uptime at least
			metric := []Metric{Metric{Name: "Uptime"}}
			systems = append(systems, System{name, commands, metric})
		}
	}
	return &systems
}
