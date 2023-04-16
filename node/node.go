package node

import (
	"io"
	"net/http"
)

type BaseNode interface {
	Register(srv *http.ServeMux)
	Endpoints() []string
}

type BaseNodeService interface {
	Load(r io.Reader) (*BaseNode, error)
	Flush(w io.Writer, node *BaseNode) error
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
