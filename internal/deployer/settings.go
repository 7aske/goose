package deployer

import (
	"../instance"
	"../port"
	"errors"
	"strconv"
)

func (d *Type) Settings(inst *instance.JSON, settings map[string]string) error {
	for k, v := range settings {
		switch k {
		case "hostname":
			if v == "" {
				return errors.New("invalid hostname")
			}
			inst.Hostname = v
		case "port":
			p, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return err
			}
			if ok, err := port.Check(int(p)); ok {
				inst.Port = uint(p)
			} else {
				if err != nil {
					return err
				}
			}
		case "backend":
			bkend := instance.Backend(v)
			if instance.IsBackendValid(bkend) {
				inst.Backend = instance.Backend(v)
			} else {
				return errors.New("invalid backend")
			}
		}
	}
	return saveInstance(*inst)
}
