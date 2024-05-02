package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/tomaslobato/lt/models"
	"github.com/tomaslobato/lt/utils"
)

func New(args []string) {
	actualTasks := utils.GetTasks()

	// Get lastTaskId
	var lastTaskId uint
	if len(actualTasks) > 0 {
		lastTaskId = actualTasks[len(actualTasks)-1].Id
	}

	if len(args) < 3 {
		fmt.Println(`Please use lt new <task title> --date <01/02/06>. 
	You can also use --date today/tomorrow/nextweek`)
		return
	}

	var TaskTitle strings.Builder
	var dateFound bool
	var dateString string

	// Looping over the arguments
	for i, arg := range args[2:] {
		if arg == "--date" {
			dateFound = true
			continue
		}
		if dateFound {
			dateString = arg
			break
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

	var date time.Time
	var err error

	//validations
	if dateString == "today" {
		today := time.Now()
		date = time.Date(today.Year(), today.Month(), today.Day()+1, 0, 0, 0, 0, time.Local)
	} else if dateString == "tomorrow" {
		today := time.Now()
		date = time.Date(today.Year(), today.Month(), today.Day()+2, 0, 0, 0, 0, time.Local)
	} else {
		date, err = time.Parse("1/2/06", dateString)
		if err != nil {
			fmt.Printf(`Date input "%s" is invalid, correct format: month/day/year (01/2/06).
	You can also use --date today/tomorrow/nextweek`, dateString)
			return
		}
	}

	timestamp := date.Unix()

	newTaskTitle := TaskTitle.String()
	for _, task := range actualTasks {
		if task.Title == newTaskTitle {
			fmt.Printf("Error: A task with the title \"%s\" already exists\n", newTaskTitle)
			return
		}
	}

	newTask := models.Task{
		Id:    lastTaskId + 1,
		Title: newTaskTitle,
		Stamp: timestamp,
		Done:  false,
	}

	newTasks := append(actualTasks, newTask)
	utils.SaveTasks(newTasks)

	fmt.Printf(`new task "%s" for %s added successfully`, newTaskTitle, dateString)
}
