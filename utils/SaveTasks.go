package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tomaslobato/local-tasker/models"
)

func SaveTasks(tasks []models.Task) {
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshaling task to JSON:", err)
		return
	}

	err = os.WriteFile("tasks.json", jsonTasks, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return
	}
}
