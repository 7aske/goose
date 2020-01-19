package server

import (
	"../config"
	"../deployer"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

var router *mux.Router = nil

func ProxyListen(host string, port int) error {
	if router != nil {
		return errors.New("proxy already started")
	}

	log.Println("starting proxy on port", port)
	router = mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.Host)
		if err != nil {
			log.Println(err)
			w.WriteHeader(404)
		}

		if host == config.Config.Deployer.Hostname {
			p := strconv.Itoa(config.Config.Deployer.Port)
			u, err := url.Parse("http://" + net.JoinHostPort(host, p))
			if err == nil {
				log.Println("proxied request for deployer ->", host)
				proxy := httputil.NewSingleHostReverseProxy(u)
				proxy.ServeHTTP(w, r)
			} else {
				log.Println("url parsing failed for", host)
			}
		} else if inst, ok := deployer.GetRunningInstanceByHost(host); ok {
			instPort := strconv.Itoa(int(inst.Port))
			u, err := url.Parse("http://" + net.JoinHostPort(host, instPort))
			if err == nil {
				log.Println("proxied request", inst.Hostname, "->", inst.Name, inst.Port)
				proxy := httputil.NewSingleHostReverseProxy(u)
				proxy.ServeHTTP(w, r)
			} else {
				log.Println("url parsing failed for", inst.Name)
			}
		} else {
			log.Println("no instance found")
			w.WriteHeader(404)
		}
	})

	return http.ListenAndServe(net.JoinHostPort(host, strconv.Itoa(port)), router)
}
