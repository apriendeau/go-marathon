package marathon

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestTasks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// 10 total version strings
	httpmock.RegisterResponder("GET", "https://api.marathon.com/v2/tasks",
		httpmock.NewStringResponder(200, `{"tasks":[{"appId":"/bridged-webapp","healthCheckResults":[{"alive":true,"consecutiveFailures":0,"firstSuccess":"2014-10-03T22:57:02.246Z","lastFailure":null,"lastSuccess":"2014-10-03T22:57:41.643Z","taskId":"bridged-webapp.eb76c51f-4b4a-11e4-ae49-56847afe9799"}],"host":"10.141.141.10","id":"bridged-webapp.eb76c51f-4b4a-11e4-ae49-56847afe9799","ports":[31000],"servicePorts":[9000],"stagedAt":"2014-10-03T22:16:27.811Z","startedAt":"2014-10-03T22:57:41.587Z","version":"2014-10-03T22:16:23.634Z"},{"appId":"/bridged-webapp","healthCheckResults":[{"alive":true,"consecutiveFailures":0,"firstSuccess":"2014-10-03T22:57:02.246Z","lastFailure":null,"lastSuccess":"2014-10-03T22:57:41.649Z","taskId":"bridged-webapp.ef0b5d91-4b4a-11e4-ae49-56847afe9799"}],"host":"10.141.141.10","id":"bridged-webapp.ef0b5d91-4b4a-11e4-ae49-56847afe9799","ports":[31001],"servicePorts":[9000],"stagedAt":"2014-10-03T22:16:33.814Z","startedAt":"2014-10-03T22:57:41.593Z","version":"2014-10-03T22:16:23.634Z"}]}`))

	c := NewClient("https://api.marathon.com")
	tasks, err := c.Task.List()
	assert.Equal(t, err, nil, "err should be nil")
	assert.Equal(t, len(tasks.Tasks), 2, "There should be two tasks")
	assert.Equal(t, tasks.Tasks[0].AppID, "/bridged-webapp", "The first task id should be /bridged-webapp ")
	assert.Equal(t, len(tasks.Tasks[0].HealthCheckResults), 1, "There should be 1 healthcheck result")
	assert.Equal(t, tasks.Tasks[0].HealthCheckResults[0].Alive, true, "It should be alive")
}
