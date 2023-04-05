package _go

import (
	"io"
	"scaffold/node"
)

type Node struct {
}

func (n *Node) Load(r io.Reader) (*node.Node, error) {
	return nil, nil
}

func (n *Node) Flush(w io.Writer, node *node.Node) error {
	return nil
}
