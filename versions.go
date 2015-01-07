package marathon

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// GetVersions return all the versions of an app
func (a App) GetVersions(name string) (VersionsResponse, error) {
	var response VersionsResponse
	url := strings.Join([]string{"/v2/apps", name, "versions"}, "/")

	res, err := a.client.Request("GET", url, nil)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	json.Unmarshal(body, &response)
	return response, nil
}

// GetVersion gets info about a single version
func (a App) GetVersion(name, v string) (App, error) {
	var response App
	url := strings.Join([]string{"/v2/apps", name, "versions", v}, "/")

	res, err := a.client.Request("GET", url, nil)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	json.Unmarshal(body, &response)
	return response, nil
}
