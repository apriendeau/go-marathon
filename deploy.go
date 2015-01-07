package marathon

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Deploy builds an app in marathon
func (a App) deploy(body App, initialDeploy bool) ([]byte, error) {
	method := "POST"
	b, _ := json.Marshal(body)
	url := "/v2/apps"
	if !initialDeploy {
		url = strings.Join([]string{url, body.ID}, "/")
		method = "PUT"
	}
	res, err := a.client.Request(method, url, b)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	complete, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return complete, nil
}
