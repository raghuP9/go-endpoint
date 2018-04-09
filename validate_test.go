package main

import (
	"testing"
)

type MockOSArgs1 struct {
}

// GetInsecureFlag function
func (a MockOSArgs1) GetInsecureFlag() bool {
	return true
}

// GetVersionFlag function
func (a MockOSArgs1) GetVersionFlag() bool {
	return false
}

// GetIntervalFlag function
func (a MockOSArgs1) GetIntervalFlag() int {
	return 10
}

// GetProtocolFlag function
func (a MockOSArgs1) GetProtocolFlag() string {
	return "https"
}

// GetHostFlag function
func (a MockOSArgs1) GetHostFlag() string {
	return "mydomain.com"
}

// GetShowFlag function
func (a MockOSArgs1) GetShowFlag() bool {
	return true
}

func TestValidate_1(t *testing.T) {
	cArgs := new(CLIArgs)
	cArgs.ARGS = new(MockOSArgs1)
	if !Validate(*cArgs) {
		t.Fatalf("Validation failed")
	}
}

type MockOSArgs2 struct {
}

// GetInsecureFlag function
func (a MockOSArgs2) GetInsecureFlag() bool {
	return true
}

// GetVersionFlag function
func (a MockOSArgs2) GetVersionFlag() bool {
	return false
}

// GetIntervalFlag function
func (a MockOSArgs2) GetIntervalFlag() int {
	return 10
}

// GetProtocolFlag function
func (a MockOSArgs2) GetProtocolFlag() string {
	return "tcp"
}

// GetHostFlag function
func (a MockOSArgs2) GetHostFlag() string {
	return "mydomain.com"
}

// GetShowFlag function
func (a MockOSArgs2) GetShowFlag() bool {
	return true
}

func TestValidate_2(t *testing.T) {
	cArgs := new(CLIArgs)
	cArgs.ARGS = new(MockOSArgs2)
	if Validate(*cArgs) {
		t.Fatalf("Validation Succeeded")
	}
}

type MockOSArgs3 struct {
}

// GetInsecureFlag function
func (a MockOSArgs3) GetInsecureFlag() bool {
	return true
}

// GetVersionFlag function
func (a MockOSArgs3) GetVersionFlag() bool {
	return false
}

// GetIntervalFlag function
func (a MockOSArgs3) GetIntervalFlag() int {
	return 10
}

// GetProtocolFlag function
func (a MockOSArgs3) GetProtocolFlag() string {
	return "https"
}

// GetHostFlag function
func (a MockOSArgs3) GetHostFlag() string {
	return ""
}

// GetShowFlag function
func (a MockOSArgs3) GetShowFlag() bool {
	return true
}

func TestValidate_3(t *testing.T) {
	cArgs := new(CLIArgs)
	cArgs.ARGS = new(MockOSArgs3)
	if Validate(*cArgs) {
		t.Fatalf("Validation Succeeded")
	}
}

type MockOSArgs4 struct {
}

// GetInsecureFlag function
func (a MockOSArgs4) GetInsecureFlag() bool {
	return true
}

// GetVersionFlag function
func (a MockOSArgs4) GetVersionFlag() bool {
	return false
}

// GetIntervalFlag function
func (a MockOSArgs4) GetIntervalFlag() int {
	return 10
}

// GetProtocolFlag function
func (a MockOSArgs4) GetProtocolFlag() string {
	return ""
}

// GetHostFlag function
func (a MockOSArgs4) GetHostFlag() string {
	return "mydomain.com"
}

// GetShowFlag function
func (a MockOSArgs4) GetShowFlag() bool {
	return true
}

func TestValidate_4(t *testing.T) {
	cArgs := new(CLIArgs)
	cArgs.ARGS = new(MockOSArgs4)
	if Validate(*cArgs) {
		t.Fatalf("Validation Succeeded")
	}
}

type MockOSArgs5 struct {
}

// GetInsecureFlag function
func (a MockOSArgs5) GetInsecureFlag() bool {
	return true
}

// GetVersionFlag function
func (a MockOSArgs5) GetVersionFlag() bool {
	return false
}

// GetIntervalFlag function
func (a MockOSArgs5) GetIntervalFlag() int {
	return -1
}

// GetProtocolFlag function
func (a MockOSArgs5) GetProtocolFlag() string {
	return "https"
}

// GetHostFlag function
func (a MockOSArgs5) GetHostFlag() string {
	return "mydomain.com"
}

// GetShowFlag function
func (a MockOSArgs5) GetShowFlag() bool {
	return true
}

func TestValidate_5(t *testing.T) {
	cArgs := new(CLIArgs)
	cArgs.ARGS = new(MockOSArgs5)
	if Validate(*cArgs) {
		t.Fatalf("Validation Succeeded")
	}
}
