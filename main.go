package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Todo App")

	command := flag.String("command", "command", "run-command")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var task, category, duedate string
	if *command == "create-task" {

		fmt.Println("Enter new task")
		scanner.Scan()
		task = scanner.Text()
		fmt.Println(task)

	} else if *command == "category" {
		fmt.Println("Enter new cate")
		scanner.Scan()
		category = scanner.Text()
		fmt.Println(category)

	} else if *command == "duedate" {
		fmt.Println("Enter new date")
		scanner.Scan()
		duedate = scanner.Text()
		fmt.Println(duedate)
	}
}
