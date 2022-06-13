package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct that contains the path and the handler.
type Route struct {
	URI          string
	Method       string
	HandlerFunc  func(w http.ResponseWriter, r *http.Request)
	Authrequired bool
}

// ConfigureRoutes configures all the routes.
func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := usersRoute
	routes = append(routes, routeLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	return r
}
