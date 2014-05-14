package start

import (
	"../config"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

var m = martini.Classic()

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
				m.Get(r.Path, r.Handler)
			case "POST":
				m.Post(r.Path, r.Handler)
			case "PATCH":
				m.Patch(r.Path, r.Handler)
			case "PUT":
				m.Put(r.Path, r.Handler)
			case "DELETE":
				m.Patch(r.Path, r.Handler)
			case "OPTIONS":
				m.Options(r.Path, r.Handler)
		}
	}

	// Listen and serve
	listen := config.Host+":"+strconv.Itoa(config.Port)
	http.ListenAndServe(listen, m)
}