package deployer

import (
	"../deployer/utils"
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
	if _, ok := GetRunningInstance(instance.FromJSONStruct(inst)); ok {
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
	err := utils.VerifyPackageJson(packageJsonPath)
	if inst.Port == 0 {
		p, _ := port.New()
		inst.Port = uint(p)
	}
	node := exec.Command("node", utils.GetPackageJsonMain(packageJsonPath))
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
	//go func() {
	//	n, _ := node.Process.Wait()
	//	d.RemoveApp(inst)
	//	d.logger.Log(fmt.Sprintf("run - app %s exited with code %d\r\n", inst.GetName(), n.ExitCode()))
	//
	//}()
	running.Pid = node.Process.Pid
	running.SetProcess(node.Process)
	inst.LastRun = time.Now()
	d.addInstance(running)
	err = d.addInstanceJSON(inst)
	if err != nil {
		_ = running.Process().Kill()
		return nil, errors.New("unable to save instance metadata json")
	}
	return running, nil
}

func (d *Type) runNpm(inst instance.JSON) (*instance.Instance, error) {
	return instance.FromJSONStruct(inst), nil
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
