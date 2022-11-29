package chess

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"github.com/dghubble/oauth1"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

const (
	playerColor = chess.White
	stateFile   = "chess.json"
)

var match *Match = nil

type Match struct {
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	EndsAt   time.Time     `json:"ends_at"`
	Game     *chess.Game   `json:"game"`
	TweetID  string        `json:"tweet_id"`

	timeout chan bool
	ticking atomic.Bool
}

func (m *Match) delay() {
	if m.ticking.Load() {
		m.timeout <- true
	}
	m.timeout = make(chan bool)

	go func() {
		m.ticking.Store(true)
		select {
		case <-time.After(match.Duration):
			m.ticking.Store(false)
			m.onTurnEnd()

		case <-match.timeout:
			m.ticking.Store(false)
			log.Println("Timeout cancled")
		}
	}()
}

func (m *Match) getMoves() (moves map[string]uint, err error) {
	tweets, err := request.Replies(m.TweetID, 100, "", "")
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		clone := match.Game.Clone()
		first := strings.Split(tweet.Text, " ")[0]
		if err = clone.MoveStr(first); err != nil {
			moves[first]++
		}
	}
	return
}

func (m *Match) onTurnEnd() {
	if m.Game.Position().Turn() == playerColor {
		// TODO: forfeit
	} else {
		moves, err := m.getMoves()
		if err != nil {
			log.Printf("Could not get tweets replies: %v", err)
			return
		}
		var (
			mostRated  string
			mostValued uint
		)
		for move, val := range moves {
			if val > mostValued {
				mostRated = move
				mostValued = val
			}
		}
		m.Game.MoveStr(mostRated)
	}
}

func (m *Match) Move(move string) error {
	if match.Game.Position().Turn() != playerColor {
		return errors.New("Cannot move in the opponent's turn")
	}
	if err := match.Game.MoveStr(move); err != nil {
		return err
	}

	m.PostGame()
	m.EndsAt = time.Now().UTC().Add(m.Duration)
	m.delay()
	return nil
}

type transportPost struct {
	bearer string
}

func (t *transportPost) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.bearer)
	return http.DefaultTransport.RoundTrip(req)
}

const (
	authorizationHeaderParam  = "Authorization"
	authorizationPrefix       = "OAuth " // trailing space is intentional
	oauthConsumerKeyParam     = "oauth_consumer_key"
	oauthNonceParam           = "oauth_nonce"
	oauthSignatureParam       = "oauth_signature"
	oauthSignatureMethodParam = "oauth_signature_method"
	oauthTimestampParam       = "oauth_timestamp"
	oauthTokenParam           = "oauth_token"
	oauthVersionParam         = "oauth_version"
	oauthCallbackParam        = "oauth_callback"
	oauthVerifierParam        = "oauth_verifier"
	defaultOauthVersion       = "1.0"
	contentType               = "Content-Type"
	formContentType           = "application/x-www-form-urlencoded"
	realmParam                = "realm"
)

// clock provides a interface for current time providers. A Clock can be used
// in place of calling time.Now() directly.
type clock interface {
	Now() time.Time
}

// auther adds an "OAuth" Authorization header field to requests.
type auther struct {
	config *oauth1.Config
	clock  clock
}

func newAuther(config *oauth1.Config) *auther {
	if config == nil {
		config = &oauth1.Config{}
	}
	if config.Noncer == nil {
		config.Noncer = oauth1.Base64Noncer{}
	}
	return &auther{
		config: config,
	}
}

// setRequestTokenAuthHeader adds the OAuth1 header for the request token
// request (temporary credential) according to RFC 5849 2.1.
func (a *auther) setRequestTokenAuthHeader(req *http.Request) error {
	oauthParams := a.commonOAuthParams()
	oauthParams[oauthCallbackParam] = a.config.CallbackURL
	params, err := collectParameters(req, oauthParams)
	if err != nil {
		return err
	}
	signatureBase := signatureBase(req, params)
	signature, err := a.signer().Sign("", signatureBase)
	if err != nil {
		return err
	}
	oauthParams[oauthSignatureParam] = signature
	if a.config.Realm != "" {
		oauthParams[realmParam] = a.config.Realm
	}
	fmt.Println(oauthParams)
	req.Header.Set(authorizationHeaderParam, authHeaderValue(oauthParams))
	return nil
}

