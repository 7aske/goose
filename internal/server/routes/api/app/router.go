package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var appRouter *mux.Router = nil

func Route(router *mux.Router) {
	appRouter = router.PathPrefix("/app").Subrouter()
	DeployRoute(appRouter)
	RemoveRoute(appRouter)
	SearchRoute(appRouter)
	RunRoute(appRouter)
	KillRoute(appRouter)
	UpdateRoute(appRouter)
	SettingsRoute(appRouter)
	LogRoute(appRouter)
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
