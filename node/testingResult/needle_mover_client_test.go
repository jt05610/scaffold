package device_test

import (
	"fmt"
	device "scaffold/node/testingResult"
	"testing"
	"time"
)

func TestClient_Home(t *testing.T) {
	c := device.Client{}
	r := c.PostHome()
	if r != nil {
		t.Fail()
	}
	r = c.PostEnable(1)
	if r != nil {
		t.Fail()
	}
	start := time.Now()
	fmt.Println("time_ms\tstallguard\ttstep")
	for {
		v, err := c.GetStallguard()
		if err != nil {
			panic(err)
		}
		tstep, err := c.GetTstep()
		if err != nil {
			panic(err)
		}
		if tstep >= 65535 {
			break
		}
		fmt.Printf("%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), v, tstep)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	r = c.PostEnable(0)
	if r != nil {
		t.Fail()
	}
}

func TestClient_Stall(t *testing.T) {
	c := device.Client{}
	r := c.PostMoveToStall()
	if r != nil {
		t.Fail()
	}
	start := time.Now()
	fmt.Println("time_ms\tstallguard\ttstep")
	for {
		v, err := c.GetStallguard()
		if err != nil {
			panic(err)
		}
		tstep, err := c.GetTstep()
		if tstep >= 65535 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), v, tstep)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	r = c.PostEnable(0)
	if r != nil {
		t.Fail()
	}
}
func TestClient_TargetPos(t *testing.T) {
	c := device.Client{}

	err := c.PostTargetVel(1000)
	if err != nil {
		t.Fail()
	}
	r, err := c.GetTargetVel()
	if err != nil {
		t.Fail()
	}
	if r != 100 {
		t.Fail()
	}
	r, _ = c.GetCurrentPos()
	if r != 0 {
		err := c.PostSetZero()
		if err != nil {
			panic(err)
		}
	}

	r, _ = c.GetCurrentPos()
	if r != 0 {
		t.Fail()
	}

	target := 300

	err = c.PostTargetPos(target)

	if err != nil {
		t.Fail()
	}
	r, err = c.GetTargetPos()
	if err != nil {
		t.Fail()
	}
	if r != uint16(target) {
		t.Fail()
	}
	err = c.PostStart()
	if err != nil {
		t.Fail()
	}
	start := time.Now()
	startPos := uint16(0)
	for curPos, _ := c.GetCurrentPos(); curPos < uint16(target); {
		curPos, _ = c.GetCurrentPos()
		fmt.Printf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), curPos-startPos)
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	err = c.PostEnable(0)
	if err != nil {
		t.Fail()
	}
}
