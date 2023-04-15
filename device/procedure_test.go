package device

import (
	"testing"
)

func TestStep_Run(t *testing.T) {
	tasks := []func(params, results map[string]interface{}){
		func(params, results map[string]interface{}) {
			results["b"] = params["b"].(int) * 2
		},
		func(params, results map[string]interface{}) {
			results["a"] = params["a"].(int) * 5
		},
	}

	s := &Step{
		Params:  map[string]interface{}{"a": 5, "b": 10},
		Results: make(map[string]interface{}),
		Tasks:   tasks,
	}
	s.Run()
	if s.Results["a"] != s.Params["a"].(int)*5 {
		t.Fail()
	}
	if s.Results["b"] != s.Params["b"].(int)*2 {
		t.Fail()
	}
}
