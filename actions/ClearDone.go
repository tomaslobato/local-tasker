package actions

import (
	"github.com/tomaslobato/lt/models"
	"github.com/tomaslobato/lt/utils"
)

func ClearDone() {
	tasks := utils.GetTasks()

	var activeTasks []models.Task
	for _, task := range tasks {
		if !task.Done {
			activeTasks = append(activeTasks, task)
		}
	}

	utils.SaveTasks(activeTasks)
}
