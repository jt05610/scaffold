package device_test

import (
	"fmt"
	"os"
	"path"
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
	fmt.Print("\n\nTARING\n\n")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("time_ms\tforce")
	start := time.Now()
	totalForces := int32(0)
	for i := 0; i < 50; i++ {
		force, err := c.GetForce()
		totalForces += force
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), force)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	zero := totalForces / 50
	for {
		force, err := c.GetForce()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), calibrateForce(zero, force))
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
	tim := time.Now()
	n := tim.String() + ".csv"
	res, err := os.Create(path.Join("data", n))
	p := device.Client{}
	c := device.ClientForcePump{}
	fmt.Print("\n\nTARING\n\n")
	_, _ = res.WriteString("\n\nTARING\n\n")

	err = p.PostStart()
	time.Sleep(1 * time.Second)
	fmt.Println("time_ms\tforce")
	_, _ = res.WriteString("time_ms\tforce\n")
	start := time.Now()
	totalForces := int32(0)
	for i := 0; i < 50; i++ {
		force, err := c.GetForce()
		totalForces += force
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), force)
		_, _ = res.WriteString(fmt.Sprintf("%v\t%v\n", time.Now().Sub(start).Milliseconds(), force))
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	zero := totalForces / 50
	err = p.PostTargetVel(300 / 15)
	if err != nil {
		t.Fail()
	}
	err = c.PostTargetVel(5573 / 30)
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
		t.Fail()
	}
	nedTarget := 1600
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
	target := 5573

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
	_, _ = res.WriteString("time_ms\tneedle_pos_µm\tsyringe_pos_µm\tforce\n")
	start = time.Now()
	startPos := uint16(0)

	for nedPos, _ := p.GetCurrentPos(); nedPos < uint16(nedTarget); {
		nedPos, _ = p.GetCurrentPos()
		curPos, _ := c.GetCurrentPos()
		force, _ := c.GetForce()
		fmt.Printf("%v\t%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), nedPos, curPos-startPos, calibrateForce(force, zero))
		_, _ = res.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\n", time.Now().Sub(start).Milliseconds(), nedPos, curPos-startPos, calibrateForce(force, zero)))
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	err = c.PostEnable(0)

	nedTarget = 1300
	err = p.PostTargetPos(nedTarget)

	err = p.PostTargetVel(1000)

	err = p.PostStart()

}

func calibrateForce(reading int32, zero int32) float32 {
	return (float32(reading) - float32(zero)) / 5920.3
}
