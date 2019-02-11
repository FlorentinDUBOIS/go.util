package http

import (
	"net"
	"net/http"
)

// Client structure hold http client and methods to interact with APIs.
type Client struct {
	inner *http.Client
}

// NewClient return a new instance of `Client`.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{
			Timeout: DefaultTimeout,
			Transport: &http.Transport{
				DisableKeepAlives: DefaultDisableKeepAlives,
				Dial: (&net.Dialer{
					Timeout: DefaultDialerTimeout,
				}).Dial,
				TLSHandshakeTimeout: DefaultTLSHandshakeTimeout,
				MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
				IdleConnTimeout:     DefaultIdleConnTimeout,
			},
		}
	}

	return &Client{
		inner: client,
	}
}

// Do execute the given request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.inner.Do(req)
}

// R return a new request instance with the current client as inner client
func (c *Client) R() *Request {
	return NewRequest(c)
}
