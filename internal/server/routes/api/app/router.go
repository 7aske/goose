package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var app *mux.Router = nil

func Route(router *mux.Router) *mux.Router {
	app = router.PathPrefix("/app").Subrouter()
	DeployRoute(app)
	RemoveRoute(app)
	return app
}

func writeErrorResponse(writer http.ResponseWriter, err error) {
	resp := struct {
		Message string `json:"message"`
	}{err.Error()}
	bytes, _ := json.Marshal(&resp)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(400)
	writer.Write(bytes)
	return
}
