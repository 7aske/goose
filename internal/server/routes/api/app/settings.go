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
	"strconv"
)

type SettingsBody struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Settings struct {
		Port     int `json:"port"`
		Hostname string `json:"hostname"`
		Backend  string `json:"backend"`
	} `json:"settings"`
	//Settings map[string]string `json:"settings"`
}
type SettingsResponse struct {
	Message  string        `json:"message"`
	Instance instance.JSON `json:"instance"`
}

func SettingsRoute(router *mux.Router) {
	router.PathPrefix("/settings").Methods("PUT").HandlerFunc(settingsPut)
}

func settingsPut(writer http.ResponseWriter, req *http.Request) {
	body := SettingsBody{}

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
			m := make(map[string]string)
			if body.Settings.Port != 0 {
				m["port"] = strconv.Itoa(body.Settings.Port)
			}
			if body.Settings.Hostname != "" {
				m["hostname"] = body.Settings.Hostname
			}
			if body.Settings.Backend != "" {
				m["backend"] = body.Settings.Backend
			}
			err = deployer.Deployer.Settings(&depl, m)
			if err != nil {
				writeErrorResponse(writer, err)
				return
			} else {
				bytes, _ := json.Marshal(&SettingsResponse{Message: fmt.Sprintf("successfuly updated %s settings", depl.Name), Instance: depl})
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
