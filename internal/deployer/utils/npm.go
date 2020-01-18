package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type PackageJsonRequired struct {
	Main    string `json:"main"`
	Scripts struct {
		Start string `json:"start"`
		Build string `json:"build"`
	} `json:"scripts"`
}

func VerifyPackageJson(fpath string) error {
	if _, err := os.Stat(fpath); err != nil {
		return errors.New("package.json not found in instance root")
	}
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	req := PackageJsonRequired{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		return errors.New("unable to parse package.json")
	}
	if req.Main == "" {
		return errors.New("'main' not found in package.json")
	}

	return nil
}
func VerifyPackageJsonFieldList(fpath string, fields []string) error {
	if _, err := os.Stat(fpath); err != nil {
		return errors.New("package.json not found in instance root")
	}
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	req := PackageJsonRequired{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		return errors.New("unable to parse package.json")
	}
	for _, field := range fields {
		switch field {
		case "main":
			if req.Main == "" {
				return errors.New("invalid field 'main' in package.json")
			}
		case "build":
			if req.Scripts.Build == "" {
				return errors.New("invalid field 'build' in package.json")
			}
		case "start":
			if req.Scripts.Start == "" {
				return errors.New("invalid field 'start' in package.json")
			}
		default:
			return errors.New(fmt.Sprintf("missing field '%s' in package.json", field))
		}
	}
	return nil
}

func RunNpmScript(script []string, root string, env []string) error {
	npm := exec.Command("npm", script...)
	npm.Dir = root
	npm.Env = append(npm.Env, env...)
	npm.Stdout = os.Stdout
	var errBuf bytes.Buffer
	npm.Stderr = &errBuf
	if err := npm.Run(); err != nil {
		errStr := string(errBuf.Bytes())
		_, _ = fmt.Fprintln(os.Stderr, errStr)
		return errors.New(errStr)
	}
	return nil
}
func StartNpmScript(script []string, root string, env []string) (*os.Process, error) {
	npm := exec.Command("npm", script...)
	npm.Dir = root
	npm.Env = append(npm.Env, env...)
	npm.Stdout = os.Stdout
	var errBuf bytes.Buffer
	npm.Stderr = &errBuf
	npm.Stderr = os.Stderr
	if err := npm.Start(); err != nil {
		errStr := string(errBuf.Bytes())
		_, _ = fmt.Fprintln(os.Stderr, errStr)
		return nil, err
	}
	return npm.Process, nil
}

func GetPackageJsonMain(fpath string) string {
	if _, err := os.Stat(fpath); err != nil {
		return ""
	}
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return ""
	}
	req := PackageJsonRequired{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		return ""
	}
	return req.Main
}
