package deployer

import (
	"../instance"
	"../utils"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path"
)

func (d *Type) Install(inst *instance.Instance) (*instance.Instance, error) {
	switch inst.Backend {
	case instance.Node:
		return d.installNode(inst)
	case instance.Python:
		return d.installPython(inst)
	case instance.Web:
		return d.installWeb(inst)
	case instance.Flask:
		return d.installFlask(inst)
	case instance.Npm:
		return d.installNpm(inst)
	}
	return inst, errors.New("backend not implemented")
}

func (d *Type) installNode(inst *instance.Instance) (*instance.Instance, error) {
	if !utils.PathExists(path.Join(inst.Root, "package.json")) {
		return nil, errors.New("package.json not found in instance root")
	}
	npm := exec.Command("npm", "install")
	npm.Dir = inst.Root
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr
	err := npm.Run()
	if err := npm.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			_, _ = fmt.Fprintln(os.Stderr, exitError)
			return nil, err
		}
	}
	err = saveInstance(instance.ToJSON(inst))
	if err != nil {
		return nil, err
	}
	return inst, nil

}

func (d *Type) installNpm(inst *instance.Instance) (*instance.Instance, error) {

	return inst, nil
}

func (d *Type) installPython(inst *instance.Instance) (*instance.Instance, error) {

	return inst, nil
}

func (d *Type) installFlask(inst *instance.Instance) (*instance.Instance, error) {

	return inst, nil
}

func (d *Type) installWeb(inst *instance.Instance) (*instance.Instance, error) {

	return inst, nil
}
