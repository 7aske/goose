package deployer

import (
	"../config"
	"../instance"
	"../port"
	sutils "../server/utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/teris-io/shortid"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func (d *Type) Deploy(inst *instance.JSON) error {
	// TODO check if runner valid

	repo, err := sutils.FixUrl(inst.Repo)

	if err != nil {
		return err
	}

	inst.Root = path.Join(config.Get().Deployer.AppRoot, inst.Name)
	if p, err := port.New(); err != nil {
		return err
	} else {
		inst.Port = uint(p)
	}

	gitCmd := exec.Command("git", "-C", config.Get().Deployer.AppRoot, "clone", repo)
	gitCmd.Stdin = nil
	gitCmd.Stdout = os.Stdout
	var errBuf bytes.Buffer
	gitCmd.Stderr = &errBuf

	if err = gitCmd.Run(); err != nil {
		errStr := string(errBuf.Bytes())

		if strings.HasSuffix(errStr, "already exists and is not an empty directory.\n") {
			return errors.New("repository already deployed")
		} else if strings.HasSuffix(errStr, "No such device or address\n") {
			return errors.New("invalid repository or private")
		} else if strings.Contains(errStr, "unable to find remote helper") {
			return errors.New(fmt.Sprintf("unable to find remote protocol helper"))
		}
		_, _ = fmt.Fprintln(os.Stderr, errStr)
		_, _ = fmt.Fprintln(os.Stderr, "error case not implemented")
		return errors.New(errStr)
	}

	inst.Id = shortid.MustGenerate()
	inst.Deployed = time.Now()
	inst.LastUpdated = time.Now()
	err = d.addInstanceJSON(*inst)
	if err != nil {
		return err
	}
	return nil
}
