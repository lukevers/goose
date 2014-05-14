package handlers

import (
	"github.com/go-martini/martini"
)

/*
|--------------------------------------------------
| Application Route Handlers
|--------------------------------------------------
|
| Application route handlers for the project should
| be defined here. Other handler files can be used
| to keep the project organized.
|
*/

func HelloWorld(params martini.Params) string {
	return "Hello, world!"
}

func HelloUser(params martini.Params) string {
	return "Hello, " + params["user"] + "!"
}
