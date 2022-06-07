package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Router is a wrapper around the Gorilla mux router.
func GerarateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRoutes(r)
}
