package app

import (
	"../../../../deployer"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type KillBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Pid  uint   `json:"pid"`
}

func KillRoute(router *mux.Router) {
	router.PathPrefix("/kill").Methods("PUT").HandlerFunc(killPut)
}

func killPut(writer http.ResponseWriter, req *http.Request) {
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
	//var depl instance.JSON
	if depl, ok := deployer.GetDeployedInstance(query); !ok {
		writeErrorResponse(writer, errors.New("instance not found"))
		return
	} else {
		if inst, ok := deployer.GetRunningInstanceById(depl.Id); ok {
			err = deployer.Deployer.Kill(inst)
			if err != nil {
				writeErrorResponse(writer, err)
				return
			}
			resp := struct {
				Message string `json:"message"`
				Query   string `json:"query"`
			}{"instance killed", query}
			bytes, _ := json.Marshal(&resp)
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(400)
			writer.Write(bytes)
			return
		} else {
			writeErrorResponse(writer, errors.New("instance not running"))
			return
		}
	}
}
