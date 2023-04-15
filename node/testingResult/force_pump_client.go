// AUTO GENERATED FILE, DO NOT CHANGE

package device

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ClientForcePump struct {
}

func NewForcePumpClient() *ClientForcePump {
	return &ClientForcePump{}
}

// GetIsMoving returns 1 if needle is currently moving, otherwise 0.
func (c *ClientForcePump) GetIsMoving() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/is_moving")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// GetStart starts needle.
func (c *ClientForcePump) GetStart() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/start")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostStart starts needle.
func (c *ClientForcePump) PostStart() error {
	buf := new(bytes.Buffer)

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/start", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetStop stops needle.
func (c *ClientForcePump) GetStop() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/stop")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostStop stops needle.
func (c *ClientForcePump) PostStop() error {
	buf := new(bytes.Buffer)

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/stop", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetHome homes needle.
func (c *ClientForcePump) GetHome() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/home")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostHome homes needle.
func (c *ClientForcePump) PostHome() error {
	buf := new(bytes.Buffer)

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/home", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetMoveToStall Moves as far forward as possible until stall position is reached.
func (c *ClientForcePump) GetMoveToStall() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/move_to_stall")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostMoveToStall Moves as far forward as possible until stall position is reached.
func (c *ClientForcePump) PostMoveToStall() error {
	buf := new(bytes.Buffer)

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/move_to_stall", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetSetZero Sets current position as 0.
func (c *ClientForcePump) GetSetZero() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/set_zero")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostSetZero Sets current position as 0.
func (c *ClientForcePump) PostSetZero() error {
	buf := new(bytes.Buffer)

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/set_zero", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetEnable returns if needle is enabled. Enable needle by writing 1.
func (c *ClientForcePump) GetEnable() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/enable")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostEnable returns if needle is enabled. Enable needle by writing 1.
func (c *ClientForcePump) PostEnable(enabled int) error {
	buf := new(bytes.Buffer)

	req := struct {
		Enabled uint16 `goClient:"enabled"`
	}{
		Enabled: uint16(enabled),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/enable", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetCurrentPos returns the current position of the needle in µm.
func (c *ClientForcePump) GetCurrentPos() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/current_pos")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// GetCurrentVel returns the current velocity of the needle in µm / s.
func (c *ClientForcePump) GetCurrentVel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/current_vel")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// GetTstep returns measured time between steps.
func (c *ClientForcePump) GetTstep() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/tstep")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// GetForce returns raw force reading from pump.
func (c *ClientForcePump) GetForce() (r int32, err error) {
	res, err := http.Get("https://127.0.0.1:8081/force_pump/force")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(res.Body)
	arrMap := make(map[string][]uint16, 0)
	err = d.Decode(&arrMap)
	if err != nil {
		panic(err)
	}
	v := arrMap["result"]
	r = (int32(v[0]) << 16) + int32(v[1])
	return r, err
}

// GetTargetPos returns target position in µm.
func (c *ClientForcePump) GetTargetPos() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/target_pos")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostTargetPos returns target position in µm.
func (c *ClientForcePump) PostTargetPos(pos int) error {
	buf := new(bytes.Buffer)

	req := struct {
		Pos uint16 `goClient:"pos"`
	}{
		Pos: uint16(pos),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/target_pos", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetTargetVel returns target velocity in µm / s.
func (c *ClientForcePump) GetTargetVel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/target_vel")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostTargetVel returns target velocity in µm / s.
func (c *ClientForcePump) PostTargetVel(vel int) error {
	buf := new(bytes.Buffer)

	req := struct {
		Vel uint16 `goClient:"vel"`
	}{
		Vel: uint16(vel),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/target_vel", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetMoveTo returns distance from target position in µm.
func (c *ClientForcePump) GetMoveTo() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/move_to")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostMoveTo returns distance from target position in µm.
func (c *ClientForcePump) PostMoveTo(pos int) error {
	buf := new(bytes.Buffer)

	req := struct {
		Pos uint16 `goClient:"pos"`
	}{
		Pos: uint16(pos),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/move_to", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetAccel returns the current acceleration of the needle in µm / s^2.
func (c *ClientForcePump) GetAccel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/accel")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostAccel returns the current acceleration of the needle in µm / s^2.
func (c *ClientForcePump) PostAccel(accel int) error {
	buf := new(bytes.Buffer)

	req := struct {
		Accel uint16 `goClient:"accel"`
	}{
		Accel: uint16(accel),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/accel", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetStallguard returns current stallguard settings from TMC2209.
func (c *ClientForcePump) GetStallguard() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/force_pump/stallguard")
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(r.Body)
	res := make(map[string]uint16, 0)
	err = d.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res["result"], err
}

// PostStallguard returns current stallguard settings from TMC2209.
func (c *ClientForcePump) PostStallguard(sgThresh int) error {
	buf := new(bytes.Buffer)

	req := struct {
		SgThresh uint16 `goClient:"sgThresh"`
	}{
		SgThresh: uint16(sgThresh),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://127.0.0.1:8081/force_pump/stallguard", "application/json", buf)
	if err != nil {
		panic(err)
	}

	respBuf := make([]byte, 2)
	_, err = resp.Body.Read(respBuf)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 && string(respBuf) == "ok" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}
