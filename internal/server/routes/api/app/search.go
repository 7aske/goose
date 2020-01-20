package app

import (
	"../../../../deployer"
	"../../../../instance"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type SearchResponse struct {
	Deployed []instance.JSON     `json:"deployed"`
	Running  []instance.Instance `json:"running"`
}

type SearchResponseRunning struct {
	Running  bool              `json:"running"`
	Instance instance.Instance `json:"instance"`
}

type SearchResponseDeployed struct {
	Running  bool          `json:"running"`
	Instance instance.JSON `json:"instance"`
}

func SearchRoute(router *mux.Router) {
	router.PathPrefix("/search").Methods("GET").HandlerFunc(searchGet)
}

func searchGet(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("query")


	var bytes []byte
	if query == "" {
		resp := SearchResponse{}
		insts := deployer.Deployer.Running
		for _, inst := range insts {
			inst.Uptime = time.Now().Sub(inst.LastRun).Milliseconds()
			resp.Running = append(resp.Running, *inst)
		}
		instsDep, err := deployer.GetDeployedInstances()
		if err == nil {
			resp.Deployed = append(resp.Deployed, instsDep...)
		}
		bytes, _ = json.Marshal(&resp)

	} else {
		if inst, ok := deployer.GetRunningInstance(query); ok {
			resp := SearchResponseRunning{}
			inst.Uptime = time.Now().Sub(inst.LastRun).Milliseconds()
			resp.Instance = *inst
			resp.Running = true
			bytes, _ = json.Marshal(&resp)

		} else if inst, ok := deployer.GetDeployedInstance(query); ok {
			resp := SearchResponseDeployed{}
			resp.Instance = inst
			resp.Running = false
			bytes, _ = json.Marshal(&resp)
		} else {
			writeErrorResponse(writer, errors.New("instance not found"))
			return
		}

	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(bytes)
}
