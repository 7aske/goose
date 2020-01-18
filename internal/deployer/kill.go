package deployer

import (
	"../instance"
	"errors"
)

func (d *Type) Kill(inst *instance.Instance) error {
	if inst.Process() != nil {
		err := inst.Process().Kill()
		if err != nil {
			return err
		}
		return d.removeInstance(inst)
	} else {
		return errors.New("instance process is nil")
	}

}
