package marathon

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// AppResponse is the response wrapper for getting information about an app
type AppResponse struct {
	App App `json:"app"`
}

// AppsResponse is the response wrapper for getting information about all apps
type AppsResponse struct {
	Apps []App `json:"apps"`
}

// App is sets the basic app structure
type App struct {
	ID              string                 `json:"id"`
	CMD             string                 `json:"cmd,omitempty" bson:"cmd,omitempty"`
	Args            []string               `json:"args,omitempty" bson:"args,omitempty"`
	CPUS            float64                `json:"cpus,omitempty" bson:"cpus"`
	Mem             float64                `json:"mem,omitempty" bson:"mem"`
	HealthChecks    []HealthCheck          `json:"healthChecks,omitempty" bson:"healthChecks"`
	Container       Container              `json:"container,omitempty" bson:"container"`
	Ports           []int                  `json:"ports,omitempty" bson:"ports,omitempty"`
	Instances       int                    `json:"instances,omitempty" bson:"instances,omitempty"`
	UpgradeStrategy UpgradeStrategy        `json:"upgradeStrategy,omitempty" bson:"upgradeStrategy,omitempty"`
	Version         string                 `json:"version,omitempty" bson:"version,omitempty"`
	Env             map[string]interface{} `json:"env" bson:"env,omitempty"`
	URIs            []string               `json:"uris" bson:"uris,omitempty"`
	DockerImage     string                 `json:"dockerImage" bson:"dockerImage,omitempty"`
	Tag             string                 `json:"tag" bson:"tag,omitempty"`
	TasksRunning    int                    `json:"tasksRunning,omitempty" bson:"tasksRunning,omitempty"`
	TasksStaged     int                    `json:"tasksStaged,omitempty" bson:"tasksStaged,omitempty"`
	Tasks           []Task                 `json:"tasks,omitempty" bson:"tasks,omitempty"`
	User            string                 `json:"user,omitempty" bson:"user,omitempty"`
	client          connInfo
}

// HealthCheck is structure for creating a health check
type HealthCheck struct {
	Protocol               string `json:"protocol"`
	Path                   string `json:"path"`
	GracePeriodSeconds     int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int    `json:"intervalSeconds,omitempty"`
	PortIndex              int    `json:"portIndex,omitempty"`
	TimeoutSeconds         int    `json:"timeoutSeconds,omitempty"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures,omitempty"`
}

// Container is the structure for creating a container object
type Container struct {
	Type    string           `json:"type"`
	Docker  DockerProperties `json:"docker"`
	Volumes Volume           `json:"volumes"`
}

// Volume is the structure for appending volumnes
type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode"`
}

// DockerProperties is everything to you need to send a docker container
type DockerProperties struct {
	Image        string        `json:"image"`
	Network      string        `json:"network,omitempty"`
	PortMappings []PortMapping `json:"portMappings,omitempty"`
}

// PortMapping is used to map ports
type PortMapping struct {
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	ServicePort   int    `json:"servicePort"`
	Protocol      string `json:"protocol"`
}

// UpgradeStrategy allsow your to set your deploy
type UpgradeStrategy struct {
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity"`
}

// All Returns information about all the apps
func (a App) All() (AppsResponse, error) {
	var apps AppsResponse
	res, err := a.client.Request("GET", "/v2/apps", nil)
	if err != nil {
		return apps, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apps, err
	}
	json.Unmarshal(body, &apps)
	return apps, nil
}

// One returns information about a single app
func (a App) One(id string) (AppResponse, error) {
	var app AppResponse
	url := strings.Join([]string{"/v2/apps", id}, "/")
	res, err := a.client.Request("GET", url, nil)
	if err != nil {
		return app, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return app, err
	}
	json.Unmarshal(body, &app)
	return app, nil
}

// Delete removes an Application from marathon
func (a App) Delete(id string) error {
	url := strings.Join([]string{"/v2/apps", id}, "/")
	res, err := a.client.Request("DELETE", url, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
