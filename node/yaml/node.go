package yaml

import (
	"gopkg.in/yaml.v3"
	"io"
	"scaffold/node/hardware"
)

type Node struct {
}

func (n *Node) Load(r io.Reader) (*hardware.Node, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	ret := &hardware.Node{}
	return ret, yaml.Unmarshal(b, ret)
}

func (n *Node) Flush(w io.Writer, node *hardware.Node) error {
	bytes, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func NewYAMLService() hardware.Service {
	return &Node{}
}
