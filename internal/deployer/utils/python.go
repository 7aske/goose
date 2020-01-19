package utils

import (
	"os"
	"os/exec"
)

func SetupPythonVenv(pth string) error {
	venv := exec.Command("python", "-m", "venv", "venv")
	venv.Env = os.Environ()
	venv.Dir = pth
	venv.Stdin = nil
	venv.Stdout = os.Stdout
	venv.Stderr = os.Stderr
	return venv.Run()
}
