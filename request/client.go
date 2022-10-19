package request

import (
	"net/http"
)

var client *http.Client

type transportWithHeader struct {
	bearer string
}

func (t *transportWithHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.bearer)
	return http.DefaultTransport.RoundTrip(req)
}

// NewClient instantiates a new http.Client for use with a bearer token authentication
func NewClient(token string) *http.Client {
	return &http.Client{Transport: &transportWithHeader{bearer: token}}
}

// SetClient sets the http.Client to use for HTTP requests
func SetClient(c *http.Client) {
	client = c
}

// Client returns the current http.Client being used to send requests
func Client() *http.Client {
	return client
}