// setAccessTokenAuthHeader sets the OAuth1 header for the access token request
// (token credential) according to RFC 5849 2.3.
func (a *auther) setAccessTokenAuthHeader(req *http.Request, requestToken, requestSecret, verifier string) error {
	oauthParams := a.commonOAuthParams()
	oauthParams[oauthTokenParam] = requestToken
	oauthParams[oauthVerifierParam] = verifier
	params, err := collectParameters(req, oauthParams)
	if err != nil {
		return err
	}
	signatureBase := signatureBase(req, params)
	signature, err := a.signer().Sign(requestSecret, signatureBase)
	if err != nil {
		return err
	}
	oauthParams[oauthSignatureParam] = signature
	req.Header.Set(authorizationHeaderParam, authHeaderValue(oauthParams))
	return nil
}

// setRequestAuthHeader sets the OAuth1 header for making authenticated
// requests with an AccessToken (token credential) according to RFC 5849 3.1.
func (a *auther) setRequestAuthHeader(req *http.Request, accessToken *oauth1.Token) error {
	oauthParams := a.commonOAuthParams()
	oauthParams[oauthTokenParam] = accessToken.Token
	params, err := collectParameters(req, oauthParams)
	if err != nil {
		return err
	}
	signatureBase := signatureBase(req, params)
	signature, err := a.signer().Sign(accessToken.TokenSecret, signatureBase)
	if err != nil {
		return err
	}
	fmt.Println(signature)
	oauthParams[oauthSignatureParam] = signature
	req.Header.Set(authorizationHeaderParam, authHeaderValue(oauthParams))
	return nil
}

// commonOAuthParams returns a map of the common OAuth1 protocol parameters,
// excluding the oauth_signature parameter. This includes the realm parameter
// if it was set in the config. The realm parameter will not be included in
// the signature base string as specified in RFC 5849 3.4.1.3.1.
func (a *auther) commonOAuthParams() map[string]string {
	params := map[string]string{
		oauthConsumerKeyParam:     a.config.ConsumerKey,
		oauthSignatureMethodParam: a.signer().Name(),
		oauthTimestampParam:       strconv.FormatInt(a.epoch(), 10),
		oauthNonceParam:           a.nonce(),
		oauthVersionParam:         defaultOauthVersion,
	}
	if a.config.Realm != "" {
		params[realmParam] = a.config.Realm
	}
	return params
}

// Returns a nonce using the configured Noncer.
func (a *auther) nonce() string {
	return a.config.Noncer.Nonce()
}

// Returns the Unix epoch seconds.
func (a *auther) epoch() int64 {
	if a.clock != nil {
		return a.clock.Now().Unix()
	}
	return time.Now().Unix()
}

// Returns the Config's Signer or the default Signer.
func (a *auther) signer() oauth1.Signer {
	if a.config.Signer != nil {
		return a.config.Signer
	}
	return &oauth1.HMACSigner{ConsumerSecret: a.config.ConsumerSecret}
}

// authHeaderValue formats OAuth parameters according to RFC 5849 3.5.1. OAuth
// params are percent encoded, sorted by key (for testability), and joined by
// "=" into pairs. Pairs are joined with a ", " comma separator into a header
// string.
// The given OAuth params should include the "oauth_signature" key.
func authHeaderValue(oauthParams map[string]string) string {
	pairs := sortParameters(encodeParameters(oauthParams), `%s="%s"`)
	return authorizationPrefix + strings.Join(pairs, ", ")
}

// encodeParameters percent encodes parameter keys and values according to
// RFC5849 3.6 and RFC3986 2.1 and returns a new map.
func encodeParameters(params map[string]string) map[string]string {
	encoded := map[string]string{}
	for key, value := range params {
		encoded[oauth1.PercentEncode(key)] = oauth1.PercentEncode(value)
	}
	return encoded
}

