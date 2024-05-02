package actions

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/tomaslobato/lt/utils"
)

func PrintAll() {
	tasks := utils.GetTasks()
	currentTime := time.Now()

	sort.Sort(utils.TaskSorter(tasks))

	fmt.Println("All tasks:")
	for _, task := range tasks {
		timestamp := time.Unix(task.Stamp, 0)
		formattedTime := timestamp.Format("2/1/06")
		if task.Done {
			color.Green("- %s | Done: %s\n", formattedTime, task.Title)
			continue
		}

		isForTomorrow := timestamp.Year() == currentTime.Year() &&
			timestamp.Month() == currentTime.Month() &&
			timestamp.Day() == currentTime.Day()+1

		if timestamp.Before(currentTime) {
			color.Red("- %s | To do: %s\n", formattedTime, task.Title)
			continue
		} else if isForTomorrow {
			color.Yellow("- %s | To do: %s\n", formattedTime, task.Title)
		} else {
			color.White("- %s | To do: %s\n", formattedTime, task.Title)
		}
	}
}
