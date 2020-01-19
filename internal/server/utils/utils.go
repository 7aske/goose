package utils

import (
	"../../instance"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
)

func JsonStructFromBody(body *io.ReadCloser) (instance.JSON, error) {
	inst := instance.JSON{}
	jsonBytes, err := ioutil.ReadAll(*body)
	if err != nil {
		return inst, err
	}
	err = json.Unmarshal(jsonBytes, &inst)
	if err != nil {
		return inst, err
	}
	return inst, nil
}

func JsonStructToBody(inst *instance.JSON) ([]byte, error) {
	out, err := json.Marshal(inst)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func GetJsonStringMap(body *io.ReadCloser) (map[string]string, error) {
	output := make(map[string]string)
	jsonBytes, err := ioutil.ReadAll(*body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
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