// sortParameters sorts parameters by key and returns a slice of key/value
// pairs formatted with the given format string (e.g. "%s=%s").
func sortParameters(params map[string]string, format string) []string {
	// sort by key
	keys := make([]string, len(params))
	i := 0
	for key := range params {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	// parameter join
	pairs := make([]string, len(params))
	for i, key := range keys {
		pairs[i] = fmt.Sprintf(format, key, params[key])
	}
	return pairs
}

// collectParameters collects request parameters from the request query, OAuth
// parameters (which should exclude oauth_signature), and the request body
// provided the body is single part, form encoded, and the form content type
// header is set. The returned map of collected parameter keys and values
// follow RFC 5849 3.4.1.3, except duplicate parameters are not supported.
func collectParameters(req *http.Request, oauthParams map[string]string) (map[string]string, error) {
	// add oauth, query, and body parameters into params
	params := map[string]string{}
	for key, value := range req.URL.Query() {
		// most backends do not accept duplicate query keys
		params[key] = value[0]
	}
	if req.Body != nil && req.Header.Get(contentType) == formContentType {
		// reads data to a []byte, draining req.Body
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		values, err := url.ParseQuery(string(b))
		if err != nil {
			return nil, err
		}
		for key, value := range values {
			// not supporting params with duplicate keys
			params[key] = value[0]
		}
		// reinitialize Body with ReadCloser over the []byte
		req.Body = ioutil.NopCloser(bytes.NewReader(b))
	}
	for key, value := range oauthParams {
		// according to 3.4.1.3.1. the realm parameter is excluded
		if key != realmParam {
			params[key] = value
		}
	}
	return params, nil
}

// signatureBase combines the uppercase request method, percent encoded base
// string URI, and normalizes the request parameters int a parameter string.
// Returns the OAuth1 signature base string according to RFC5849 3.4.1.
func signatureBase(req *http.Request, params map[string]string) string {
	method := strings.ToUpper(req.Method)
	baseURL := baseURI(req)
	parameterString := normalizedParameterString(params)
	// signature base string constructed accoding to 3.4.1.1
	baseParts := []string{method, oauth1.PercentEncode(baseURL), oauth1.PercentEncode(parameterString)}
	return strings.Join(baseParts, "&")
}

// baseURI returns the base string URI of a request according to RFC 5849
// 3.4.1.2. The scheme and host are lowercased, the port is dropped if it
// is 80 or 443, and the path minus query parameters is included.
func baseURI(req *http.Request) string {
	scheme := strings.ToLower(req.URL.Scheme)
	host := strings.ToLower(req.URL.Host)
	if hostPort := strings.Split(host, ":"); len(hostPort) == 2 && (hostPort[1] == "80" || hostPort[1] == "443") {
		host = hostPort[0]
	}
	// TODO: use req.URL.EscapedPath() once Go 1.5 is more generally adopted
	// For now, hacky workaround accomplishes the same internal escaping mode
	// escape(u.Path, encodePath) for proper compliance with the OAuth1 spec.
	path := req.URL.Path
	if path != "" {
		path = strings.Split(req.URL.RequestURI(), "?")[0]
	}
	return fmt.Sprintf("%v://%v%v", scheme, host, path)
}

// parameterString normalizes collected OAuth parameters (which should exclude
// oauth_signature) into a parameter string as defined in RFC 5894 3.4.1.3.2.
// The parameters are encoded, sorted by key, keys and values joined with "&",
// and pairs joined with "=" (e.g. foo=bar&q=gopher).
func normalizedParameterString(params map[string]string) string {
	return strings.Join(sortParameters(encodeParameters(params), "%s=%s"), "&")
}

type fixedClock struct {
	now time.Time
}

func (c *fixedClock) Now() time.Time {
	return c.now
}

func SendPost(url string, bodyJSON []byte, contentType string) {
	config := &oauth1.Config{
		ConsumerKey:    "byD0Zoa91L61qlRedu3sfCt4w",
		ConsumerSecret: "2FWir4QktLvOSgqyle4vXxFKFfDOEDGeYeeTTVUbBZoX8uf93x",
		CallbackURL:    "http://team14.hjkl.gq/",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
		Noncer: oauth1.Base64Noncer{},
	}
	//
	auther := &auther{config, &fixedClock{time.Now()}}
	token := oauth1.NewToken("1582385809914740736-sZ120gHzfV0x4JoAo9gz0s91X5V0do", "zE9cllCHaJigZaRvUXlmALCnyMXgbGbV45J6hj33gSkRg")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	err = auther.setRequestAuthHeader(req, token)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func (m *Match) PostGame() {
	image, err := m.Image()
	if err != nil {
		log.Printf("WARN: Could not generate chessboard picture: %v", err)
		return
	}
	myurl := "https://api.twitter.com/2/tweets"
	// mediaStr := []byte(`{"media_data":image}`)
	config := &oauth1.Config{
		ConsumerKey:    "byD0Zoa91L61qlRedu3sfCt4w",
		ConsumerSecret: "2FWir4QktLvOSgqyle4vXxFKFfDOEDGeYeeTTVUbBZoX8uf93x",
		CallbackURL:    "http://team14.hjkl.gq/",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
		Noncer: oauth1.Base64Noncer{},
	}

	auther := &auther{config, &fixedClock{time.Now()}}
	token := oauth1.NewToken("1582385809914740736-sZ120gHzfV0x4JoAo9gz0s91X5V0do", "zE9cllCHaJigZaRvUXlmALCnyMXgbGbV45J6hj33gSkRg")
	hc := http.Client{}

	form := url.Values{}
	// form.Add("media_data", string(image[:]))
	form.Add("media", string(image))
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	part, err := mp.CreateFormFile("media", "prova")
	io.Copy(part, bytes.NewReader(image))

	// req, err := http.NewRequest("POST", "https://upload.twitter.com/1.1/media/upload.json", body)
	req, err := http.NewRequest("POST", "https://webhook.site/2421176f-aa9d-4a7b-ac82-b25470114032/1.1/media/upload.json", body)
	err = auther.setRequestAuthHeader(req, token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body2, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body2))

	jsonStr := []byte(`{"text":"Buy cheese and bread for breakfast again."}`)
	SendPost(myurl, jsonStr, "application/json")

	// media, err := request.UploadMedia(image)
	// if err != nil {
	// 	log.Printf("WARN: Could not upload image to twitter: %v", err)
	// 	return
	// }

	// board := m.ASCII()
	// tweet, err := request.PostCustom(fakeClient, request.TweetRequest{
	// 	Text: board,
	// 	// Media: request.TweetRequestMedia{
	// 	// 	MediaIDs: []string{media.MediaID},
	// 	// },
	// })
	// if err != nil {
	// 	log.Printf("WARN: Could not post a new tweet: %v", err)
	// 	return
	// }
	// m.TweetID = tweet.ID
}

func (m *Match) Image() (buf []byte, err error) {
	dest := bytes.NewBuffer(buf)
	if err = image.SVG(dest, m.Game.Position().Board()); err != nil {
		return
	}

	return dest.Bytes(), nil
}

func (m *Match) ASCII() string {
	str := m.Game.Position().Board().Draw()
	strings.ReplaceAll(str, "-", " - ")
	return str
}

type SerializedMatch struct {
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	EndsAt   time.Time     `json:"ends_at"`
	Game     string        `json:"game"`
}

func (m Match) Serialized() SerializedMatch {
	return SerializedMatch{
		Code:     m.Code,
		Duration: m.Duration,
		EndsAt:   m.EndsAt,
		Game:     m.Game.FEN(),
	}
}

func Store() (err error) {
	buf, err := json.Marshal(match)
	err = ioutil.WriteFile(stateFile, buf, 0666)
	return
}

func Resume() (err error) {
	var m Match
	buf, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return
	}
	if err = json.Unmarshal(buf, &m); err != nil {
		return err
	}
	match = &m
	return
}

func SetMatch(m *Match) {
	match = m
	if err := Store(); err != nil {
		log.Printf("Could not store chess match state: %v", err)
	}
}

func GetMatch() *Match {
	return match
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func code(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewMatch(duration time.Duration) Match {
	m := Match{
		Code:     code(6),
		Duration: duration,
		EndsAt:   time.Now().Add(duration).UTC(),

		Game: chess.NewGame(),
	}
	go func() {
		<-m.timeout
	}()

	m.delay()
	return m
}
