package marathon

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.marathon.com/v2/apps",
		httpmock.NewStringResponder(200, `{"id":"/utopia","cmd":null,"args":["production"],"user":null,"env":{},"instances":2,"cpus":2,"mem":1024,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10003],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[{"containerPath":"/var/run/docker.sock","hostPath":"/var/run/docker.sock","mode":"RW"}],"docker":{"image":"docker-prd.itriagehealth.com/utopia:0.1.6","network":null,"portMappings":null}},"healthChecks":[{"path":"/heartbeat","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":10,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":10}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-07T18:59:38.310Z"}`))
	health := createHealthCheck()

	docker := DockerProperties{
		Image: "docker-prd.itriagehealth.com/utopia:0.1.6",
	}
	container := Container{
		Type:   "DOCKER",
		Docker: docker,
	}
	testApp := App{
		ID:           "/utopia",
		Args:         []string{"production"},
		CMD:          "",
		Instances:    2,
		CPUS:         2,
		Mem:          1024,
		HealthChecks: health,
		Container:    container,
	}
	c := NewClient("https://api.marathon.com")
	app, err := c.App.Create(testApp)
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, app.Instances, 2, "There should be 2 instances")
	assert.Equal(t, app.CPUS, 2, "There should be 2 cpus")
	assert.Equal(t, app.Mem, 1024, "There should be 1024mb")
}

func TestUpdateApp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.marathon.com/v2/apps/utopia",
		httpmock.NewStringResponder(200, `{"id":"/utopia","cmd":null,"args":["production"],"user":null,"env":{},"instances":1,"cpus":2,"mem":1024,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10003],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[{"containerPath":"/var/run/docker.sock","hostPath":"/var/run/docker.sock","mode":"RW"}],"docker":{"image":"docker-prd.itriagehealth.com/utopia:0.1.6","network":null,"portMappings":null}},"healthChecks":[{"path":"/heartbeat","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":10,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":10}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-07T18:59:38.310Z"}`))

	testApp := App{
		Instances: 1,
	}
	c := NewClient("https://api.marathon.com")
	app, err := c.App.Update("utopia", testApp)
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, app.Instances, 1, "There should be 2 instances")
	assert.Equal(t, app.Mem, 1024, "There should be 1024mb of ram")
}

func createHealthCheck() []HealthCheck {
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
