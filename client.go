package marathon

import (
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Client is the object for containing the client for
type Client struct {
	BaseURI string       `json:"uri"`
	HTTP    *http.Client `json:"-"`
}

// BasicAuth sets the basic auth for the marathon requests if you need
type BasicAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// NewClient returns the Client object to interact with marathon
func NewClient(uri string) Client {
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar: cookieJar,
	}

	client := Client{
		BaseURI: cleanBase(uri),
		HTTP:    c,
	}
	return client
}

func (c Client) createMarathonURL(endpoint string) string {
	built := (c.BaseURI + endpoint)
	return built
}

func cleanBase(uri string) string {
	u := strings.TrimSuffix(uri, "/")
	return u
}
