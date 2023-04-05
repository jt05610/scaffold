package wire

import (
	"errors"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
	"go.uber.org/zap"
)

type SerialOpt struct {
	Port     string
	VID      string
	PID      string
	Baud     int
	Parity   serial.Parity
	DataBits int
	StopBits serial.StopBits
}

type Serial struct {
	rx  serial.Port
	tx  serial.Port
	log *zap.Logger
}

var DefaultSerial = &SerialOpt{
	Port:     "",
	VID:      "1A86",
	PID:      "7523",
	Baud:     19200,
	Parity:   serial.NoParity,
	DataBits: 8,
	StopBits: serial.TwoStopBits,
}

var NoPort = errors.New("no port given to NewSerial")

func NewSerial(opt *SerialOpt, log *zap.Logger) (*Serial, error) {
	s := &Serial{log: log}
	var err error
	if opt.Port == "" {
		if opt.PID == "" && opt.VID == "" {
			return nil, NoPort
		}
		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			panic(err)
		}
		vCheck := len(opt.PID) > 0
		for _, port := range ports {
			if port.PID == opt.PID {
				opt.Port = port.Name
				if vCheck {
					if port.VID == opt.VID {
						opt.Port = port.Name
						s.log.Info("found port", zap.String("port", port.Name))
						break
					}
				} else {
					s.log.Info("found port", zap.String("port", port.Name))
					break
				}
			}
		}
	}
	if len(opt.Port) == 0 {
		panic(errors.New("no port found"))
	}
	mode := &serial.Mode{
		BaudRate: opt.Baud,
		Parity:   opt.Parity,
		DataBits: opt.DataBits,
		StopBits: opt.StopBits,
	}
	var portName string
	if opt.Port[:8] == "/dev/cu." {
		portName = opt.Port[8:]
	} else {
		portName = opt.Port[9:]
	}

	s.rx, err = serial.Open("/dev/cu."+portName, mode)

	if err != nil {
		panic(err)
	}
	s.log.Info("connected rx", zap.String("port", opt.Port))
	s.tx = s.rx
	return s, err
}

func (s *Serial) Read(p []byte) (n int, err error) {
	n, err = s.rx.Read(p)
	s.log.Info("received", zap.ByteString("pdu", p[:n]))
	return n, err
}

func (s *Serial) Write(p []byte) (n int, err error) {
	s.log.Info("sending", zap.ByteString("pdu", p))
	n, err = s.tx.Write(p)
	return n, err
}
