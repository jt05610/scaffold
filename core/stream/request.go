package stream

import (
	"context"
	"errors"
	"io"
)

type Req[T any] struct {
	v       T
	name    string
	err     error
	call    func() (T, error)
	resChan chan T
}

func (r *Req[T]) Request(ctx context.Context) (err error) {
	go func() {
		v, err := r.call()
		if err != nil {
			r.err = err
		}
		r.resChan <- v
	}()
	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout")
		case r.v = <-r.resChan:
			return r.err
		}
	}
}

func (r *Req[T]) Value() (interface{}, error) {
	return r.v, r.err
}

func (r *Req[T]) Name() string {
	return r.name
}

func NewReq[T any](name string, call func() (T, error)) Requester[T] {
	return &Req[T]{name: name, call: call, resChan: make(chan T)}
}

type Requester[T any] interface {
	Request(ctx context.Context) error
	Value() (interface{}, error)
	Name() string
}

type RequestService interface {
	Load(r io.Reader) []*Req[any]
	Flush(w io.Writer, reqs []*Req[any])
}
