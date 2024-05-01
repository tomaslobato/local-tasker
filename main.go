package main

import (
	"fmt"
	"os"

	"github.com/tomaslobato/local-tasker/actions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Usage: [all|todo|new|today|done|delete|edit]`)
		return
	}

	switch os.Args[1] {
	case "all":
		actions.PrintAll()
	case "todo":
		actions.PrintTodos()
	case "new":
		actions.New(os.Args)
	}

}
