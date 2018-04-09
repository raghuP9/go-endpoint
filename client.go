package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Http RealWebClient
)

type RealWebClient struct {
	Transport *CustomTransport
}

type WebClient interface {
	Get(string) ([]byte, error)
}

type Client struct {
	Http WebClient
}

func (r RealWebClient) Get(url string) ([]byte, error) {
	var body []byte
	client := &http.Client{Transport: r.Transport}

	if resp, err := client.Get(url); nil != err {
		log.Printf(err.Error())
		return nil, err
	} else {
		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)
	}
	return body, nil
}
