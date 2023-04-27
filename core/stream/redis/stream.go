package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"scaffold"
	"time"
)

type Requester[T any] interface {
	Request(ctx context.Context) error
	Value() (interface{}, error)
	Name() string
}

type Stream struct {
	name     string
	client   *redis.Client
	interval time.Duration
	requests []Requester[any]
}

func (s *Stream) open() error {
	s.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return nil
}

func (s *Stream) close() {
	_ = s.client.Close()
}

func (s *Stream) doRequests(ctx context.Context) (*redis.XAddArgs, error) {
	data := make(map[string]interface{})
	for _, r := range s.requests {
		err := r.Request(ctx)
		if err != nil {
			return nil, err
		}
		v, err := r.Value()
		if err != nil {
			return nil, err
		}
		data[r.Name()] = v
	}
	return &redis.XAddArgs{
		Stream:     s.name,
		NoMkStream: false,
		MaxLen:     10000,
		MinID:      "",
		Approx:     false,
		Limit:      0,
		ID:         "",
		Values:     data,
	}, nil
}

func (s *Stream) Stream(ctx context.Context) error {
	err := s.open()
	if err != nil {
		return err
	}
	defer s.close()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(s.interval):
			args, err := s.doRequests(ctx)
			if err != nil {
				return err
			}
			s.client.XAdd(ctx, args)
		}
	}
}

func NewRedisStream(name string, interval time.Duration, requests []Requester[any]) scaffold.Streamer {
	return &Stream{
		name:     name,
		interval: interval,
		requests: requests,
	}
}

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
