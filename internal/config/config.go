package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const defaultConfig = `router:
  port: 8080
  hostname: 0.0.0.0

deployer:
  port: 5000
  hostname: 127.0.0.1

auth:
  enable: true
  secret: secret
  user: admin
  pass: admin`

var Config *Type = nil

type Type struct {
	Deployer struct {
		AppRoot  string `yaml:"approot,omitempty"` // Root directory for instances
		LogRoot  string `yaml:"logroot,omitempty"` // Root directory for logs
		Root     string `yaml:"root,omitempty"`    // Root directory for instances and logs
		Port     int    `yaml:"port"`              // Port on which deployer server is running
		Hostname string `yaml:"hostname"`          // Hostname for which proxy will redirect requests to the deployer
	}
	Router struct {
		Port     int    `yaml:"port"`               // Port at which proxy is running (80 for HTTP)
		Hostname string `yaml:"hostname"`           // IP on which the proxy server listens to (0.0.0.0 to make server public)
		RootHost string `yaml:"rootHost,omitempty"` // Redirects root domain request to this sub-domain
	}
	Auth struct {
		Enable bool   `yaml:"enable"` // Enable authenticating requests with JWT tokens
		Secret string `yaml:"secret"` // Secret for generating JWT tokens
		User   string `yaml:"user"`   // Deployer username
		Pass   string `yaml:"pass"`   // Deployer password
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
	home := os.Getenv("HOME")
	configPath := path.Join(home, ".config/goose/config.yaml")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err)
		log.Println("using default config")
		log.Println(defaultConfig)
		data = []byte(defaultConfig)
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

	if t.Deployer.LogRoot == "" {
		t.Deployer.LogRoot = path.Join(t.Deployer.Root, "logs")
		log.Println("deployer 'logroot'  not set using default value - ", t.Deployer.LogRoot)
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

	if t.Router.RootHost == "" {
		t.Router.RootHost = t.Deployer.Hostname
		log.Println("router   'rootHost' not set using default value - ", t.Router.RootHost)
	}

	Config = t
	return t
}
