package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tomaslobato/local-tasker/models"
)

func GetTasks() ([]models.Task, error) {
	file, err := ioutil.ReadFile("tasks.json")

	var tasks []models.Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
