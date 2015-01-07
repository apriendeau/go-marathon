package marathon

import (
	"os"
	"strings"
)

// Constructor is a type that the minimum fields required to construct an app
type Constructor struct {
	Mem         float64                `json:"mem"`
	CPUs        float64                `json:"cpus"`
	Env         map[string]interface{} `json:"env"`
	Args        []string               `json:"args"`
	Cmd         string                 `json:"cmd"`
	Project     string                 `json:"project"`
	Repo        string                 `json:"repo"`
	Name        string                 `json:"name"`
	Branch      string                 `json:"branch"`
	Instances   int                    `json:"instances"`
	DockerImage string                 `json:"dockerImage"`
	Tag         string                 `json:"tag"`
}

// Construct is an app that creates the docker properties
func (c Constructor) Construct(kind string) App {
	var a = App{
		ID:           c.Name,
		Mem:          c.Mem,
		CPUS:         c.CPUs,
		Env:          c.Env,
		Instances:    c.Instances,
		DockerImage:  c.DockerImage,
		Tag:          c.Tag,
		HealthChecks: []HealthCheck{},
		Ports:        []int{},
		Args:         c.Args,
		URIs:         []string{"/root/.dockercfg"},
		CMD:          c.Cmd,
	}
	a.Container = Container{
		Type:   "DOCKER",
		Docker: CreateDockerProperties(c),
	}
	switch {
	case kind == "task":
	// skip healthcheck
	default:
		a.HealthChecks = CreateHealthCheck()
	}

	return a
}

// CreateHealthCheck returns an array of required healthchecks
func CreateHealthCheck() []HealthCheck {
	var healthCheck = HealthCheck{
		Protocol:               "HTTP",
		Path:                   "/heartbeat",
		GracePeriodSeconds:     5,
		IntervalSeconds:        10,
		TimeoutSeconds:         10,
		MaxConsecutiveFailures: 3,
		PortIndex:              0,
	}
	healthChecks := []HealthCheck{healthCheck}
	return healthChecks
}

// CreateDockerProperties creates marathon docker properties
func CreateDockerProperties(c Constructor) DockerProperties {
	image := createDevRegistryTag(c.DockerImage, c.Tag)
	dockerProps := DockerProperties{
		Image: image,
	}
	return dockerProps
}

// CreateArchiveURL creates the download link for stash repo
func CreateArchiveURL(project, repo, at string) string {
	s := os.Getenv("STASH_URL")
	strArray := []string{
		s,
		"/plugins/servlet/archive/projects/",
		project,
		"/repos/",
		repo,
		"?format=tar.gz",
	}
	url := strings.Join(strArray, "")
	url = strings.Join([]string{url, "&at=", at}, "")
	return url
}

func createDevRegistryTag(dockerImage, tag string) string {
	s := strings.Join([]string{dockerImage, ":", tag}, "")
	return s
}
