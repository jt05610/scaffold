package yaml

import (
	"scaffold/device"
	"scaffold/modbus"
	"scaffold/node/hardware"
)

func NewYAMLDevice(addr string, port int, client *modbus.Client) *device.Device {
	return &device.Device{
		Address:     addr,
		Port:        port,
		Nodes:       make([]*hardware.Node, 0),
		NodeService: NewYAMLService(),
		Client:      client,
	}
}
