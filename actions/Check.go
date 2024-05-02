package actions

import (
	"strings"

	"github.com/tomaslobato/lt/utils"
)

func Check(args []string) {
	tasks := utils.GetTasks()

	argstring := strings.Join(args[2:], " ")
	for i, task := range tasks {
		if task.Title == argstring {
			tasks[i].Done = !task.Done
			utils.SaveTasks(tasks)
			break
		}
	}
}
