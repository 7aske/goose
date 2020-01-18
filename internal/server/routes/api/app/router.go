package app

import (
	"github.com/gorilla/mux"
)

var app *mux.Router = nil

func Route(router *mux.Router) *mux.Router {
	app = router.PathPrefix("/app").Subrouter()
	DeployRoute(app)
	return app
}
