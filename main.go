package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/tomaslobato/lt/actions"
)

func wrongUseMsg() {
	fmt.Println(`Usage: [all|todos|new|check|delete|cleardone]`)
	color.White("- White: To do")
	color.Yellow("- Yellow: To do today")
	color.Green("- Green: Done")
	color.Red("- Red: Pending task")
	fmt.Println("Note: Today's task have the date of tomorrow because they are for tomorrow and you gotta do them today.")
}

func main() {
	if len(os.Args) < 2 {
		wrongUseMsg()
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
	default:
		wrongUseMsg()
	}
}
