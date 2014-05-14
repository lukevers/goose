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
| We are defining a slice of `Route`s and will make
| sure that all routes defined here will be in the
| project. `Type` can be one of the following:
|
|     GET, POST, PATCH, PUT, DELETE, OPTIONS
|
| `Path` is a [martini] route. Goose uses [martini] 
| for routing. For a more comprehensive explanation
| of routing possible, view their documentation.
|
|     https://github.com/go-martini/martini
|
| `Handler` is the func in our `app/handlers` folder
| that handles the route and returns the view. 
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
