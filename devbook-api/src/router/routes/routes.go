package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Description string
	Handler     func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
