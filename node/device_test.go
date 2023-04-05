package node_test

import (
	"go.uber.org/zap"
	"scaffold/modbus"
	"scaffold/node"
	"scaffold/node/yaml"
	"testing"
)

func TestNewDevice(t *testing.T) {
	node.NewDevice("fake_device", "./testingResult")

}

func TestNewNode(t *testing.T) {
	srv := yaml.NewYAMLService()
	node.NewNode("fake_node", "./testingResult/fake_device/nodes", srv)
}

func TestLoadNodes(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	d := yaml.NewYAMLDevice("127.0.0.1", 9000, modbus.DefaultClient(logger))
	d.Load("./testingResult/fake_device/nodes")
	d.Serve()
}
