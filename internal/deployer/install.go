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

func (d *Type) Install(inst instance.JSON) (instance.JSON, error) {
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

func (d *Type) installNode(inst instance.JSON) (instance.JSON, error) {
	if !utils.PathExists(path.Join(inst.Root, "package.json")) {
		return instance.JSON{}, errors.New("package.json not found in instance root")
	}
	npm := exec.Command("npm", "install")
	npm.Dir = inst.Root
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr
	err := npm.Run()
	if err := npm.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			_, _ = fmt.Fprintln(os.Stderr, exitError)
			return instance.JSON{}, err
		}
	}
	err = saveInstance(inst)
	if err != nil {
		return instance.JSON{}, err
	}
	return inst, nil

}

func (d *Type) installNpm(inst instance.JSON) (instance.JSON, error) {

	return inst, nil
}

func (d *Type) installPython(inst instance.JSON) (instance.JSON, error) {

	return inst, nil
}

func (d *Type) installFlask(inst instance.JSON) (instance.JSON, error) {

	return inst, nil
}

func (d *Type) installWeb(inst instance.JSON) (instance.JSON, error) {

	return inst, nil
}
