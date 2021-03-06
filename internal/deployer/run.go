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
	if _, ok := GetRunningInstanceById(inst.Id); ok {
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
	wr, _ := dutils.SetUpLog(Config.Deployer.LogRoot, inst.Name, "run_out", os.Stdout)
	wre, _ := dutils.SetUpLog(Config.Deployer.LogRoot, inst.Name, "run_err", os.Stderr)
	node.Stdout = wr
	node.Stderr = wre

	node.Dir = inst.Root
	node.Env = os.Environ()
	node.Env = append(node.Env, fmt.Sprintf("PORT=%d", inst.Port))

	err = node.Start()
	if err != nil {
		return nil, err
	}
	running := instance.FromJSONStruct(inst)

	running.Pid = node.Process.Pid
	running.SetProcess(node.Process)
	running.LastRun = time.Now()
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
	running.LastRun = time.Now()
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

func (d *Type) runPython(inst instance.JSON) (*instance.Instance, error) {
	return nil, errors.New("backend not implemented")
}

func (d *Type) runFlask(inst instance.JSON) (*instance.Instance, error) {
	python := exec.Command(fmt.Sprintf("%s/venv/bin/flask", inst.Root), "run", "--host=0.0.0.0")
	python.Dir = inst.Root
	python.Env = os.Environ()
	python.Env = append(python.Env, fmt.Sprintf("FLASK_RUN_PORT=%d", inst.Port))
	dutils.SourceVenv(inst.Root, &python.Env)

	wr, _ := dutils.SetUpLog(Config.Deployer.LogRoot, inst.Name, "run_out", os.Stdout)
	wre, _ := dutils.SetUpLog(Config.Deployer.LogRoot, inst.Name, "run_out", os.Stderr)
	python.Stdout = wr
	python.Stderr = wre

	err := python.Start()
	if err != nil {
		return nil, err
	}
	running := instance.FromJSONStruct(inst)

	running.Pid = python.Process.Pid
	running.SetProcess(python.Process)
	running.LastRun = time.Now()
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

func (d *Type) runWeb(inst instance.JSON) (*instance.Instance, error) {
	return nil, errors.New("backend not implemented")
}
