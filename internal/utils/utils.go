package utils

import (
	"../instance"
	"os"
	"os/exec"
	"path"
	"strings"
)

func IndexOfString(q string, s *[]string) int {
	for i, str := range *s {
		if str == q {
			return i
		}
	}
	return -1
}

func ContainsFile(q string, dir *[]os.FileInfo) bool {
	for _, file := range *dir {
		if file.Name() == q {
			return true
		}
	}
	return false
}

func IndexOfInstance(instance instance.JSON, instances *[]instance.JSON) int {
	for i, inst := range *instances {
		if inst.Name == instance.Name || inst.Id == instance.Id {
			return i
		}
	}
	return -1
}

func PathExists(pth string) bool {
	_, err := os.Stat(pth)
	return err == nil
}

func MakeDirIfNotExist(pth string) error {
	if !PathExists(pth) {
		return os.MkdirAll(pth, 0775)
	}
	return nil
}

func GetAbsDir(pth string) string {
	if path.IsAbs(pth) {
		source, err := os.Readlink(pth)
		if err != nil {
			return os.Args[0]
		} else {
			return source
		}
	} else {
		out, err := exec.Command("which", pth).Output()
		if err != nil {
			return ""
		}
		link := strings.TrimRight(string(out), "\n")
		source, err := os.Readlink(link)
		if err != nil {
			return link
		} else {
			return source
		}
	}
}
