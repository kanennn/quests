package main

import (
	"fmt"
	"os"

	"github.com/kanennn/quests/internal/leo"
)

func main() {
	cli()
}

func cli() {
	if len(os.Args) < 2 {
		println("A subcommand is required") 
	} else {
		subcommand := os.Args[1]
		switch subcommand {
		// case "new":
		// 	new()
		case "ls":
			list()
		default:
			fmt.Printf("%s: unknown command\n", subcommand)
		}
	}
}

// func new() {
// 	//probably should check what the error is ,
// 	fmt.Println("Enter the command name")
//  	name, _ := io.ReadAll(os.Stdin)
// 	task1 := task{
// 		name: string(name),
// 		description: "wwwww",
// 		priority: high,
// 		precedence: 0,
// 	}
// 	output(task1)
// }

func list() {
	//this proves we can read data
	files := leo.ReadFiles()
	tasks := leo.ParseFiles(files)
	// for _, task := range tasks {
	// 	fmt.Println("that's the name")
	// 	fmt.Println(string(task.Name))
	// 	fmt.Println(string(task.Description))
	// 	fmt.Println(string(task.Context))
	// 	fmt.Println(string(task.POI))
	// }
	leo.CreateTasks(tasks)
}