package main

import (
	"github.com/martini-contrib/render"
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

func ReportMetrics(systems *[]System, r render.Render) {
	result := make(map[string]interface{})
	for _, system := range *systems {
		log.Println(system.Name)
		result[system.Name] = system.Metrics
	}
	r.JSON(200, result)
}
