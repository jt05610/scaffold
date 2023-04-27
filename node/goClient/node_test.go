package goClient_test

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"scaffold/modbus"
	"scaffold/node/goClient"
	"scaffold/node/yaml"
	"testing"
)

func TestNode_Flush(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	device := yaml.NewYAMLDevice("127.0.0.1", 8081, modbus.DefaultClient(logger))
	device.Load("../testing")
	srv := goClient.NewService("https://127.0.0.1:8081")
	for _, n := range device.Nodes {
		out, err := os.Create(fmt.Sprintf("../testingResult/%s_client.go", n.Node))
		if err != nil {
			t.Error(err)
		}
		err = srv.Flush(out, n)
		if err != nil {
			t.Error(err)
		}
	}
}
