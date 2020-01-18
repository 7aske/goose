package app

import (
	"../../../../deployer"
	"../../../utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var deploy *mux.Router = nil

func DeployRoute(router *mux.Router) *mux.Router {
	deploy = router.PathPrefix("/deploy").Subrouter()
	deploy.Methods("POST").HandlerFunc(deployPost)
	return deploy
}

func deployPost(writer http.ResponseWriter, request *http.Request) {
	body := utils.GetJsonMap(&request.Body)
	fmt.Println(deployer.Deployer.Instances)
	for k, v := range body {
		fmt.Printf("%v %v\n", k, v)
	}
	writer.WriteHeader(200)
}
