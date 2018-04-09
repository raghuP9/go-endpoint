package main

import (
  "testing"
	"time"
)

var insecure = true

func TestReqDuration(t *testing.T) {
  tp := NewTransport(insecure)
  tp.reqStart = time.Now()
  tp.reqEnd = tp.reqStart.Add(time.Second * 30)
  tp.connStart = tp.reqStart.Add(time.Second * 10)
  tp.connEnd = tp.connStart.Add(time.Second * 10)

  if tp.Duration() != (time.Second * 30) {
    t.Fatalf("Wrong duration reported!")
  }

  if tp.ConnDuration() != (time.Second * 10) {
    t.Fatalf("Wrong connection duration reported!")
  }

  if tp.ReqDuration() != (time.Second * 20) {
    t.Fatalf("Wrong request duration reported!")
  }
}
