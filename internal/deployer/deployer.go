package deployer

import (
	"../config"
	"../instance"
	"../utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Type struct {
	Running []*instance.Instance
}

var Deployer *Type = nil
var Config *config.Type = nil

func New() *Type {
	ret := new(Type)
	Config = config.Get()
	err := updateInstancesFile()
	if err != nil {
		log.Fatal(err)
	}
	Deployer = ret
	return ret
}

func GetDeployedInstances() ([]instance.JSON, error) {
	pth := path.Join(Config.Deployer.Root, "instances.json")
	f, err := ioutil.ReadFile(pth)
	if err != nil {
		return nil, err
	}
	instances := instance.File{}
	err = json.Unmarshal(f, &instances)
	if err != nil {
		return nil, err
	}
	return instances.Instances, nil
}

// Compares between JSON file and deployed instances in instance folder and removes set difference of both from each.
// Should be called whenever instance is added or removed
func updateInstancesFile() error {
	pth := path.Join(Config.Deployer.Root, "instances.json")
	folders, err := ioutil.ReadDir(path.Join(Config.Deployer.AppRoot))
	if err != nil {
		return err
	}
	instances, err := GetDeployedInstances()
	if err != nil {
		return err
	}
	for i, inst := range instances {
		if !utils.ContainsFile(inst.Name, &folders) {
			fmt.Println("not found ", inst.Name)
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	for _, folder := range folders {
		if !utils.ContainsInstance(folder.Name(), &instances) {
			fmt.Println("not found ", folder.Name())
			err = os.RemoveAll(path.Join(Config.Deployer.AppRoot, folder.Name()))
			if err != nil {
				return err
			}
		}
	}
	updated, err := json.Marshal(instance.File{Instances: instances})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pth, updated, 0775)
}

func saveInstance(inst instance.JSON) error {
	pth := path.Join(Config.Deployer.Root, "instances.json")

	instances, _ := GetDeployedInstances()
	if pos := utils.IndexOfInstance(inst, &instances); pos == -1 {
		instances = append(instances, inst)
	} else {
		instances = append(instances[:pos], instances[pos+1:]...)
		instances = append(instances, inst)
	}
	apps, err := json.Marshal(instance.File{Instances: instances})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pth, apps, 0775)
}

func removeInstance(inst instance.JSON) error {
	pth := path.Join(Config.Deployer.Root, "instances.json")

	instances, _ := GetDeployedInstances()
	if pos := utils.IndexOfInstance(inst, &instances); pos == -1 {
		return errors.New("cannot remove instance - doesn't exist")
	} else {
		instances = append(instances[:pos], instances[pos+1:]...)
	}
	apps, err := json.Marshal(instance.File{Instances: instances})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pth, apps, 0775)
}

func (d *Type) RemoveInstanceJSON(instance instance.JSON) error {
	instances, err := GetDeployedInstances()
	if err != nil {
		return err
	}
	for _, inst := range instances {
		if inst.Name == instance.Name || instance.Id == inst.Id {
			return removeInstance(inst)
		}
	}
	return errors.New("instance not found")
}

func (d *Type) AddInstanceJSON(instance instance.JSON) error {
	return saveInstance(instance)
}

// Removes RUNNING instance
func (d *Type) RemoveInstance(instance *instance.Instance) error {
	for i, inst := range d.Running {
		if inst.Name == instance.Name || instance.Id == inst.Id {
			d.Running = append(d.Running[:i], d.Running[i+1:]...)
			return nil
		}
	}
	return errors.New("instance not found")
}

// Adds RUNNING instance
func (d *Type) AddInstance(instance *instance.Instance) {
	d.Running = append(d.Running, instance)
}
