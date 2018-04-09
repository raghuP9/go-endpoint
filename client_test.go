package main

import (
	"fmt"
	"testing"
)

var (
	url = "https://mydomain.com"
)

// Mock successful client
type SuccessClientMock struct {
}

// Mock failed client
type FailClientMock struct {
}

// Mock http response body
var testResp = `{"status": "active"}`

// Mock Successful Get method
func (c *SuccessClientMock) Get(ep string) ([]byte, error) {
	return []byte(testResp), nil
}

// Mock Failed Get method
func (c *FailClientMock) Get(ep string) ([]byte, error) {
	err := fmt.Errorf("GET request failed")
	return nil, err
}

func TestClient_1(t *testing.T) {
	client := new(Client)
	client.HTTP = new(SuccessClientMock)
	body, err := client.HTTP.Get(host)

	if string(body) != string([]byte(testResp)) {
		t.Fatalf("Expected response [%x] to equal [%+v]", body, testResp)
	}
	if err != nil {
		t.Fatalf("Expected error to be nil but received: [%s]", err.Error())
	}
}

func TestClient_2(t *testing.T) {
	client := new(Client)
	client.HTTP = new(FailClientMock)
	body, err := client.HTTP.Get(host)

	if body != nil {
		t.Fatalf("Expected response [%x] to be nil", body)
	}
	if err == nil {
		t.Fatalf("Expected error not to be nil")
	}
}
