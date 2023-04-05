package yaml

import (
	"scaffold/modbus"
	"scaffold/node"
)

func NewYAMLDevice(addr string, port int, client *modbus.Client) *node.Device {
	return &node.Device{
		Address:     addr,
		Port:        port,
		Nodes:       make([]*node.Node, 0),
		NodeService: NewYAMLService(),
		Client: client,
	}
}
