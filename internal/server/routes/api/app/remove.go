package app

import (
	"../../../../deployer"
	"encoding/json"
	"errors"
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

func RemoveRoute(router *mux.Router) {
	router.PathPrefix("/remove").Methods("DELETE").HandlerFunc(removeDelete)
}

func removeDelete(writer http.ResponseWriter, req *http.Request) {
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
		err = deployer.Deployer.Remove(inst)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			writer.WriteHeader(500)
			return
		} else {
			resp := struct {
				Message string `json:"message"`
				Query   string `json:"query"`
			}{"instance removed", query}
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
		}{"instance not found", query}
		bytes, _ := json.Marshal(&resp)
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(400)
		writer.Write(bytes)
	}
}
