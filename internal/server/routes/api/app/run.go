package app

import (
	"../../../../deployer"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type RunBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var run *mux.Router = nil

func RunRoute(router *mux.Router) *mux.Router {
	run = router.PathPrefix("/run").Subrouter()
	run.Methods("POST").HandlerFunc(runPost)
	return run
}

func runPost(writer http.ResponseWriter, req *http.Request) {
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
		writeErrorResponse(writer, errors.New("invalid arguments"))
		return
	}
	if inst, ok := deployer.GetDeployedInstance(query); ok {
		running, err := deployer.Deployer.Run(inst)
		if err != nil {
			writeErrorResponse(writer, err)
			return
		}
		bytes, _ := json.Marshal(running)
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(200)
		writer.Write(bytes)
		return
	} else {
		writeErrorResponse(writer, errors.New("instance not found"))
		return
	}

}
