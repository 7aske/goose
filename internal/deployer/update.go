package deployer

import (
	"../config"
	"../instance"
	"./utils"
	"bytes"
	"errors"
	"fmt"
	"io"
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
	var errBuf bytes.Buffer
	wr, _ := utils.SetUpLog(config.Config.Deployer.LogRoot, inst.Name, "update_out", os.Stdout)
	wre, _ := utils.SetUpLog(config.Config.Deployer.LogRoot, inst.Name, "update_err", os.Stderr)
	gitCmd.Stdout = wr
	mw := io.MultiWriter(wre, &errBuf)
	gitCmd.Stderr = mw
	if err := gitCmd.Run(); err != nil {
		errStr := string(errBuf.Bytes())
		_, _ = fmt.Fprintln(os.Stderr, errStr)
		return errors.New(errStr)
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
