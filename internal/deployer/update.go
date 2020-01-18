package deployer

import (
	"../config"
	"../instance"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

func (d *Type) Update(inst instance.JSON) (instance.JSON, error) {
	// TODO check if runner valid
	// TODO instance new
	// TODO check if already deployed
	gitCmd := exec.Command("git", "-C", path.Join(config.Get().Deployer.AppRoot, inst.Name), "pull")
	gitCmd.Stdin = nil
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			_, _ = fmt.Fprintln(os.Stderr, exitError)
			return instance.JSON{}, err
		}
	}
	inst.LastUpdated = time.Now()

	inst, err := Deployer.Install(inst)
	if err != nil {
		return instance.JSON{}, err
	}

	err = saveInstance(inst)
	if err != nil {
		return instance.JSON{}, err
	}
	return inst, nil
}
