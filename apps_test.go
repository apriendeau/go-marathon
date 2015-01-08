package marathon

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchingApps(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.marathon.com/v2/apps",
		httpmock.NewStringResponder(200, `{"apps":[{"id":"/utopia","cmd":null,"args":["production"],"user":null,"env":{"TEST":"AWESOME"},"instances":2,"cpus":2,"mem":1024,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10003],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[{"containerPath":"/var/run/docker.sock","hostPath":"/var/run/docker.sock","mode":"RW"}],"docker":{"image":"docker-prd.itriagehealth.com/utopia:0.1.6","network":null,"portMappings":null}},"healthChecks":[{"path":"/heartbeat","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":10,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":10}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-07T18:59:38.310Z"},{"id":"/site","cmd":"make run","args":null,"user":null,"env":{},"instances":1,"cpus":1,"mem":512,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10000],"requirePorts":false,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[],"docker":{"image":"docker-prd.itriagehealth.com/splash:0.0.11","network":null,"portMappings":null}},"healthChecks":[{"path":"/","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":5,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":3}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-06T06:10:11.441Z","deployments":[],"tasksStaged":0,"tasksRunning":1,"backoffSeconds":1}]}`))

	c := NewClient("https://api.marathon.com")
	a, err := c.App.All()
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, len(a.Apps), 2, "There should be two apps")
	assert.Equal(t, a.Apps[0].ID, "/utopia", "The first app should be utopia")
	assert.Equal(t, a.Apps[1].ID, "/site", "The first app should be site")
}

func TestFetchingApp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.marathon.com/v2/apps/utopia",
		httpmock.NewStringResponder(200, `{"app":{"id":"/utopia","cmd":null,"args":["production"],"user":null,"env":{"TEST":"AWESOME"},"instances":2,"cpus":2,"mem":1024,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10003],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[{"containerPath":"/var/run/docker.sock","hostPath":"/var/run/docker.sock","mode":"RW"}],"docker":{"image":"docker-prd.itriagehealth.com/utopia:0.1.6","network":null,"portMappings":null}},"healthChecks":[{"path":"/heartbeat","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":10,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":10}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-07T18:59:38.310Z"}}`))

	c := NewClient("https://api.marathon.com")
	a, err := c.App.One("utopia")
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, a.App.ID, "/utopia", "The first app should be utopia")
	assert.Equal(t, len(a.App.Args), 1, "the length of args should be 1")
	assert.Equal(t, a.App.Instances, 2, "there should be 2 instances of the app")
}

func TestDeleteApp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.marathon.com/v2/apps/utopia",
		httpmock.NewStringResponder(204, ``))

	c := NewClient("https://api.marathon.com")
	err := c.App.Delete("utopia")
	assert.Equal(t, err, nil, "err should be nil")
}
