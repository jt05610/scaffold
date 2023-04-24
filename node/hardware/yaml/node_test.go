package yaml_test

import (
	"os"
	"scaffold/node/hardware/yaml"
	"testing"
)

func TestNode_LoadFlush(t *testing.T) {
	srv := yaml.NewYAMLHardwareService()
	df, err := os.Open("../testing/node.yaml")
	if err != nil {
		t.Error(err)
	}
	node, err := srv.Load(df)
	if err != nil {
		t.Error(err)
	}
	new, err := os.Create("../testing/node_written.yaml")
	err = srv.Flush(new, node)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
