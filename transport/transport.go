package transport

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// Time Constants
const (
	ConntectionTimeout = 60
	KeepAliveTime      = 60
	HandshakeTimeout   = 20
)

// CustomTransport struct
type CustomTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

// NewTransport returns a transport object
func NewTransport(insecure bool) *CustomTransport {

	tr := &CustomTransport{
		dialer: &net.Dialer{
			Timeout:   ConntectionTimeout * time.Second,
			KeepAlive: KeepAliveTime * time.Second,
		},
	}
	tr.rtp = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dial,
		TLSHandshakeTimeout: HandshakeTimeout * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
	}
	return tr
}

// RoundTrip gives round trip time of a request
func (tr *CustomTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

// dial to addr
func (tr *CustomTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

// ReqDuration returns duration of the request
func (tr *CustomTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

// ConnDuration returns duration of the connection
func (tr *CustomTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

// Duration returns entire duration of the connection
// and request
func (tr *CustomTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}
