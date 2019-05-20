package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// RealWebClient struct
type RealWebClient struct {
	Transport *CustomTransport
}

// WebClient interface
type WebClient interface {
	Get(string) ([]byte, error)
}

// Client struct
type Client struct {
	HTTP WebClient
}

// Get method for interface
func (r RealWebClient) Get(url string) ([]byte, error) {
	var body []byte
	client := &http.Client{Transport: r.Transport}

	resp, err := client.Get(url)
	if nil != err {
		log.Printf(err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return body, nil
}
