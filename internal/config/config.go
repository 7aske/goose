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
		AppRoot  string `yaml:"approot,omitempty"`
		Root     string `yaml:"root,omitempty"`
		Port     int    `yaml:"port"`
		Hostname string `yaml:"hostname"`
	}
	Router struct {
		Port     int    `yaml:"port"`
		Hostname string `yaml:"hostname"`
	}
	Auth struct {
		Enable bool   `yaml:"enable"`
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

	if t.Deployer.Hostname == "" {
		t.Deployer.Hostname = "0.0.0.0"
		log.Println("deployer 'hostname' not set using default value - ", t.Deployer.Hostname)
	}

	if t.Deployer.Port == 0 {
		t.Deployer.Port = 5000
		log.Println("deployer 'port'     not set using default value - ", t.Deployer.Port)
	}

	if t.Router.Hostname == "" {
		t.Router.Hostname = "0.0.0.0"
		log.Println("router   'hostname' not set using default value - ", t.Router.Hostname)
	}

	if t.Router.Port == 0 {
		t.Router.Port = 8080
		log.Println("router   'port'     not set using default value - ", t.Router.Port)
	}

	if t.Deployer.Root == "" {
		home := os.Getenv("HOME")
		t.Deployer.Root = path.Join(home, ".local/share/goose")
		log.Println("deployer 'root'     not set using default value - ", t.Deployer.Root)
	}

	if t.Deployer.AppRoot == "" {
		t.Deployer.AppRoot = path.Join(t.Deployer.Root, "instances")
		log.Println("deployer 'approot'  not set using default value - ", t.Deployer.AppRoot)
	}

	if t.Auth.User == "" {
		t.Auth.User = "admin"
		log.Println("auth     'user'     not set using default value - ", t.Auth.User)
	}

	if t.Auth.Pass == "" {
		t.Auth.Pass = "admin"
		log.Println("auth     'pass'     not set using default value - ", t.Auth.Pass)
	}

	if t.Auth.Secret == "" {
		t.Auth.Secret = "secret"
		log.Println("auth     'secret'   not set using default value - ", t.Auth.Secret)
	}

	if t.Auth.Enable == false {
		log.Println("WARNING authentication disabled")
	}

	if t.Deployer.Port == t.Router.Port {
		log.Fatal("deployer and router ports cannot be same", t.Deployer.Port)
	}

	Config = t
	return t
}
