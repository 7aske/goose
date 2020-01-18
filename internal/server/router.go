package server

import (
	"./routes/api"
	"github.com/gorilla/mux"
	"net/http"
)

func Route() {
	r := mux.NewRouter()
	_ = api.Route(r)
	_ = http.ListenAndServe(":5000", r)
}

