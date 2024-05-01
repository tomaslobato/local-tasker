package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tomaslobato/local-tasker/models"
	"github.com/tomaslobato/local-tasker/utils"
)

func New(args []string) {
	//get tasks
	var actualTasks []models.Task
	actualTasks, err := utils.GetTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	//get lastTaskId
	var lastTaskId uint
	if len(actualTasks) > 0 {
		lastTaskId = actualTasks[len(actualTasks)-1].Id
	}

	if len(args) < 3 {
		fmt.Println("Please use lt new <task title> --date <01/02/06>")
	}

	var TaskTitle strings.Builder
	var dateFound bool
	var dateString string

	//looping over the arguments
	for i, arg := range args[2:] {
		if arg == "--date" {
			dateFound = true
			continue //set dateFound to true so dateString gets the next arg
		}

		if dateFound {
			dateString = arg
			break //date is the last argument
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

	date, err := time.Parse("01/02/06", dateString)
	if err != nil {
		fmt.Printf(`Date input "%s" is invalid, correct format: day/month/year (01/02/06)`, dateString)
		return
	}

	timestamp := date.Unix()

	//set newTask
	newTask := models.Task{
		Id:    lastTaskId + 1,
		Title: TaskTitle.String(),
		Stamp: timestamp,
		Done:  false,
	}
	newTasks := append(actualTasks, newTask)
	jsonTasks, err := json.Marshal(newTasks)
	if err != nil {
		fmt.Println("Error marshaling task to JSON:", err)
		return
	}

	err = os.WriteFile("tasks.json", jsonTasks, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return
	}

	fmt.Printf(`new task "%s, Date: %d" added successfully`, TaskTitle.String(), timestamp)
}
