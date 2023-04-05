package pdu_test

import (
	"scaffold/modbus/pdu"
	"sync"
	"testing"
)

func TestRW(t *testing.T) {
	for _, tc := range []struct {
		Name     string
		Call     func(uint16, uint16) *pdu.ModbusPDU
		First    uint16
		Second   uint16
		Expected []byte
	}{
		{"Coils", pdu.ReadCoils, 0xFEED, 0xBEAD, []byte{0x01, 0xFE, 0xED, 0xBE, 0xAD}},
		{"DiscreteInputs", pdu.ReadDiscreteInputs, 0xFEED, 0xBEAD, []byte{0x02, 0xFE, 0xED, 0xBE, 0xAD}},
		{"HoldingRegisters", pdu.ReadHoldingRegisters, 0xFEED, 0xBEAD, []byte{0x03, 0xFE, 0xED, 0xBE, 0xAD}},
		{"InputRegisters", pdu.ReadInputRegisters, 0xFEED, 0xBEAD, []byte{0x04, 0xFE, 0xED, 0xBE, 0xAD}},
		{"WriteCoilOn", pdu.WriteCoil, 0xFEED, 0x1234, []byte{0x05, 0xFE, 0xED, 0xFF, 0x00}},
		{"WriteCoilOff", pdu.WriteCoil, 0xFEED, 0x0000, []byte{0x05, 0xFE, 0xED, 0x00, 0x00}},
		{"WriteRegister", pdu.WriteRegister, 0xFEED, 0xBEAD, []byte{0x06, 0xFE, 0xED, 0xBE, 0xAD}},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				m := tc.Call(tc.First, tc.Second)
				buf := make([]byte, len(tc.Expected))
				n, err := m.Read(buf)
				if err != nil {
					t.Fatal(err)
				}
				if n != len(tc.Expected) {
					t.Logf("expected %v bytes but received %v", len(tc.Expected), n)
					t.Fail()
				}
				for i := 0; i < len(buf); i++ {
					if buf[i] != tc.Expected[i] {
						t.Logf("failure at position %v: expected %v bytes but received %v", i, tc.Expected[i], buf[i])
						t.Fail()
					}
				}
				wg.Done()
			}()
			go func() {
				m := &pdu.ModbusPDU{}
				n, err := m.Write(tc.Expected)
				if err != nil {
					t.Fatal(err)
				}
				if n != len(tc.Expected) {
					t.Logf("expected %v bytes but received %v", len(tc.Expected), n)
					t.Fail()
				}
				if byte(m.FuncCode) != tc.Expected[0] {

					t.Logf("expected func code %v but received %v", tc.Expected[0], m.FuncCode)
					t.Fail()
				}
				for i := 1; i < len(tc.Expected); i++ {
					if m.Data[i-1] != tc.Expected[i] {
						t.Logf("failure at position %v: expected %v bytes but received %v", i, tc.Expected[i], m.Data[i])
						t.Fail()
					}
				}
				wg.Done()
			}()
			wg.Wait()
		})
	}

}

func TestEcho(t *testing.T) {
	m := pdu.Echo(0x12, 0x34, 0x32, 0x10)
	expected := []byte{0x00, 0x00, 0x12, 0x34, 0x32, 0x10}
	for i := 0; i < 4; i++ {
		if m.Data[i] != expected[i] {
			t.Fail()
			return
		}
	}
}
