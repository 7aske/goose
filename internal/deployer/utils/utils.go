package utils

import "regexp"

func GetNameFromRepo(repo string) string {
	reg, _ := regexp.Compile("([^/]+$)")
	return string(reg.Find([]byte(repo)))
}