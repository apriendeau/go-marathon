package marathon

import (
	"net/http"
	"strings"
)

// Client is the object for containing the client for
type Client struct {
	BaseURI string        `json:"uri"`
	HTTP    *http.Request `json:"-"`
}

// BasicAuth sets the basic auth for the marathon requests if you need
type BasicAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// NewClient returns the Client object to interact with marathon
func NewClient(uri string) Client {
	req, _ := http.NewRequest("GET", uri, nil)
	var client = Client{
		BaseURI: cleanBase(uri),
		HTTP:    req,
	}
	return client
}

// NewClientWithBasicAuth returns the Client object to interact with marathon
// but also sets basic Auth
func NewClientWithBasicAuth(uri string, auth BasicAuth) Client {
	client := NewClient(uri)
	client.HTTP.SetBasicAuth(auth.User, auth.Password)
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
