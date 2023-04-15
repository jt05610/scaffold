package node

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"io"
	"net/http"
	"net/url"
	"scaffold/modbus"
	"scaffold/modbus/pdu"
)

type MetaData struct {
	Node    string
	Author  string
	Address byte
	Date    string
}

type ParamType string

type Param struct {
	Type        ParamType
	Description string `yaml:"desc"`
}

type Handler struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"desc"`
	Params      []map[string]*Param `yaml:"params,omitempty"`
}

type Node struct {
	MetaData    `yaml:"meta"`
	Tables      map[string][]*Handler `yaml:"tables"`
	Diag        []*Handler            `yaml:"diag"`
	client      *modbus.Client
	rfLookup    map[string]map[string]func(uint16, uint16) *pdu.ModbusPDU
	addrLookup  map[string]uint16
	paramLookup map[string]string
}

var lookup = map[string]map[string]func(uint16, uint16) *pdu.ModbusPDU{
	http.MethodGet: {
		"coils":             pdu.ReadCoils,
		"discrete_inputs":   pdu.ReadDiscreteInputs,
		"holding_registers": pdu.ReadHoldingRegisters,
		"input_registers":   pdu.ReadInputRegisters,
	},
	http.MethodPost: {
		"coils":             pdu.WriteCoil,
		"holding_registers": pdu.WriteRegister,
	},
}

type EndpointParam struct {
	Name        string
	NameCap     string
	Type        string
	Description string
	Tag         string
}

type Endpoint struct {
	Func        string         `goClient:"func"`
	Route       string         `goClient:"route"`
	Method      string         `goClient:"method"`
	Description string         `goClient:"description"`
	Param       *EndpointParam `goClient:"params,omitempty"`
}

func (n *Node) Endpoints(baseURL string) []*Endpoint {
	res := make([]*Endpoint, 0)
	n.paramLookup = make(map[string]string)
	for name, handlers := range n.Tables {
		for _, h := range handlers {
			route, err := url.JoinPath(baseURL, n.Node, h.Name)
			if err != nil {
				panic(err)
			}
			res = append(res, &Endpoint{
				Route:       route,
				Method:      http.MethodGet,
				Description: h.Description,
				Param:       nil,
				Func:        strcase.ToCamel(fmt.Sprintf("Get %s", h.Name)),
			})
			if name == "coils" || name == "holding_registers" {
				var param *EndpointParam
				if len(h.Params) > 0 {
					param = &EndpointParam{}
					for paramName, p := range h.Params[0] {
						param.Name = strcase.ToLowerCamel(paramName)
						param.Type = string(p.Type)
						param.Description = p.Description
						param.NameCap = strcase.ToCamel(paramName)
						param.Tag = fmt.Sprintf("`goClient:\"%s\"`", param.Name)
					}
					n.paramLookup[route] = param.Name
				} else {
					param = nil
				}
				res = append(res, &Endpoint{
					Route:       route,
					Method:      http.MethodPost,
					Description: h.Description,
					Param:       param,
					Func:        strcase.ToCamel(fmt.Sprintf("Post %s", h.Name)),
				})
			}
		}
	}
	return res
}

