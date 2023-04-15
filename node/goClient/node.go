package goClient

import (
	"io"
	"net/http"
	"scaffold/node"
	"text/template"
)

type Node struct {
	baseUrl string
}

func (n *Node) Load(r io.Reader) (*node.Node, error) {
	return nil, nil
}

var clientTemplate = `
// AUTO GENERATED FILE, DO NOT CHANGE

package device

import (
	"bytes"
	"encoding/goClient"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
}

func NewClient() *Client{
	return &Client{}
}
`

var readTemplate = `
// {{.Func}} {{.Description}}.
func (c *Client){{.Func}}() (uint16, error) {
	r, err := http.Get("{{.Route}}")
	if err != nil {
		panic(err)
	}
	d := goClient.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}
`

var writeTemplate = `
// {{.Func}} {{.Description}}.
func (c *Client){{.Func}}({{if .Param}}{{.Param.Name}} {{.Param.Type}}{{end}}) error {
	buf := new(bytes.Buffer)
	{{if .Param}}
	
	req := struct{
		{{.Param.NameCap}} uint16	{{.Param.Tag}}
	}{
		{{.Param.NameCap}}: uint16({{.Param.Name}}),
	}
	err := goClient.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	{{end}}
	resp, err := http.Post("{{.Route}}", "application/goClient", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}
`

func (n *Node) Flush(w io.Writer, node *node.Node) error {
	_, err := w.Write([]byte(clientTemplate))

	if err != nil {
		panic(err)
	}
	tmpMap := map[string]string{
		http.MethodGet:  readTemplate,
		http.MethodPost: writeTemplate,
	}
	for _, e := range node.Endpoints(n.baseUrl) {
		var tmpl *template.Template
		t := tmpMap[e.Method]
		tmpl, err = template.New(e.Method).Parse(t)
		if err != nil {
			return err
		}
		err = tmpl.Execute(w, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewService(baseUrl string) node.Service {
	return &Node{baseUrl: baseUrl}
}
