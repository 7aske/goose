package app

import (
	"../../../../config"
	"../../../../deployer"
	"../../../../instance"
	"../../../utils"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func DeployRoute(router *mux.Router) {
	router.PathPrefix("/deploy").Methods("POST").HandlerFunc(deployPost)
}

func deployPost(writer http.ResponseWriter, req *http.Request) {
	jsonBody, err := utils.GetJsonStringMap(&req.Body)
	if err != nil {
		writeErrorResponse(writer, err)
		return
	}
	repo := jsonBody["repo"]
	backend := jsonBody["backend"]
	hostname := jsonBody["hostname"]

	if repo == "" || backend == "" || hostname == "" {
		writeErrorResponse(writer, errors.New("invalid arguments"))
		return
	}

	inst := instance.ToJSONStruct(instance.New(repo, hostname, instance.Backend(backend)))

	err = deployer.Deployer.Deploy(inst)
	if err != nil {
		if strings.HasPrefix(inst.Root, config.Config.Deployer.AppRoot) &&
			err.Error() != "repository already deployed" {
			_ = deployer.Deployer.Remove(*inst)
		}
		writeErrorResponse(writer, err)
		return
	}

	err = deployer.Deployer.Install(inst)
	if err != nil {
		if strings.HasPrefix(inst.Root, config.Config.Deployer.AppRoot) {
			_ = deployer.Deployer.Remove(*inst)
		}
		writeErrorResponse(writer, err)
		return
	}

	resp := struct {
		Message  string         `json:"message"`
		Instance *instance.JSON `json:"instance"`
	}{"instance deployed", inst}
	bytes, _ := json.Marshal(&resp)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(bytes)
}
