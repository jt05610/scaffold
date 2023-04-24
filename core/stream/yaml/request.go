package yaml

import (
	"gopkg.in/yaml.v3"
	"io"
	"os/user"
	"scaffold/core/stream"
	"time"
)

type StreamNode struct {
	Name     string             `yaml:"name"`
	Author   string             `yaml:"author"`
	Date     string             `yaml:"date"`
	Requests []*stream.Req[any] `yaml:"requests"`
}

func (n *StreamNode) Load(r io.Reader) []*stream.Req[any] {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	ret := make([]*stream.Req[any], 0)
	err = yaml.Unmarshal(b, ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (n *StreamNode) Flush(w io.Writer, reqs []*stream.Req[any]) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	n.Name = "stream"
	n.Author = u.Name
	n.Date = time.Now().String()
	n.Requests = reqs
	bytes, err := yaml.Marshal(n)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func NewYAMLReqService() stream.RequestService {
	return &StreamNode{}
}
