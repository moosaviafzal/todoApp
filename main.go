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

var userStorage []User

func main() {
	fmt.Println("Welcome To App")
	command := flag.String("command", "no command", "run command")
	flag.Parse()

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the email")
	scn.Scan()
	email := scn.Text()
	fmt.Println("please enter the password")
	scn.Scan()
	password := scn.Text()

	notfound := true
	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			notfound = false
			fmt.Println("Youre logged in")
		} else {
			fmt.Println("Your email or pass incorrect")
		}
	}
	if notfound {
		fmt.Println("The email or password is incorrect")
		return
	}

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter another Command")
		scanner.Scan()
		*command = scanner.Text()
	}
}
func runCommand(command string) {

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)

	default:
		fmt.Println("command is not valid:", command)

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
	fmt.Println("\ntask name is:", title, "\ntask category is:", category, "\ntask Due is:", duedate)
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
	var email, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter user email")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Enter user password")
	scanner.Scan()
	password = scanner.Text()
	fmt.Println("\nuser email is:", email, "\nuser password is:", password)
}
