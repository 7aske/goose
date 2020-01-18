package goose

import (
	"../../config"
	"../../deployer"
	"../../instance"
	"../../server"
	"../../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var goose *Goose = nil

type Goose struct {
}

var Config *config.Type = nil
var Deployer *deployer.Type = nil

//https://github.com/7aske/player-database
func New() *Goose {
	Config = config.Get()
	setupDirs()
	initInstanceFile()
	Deployer = deployer.New()
	//err := os.RemoveAll(config.Config.Deployer.AppRoot)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//inst := instance.New("https://github.com/7aske/player-database", "127.0.0.1", instance.Node)
	//inst, _ = Deployer.Deploy(inst)
	//_, _ = Deployer.Install(inst)
	server.Route()
	if goose == nil {
		goose = new(Goose)
		return goose
	}
	return goose
}

func setupDirs() error {
	err := utils.MakeDirIfNotExist(Config.Deployer.Root)
	if err != nil {
		return err
	}
	return utils.MakeDirIfNotExist(Config.Deployer.AppRoot)
}

func initInstanceFile() error {
	pth := path.Join(Config.Deployer.Root, "instances.json")
	if _, err := os.Stat(pth); err != nil {
		fmt.Println("initializing instances file")
		emptyArr, _ := json.Marshal(&instance.File{Instances: []instance.JSON{}})
		err := ioutil.WriteFile(pth, emptyArr, 0775)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
