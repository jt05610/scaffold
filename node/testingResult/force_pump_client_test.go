package device_test

import (
	"fmt"
	device "scaffold/node/testingResult"
	"testing"
	"time"
)

func TestClientForcePump_PostHome(t *testing.T) {
	c := device.ClientForcePump{}
	r := c.PostHome()
	if r != nil {
		t.Fail()
	}
	start := time.Now()
	fmt.Println("time_ms\tstallguard\ttstep\tforce")
	for isHome, _ := c.GetHome(); isHome != 1; {
		isHome, _ = c.GetHome()
		v, err := c.GetStallguard()
		if err != nil {
			panic(err)
		}
		tstep, err := c.GetTstep()
		if err != nil {
			panic(err)
		}
		force, err := c.GetForce()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), v, tstep, force)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	r = c.PostEnable(0)
	if r != nil {
		t.Fail()
	}
}

func TestClientForceReading(t *testing.T) {
	c := device.NewForcePumpClient()
	fmt.Println("time_ms\tforce")
	start := time.Now()
	for {
		force, err := c.GetForce()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), force)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func TestClientForcePump_Stall(t *testing.T) {
	c := device.ClientForcePump{}
	r := c.PostMoveToStall()
	if r != nil {
		t.Fail()
	}
	start := time.Now()
	fmt.Println("time_ms\tstallguard\ttstep\tforce")
	for isHome, _ := c.GetMoveToStall(); isHome != 1; {
		isHome, _ = c.GetMoveToStall()
		v, err := c.GetStallguard()
		if err != nil {
			panic(err)
		}
		tstep, err := c.GetTstep()
		if err != nil {
			panic(err)
		}
		force, err := c.GetForce()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), v, tstep, force)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	r = c.PostEnable(0)
	if r != nil {
		t.Fail()
	}
}
func TestForcePumpClient_GoTo(t *testing.T) {
	p := device.Client{}
	c := device.ClientForcePump{}
	err := p.PostTargetVel(50)
	if err != nil {
		t.Fail()
	}
	err = c.PostTargetVel(100)
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
	r, _ = p.GetCurrentPos()
	if r != 0 {
		err := p.PostSetZero()
		if err != nil {
			panic(err)
		}
	}

	r, _ = p.GetCurrentPos()
	if r != 0 {
		t.Fail()
	}
	nedTarget := 1500
	err = p.PostTargetPos(nedTarget)

	if err != nil {
		t.Fail()
	}
	r, err = p.GetTargetPos()
	if err != nil {
		t.Fail()
	}
	if r != uint16(nedTarget) {
		t.Fail()
	}
	target := 10000

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
	err = p.PostStart()
	if err != nil {
		t.Fail()
	}
	fmt.Println("time_ms\tneedle_pos_µm\tsyringe_pos_µm\tforce")
	start := time.Now()
	startPos := uint16(0)

	for nedPos, _ := p.GetCurrentPos(); nedPos < uint16(nedTarget); {
		nedPos, _ = p.GetCurrentPos()
		curPos, _ := c.GetCurrentPos()
		force, err := c.GetForce()
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\t%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), nedPos, curPos-startPos, force)
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	err = c.PostEnable(0)
	if err != nil {
		t.Fail()
	}
}
