package actions

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/tomaslobato/local-tasker/utils"
)

func PrintTodos() {
	tasks := utils.GetTasks()

	fmt.Println("To-do's:")
	for _, task := range tasks {
		if task.Done == true {
			continue
		}

		timestamp := time.Unix(task.Stamp, 0)
		formattedTime := timestamp.Format("2/1/06")

		currentTime := time.Now()
		if timestamp.Before(currentTime) {
			color.Red("- %s | %s \n", formattedTime, task.Title)
		} else {
			color.White("- %s | %s \n", formattedTime, task.Title)
		}
	}
}
