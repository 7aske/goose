package app

import (
	"../../../../deployer"
	"../../../../instance"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UpdateBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type UpdateResponse struct {
	Message  string        `json:"message"`
	Instance instance.JSON `json:"instance"`
}

func UpdateRoute(router *mux.Router) {
	router.PathPrefix("/update").Methods("PUT").HandlerFunc(updatePut)
}

func updatePut(writer http.ResponseWriter, req *http.Request) {
	body := UpdateBody{}

	jsonBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writer.WriteHeader(400)
		return
	}

	err = json.Unmarshal(jsonBytes, &body)
	if err != nil {
		writeErrorResponse(writer, errors.New("invalid arguments"))
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

	if depl, ok := deployer.GetDeployedInstance(query); ok {
		if _, ok := deployer.GetRunningInstanceById(depl.Id); ok {
			writeErrorResponse(writer, errors.New("instance is running"))
			return
		} else {
			err := deployer.Deployer.Update(&depl)
			if err != nil {
				writeErrorResponse(writer, err)
				return
			} else {
				bytes, _ := json.Marshal(&UpdateResponse{Message: fmt.Sprintf("successfuly updated %s", depl.Name), Instance: depl})
				writer.Header().Add("Content-Type", "application/json")
				writer.WriteHeader(200)
				writer.Write(bytes)
			}
		}
	} else {
		writeErrorResponse(writer, errors.New("instance not found"))
		return
	}

}
