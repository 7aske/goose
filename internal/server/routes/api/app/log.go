package app

import (
	"../../../../config"
	"../../../../deployer"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path"
)

func LogRoute(router *mux.Router) {
	router.PathPrefix("/log").Methods("GET").HandlerFunc(logGet)
}

type logFilesResponse struct {
	Name  string   `json:"name"`
	Files []string `json:"files"`
}

type logOutputResponse struct {
	Name    string `json:"name"`
	LogName string `json:"log_name"`
	Content string `json:"content"`
}

func logGet(writer http.ResponseWriter, req *http.Request) {
	instanceQuery := req.URL.Query().Get("instance")
	logType := req.URL.Query().Get("type")
	if instanceQuery != "" && logType != "" {
		if inst, ok := deployer.GetDeployedInstance(instanceQuery); ok {
			logBytes, err := ioutil.ReadFile(path.Join(config.Config.Deployer.LogRoot, inst.Name, logType))
			if err != nil {
				writeErrorResponse(writer, errors.New("failed reading log file"))
				return
			}

			resp := logOutputResponse{}
			resp.Name = inst.Name
			resp.Content = string(logBytes)
			resp.LogName = logType

			respBytes, err := json.Marshal(&resp)
			if err != nil {
				writeErrorResponse(writer, errors.New("something went wrong"))
				return
			}

			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(respBytes)
		}
	} else if instanceQuery != "" {
		if inst, ok := deployer.GetDeployedInstance(instanceQuery); ok {
			files, err := ioutil.ReadDir(path.Join(config.Config.Deployer.LogRoot, inst.Name))
			if err != nil {
				writeErrorResponse(writer, errors.New("instance not found"))
				return
			}

			resp := logFilesResponse{}
			resp.Name = inst.Name
			resp.Files = []string{}
			for _, file := range files {
				resp.Files = append(resp.Files, file.Name())
			}

			respBytes, err := json.Marshal(&resp)
			if err != nil {
				writeErrorResponse(writer, errors.New("something went wrong"))
				return
			}

			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(respBytes)
		}
	}
}
