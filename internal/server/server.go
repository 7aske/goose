package server

import (
	"../config"
	"./routes/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
)

func Listen(host string, port int) error {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	api.Route(router)
	log.Println("goose started")
	go ProxyListen(config.Config.Router.Hostname, config.Config.Router.Port)
	return http.ListenAndServe(net.JoinHostPort(host, strconv.Itoa(port)), handlers.CORS()(router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
