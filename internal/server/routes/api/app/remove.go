package app

import (
	"../../../../deployer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type RemoveBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var remove *mux.Router = nil

func RemoveRoute(router *mux.Router) *mux.Router {
	remove = router.PathPrefix("/remove").Subrouter()
	remove.Methods("POST").HandlerFunc(removePost)
	return remove
}

func removePost(writer http.ResponseWriter, req *http.Request) {
	body := RemoveBody{}
	jsonBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	err = json.Unmarshal(jsonBytes, &body)
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	var query string
	if body.Name != "" {
		query = body.Name
	} else if body.Id != "" {
		query = body.Id
	} else {
		writer.WriteHeader(400)
		return
	}

	if inst, ok := deployer.GetDeployedInstance(query); ok {
		err = deployer.Deployer.Remove(inst)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			writer.WriteHeader(500)
			return
		} else {
			resp := struct {
				Message string `json:"message"`
				Query   string `json:"query"`
			}{"Instance removed", query}
			bytes, _ := json.Marshal(&resp)
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(200)
			writer.Write(bytes)
			return
		}

	} else {
		resp := struct {
			Message string `json:"message"`
			Query   string `json:"query"`
		}{"Instance not found.", query}
		bytes, _ := json.Marshal(&resp)
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(400)
		writer.Write(bytes)
	}
}
