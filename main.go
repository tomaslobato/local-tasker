package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/tomaslobato/local-tasker/actions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Usage: [all|todo|new|today|done|cleardone]`)
		color.White("- White: To do")
		color.Red("- Red: Pending task")
		color.Green("- Green: Done")
		return
	}

	switch os.Args[1] {
	case "all":
		actions.PrintAll()
	case "todos":
		actions.PrintTodos()
	case "new":
		actions.New(os.Args)
	case "check":
		actions.Check(os.Args)
	case "cleardone":
		actions.ClearDone()
	case "delete":
		actions.Delete(os.Args)
	}
}
