package modbus

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"scaffold/modbus/pdu"
	"scaffold/modbus/serial"
	"scaffold/modbus/wire"
	"sync"
)

type Client struct {
	logger *zap.Logger
	dl     *serial.DataLink
	mu     sync.Mutex
}

func (c *Client) Request(ctx context.Context, addr byte, req *pdu.ModbusPDU) (*pdu.ModbusPDU, error) {
	s := pdu.NewSerialPDU(addr, req)
	c.mu.Lock()
	_, err := c.dl.Send(s)
	if err != nil {
		panic(err)
	}
	resCh := make(chan *pdu.SerialPDU)
	go func() {
		defer close(resCh)
		res := &pdu.SerialPDU{}
		_, err := c.dl.Recv(res)
		if err != nil {
			panic(err)
		}
		resCh <- res
	}()
	for {
		select {
		case res := <-resCh:
			c.mu.Unlock()
			return res.PDU, nil
		case <-ctx.Done():
			return nil, errors.New("timeout during request")
		}
	}
}

func DefaultClient(logger *zap.Logger) *Client {
	ser, err := wire.NewSerial(wire.DefaultSerial, logger)
	if err != nil {
		panic(err)
	}
	return &Client{
		logger: logger,
		dl:     serial.NewDataLink(ser),
	}
}
