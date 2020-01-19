package server

import (
	"../config"
	"../deployer"
	"./routes/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func Listen(host string, port int) error {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	api.Route(router)
	log.Println("goose started")
	go ProxyListen(config.Config.Router.Hostname, config.Config.Router.Port)
	return http.ListenAndServe(net.JoinHostPort(host, strconv.Itoa(port)), router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("THEY SEE ME LOGGIN THEY HATEIN")
		host, port, err := net.SplitHostPort(r.Host)
		if err != nil {
			log.Println(err)
		}

		log.Println("PROXY", host, port)
		if inst, ok := deployer.GetRunningInstanceByHost(host); ok {
			p := strconv.Itoa(int(inst.Port))
			u, err := url.Parse("http://" + net.JoinHostPort(host, p))
			if err == nil {
				log.Println(u)
				proxy := httputil.NewSingleHostReverseProxy(u)
				proxy.ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
