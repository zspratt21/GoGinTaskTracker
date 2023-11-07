package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskJSONMarshaling(t *testing.T) {
	task := Task{
		Name:        "Task 1",
		IsCompleted: false,
	}

	jsonBytes, err := json.Marshal(task)
	assert.NoError(t, err, "marshaling to JSON should not produce an error")

	var decodedTask Task
	err = json.Unmarshal(jsonBytes, &decodedTask)
	assert.NoError(t, err, "unmarshaling JSON should not produce an error")

	assert.Equal(t, task, decodedTask, "the decoded task should match the original")
}

func TestTaskJSONUnmarshaling(t *testing.T) {
	jsonString := `{"name":"Task 2","isCompleted": true}`

	var task Task
	err := json.Unmarshal([]byte(jsonString), &task)
	assert.NoError(t, err, "unmarshaling JSON should not produce an error")

	assert.Equal(t, "Task 2", task.Name, "the name should be set correctly")
	assert.Equal(t, true, task.IsCompleted, "the hex value should be set correctly")
}
