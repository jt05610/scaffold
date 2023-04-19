package device

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"scaffold/modbus"
	"scaffold/node/hardware"
	"time"
)

type Device struct {
	Client      *modbus.Client
	Nodes       []*hardware.Node
	Address     string
	Port        int
	NodeDir     string
	NodeService hardware.Service
}

func (d *Device) Serve() {
	mux := http.NewServeMux()
	for _, n := range d.Nodes {
		n.Register(mux)
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%v", d.Address, d.Port),
		Handler:           mux,
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (d *Device) Load(nodeDir string) {
	base, err := os.Open(path.Join(nodeDir, "base.yaml"))
	bNode, err := d.NodeService.Load(base)
	items, err := os.ReadDir(nodeDir)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		if item.IsDir() {
			df, err := os.Open(path.Join(nodeDir, item.Name(), "node.yaml"))
			if err != nil {
				panic(err)
			}
			node, err := d.NodeService.Load(df)
			node.Diag = bNode.Diag
			node.Client = d.Client
			if err != nil {
				panic(err)
			}
			d.Nodes = append(d.Nodes, node)
		}
	}
}

func NewDevice(name string, dest string) {
	for _, p := range []string{
		path.Join(dest, name),
		path.Join(dest, name, "nodes"),
		path.Join(dest, name, "modules"),
		path.Join(dest, name, ".drivers"),
		path.Join(dest, name, "clients"),
		path.Join(dest, name, "procedures"),
		path.Join(dest, name, "data"),
	} {
		err := os.MkdirAll(p, 0777)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}
}

func NewNode(name, kind, parent string, srv hardware.Service) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	existing, err := os.ReadDir(parent)
	if err != nil {
		panic(err)
	}

	n := &hardware.Node{
		MetaData: hardware.MetaData{
			Node:    name,
			Author:  u.Name,
			Address: byte(len(existing)),
			Date:    time.Now().Format(time.RFC822),
		},
		Tables: map[string][]*hardware.Handler{
			"discrete_inputs": {
				&hardware.Handler{
					Name:        "discrete_input_1",
					Description: "Discrete inputs are binary read only registers",
					Params:      nil,
				},
				&hardware.Handler{
					Name:        "discrete_input_2",
					Description: "Add a new one like this",
					Params:      nil,
				},
			},
			"coils": {
				&hardware.Handler{
					Name:        "coil_1",
					Description: "Coils are binary read/write registers. They are used to execute on/off functions on the target device",
					Params: []map[string]*hardware.Param{
						{
							"value": &hardware.Param{
								Type:        "int",
								Description: "Writing 0 will write 0 to the device. Writing anything else will write 1.",
							},
						},
					},
				},
				&hardware.Handler{
					Name:        "coil_2",
					Description: "They can also be parameter-less if you want",
					Params:      nil,
				},
			},
			"input_registers": {
				&hardware.Handler{
					Name:        "input_register_1",
					Description: "Input registers are 16-bit read only registers",
					Params:      nil,
				},
				&hardware.Handler{
					Name:        "input_register_2",
					Description: "Add a new one like this",
					Params:      nil,
				},
			},
			"holding_registers": {
				&hardware.Handler{
					Name:        "holding_register_1",
					Description: "Holding registers are 16-bit read/write registers. They are used to set variables on the device, and can execute a if desired",
					Params: []map[string]*hardware.Param{
						{
							"value": &hardware.Param{
								Type:        "int",
								Description: "Whatever you write will be converted to a uint16",
							},
						},
					},
				},
				&hardware.Handler{
					Name:        "holding_register_2",
					Description: "Add a new one like this",
					Params: []map[string]*hardware.Param{
						{
							"value": &hardware.Param{
								Type:        "int",
								Description: "You technically don't need to include a parameter but you probably should",
							},
						},
					},
				},
			},
		},
	}
	base := path.Join(parent, name)
	err = os.Mkdir(base, 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	out, err := os.Create(path.Join(base, "node.yaml"))
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = srv.Flush(out, n)
	if err != nil {
		panic(err)
	}
}
