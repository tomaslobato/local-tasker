package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tomaslobato/lt/models"
)

func GetTasks() []models.Task {
	file, err := os.ReadFile("tasks.json")

	var tasks []models.Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("error getting tasks:", err)
		return nil
	}

	return tasks
}
