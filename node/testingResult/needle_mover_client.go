
// AUTO GENERATED FILE, DO NOT CHANGE

package device

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
}

func NewClient() *Client{
	return &Client{}
}

// GetTargetPos returns target position in µm.
func (c *Client)GetTargetPos() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/target_pos")
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
func (c *Client)PostTargetPos(pos int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		Pos uint16	`json:"pos"`
	}{
		Pos: uint16(pos),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/target_pos", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetTargetVel returns target velocity in µm / s.
func (c *Client)GetTargetVel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/target_vel")
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
func (c *Client)PostTargetVel(vel int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		Vel uint16	`json:"vel"`
	}{
		Vel: uint16(vel),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/target_vel", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetMoveTo returns distance from target position in µm.
func (c *Client)GetMoveTo() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/move_to")
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
func (c *Client)PostMoveTo(pos int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		Pos uint16	`json:"pos"`
	}{
		Pos: uint16(pos),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/move_to", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetAccel returns the current acceleration of the needle in µm / s^2.
func (c *Client)GetAccel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/accel")
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
func (c *Client)PostAccel(accel int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		Accel uint16	`json:"accel"`
	}{
		Accel: uint16(accel),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/accel", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetStallguard returns current stallguard settings from TMC2209.
func (c *Client)GetStallguard() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/stallguard")
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
func (c *Client)PostStallguard(sgThresh int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		SgThresh uint16	`json:"sgThresh"`
	}{
		SgThresh: uint16(sgThresh),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/stallguard", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetIsMoving returns 1 if needle is currently moving, otherwise 0.
func (c *Client)GetIsMoving() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/is_moving")
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
func (c *Client)GetStart() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/start")
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
func (c *Client)PostStart() error {
	buf := new(bytes.Buffer)
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/start", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetStop stops needle.
func (c *Client)GetStop() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/stop")
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
func (c *Client)PostStop() error {
	buf := new(bytes.Buffer)
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/stop", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetHome homes needle.
func (c *Client)GetHome() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/home")
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
func (c *Client)PostHome() error {
	buf := new(bytes.Buffer)
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/home", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetMoveToStall Moves as far forward as possible until stall position is reached.
func (c *Client)GetMoveToStall() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/move_to_stall")
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
func (c *Client)PostMoveToStall() error {
	buf := new(bytes.Buffer)
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/move_to_stall", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetSetZero Sets current position as 0.
func (c *Client)GetSetZero() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/set_zero")
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
func (c *Client)PostSetZero() error {
	buf := new(bytes.Buffer)
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/set_zero", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetEnable returns if needle is enabled. Enable needle by writing 1.
func (c *Client)GetEnable() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/enable")
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
func (c *Client)PostEnable(enabled int) error {
	buf := new(bytes.Buffer)
	
	
	req := struct{
		Enabled uint16	`json:"enabled"`
	}{
		Enabled: uint16(enabled),
	}
	err := json.NewEncoder(buf).Encode(&req)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post("https://127.0.0.1:8081/needle_mover/enable", "application/json", buf)
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
	}  else {
		return errors.New(fmt.Sprintf("request failed with status code %v", resp.StatusCode))
	}
}

// GetCurrentPos returns the current position of the needle in µm.
func (c *Client)GetCurrentPos() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/current_pos")
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
func (c *Client)GetCurrentVel() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/current_vel")
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
func (c *Client)GetTstep() (uint16, error) {
	r, err := http.Get("https://127.0.0.1:8081/needle_mover/tstep")
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
