package actions

import (
	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
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
