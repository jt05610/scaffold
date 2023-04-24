package node

import (
	"io"
	"net/http"
)

type Node interface {
	Register(srv *http.ServeMux)
	Endpoints(base string) []*Endpoint
}

type Service interface {
	Load(r io.Reader) (Node, error)
	Flush(w io.Writer, node Node) error
}

type EndpointParam struct {
	Name        string
	NameCap     string
	Type        string
	Description string
	Tag         string
}

type Endpoint struct {
	Func        string         `json:"func"`
	Route       string         `json:"route"`
	Method      string         `json:"method"`
	Description string         `json:"description"`
	Param       *EndpointParam `json:"params,omitempty"`
}
