package utils

import (
	"log"
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
	log.Println(py.Env)
	py.Dir = pth
	py.Stdin = nil
	py.Stdout = os.Stdout
	py.Stderr = os.Stderr
	return py.Run()
}
