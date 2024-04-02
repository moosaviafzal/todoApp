package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	Id         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var userStorage []User
var authenticatedUser *User
var taskStorage []Task
var categoryStorage []Category

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
	case "list-cat":
		listCategory()
	case "login":
		login()
	case "exit":
		fmt.Println("You choice exit ")
		os.Exit(0)

	default:
		fmt.Println("command is not valid:", command)

	}

}

func listCategory() {
	for _, listcat := range categoryStorage {
		if listcat.ID == authenticatedUser.ID {
			fmt.Printf("%+v\n", listcat)
		}
	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Printf("%+v\n", task)
		}
	}
}
func createTask() {
	var title, category, duedate string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Task title")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("Enter Task category ID")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category-id is not valid integer , %v\n", err)

		return
	}

	isfound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isfound = true

			break
		}
	}
	if !isfound {
		fmt.Printf("category-id is not valid\n")

		return
	}

	fmt.Println("Enter Task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	t := Task{
		Id:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}
	taskStorage = append(taskStorage, t)

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

	c := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}
	categoryStorage = append(categoryStorage, c)

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
	}
	if authenticatedUser == nil {
		fmt.Println("The email or password is incorrect , try again 0r register-user")

	}
}
