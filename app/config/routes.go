package config

type Route struct {
	Type 		string
	Path		string
	Handler 	func() string
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
		Handler: test,
	},
}


func test() string {
	return "hi world"
}