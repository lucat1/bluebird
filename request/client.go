package request

import (
	"net/http"
	"net/url"
	"time"

	"git.hjkl.gq/luca.tagliavini5/oauth1"
)

type RequestClient struct {
	URL  *url.URL
	HTTP *http.Client
}

type RequestClientV1 struct {
	UploadURL *url.URL
	APIURL    *url.URL
	HTTP      *http.Client
}

var (
	client   *RequestClient
	v1Client *RequestClientV1
)

type transportWithHeader struct {
	bearer string
}

func (t *transportWithHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.bearer)
	return http.DefaultTransport.RoundTrip(req)
}

// NewClient instantiates a new http.Client for use with a bearer token authentication
func NewClient(URL, token string) (client *RequestClient, err error) {
	url, err := url.Parse(URL)
	if err != nil {
		return
	}

	client = &RequestClient{
		URL:  url,
		HTTP: &http.Client{Transport: &transportWithHeader{bearer: token}},
	}
	return
}

type transportWithV1Header struct {
	token  *oauth1.Token
	auther *oauth1.Auther
}

func (t *transportWithV1Header) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.auther.SetRequestAuthHeader(req, t.token); err != nil {
		return nil, err
	}
	return http.DefaultTransport.RoundTrip(req)
}

// NewV1Client instantiates a new http.Client for use with the v1 authentication system
func NewV1Client(uploadURL, apiURL, consumerKey, consumerSecret, oauthToken, oauthSecret string) (client *RequestClientV1, err error) {
	uu, err := url.Parse(uploadURL)
	if err != nil {
		return
	}
	au, err := url.Parse(apiURL)
	if err != nil {
		return
	}

	var transport http.RoundTripper = nil
	if consumerKey != "" {
		config := &oauth1.Config{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    "http://team14.hjkl.gq/",
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: "https://api.twitter.com/oauth/request_token",
				AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
				AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
			},
			Noncer: oauth1.Base64Noncer{},
		}
		clock := oauth1.NewFixedClock(time.Now())
		auther := &oauth1.Auther{
			Config: config,
			Clock:  &clock,
		}
		token := oauth1.NewToken(oauthToken, oauthSecret)

		t := transportWithV1Header{token, auther}
		transport = &t
	}
	client = &RequestClientV1{
		UploadURL: uu,
		APIURL:    au,
		HTTP:      &http.Client{Transport: transport},
	}
	return
}

// SetClient sets the http.Client to use for HTTP requests
func SetClient(c *RequestClient) {
	client = c
}

// SetV1Client sets the http.Client to use for HTTP requests (for the V1 api)
func SetV1Client(c *RequestClientV1) {
	v1Client = c
}

// Client returns the current http.Client being used to send requests
func Client() *RequestClient {
	return client
}

// Client returns the current http.Client being used to send requests
func V1Client() *RequestClientV1 {
	return v1Client
}
