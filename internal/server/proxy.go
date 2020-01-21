package server

import (
	"../auth"
	"../config"
	"../deployer"
	"./responses"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

var router *mux.Router = nil

func ProxyListen(host string, port int) error {
	if router != nil {
		return errors.New("proxy already started")
	}

	log.Println("starting proxy on port", port)
	router = mux.NewRouter()

	router.Path("/auth").Methods("POST").HandlerFunc(authRoute)
	router.Path("/validate").Methods("POST").HandlerFunc(validateRoute)

	router.PathPrefix("/").HandlerFunc(proxyRoute)

	return http.ListenAndServe(net.JoinHostPort(host, strconv.Itoa(port)), router)
}

func proxyRoute(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(responses.NotFound))
		return
	}

	if host == config.Config.Deployer.Hostname {
		p := strconv.Itoa(config.Config.Deployer.Port)
		u, err := url.Parse("http://" + net.JoinHostPort(host, p))
		if err != nil {
			log.Println("url parsing failed for", host)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(responses.BadRequest))
			return
		}

		if authorizeRequest(r) {
			log.Println("proxied request for deployer ->", host)
			proxy := httputil.NewSingleHostReverseProxy(u)
			proxy.ServeHTTP(w, r)
		} else {
			log.Println("unauthorized request for deployer ->", host)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(responses.Unauthorized))
			return
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(responses.BadRequest))
		}
	} else {
		log.Println("no instance found")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(responses.NotFound))
	}
}

type authBody struct {
	Token string `json:"token"`
}

func authorizeRequest(r *http.Request) bool {
	if r.URL.Path == "/auth" {
		return true
	}

	if !config.Config.Auth.Enable {
		return true
	}

	if cookie, err := r.Cookie("Authorization"); err == nil {
		if strings.HasPrefix(cookie.Value, "Bearer") {
			token := strings.Split(cookie.Value, "Bearer ")[1]
			if auth.VerifyToken(token) {
				return true
			}
		}
	} else if header := r.Header.Get("Authorization"); header != "" {
		if strings.HasPrefix(header, "Bearer") {
			token := strings.Split(header, "Bearer ")[1]
			if auth.VerifyToken(token) {
				return true
			}
		}
	} else {
		body := authBody{}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		log.Println(bodyBytes)
		if err == nil {
			err = json.Unmarshal(bodyBytes, &body)
			r.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
			if err == nil {
				token := body.Token
				if auth.VerifyToken(token) {
					return true

				}
			}
		}
	}

	return false
}

type loginBody struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func validateRoute(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(r) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responses.OK))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(responses.Unauthorized))
	}
}

func authRoute(w http.ResponseWriter, r *http.Request) {
	body := loginBody{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(responses.Unauthorized))
		return
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(responses.Unauthorized))
		return
	}

	if auth.VerifyCredentials(body.User, body.Pass) {
		token := auth.GenerateToken()
		log.Println(token)
		resp := loginResponse{}
		resp.Token = token

		respBytes, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(responses.Unauthorized))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(respBytes)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(responses.Unauthorized))
		return
	}
}
