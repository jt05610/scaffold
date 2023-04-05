package yaml

import (
	"gopkg.in/yaml.v3"
	"io"
	"scaffold/node"
)

type Node struct {
}

func (n *Node) Load(r io.Reader) (*node.Node, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	ret := &node.Node{}
	return ret, yaml.Unmarshal(b, ret)
}

func (n *Node) Flush(w io.Writer, node *node.Node) error {
	bytes, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func NewYAMLService() node.Service {
	return &Node{}
}