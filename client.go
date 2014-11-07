package marathon

import "net/http"

// Client is the object for containing the client for
type Client struct {
	URI  string        `json:"uri"`
	HTTP *http.Request `json:"-"`
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
		URI:  uri,
		HTTP: req,
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
