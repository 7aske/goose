package app

import (
	"../../../../deployer"
	"../../../../instance"
	"../../../utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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
	}
	repo := jsonBody["repo"]
	backend := jsonBody["backend"]
	hostname := jsonBody["hostname"]

	if repo == "" || backend == "" || hostname == "" {
		writeErrorResponse(writer, errors.New("invalid arguments."))
	}

	inst := instance.New(repo, hostname, instance.Backend(backend))

	inst, err = deployer.Deployer.Deploy(inst)
	if err != nil {
		writeErrorResponse(writer, err)
	}

	inst, err = deployer.Deployer.Install(inst)
	if err != nil {
		writeErrorResponse(writer, err)
	}

	instJson := instance.ToJSONStruct(inst)
	resp, err := utils.JsonStructToBody(&instJson)
	if err != nil {
		writeErrorResponse(writer, err)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(resp)
}
