package deployer

import (
	dutils "../deployer/utils"
	"../instance"
	"../port"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

func (d *Type) Run(inst instance.JSON) (*instance.Instance, error) {
	depl := instance.FromJSONStruct(inst)
	if _, ok := GetRunningInstance(depl.Id); ok {
		return nil, errors.New("instance already running")
	}
	switch inst.Backend {
	case instance.Node:
		return d.runNode(inst)
	case instance.Python:
		return d.runPython(inst)
	case instance.Web:
		return d.runWeb(inst)
	case instance.Flask:
		return d.runFlask(inst)
	case instance.Npm:
		return d.runNpm(inst)
	}
	return nil, errors.New("backend not implemented")
}

func (d *Type) runNode(inst instance.JSON) (*instance.Instance, error) {
	packageJsonPath := path.Join(inst.Root, "package.json")
	err := dutils.VerifyPackageJson(packageJsonPath)
	if inst.Port == 0 {
		p, _ := port.New()
		inst.Port = uint(p)
	}
	node := exec.Command("node", dutils.GetPackageJsonMain(packageJsonPath))
	node.Dir = inst.Root
	node.Env = os.Environ()
	node.Env = append(node.Env, fmt.Sprintf("PORT=%d", inst.Port))
	node.Stdout = os.Stdout
	node.Stderr = os.Stderr
	err = node.Start()
	if err != nil {
		return nil, err
	}
	running := instance.FromJSONStruct(inst)

	running.Pid = node.Process.Pid
	running.SetProcess(node.Process)
	inst.LastRun = time.Now()
	d.addInstance(running)
	err = d.addInstanceJSON(inst)
	if err != nil {
		_ = running.Process().Kill()
		return nil, errors.New("unable to save instance metadata json")
	}
	go AddExitListener(running, d)
	return running, nil
}

func (d *Type) runNpm(inst instance.JSON) (*instance.Instance, error) {
	packageJsonPath := path.Join(inst.Root, "package.json")
	err := dutils.VerifyPackageJsonFieldList(packageJsonPath, []string{"start"})
	if err != nil {
		return nil, err
	}

	env := os.Environ()
	env = append(env, []string{fmt.Sprintf("PORT=%d", inst.Port)}...)
	proc, err := dutils.StartNpmScript([]string{"run", "start"}, inst.Root, env)
	if err != nil {
		return nil, err
	}

	running := instance.FromJSONStruct(inst)
	running.SetProcess(proc)
	running.Pid = proc.Pid
	d.addInstance(running)
	err = d.addInstanceJSON(inst)
	if err != nil {
		_ = running.Process().Kill()
		return nil, errors.New("unable to save instance metadata json")
	}
	go AddExitListener(running, d)
	return running, nil
}

func (d *Type) runPython(inst instance.JSON) (*instance.Instance, error) {
	return instance.FromJSONStruct(inst), nil
}

func (d *Type) runFlask(inst instance.JSON) (*instance.Instance, error) {
	return instance.FromJSONStruct(inst), nil
}

func (d *Type) runWeb(inst instance.JSON) (*instance.Instance, error) {
	return instance.FromJSONStruct(inst), nil
}
