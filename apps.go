package marathon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
	UpgradeStrategy UpgradeStrategy        `json:"upgradeStrategy, omitempty" bson:"upgradeStrategy,omitempty"`
	Version         string                 `json:"version,omitempty" bson:"version,omitempty"`
	Env             map[string]interface{} `json:"env" bson:"env,omitempty"`
	URIs            []string               `json:"uris" bson:"uris,omitempty"`
	DockerImage     string                 `json:"dockerImage" bson:"dockerImage,omitempty"`
	Tag             string                 `json:"tag" bson:"tag,omitempty"`
}

// HealthCheck is structure for creating a health check
type HealthCheck struct {
	Protocol               string `json:"protocol"`
	Path                   string `json:"path"`
	GracePeriodSeconds     int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int    `json:"intervalSeconds, omitempty"`
	PortIndex              int    `json:"portIndex, omitempty"`
	TimeoutSeconds         int    `json:"timeoutSeconds, omitempty"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures, omitempty"`
}

// Container is the structure for creating a container object
type Container struct {
	Type   string           `json:"type"`
	Docker DockerProperties `json:"docker"`
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

// Deploy builds an app in marathon
func (c Client) Deploy(body App, initialDeploy bool) []byte {
	var req *http.Request
	var url string
	var b []byte
	method := "POST"
	b, _ = json.Marshal(body)
	if initialDeploy {
		url = c.createMarathonURL("/v2/apps")
	} else {
		s := strings.Join([]string{"/v2/apps", body.ID}, "/")
		url = c.createMarathonURL(s)
		method = "PUT"
	}

	req, _ = http.NewRequest(method, url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	client := c.HTTP

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	complete, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return complete
}
