package instance

import (
	"../deployer/utils"
	"os"
	"time"
)

type Instance struct {
	Id          string    `json:"id"`
	Repo        string    `json:"repo"`
	Name        string    `json:"name"`
	Root        string    `json:"root"`
	Port        uint      `json:"port"`
	Hostname    string    `json:"hostname"`
	Deployed    time.Time `json:"deployed"`
	LastUpdated time.Time `json:"last_updated"`
	LastRun     time.Time `json:"last_run"`
	Uptime      int64     `json:"uptime"`
	Backend     Backend   `json:"backend"`
	Pid         int       `json:"pid"`
	process     *os.Process
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
	inst.process = nil
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

func (d *Instance) Process() *os.Process {
	return d.process
}

func (d *Instance) SetProcess(proc *os.Process) {
	d.process = proc
}
