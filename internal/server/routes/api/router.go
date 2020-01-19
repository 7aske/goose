package api

import (
	"./app"
	"github.com/gorilla/mux"
)

var api *mux.Router = nil

func Route(router *mux.Router) {
	app.Route(router.PathPrefix("/api").Subrouter())
}
