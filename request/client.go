package request

import (
	"net/http"
	"net/url"
)

type RequestClient struct {
	URL       *url.URL
	UploadURL *url.URL
	HTTP      *http.Client
}

var client *RequestClient

type transportWithHeader struct {
	bearer string
}

func (t *transportWithHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.bearer)
	return http.DefaultTransport.RoundTrip(req)
}

// NewClient instantiates a new http.Client for use with a bearer token authentication
func NewClient(URL string, UploadURL string, token string) (client *RequestClient, err error) {
	url, err := url.Parse(URL)
	if err != nil {
		return
	}

	uploadURL, err := url.Parse(UploadURL)
	if err != nil {
		return
	}

	client = &RequestClient{
		URL:       url,
		UploadURL: uploadURL,
		HTTP:      &http.Client{Transport: &transportWithHeader{bearer: token}},
	}
	return
}

// SetClient sets the http.Client to use for HTTP requests
func SetClient(c *RequestClient) {
	client = c
}

// Client returns the current http.Client being used to send requests
func Client() *RequestClient {
	return client
}
