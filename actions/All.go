package actions

import (
	"time"

	"github.com/fatih/color"
	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func PrintAll() {
	var tasks []models.Task
	tasks = utils.GetTasks()

	for _, task := range tasks {
		timestamp := time.Unix(task.Stamp, 0) // Convert Unix timestamp to time.Time value
		formattedTime := timestamp.Format("2/1/06")

		if task.Done {
			color.Green("- %s | Done: %s\n", formattedTime, task.Title)
			continue
		}

		currentTime := time.Now()
		if timestamp.Before(currentTime) {
			color.Red("- %s | To do: %s\n", formattedTime, task.Title)
			continue
		}
		color.White("- %s | To do: %s\n", formattedTime, task.Title)
	}
}
