package marathon

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Client is the object for containing the client for
type Client struct {
	App App `json:"app"`
}

type connInfo struct {
	BaseURI string       `json:"uri"`
	HTTP    *http.Client `json:"-"`
}

// NewClient returns the Client object to interact with marathon
func NewClient(uri string) Client {
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar: cookieJar,
	}
	conn := connInfo{
		BaseURI: cleanBase(uri),
		HTTP:    c,
	}
	client := Client{
		App: App{client: conn},
	}
	return client
}

func (c connInfo) createMarathonURL(endpoint string) string {
	return strings.Join([]string{c.BaseURI, endpoint}, "")
}

// Request sets up the http calls to marathon
func (c connInfo) Request(method, url string, body []byte) (*http.Response, error) {
	var res *http.Response
	u := c.createMarathonURL(url)
	req, err := http.NewRequest(method, u, bytes.NewBuffer(body))
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := c.HTTP
	res, err = client.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func cleanBase(uri string) string {
	u := strings.TrimSuffix(uri, "/")
	return u
}
