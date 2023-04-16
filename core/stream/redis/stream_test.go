package redis

import (
	"context"
	"testing"
	"time"
)

func TestStream_Stream(t *testing.T) {
	np := NewReq[uint16]("NeedlePosition", func() (uint16, error) {
		return 123, nil
	})
	vd := NewReq[float32]("VolumeDispensed", func() (float32, error) {
		return 123.45, nil
	})
	pf := NewReq[float32]("PlungerForce", func() (float32, error) {
		return 678.9, nil
	})
	rr := []Requester[any]{
		np, vd, pf,
	}
	s := NewRedisStream("test", time.Duration(100)*time.Millisecond, rr)
	ctx := context.Background()
	err := s.Stream(ctx)
	if err != nil {
		t.Error(err)
	}
}
