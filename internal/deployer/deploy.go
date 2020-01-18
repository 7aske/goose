package deployer

import (
	"../config"
	"../instance"
	"../port"
	sutils "../server/utils"
	"fmt"
	"github.com/teris-io/shortid"
	"os"
	"os/exec"
	"path"
	"time"
)

func (d *Type) Deploy(inst *instance.Instance) (*instance.Instance, error) {
	// TODO check if runner valid

	repo, err := sutils.FixUrl(inst.Repo)
	fmt.Println(repo)
	if err != nil {
		return nil, err
	}

	inst.Root = path.Join(config.Get().Deployer.AppRoot, inst.Name)
	if p, err := port.New(); err != nil {
		return nil, err
	} else {
		inst.Port = uint(p)
	}

	// TODO instance new
	// TODO check if already deployed
	gitCmd := exec.Command("git", "-C", config.Get().Deployer.AppRoot, "clone", repo)
	gitCmd.Stdin = nil
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// repo already exists
			_, _ = fmt.Fprintln(os.Stderr, exitError)
			if exitError.ExitCode() == 128 {
				return inst, err
			}
			return nil, err
		}
	}
	inst.Id = shortid.MustGenerate()
	inst.Deployed = time.Now()
	err = d.AddInstanceJSON(instance.ToJSON(inst))
	if err != nil {
		return nil, err
	}
	// TODO save app
	return inst, nil
}
