package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type PackageJsonRequired struct {
	Main string `json:"main"`
}

func VerifyPackageJson(fpath string) error {
	if _, err := os.Stat(fpath); err != nil {
		return errors.New("package.json not found in instance root")
	}
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	req := PackageJsonRequired{}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		return errors.New("unable to parse package.json")
	}
	if req.Main == "" {
		return errors.New("'main' not found in package.json")
	}

	return nil
}

func GetPackageJsonMain(fpath string) string {
	if _, err := os.Stat(fpath); err != nil {
		return ""
	}
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return ""
	}
	req := PackageJsonRequired{}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		return ""
	}
	return req.Main
}