func (n *Node) handlers() map[string]http.HandlerFunc {
	_ = n.Endpoints("/")
	res := make(map[string]http.HandlerFunc)
	n.rfLookup = make(map[string]map[string]func(uint16, uint16) *pdu.ModbusPDU)
	n.rfLookup[http.MethodGet] = make(map[string]func(uint16, uint16) *pdu.ModbusPDU)
	n.rfLookup[http.MethodPost] = make(map[string]func(uint16, uint16) *pdu.ModbusPDU)
	n.addrLookup = make(map[string]uint16)
	for name, handlers := range n.Tables {
		for i, h := range handlers {
			endpoint := fmt.Sprintf("/%s/%s", n.Node, h.Name)
			if reqFunc, ok := lookup[http.MethodGet][name]; ok {
				n.rfLookup[http.MethodGet][endpoint] = reqFunc
				n.addrLookup[endpoint] = uint16(i)
			} else {
				panic(errors.New("failed to find get request formatter"))
			}
			if name == "coils" || name == "holding_registers" {
				if rf, ok := lookup[http.MethodPost][name]; ok {
					n.rfLookup[http.MethodPost][endpoint] = rf
				} else {
					panic(errors.New("failed to find post request formatter"))
				}
			}
			res[endpoint] = func(w http.ResponseWriter, r *http.Request) {
				if reqFunc, ok := n.rfLookup[r.Method][r.RequestURI]; ok {
					if reqFunc == nil {
						http.Error(w, "internal server error", http.StatusInternalServerError)
					} else {
						var bytes []byte
						if r.Method == http.MethodGet {
							res, err := n.client.Request(r.Context(), n.Address, reqFunc(n.addrLookup[r.RequestURI], 1))
							if err != nil || res == nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							if res != nil && res.FuncCode < pdu.ReadHoldingRegistersFC {
								bytes, err = json.Marshal(map[string]uint16{
									"result": uint16(res.Data[1]),
								})

							} else {
								if res.Data[1] == 2 {
									bytes, err = json.Marshal(map[string]uint16{
										"result": binary.BigEndian.Uint16(res.Data[2:]),
									})
								} else {
									rr := make([]uint16, 0)
									for i := uint8(0); i < res.Data[1]; i += 2 {
										rr = append(rr, binary.BigEndian.Uint16(res.Data[2+i:]))
									}
									bytes, err = json.Marshal(map[string][]uint16{
										"result": rr,
									})
								}
							}

						} else {
							dec := json.NewDecoder(r.Body)
							req := make(map[string]uint16)
							err := dec.Decode(&req)
							if err != nil {
								if _, found := n.paramLookup[r.RequestURI]; found {
									http.Error(w, "bad request", http.StatusBadRequest)
									return
								} else {
									if err.Error() != "EOF" {
										http.Error(w, "bad request", http.StatusBadRequest)
										return
									}
								}
							}
							var reqPDU *pdu.ModbusPDU
							if len(req) > 0 {
								if param, found := n.paramLookup[r.RequestURI]; !found {
									http.Error(w,
										fmt.Sprintf("must include %s", param),
										http.StatusBadRequest)
								} else {
									if p, found := req[param]; found {
										reqPDU = reqFunc(n.addrLookup[r.RequestURI], p)
									} else {
										http.Error(w,
											fmt.Sprintf("must include %s", param),
											http.StatusBadRequest)
									}
								}
							} else {
								reqPDU = reqFunc(n.addrLookup[r.RequestURI], 1)
							}

							res, err := n.client.Request(r.Context(), n.Address, reqPDU)
							if err != nil {
								http.Error(w, "server error", http.StatusInternalServerError)
								return
							}
							if res.FuncCode != reqPDU.FuncCode {
								http.Error(w, "Modbus server error", http.StatusInternalServerError)
							} else {
								bytes = []byte("ok")
							}
						}
						w.WriteHeader(http.StatusOK)
						_, err := w.Write(bytes)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
	return res
}

func (n *Node) Register(srv *http.ServeMux) {
	for name, handler := range n.handlers() {
		srv.HandleFunc(name, handler)
	}
	for _, handler := range n.Diag {
		if handler.Name == "echo" {
			endpoint := fmt.Sprintf("/%s/%s", n.Node, "echo")
			srv.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
				dec := json.NewDecoder(r.Body)
				body := make(map[string][]byte)
				err := dec.Decode(&body)
				if err != nil {
					http.Error(w, "bad request", http.StatusBadRequest)
				}
				req := pdu.Echo(body["message"]...)
				res, err := n.client.Request(r.Context(), n.Address, req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				var bytes []byte
				if res.FuncCode != req.FuncCode {
					http.Error(w, "Modbus server error", http.StatusInternalServerError)
				} else {
					bytes = res.Data
				}
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				if err != nil {
					panic(err)
				}
			})
		}
	}
}

type Service interface {
	Load(r io.Reader) (*Node, error)
	Flush(w io.Writer, node *Node) error
}
