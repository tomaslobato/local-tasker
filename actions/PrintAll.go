package actions

import (
	"fmt"
	"time"

	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func PrintAll() {
	var tasks []models.Task
	tasks, err := utils.GetTasks()
	if err != nil {
		fmt.Println(err)
	}

	for _, task := range tasks {
		timestamp := time.Unix(task.Stamp, 0)             // Convert Unix timestamp to time.Time value
		formattedTime := timestamp.Format("2/1/06 15:04") // Format the time as "YYYY-MM-DD HH:MM:SS"

		done := "To Do:"
		if task.Done {
			done = "Done:"
		}

		fmt.Printf("%s %s | %s\n", done, task.Title, formattedTime)
	}
}
