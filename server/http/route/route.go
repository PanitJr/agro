package route

import (
	"database/sql"
	"net/http"

	"agro/components/user"
	"agro/server/http/responseHelper"
	"agro/server/logger"

	"github.com/gorilla/mux"
)

type APIProvider struct {
	APIUser *user.APIUser
	routes  Routes
}

func (api *APIProvider) InitComponentsUsed(db *sql.DB) {
	api.APIUser = user.NewUserAPI(db)
}

func (api *APIProvider) GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range api.initRoutes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	responseHelper.HTTPJSONResponse("Hello, Welcome Back", nil, w)
}
