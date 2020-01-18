package instance

import (
	"time"
)

type JSON struct {
	Id          string    `json:"id"`
	Repo        string    `json:"repo"`
	Name        string    `json:"name"`
	Root        string    `json:"root"`
	Port        uint      `json:"port"`
	Hostname    string    `json:"hostname"`
	Deployed    time.Time `json:"deployed"`
	LastUpdated time.Time `json:"last_updated"`
	LastRun     time.Time `json:"last_run"`
	//Uptime      int64       `json:"uptime"`
	Backend Backend `json:"backend"`
	//Pid         int         `json:"pid"`
	//Process     *os.Process `json:"process"`
}

type File struct {
	Instances []JSON
}

type IFile interface {
	RemoveInstanceJSON(instance JSON) error
	AddInstanceJSON(instance JSON) error
}
