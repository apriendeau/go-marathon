package marathon

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Create builds an app in marathon
func (a App) Create(body App) (App, error) {
	deploy, err := a.deploy(body, true, "")
	return deploy, err
}

// Update builds an app in marathon
func (a App) Update(id string, body App) (App, error) {
	deploy, err := a.deploy(body, false, id)
	return deploy, err
}

// Deploy builds an app in marathon
func (a App) deploy(body App, initialDeploy bool, id string) (App, error) {
	var app App
	method := "POST"
	b, _ := json.Marshal(body)
	url := "/v2/apps"
	if !initialDeploy {
		url = strings.Join([]string{url, id}, "/")
		method = "PUT"
	}
	res, err := a.client.Request(method, url, b)
	if err != nil {
		return app, err
	}

	defer res.Body.Close()
	parsedBody, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return app, err
	}
	json.Unmarshal(parsedBody, &app)
	return app, nil
}
