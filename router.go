package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rojoherrero/learning_go_web/handlers"
)

// NewRouter comment
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = handlers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
