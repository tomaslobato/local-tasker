package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func New(args []string) {
	// Get tasks
	actualTasks := utils.GetTasks()

	// Get lastTaskId
	var lastTaskId uint
	if len(actualTasks) > 0 {
		lastTaskId = actualTasks[len(actualTasks)-1].Id
	}

	if len(args) < 3 {
		fmt.Println("Please use lt new <task title> --date <01/02/06>")
		return
	}

	var TaskTitle strings.Builder
	var dateFound bool
	var dateString string

	// Looping over the arguments
	for i, arg := range args[2:] {
		if arg == "--date" {
			dateFound = true
			continue // Set dateFound to true so dateString gets the next arg
		}
		if dateFound {
			dateString = arg
			break // Date is the last argument
		}
		if i > 0 {
			TaskTitle.WriteString(" ")
		}
		TaskTitle.WriteString(arg)
	}

	if !dateFound {
		fmt.Println("--date flag not found")
		return
	}

	//validations
	date, err := time.Parse("1/2/06", dateString)
	if err != nil {
		fmt.Printf(`Date input "%s" is invalid, correct format: month/day/year (01/2/06)`, dateString)
		return
	}
	timestamp := date.Unix()

	newTaskTitle := TaskTitle.String()
	for _, task := range actualTasks {
		if task.Title == newTaskTitle {
			fmt.Printf("Error: A task with the title \"%s\" already exists\n", newTaskTitle)
			return
		}
	}

	// Set newTask
	newTask := models.Task{
		Id:    lastTaskId + 1,
		Title: newTaskTitle,
		Stamp: timestamp,
		Done:  false,
	}

	newTasks := append(actualTasks, newTask)
	utils.SaveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving new task: ", err)
		return
	}

	fmt.Printf(`new task "%s, Date: %d" added successfully`, newTaskTitle, timestamp)
}
