package serial

import (
	"scaffold/modbus/pdu"
	"scaffold/modbus/wire"
)

type DataLink struct {
	serial *wire.Serial
}

func (d *DataLink) Send(pdu *pdu.SerialPDU) (int, error) {
	bytes := make([]byte, len(pdu.PDU.Data)+4)
	_, err := pdu.Read(bytes)
	if err != nil {
		panic(err)
	}
	return d.serial.Write(bytes)
}

func (d *DataLink) Recv(pdu *pdu.SerialPDU) (int, error) {
	bytes := make([]byte, 256)
	n, err := d.serial.Read(bytes)
	if err != nil {
		panic(err)
	}
	return pdu.Write(bytes[:n])
}

func NewDataLink(serial *wire.Serial) *DataLink {
	return &DataLink{serial: serial}
}
