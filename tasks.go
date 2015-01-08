package marathon

import (
	"encoding/json"
	"io/ioutil"
)

// Task is the structure of a task in app response
type Task struct {
	AppID              string              `json:"appId" bson:"appId"`
	Host               string              `json:"host" bson:"host"`
	ID                 string              `json:"id" bson:"id"`
	Ports              []int               `json:"ports" bson:"ports"`
	StagedAt           string              `json:"stagedAt" bson:"stagedAt"`
	StartedAt          string              `json:"startedAt" bson:"startedAt"`
	Version            string              `json:"version" bson:"version"`
	HealthCheckResults []HealthCheckResult `json:"healthCheckResults" bson:"healthCheckResults"`
	client             connInfo
}

// HealthCheckResult is the structure of the status check
type HealthCheckResult struct {
	Alive               bool   `json:"alive" bson:"alive"`
	ConsecutiveFailures int    `json:"consecutiveFailures" bson:"consecutiveFailures"`
	FirstSuccess        string `json:"firstSuccess" bson:"firstSuccess"`
	LastFailure         string `json:"lastFailure,omitempty" bson:"lastFailure"`
	LastSuccess         string `json:"lastSuccess" bson:"lastSuccess"`
	TaskID              string `json:"taskId" bson:"taskId"`
}

// TasksResponse is the object returned from marathon
type TasksResponse struct {
	Tasks []Task `json:"tasks" bson:"tasks"`
}

// List lists all the current running tasks
func (t Task) List() (TasksResponse, error) {
	var tasks TasksResponse
	res, err := t.client.Request("GET", "/v2/tasks", nil)
	if err != nil {
		return tasks, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tasks, err
	}
	json.Unmarshal(body, &tasks)
	return tasks, nil
}
