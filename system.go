package main

import (
	"github.com/antonholmquist/jason"
	"log"
)

// One data point
type DataPoint struct {
	Timestamp int
	Status    uint64
}

// One value's history
type Metric struct {
	Name string
	Data [144]DataPoint
}

// One system's status history
type System struct {
	Name     string
	Commands []string
	Metrics  []Metric
}

// All system's status history
type Status struct {
	Systems []System
}

func ParseSystems(systemCfg []*jason.Object) Status {
	var status Status
	for _, system := range systemCfg {
		name, err1 := system.GetString("name")
		commands, err2 := system.GetStringArray("commands")
		if err1 == nil && err2 == nil {
			var metric []Metric
			status.Systems = append(status.Systems, System{name, commands, metric})
		}
	}
	return status
}

func ReportMetrics(status Status) {
	for _, system := range status.Systems {
		log.Println(system.Name)
	}
}
