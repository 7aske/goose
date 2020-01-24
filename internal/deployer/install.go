package deployer

import (
	"../instance"
	dutils "./utils"
	"errors"
	"log"
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
	return d.installNpm(inst)
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
		log.Println("package.json missing build 'script'")
		return nil
	} else {
		return dutils.RunNpmScript([]string{"run", "build"}, inst.Root, []string{})
	}
}

func (d *Type) installPython(inst *instance.JSON) error {
	return errors.New("backend not implemented")
}

func (d *Type) installFlask(inst *instance.JSON) error {
	err := dutils.SetupPythonVenv(inst.Root)
	if err != nil {
		return err
	}
	return dutils.InstallPythonRequirements(inst.Root)
}

func (d *Type) installWeb(inst *instance.JSON) error {
	return errors.New("backend not implemented")
}
