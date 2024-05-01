package actions

import (
	"fmt"
	"time"

	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func PrintTodos() {
	var tasks []models.Task
	tasks, err := utils.GetTasks()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("To-do's:")
	for _, task := range tasks {
		if !task.Done {
			continue
		}

		timestamp := time.Unix(task.Stamp, 0)
		formattedTime := timestamp.Format("2/1/06 15:04")

		fmt.Printf("- %s | %s\n", task.Title, formattedTime)
	}

}
