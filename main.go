package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	Id       int
	Title    string
	DueDate  string
	Category string
	IsDone   bool
	UserId   int
}

var userStorage []User
var authenticatedUser *User
var taskStorage []Task

func main() {
	fmt.Println("Welcome To App")
	command := flag.String("command", "no command", "run command")
	flag.Parse()

	//if *command != "register-user" && *command != "exit" && *command != "login" && authenticatedUser == nil {
	//	login()
	//
	//}

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter another Command")
		scanner.Scan()
		*command = scanner.Text()
	}
}
func runCommand(command string) {
	if command != "register-user" && command != "exit" && command != "login" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {

			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		fmt.Println("You choice exit ")
		os.Exit(0)

	default:
		fmt.Println("command is not valid:", command)

	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.UserId == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
func createTask() {
	var title, category, duedate string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Task title")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("Enter Task category")
	scanner.Scan()
	category = scanner.Text()
	fmt.Println("Enter Task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		Id:       len(taskStorage) + 1,
		Title:    title,
		DueDate:  duedate,
		Category: category,
		IsDone:   false,
		UserId:   authenticatedUser.ID,
	}
	taskStorage = append(taskStorage, task)

}
func createCategory() {
	var title, color string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter category title")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("Enter color category")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("\ncateogry name is:", title, "\ncolor category is:", color)
}
func registerUser() {
	var name, email, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter user name")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Enter user email")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Enter user password")
	scanner.Scan()
	password = scanner.Text()

	//fmt.Println("\nuser id is:", id, "\nuser email is:", email, "\nuser password is:", password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)
	fmt.Printf("userStorage: %+v\n", userStorage)

}
func login() {
	fmt.Println("Log in process")
	var email, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter user email")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Enter user password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
		if authenticatedUser == nil {
			fmt.Println("The email or password is incorrect")

		}

	}
}
