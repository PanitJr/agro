package route

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes array is a
type Routes []Route

func (api *APIProvider) initRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},

		Route{"CreateUser", "POST", "/user", api.APIUser.Create},
		Route{"UpdateUser", "PUT", "/user", api.APIUser.Update},
		Route{"DeleteUser", "DELETE", "/user", api.APIUser.Delete},
		Route{"GetUser", "GET", "/user/{id}", api.APIUser.Get},
		Route{"ListUser", "GET", "/user", api.APIUser.List},
	}
	return routes
}
