package scour

import (
	"errors"
	"net/http"
	"net/url"
)

type Client struct {
	*http.Client

	req  *http.Request
	resp *http.Response
}

// Create a new Client with and embedded
// http.Client type
func NewClient() *Client {
	return &Client{&http.Client{}, &http.Request{}, &http.Response{}}
}

// Request is used to initiate a client request
// and return a crawler object
func (c *Client) Request(method string, url string) (ret *Crawler, err error) {
	c.req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	c.resp, err = c.Do(c.req)
	if err != nil {
		return nil, err
	}

	crawler := NewCrawler(c.resp)

	return crawler, nil
}

// Click is a convenience function simulating a client browser link click
// Returns a scour crawler object
func (c *Client) Click(l *Link) (ret *Crawler, err error) {
	if l == nil {
		return nil, errors.New("Can't click without a link!")
	}
	return c.Request(l.Method(), l.Url())
}

// SetProxy is a convenience function for setting
// a proxy url used by the client object
func (c *Client) SetProxy(proxy_str string) error {
	proxy, err := url.Parse(proxy_str)
	if err != nil {
		return errors.New("Problem parsing proxy url")
	}

	c.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}

	return nil
}
