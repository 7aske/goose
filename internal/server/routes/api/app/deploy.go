package app

import (
	"../../../../deployer"
	"../../../../instance"
	"../../../utils"
	"github.com/gorilla/mux"
	"errors"
	"net/http"
)

var deploy *mux.Router = nil

func DeployRoute(router *mux.Router) *mux.Router {
	deploy = router.PathPrefix("/deploy").Subrouter()
	deploy.Methods("POST").HandlerFunc(deployPost)
	return deploy
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
		writeErrorResponse(writer, errors.New("invalid arguments."))
		return
	}

	inst := instance.ToJSONStruct(instance.New(repo, hostname, instance.Backend(backend)))

	err = deployer.Deployer.Deploy(inst)
	if err != nil {
		writeErrorResponse(writer, err)
		return
	}
	err = deployer.Deployer.Install(inst)
	if err != nil {
		writeErrorResponse(writer, err)
		return
	}

	resp, err := utils.JsonStructToBody(inst)
	if err != nil {
		writeErrorResponse(writer, err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(resp)
}
