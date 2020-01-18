package app

import (
	"../../../../deployer"
	"../../../../instance"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type SearchBody struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Running bool   `json:"running"`
}
type SearchResponse struct {
	Deployed []instance.JSON     `json:"deployed"`
	Running  []instance.Instance `json:"running"`
}

var search *mux.Router = nil

func SearchRoute(router *mux.Router) *mux.Router {
	search = router.PathPrefix("/search").Subrouter()
	search.Methods("GET").HandlerFunc(searchGet)
	return search
}
func searchGet(writer http.ResponseWriter, req *http.Request) {
	body := SearchBody{}

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

	resp := SearchResponse{}

	insts := deployer.Deployer.Running
	for _, inst := range insts {
		resp.Running = append(resp.Running, *inst)
	}
	instsDep, err := deployer.GetDeployedInstances()
	if err == nil {
		resp.Deployed = append(resp.Deployed, instsDep...)
	}

	bytes, _ := json.Marshal(&resp)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(bytes)

}