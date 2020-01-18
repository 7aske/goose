package deployer

import (
	"../instance"
	"github.com/pkg/errors"
)

func (d *Type) Kill(inst *instance.Instance) error {
	if inst, ok := GetRunningInstance(inst); ok {
		err := inst.Process().Kill()
		if err != nil {
			return err
		}
		return d.removeInstance(inst)
	} else {
		return errors.New("instance not running")
	}
}
