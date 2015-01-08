package marathon

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMarathonVersions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// 10 total version strings
	httpmock.RegisterResponder("GET", "https://api.marathon.com/v2/apps/utopia/versions",
		httpmock.NewStringResponder(200, `{"versions":["2015-01-07T18:59:38.310Z","2015-01-07T18:09:40.030Z","2015-01-07T05:27:06.642Z","2015-01-07T04:45:03.861Z","2015-01-06T23:05:25.821Z","2015-01-06T23:02:23.488Z","2015-01-06T22:51:35.799Z","2015-01-05T23:16:50.419Z","2014-12-31T21:58:05.656Z","2014-12-31T21:48:59.181Z"]}`))

	c := NewClient("https://api.marathon.com")
	v, err := c.App.GetVersions("utopia")
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, len(v.Versions), 10, "there should be 10 versions in the array")
}

func TestVersion(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.marathon.com/v2/apps/utopia/versions/2015-01-07T18:59:38.310Z",
		httpmock.NewStringResponder(200, `{"id":"/utopia","cmd":null,"args":["production"],"user":null,"env":{"TEST":"AWESOME"},"instances":2,"cpus":2,"mem":1024,"disk":0,"executor":"","constraints":[],"uris":["/root/.dockercfg"],"storeUrls":[],"ports":[10003],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"container":{"type":"DOCKER","volumes":[{"containerPath":"/var/run/docker.sock","hostPath":"/var/run/docker.sock","mode":"RW"}],"docker":{"image":"docker-prd.itriagehealth.com/utopia:0.1.6","network":null,"portMappings":null}},"healthChecks":[{"path":"/heartbeat","protocol":"HTTP","portIndex":0,"command":null,"gracePeriodSeconds":10,"intervalSeconds":10,"timeoutSeconds":10,"maxConsecutiveFailures":10}],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1},"version":"2015-01-07T18:59:38.310Z"}`))

	c := NewClient("https://api.marathon.com")
	v, err := c.App.GetVersion("utopia", "2015-01-07T18:59:38.310Z")

	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, v.ID, "/utopia", "id should be equal to /utopia")
}
