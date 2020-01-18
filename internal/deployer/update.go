package deployer

import (
	"../config"
	"../instance"
	"bytes"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path"
	"time"
)

func (d *Type) Update(inst *instance.JSON) error {
	// TODO check if runner valid
	// TODO instance new
	// TODO check if already deployed
	gitCmd := exec.Command("git", "-C", path.Join(config.Get().Deployer.AppRoot, inst.Name), "pull")
	gitCmd.Stdin = nil
	gitCmd.Stdout = os.Stdout
	var errBuf bytes.Buffer
	gitCmd.Stderr = &errBuf
	if err := gitCmd.Run(); err != nil {
		return errors.New(string(errBuf.Bytes()))
	}
	inst.LastUpdated = time.Now()

	err := Deployer.Install(inst)
	if err != nil {
		return err
	}

	err = saveInstance(*inst)
	if err != nil {
		return err
	}
	return nil
}
