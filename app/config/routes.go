package config

import (
	"../handlers"
	"github.com/go-martini/martini"
)

/*
|--------------------------------------------------
| Route Type
|--------------------------------------------------
|
| The `Route` type that is defined here is used to
| control routes 
|
*/

type Route struct {
	Type 		string
	Path		string
	Handler 	func(martini.Params) string
}

/*
|--------------------------------------------------
| Application Routes
|--------------------------------------------------
|
| Application routes for the project should be
| defined here.
|
*/

var Routes = []*Route{
	&Route{
		Type: "GET",
		Path: "/",
		Handler: handlers.HelloWorld,
	},
	&Route{
		Type: "GET",
		Path: "/:user",
		Handler: handlers.HelloUser,
	},
}
