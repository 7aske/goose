package deployer

import (
	"../instance"
	"errors"
)

func (d *Type) Run(inst *instance.Instance) (*instance.Instance, error) {
	switch inst.Backend {
	case instance.Node:
		return runNode(inst)
	case instance.Python:
		return runPython(inst)
	case instance.Web:
		return runWeb(inst)
	case instance.Flask:
		return runFlask(inst)
	case instance.Npm:
		return runNpm(inst)
	}
	return inst, errors.New("backend not implemented")
}

func runNode(inst *instance.Instance) (*instance.Instance, error) {
	return inst, nil
}

func runNpm(inst *instance.Instance) (*instance.Instance, error) {
	return inst, nil
}

func runPython(inst *instance.Instance) (*instance.Instance, error) {
	return inst, nil
}

func runFlask(inst *instance.Instance) (*instance.Instance, error) {
	return inst, nil
}

func runWeb(inst *instance.Instance) (*instance.Instance, error) {
	return inst, nil
}
