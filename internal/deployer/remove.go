package deployer

import (
	"../instance"
	"os"
	"path"
)

func (d *Type) Remove(inst instance.JSON) error {
	err := os.RemoveAll(path.Join(Config.Deployer.AppRoot, inst.Name))
	if err != nil {
		return err
	}
	return d.removeInstanceJSON(inst)
}
