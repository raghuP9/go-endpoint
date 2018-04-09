package main

import (
	"testing"
)

var (
	host     = "mydomain.com"
	protocol = "https"
	show     = true
)

func TestMonitor_1(t *testing.T) {
	pass := make(chan string, 1)
	fail := make(chan string, 1)

	passmsg, failmsg := "", ""

	client := new(Client)
	client.HTTP = new(SuccessClientMock)

	go monitor(client, protocol, host, show, pass, fail)
	select {
	case passmsg = <-pass:
	case failmsg = <-fail:
	}

	if failmsg != "" {
		t.Fatalf("Expected monitor to pass")
	}
	if passmsg == "" {
		t.Fatalf("Expected monitor to pass")
	}
}

func TestMonitor_2(t *testing.T) {
	pass := make(chan string, 1)
	fail := make(chan string, 1)

	passmsg, failmsg := "", ""

	client := new(Client)
	client.HTTP = new(FailClientMock)

	go monitor(client, "https", "mydomain.com", false, pass, fail)
	select {
	case passmsg = <-pass:
	case failmsg = <-fail:
	}

	if failmsg == "" {
		t.Fatalf("Expected monitor to fail")
	}
	if passmsg != "" {
		t.Fatalf("Expected monitor to fail")
	}
}
