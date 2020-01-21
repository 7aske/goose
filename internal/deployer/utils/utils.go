package utils

import (
	"io"
	"log"
	"os"
	"path"
	"regexp"
)

func GetNameFromRepo(repo string) string {
	reg, _ := regexp.Compile("([^/]+$)")
	return string(reg.Find([]byte(repo)))
}

func SetUpLog(root string, name string, logType string, tee io.Writer) (io.Writer, error) {
	folderPath := path.Join(root, name)
	filePath := path.Join(folderPath, logType+".log")
	log.Println(filePath)
	err := os.MkdirAll(folderPath, 0775)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return nil, err
	}
	mw := io.MultiWriter(tee, file)
	return mw, nil
}
