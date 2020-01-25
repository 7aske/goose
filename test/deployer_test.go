package deployer

import (
	"../internal/deployer"
	"../internal/instance"
	"testing"
)

func TestInstall(t *testing.T) {
	bkend := "golang" // invalid backend
	inst := instance.New(
		"https://github.com/7aske/goose",
		"goose.7aske.com",
		instance.Backend(bkend))
	json := instance.ToJSONStruct(inst)
	err := deployer.Deployer.Install(json)
	if err == nil {
		t.Errorf("Backend check failed, invalid backend %s was installed", bkend)
	}
}