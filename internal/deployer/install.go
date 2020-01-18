package deployer

import (
	"../instance"
	"../utils"
	"fmt"
	"errors"
	"os"
	"os/exec"
	"path"
)

func (d *Type) Install(inst *instance.JSON) error {
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
	return errors.New("backend not implemented")
}

func (d *Type) installNode(inst *instance.JSON) error {
	if !utils.PathExists(path.Join(inst.Root, "package.json")) {
		return errors.New("package.json not found in instance root")
	}
	npm := exec.Command("npm", "install")
	npm.Dir = inst.Root
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr
	err := npm.Run()
	if err := npm.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			_, _ = fmt.Fprintln(os.Stderr, exitError)
			return err
		}
	}
	err = saveInstance(*inst)
	if err != nil {
		return err
	}
	return nil

}

func (d *Type) installNpm(inst *instance.JSON) error {

	return nil
}

func (d *Type) installPython(inst *instance.JSON) error {

	return nil
}

func (d *Type) installFlask(inst *instance.JSON) error {

	return nil
}

func (d *Type) installWeb(inst *instance.JSON) error {

	return nil
}
