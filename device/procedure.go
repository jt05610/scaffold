package device

import (
	"fmt"
	"sync"
)

type Step struct {
	wg      sync.WaitGroup
	Params  map[string]interface{}
	Results map[string]interface{}
	Tasks   []func(params, results map[string]interface{})
}

type Proc struct {
	Steps []*Step
}

func (s *Step) Run() {
	for _, t := range s.Tasks {
		s.wg.Add(1)
		go func(task func(p, r map[string]interface{})) {
			defer s.wg.Done()
			task(s.Params, s.Results)
		}(t)
	}
	s.wg.Wait()
}

func (p *Proc) Run() {
	for i, s := range p.Steps {
		fmt.Printf("Starting step %v/%v\n", i, len(p.Steps))
		s.Run()
	}
}
