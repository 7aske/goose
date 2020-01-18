package deployer

import (
	"../instance"
	dutils "./utils"
	"bytes"
	"errors"
	"fmt"
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
	err := dutils.VerifyPackageJson(path.Join(inst.Root, "package.json"))
	if err != nil {
		return err
	}
	npm := exec.Command("npm", "install")
	npm.Dir = inst.Root
	npm.Stdout = os.Stdout
	var errBuf bytes.Buffer
	npm.Stderr = &errBuf
	err = npm.Run()
	if err := npm.Run(); err != nil {
		errStr := string(errBuf.Bytes())
		_, _ = fmt.Fprintln(os.Stderr, errStr)
		return errors.New(errStr)
	}
	err = saveInstance(*inst)
	if err != nil {
		return err
	}
	return nil

}

func (d *Type) installNpm(inst *instance.JSON) error {
	return d.installNode(inst)
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
