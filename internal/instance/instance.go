package instance

import (
	"../deployer/utils"
	"os"
	"time"
)

type Instance struct {
	Id          string
	Repo        string
	Name        string
	Root        string
	Port        uint
	Hostname    string
	Deployed    time.Time
	LastUpdated time.Time
	LastRun     time.Time
	Uptime      int64
	Backend     Backend
	Pid         int
	Process     *os.Process
}

func New(repo string, hostname string, backend Backend) *Instance {
	inst := new(Instance)
	inst.Repo = repo
	inst.Name = utils.GetNameFromRepo(repo)
	inst.Hostname = hostname
	inst.Backend = backend
	return inst
}

func FromJSONStruct(json JSON) *Instance {
	inst := new(Instance)
	inst.Id = json.Id
	inst.Repo = json.Repo
	inst.Name = json.Name
	inst.Root = json.Root
	inst.Port = json.Port
	inst.Hostname = json.Hostname
	inst.Deployed = json.Deployed
	inst.LastUpdated = json.LastUpdated
	inst.LastRun = json.LastRun
	inst.Uptime = 0
	inst.Backend = json.Backend
	inst.Pid = -1
	inst.Process = nil
	return inst
}

func ToJSONStruct(inst *Instance) *JSON {
	json := new(JSON)
	json.Id = inst.Id
	json.Repo = inst.Repo
	json.Name = inst.Name
	json.Root = inst.Root
	json.Port = inst.Port
	json.Hostname = inst.Hostname
	json.Deployed = inst.Deployed
	json.LastUpdated = inst.LastUpdated
	json.LastRun = inst.LastRun
	json.Backend = inst.Backend
	return json
}
