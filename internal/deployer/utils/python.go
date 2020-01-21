package utils

import (
	"../../config"
	"os"
	"os/exec"
	"path"
	"strings"
)

func SourceVenv(pth string, env *[]string) {
	var OldPath string
	for i, e := range *env {
		if strings.HasPrefix(e, "PATH=") {
			OldPath = strings.TrimPrefix(e, "PATH=")
			*env = append((*env)[:i], (*env)[i+1:]...)
		}
	}
	VirtualEnv := path.Join(pth, "venv")
	NewPath := "PATH=" + VirtualEnv + "/bin:" + OldPath
	*env = append([]string{NewPath}, *env...)
}

func SetupPythonVenv(pth string) error {
	venv := exec.Command("env", "python3", "-m", "venv", "venv")
	venv.Env = os.Environ()
	venv.Dir = pth
	venv.Stdin = nil
	venv.Stdout = os.Stdout
	venv.Stderr = os.Stderr
	return venv.Run()
}

func InstallPythonRequirements(pth string) error {
	py := exec.Command("env", "python3", "-m", "pip", "install", "-r", "requirements.txt")
	py.Env = os.Environ()
	SourceVenv(pth, &py.Env)

	wr, _ := SetUpLog(config.Config.Deployer.LogRoot, path.Base(pth), "install_out", os.Stdout)
	wre, _ := SetUpLog(config.Config.Deployer.LogRoot, path.Base(pth), "install_err", os.Stderr)
	py.Stdout = wr
	py.Stderr = wre
	py.Stdin = nil
	py.Dir = pth
	return py.Run()
}
