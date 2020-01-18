package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

/*
router:
  port: 8080
  host: hostname

auth:
  secret: secret
  user: admin
  pass: admin

deployer:
  root: apps
*/

var Config *Type = nil

const configPath = "api/config.yaml"

type Type struct {
	Deployer struct {
		AppRoot string `yaml:"approot,omitempty"`
		Root    string `yaml:"root,omitempty"`
	}
	Router struct {
		Port     int    `yaml:"port"`
		Hostname string `yaml:"hostname"`
	}
	Auth struct {
		Secret string `yaml:"secret"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
	}
}

func Get() *Type {
	if Config == nil {
		return Parse()
	}
	return Config
}

func Parse() *Type {
	t := new(Type)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if t.Deployer.Root == "" {
		home := os.Getenv("HOME")
		t.Deployer.Root = path.Join(home, ".local/share/goose")
	}
	if t.Deployer.AppRoot == "" {
		t.Deployer.AppRoot = path.Join(t.Deployer.Root, "instances")
	}

	Config = t
	return t
}
