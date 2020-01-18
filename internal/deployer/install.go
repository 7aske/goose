package deployer

import (
	"../instance"
	dutils "./utils"
	"errors"
	"fmt"
	"os"
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
	err = dutils.RunNpmScript([]string{"install"}, inst.Root, []string{})
	err = saveInstance(*inst)
	if err != nil {
		return err
	}
	return nil

}

func (d *Type) installNpm(inst *instance.JSON) error {
	packageJsonPath := path.Join(inst.Root, "package.json")
	err := dutils.VerifyPackageJson(packageJsonPath)
	if err != nil {
		return err
	}

	err = dutils.RunNpmScript([]string{"install"}, inst.Root, []string{})
	if err != nil {
		return err
	}

	if err = dutils.VerifyPackageJsonFieldList(packageJsonPath, []string{"build"}); err != nil {
		_, _ = fmt.Fprint(os.Stderr, "package.json missing build 'script'")
	} else {
		err = dutils.RunNpmScript([]string{"run", "build"}, inst.Root, []string{})
		if err != nil {
			return err
		}
	}
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
