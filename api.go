package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

type API struct {
	M       *martini.ClassicMartini
	Systems *[]System
}

func NewAPI(systems *[]System) *API {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	// m.Use(staticbin.Static("web", Asset))
	return &API{m, systems}
}

func (api *API) AddRoutes() {
	//
}

func (api *API) Run(port string) {
	api.M.RunOnAddr(port)
}
