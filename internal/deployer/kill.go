package deployer

import (
	"../instance"
	"errors"
	"fmt"
	"os"
	"time"
)

func (d *Type) Kill(inst *instance.Instance) error {
	if inst.Process() != nil {
		err := inst.Process().Signal(os.Interrupt)
		go func() {
			time.Sleep(time.Millisecond * 1000)
			if inst != nil && inst.Process() != nil {
				err = inst.Process().Signal(os.Kill)
				if err != nil {
					_, _ = fmt.Fprintln(os.Stderr, err)
				}
			}
		}()
		if err != nil {
			return err
		}
		return d.removeInstance(inst)
	} else {
		return errors.New("instance process is nil")
	}

}
