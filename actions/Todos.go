package actions

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/tomaslobato/lt/utils"
)

func PrintTodos() {
	tasks := utils.GetTasks()

	sort.Sort(utils.TaskSorter(tasks))

	fmt.Println("To-do's:")
	for _, task := range tasks {
		if task.Done == true {
			continue
		}

		timestamp := time.Unix(task.Stamp, 0)
		formattedTime := timestamp.Format("2/1/06")

		currentTime := time.Now()
		isForTomorrow := timestamp.Year() == currentTime.Year() &&
			timestamp.Month() == currentTime.Month() &&
			timestamp.Day() == currentTime.Day()+1

		if timestamp.Before(currentTime) {
			color.Red("- %s | %s \n", formattedTime, task.Title)
		} else if isForTomorrow {
			color.Yellow("- %s | To do: %s\n", formattedTime, task.Title)
		} else {
			color.White("- %s | To do: %s\n", formattedTime, task.Title)
		}
	}
}
