package device_test

import (
	"go.uber.org/zap"
	"scaffold/device"
	"scaffold/modbus"
	"scaffold/node/yaml"
	"testing"
)

func TestNewDevice(t *testing.T) {
	device.NewDevice("fake_device", "./testingResult")

}

func TestNewNode(t *testing.T) {
	srv := yaml.NewYAMLService()
	device.NewNode("fake_node", "software", "./testingResult/fake_device/nodes", srv)
}

func TestLoadNodes(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	d := yaml.NewYAMLDevice("127.0.0.1", 9000, modbus.DefaultClient(logger))
	d.Load("./testingResult/fake_device/nodes")
	d.Serve()
}
