package api

import (
	"./app"
	"github.com/gorilla/mux"
)

var api *mux.Router = nil

func Route(router *mux.Router) *mux.Router {
	api = router.PathPrefix("/api").Subrouter()
	app.Route(api)
	return api
}