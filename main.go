package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Todo App")

	scanner := bufio.NewScanner(os.Stdin)
	command := flag.String("command", "command", "run-command")
	flag.Parse()

	if *command == "create-task" {
		var name, category string

		fmt.Println("Enter task name")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("Enter category name")
		scanner.Scan()
		category = scanner.Text()
		fmt.Println("name task", name, "category name", category)

	} else if *command == "category" {
		var name, color string

		fmt.Println("Enter cate name")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("Enter color category")
		scanner.Scan()
		color = scanner.Text()
		fmt.Println("cate name:", name, "cate color:", color)

	} else if *command == "duedate" {
		var date string

		fmt.Println("Enter new date")
		scanner.Scan()
		date = scanner.Text()
		fmt.Println("date is :", date)
	}
}
