package goose

import (
	"../../config"
	"../../deployer"
	"../../instance"
	"../../server"
	"../../utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var goose *Goose = nil

type Goose struct {
}

var Config *config.Type = nil
var Deployer *deployer.Type = nil
var host = "0.0.0.0"

//var port = 5000

func New() {
	Config = config.Get()

	err := setupDirs()
	if err != nil {
		log.Fatal("failed setting up dirs ", err)
	}

	err = initInstanceFile()
	if err != nil {
		log.Fatal(err)
	}

	Deployer = deployer.New()

	err = server.Listen(Config.Deployer.Hostname, Config.Deployer.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func setupDirs() error {
	err := utils.MakeDirIfNotExist(Config.Deployer.LogRoot)
	if err != nil {
		return err
	}
	err = utils.MakeDirIfNotExist(Config.Deployer.Root)
	if err != nil {
		return err
	}
	return utils.MakeDirIfNotExist(Config.Deployer.AppRoot)
}

func initInstanceFile() error {
	pth := path.Join(Config.Deployer.Root, "instances.json")
	if _, err := os.Stat(pth); err != nil {
		log.Println("initializing instances file")
		emptyArr, _ := json.Marshal(&instance.File{Instances: []instance.JSON{}})
		err := ioutil.WriteFile(pth, emptyArr, 0775)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
