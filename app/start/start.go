package start

import (
	"../config"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

var Route = martini.Classic()

func Start() {
	// Configure databases
	ConfigureDatabase()

	// Start webserver
	StartWebserver()
}

func StartWebserver() {
	// Add application routes
	for _, r := range config.Routes {
		switch r.Type {
			case "GET":
				Route.Get(r.Path, r.Handler)
			case "POST":
				Route.Post(r.Path, r.Handler)
			case "PATCH":
				Route.Patch(r.Path, r.Handler)
			case "PUT":
				Route.Put(r.Path, r.Handler)
			case "DELETE":
				Route.Patch(r.Path, r.Handler)
			case "OPTIONS":
				Route.Options(r.Path, r.Handler)
		}
	}

	// Listen and serve
	listen := config.Host+":"+strconv.Itoa(config.Port)
	http.ListenAndServe(listen, Route)
}