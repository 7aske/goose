package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
)

func GetJsonMap(body *io.ReadCloser) map[interface{}]interface{} {
	output := make(map[interface{}]interface{})
	jsonBytes, _ := ioutil.ReadAll(*body)
	err := json.Unmarshal(jsonBytes, &output)
	if err != nil {
		fmt.Println(err)
	}
	return output
}

func FixUrl(uString string) (string, error) {
	u, err := url.Parse(uString)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.String(), nil
}
