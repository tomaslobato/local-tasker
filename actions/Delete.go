package actions

import (
	"fmt"
	"strings"

	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func Delete(args []string) {
	tasks := utils.GetTasks()

	if len(args) < 3 {
		fmt.Println("Please use lt delete <task title>")
		return
	}

	taskTitle := strings.Join(args[2:], " ")

	var updatedTasks []models.Task
	found := false
	for _, task := range tasks {
		if task.Title != taskTitle {
			updatedTasks = append(updatedTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Task with title '%s' not found\n", taskTitle)
		return
	}

	utils.SaveTasks(updatedTasks)
	fmt.Println("deleted:", taskTitle)
}
