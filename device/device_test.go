package device_test

import (
	"go.uber.org/zap"
	"scaffold/device"
	"scaffold/modbus"
	yaml2 "scaffold/node/hardware/yaml"
	"testing"
)

func TestNewDevice(t *testing.T) {
	device.NewDevice("fake_device", "./testingResult")

}

func TestNewNode(t *testing.T) {
	srv := yaml2.NewYAMLHardwareService()
	device.NewNode("fake_node", "software", "./testingResult/fake_device/nodes", srv)
}

func TestLoadNodes(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	d := yaml2.NewYAMLDevice("127.0.0.1", 9000, modbus.DefaultClient(logger))
	d.Load("./testingResult/fake_device/nodes")
	d.Serve()
}
