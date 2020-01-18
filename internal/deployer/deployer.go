package deployer

import (
	"../config"
	"../instance"
	"../utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

type Type struct {
	Instances []instance.JSON
	Running   []*instance.Instance
}

var Deployer *Type = nil
var Config *config.Type = nil

func New() *Type {
	ret := new(Type)
	Config = config.Get()
	ret.Instances = []instance.JSON{}
	ret.updateInstancesFile()
	Deployer = ret
	return ret
}
func (d *Type) GetDeployedInstances() ([]instance.JSON, error) {
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

func (d *Type) updateInstancesFile() error {
	pth := path.Join(Config.Deployer.Root, "instances.json")
	folders, _ := ioutil.ReadDir(path.Join(Config.Deployer.AppRoot))
	instances, err := d.GetDeployedInstances()
	if err != nil {
		return err
	}
	for i, inst := range instances {
		if !utils.ContainsFile(inst.Name, &folders) {
			fmt.Println("not found ", inst.Name)
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	updated, err := json.Marshal(instance.File{Instances: instances})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(pth, updated, 0775)
	if err != nil {
		return err
	}
	return nil
}

func (d *Type) saveInstance(inst instance.JSON) error {
	pth := path.Join(Config.Deployer.Root, "instances.json")

	instances, _ := d.GetDeployedInstances()
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
	err = ioutil.WriteFile(pth, apps, 0775)
	if err != nil {
		return err
	}
	return nil
}

func (d *Type) RemoveInstanceJSON(instance instance.JSON) error {
	for i, inst := range d.Instances {
		if inst.Name == instance.Name || instance.Id == inst.Id {
			d.Instances = append(d.Instances[:i], d.Instances[i+1:]...)
			return nil
		}
	}
	return errors.New("instance not found")
}

func (d *Type) AddInstanceJSON(instance instance.JSON) error {
	d.Instances = append(d.Instances, instance)
	return d.updateInstancesFile()
}

func (d *Type) RemoveInstance(instance *instance.Instance) error {
	for i, inst := range d.Running {
		if inst.Name == instance.Name || instance.Id == inst.Id {
			d.Running = append(d.Running[:i], d.Running[i+1:]...)
			return nil
		}
	}
	return errors.New("instance not found")
}

func (d *Type) AddInstance(instance *instance.Instance) {
	d.Running = append(d.Running, instance)
}
