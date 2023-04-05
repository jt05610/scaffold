package pdu

import (
	"encoding/binary"
)

type FuncCode uint8

const (
	ReadCoilsFC            FuncCode = 0x01
	ReadDiscreteInputsFC   FuncCode = 0x02
	ReadHoldingRegistersFC FuncCode = 0x03
	ReadInputRegistersFC   FuncCode = 0x04
	WriteCoilFC            FuncCode = 0x05
	WriteHoldingRegisterFC FuncCode = 0x06
	DiagFC                 FuncCode = 0x08
)

type ModbusPDU struct {
	FuncCode
	Data []byte
}

func (m *ModbusPDU) Read(p []byte) (n int, err error) {
	p[0] = byte(m.FuncCode)
	for i := 1; i < len(p); i++ {
		p[i] = m.Data[i-1]
	}
	return len(p), nil
}

func (m *ModbusPDU) Write(p []byte) (n int, err error) {
	m.FuncCode = FuncCode(p[0])
	if len(p) > 1 {
		m.Data = p[1:]
	}
	return len(p), nil

}

func ReadCoils(addr uint16, quantity uint16) *ModbusPDU {
	m := &ModbusPDU{
		FuncCode: ReadCoilsFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(quantity))
	return m
}

func ReadDiscreteInputs(addr uint16, quantity uint16) *ModbusPDU {
	m := &ModbusPDU{
		FuncCode: ReadDiscreteInputsFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(quantity))
	return m
}

func ReadHoldingRegisters(addr uint16, quantity uint16) *ModbusPDU {
	m := &ModbusPDU{
		FuncCode: ReadHoldingRegistersFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(quantity))
	return m
}

func ReadInputRegisters(addr uint16, quantity uint16) *ModbusPDU {
	m := &ModbusPDU{
		FuncCode: ReadInputRegistersFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(quantity))
	return m
}

func WriteCoil(addr uint16, value uint16) *ModbusPDU {
	if value != 0 {
		value = 0xFF00
	}
	m := &ModbusPDU{
		FuncCode: WriteCoilFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(value))
	return m
}

func WriteRegister(addr uint16, value uint16) *ModbusPDU {
	m := &ModbusPDU{
		FuncCode: WriteHoldingRegisterFC,
		Data:     make([]byte, 4),
	}
	binary.BigEndian.PutUint32(m.Data, (uint32(addr)<<16)+uint32(value))
	return m
}

func Echo(data ...byte) *ModbusPDU {
	return &ModbusPDU{
		FuncCode: DiagFC,
		Data:     append([]byte{0x00, 0x00}, data...),
	}
}
