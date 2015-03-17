package main

import (
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

func ReportMetrics(systems *[]System) {
	for _, system := range *systems {
		log.Println(system.Name)
	}
}
