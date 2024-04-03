package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var (
	userStorage       []User
	taskStorage       []Task
	categoryStorage   []Category
	authenticatedUser *User
	serializationMode string
)

const (
	userStoragePath               = "user.txt"
	ManDarAvordiSerializationMode = "mandaravordi"
	JsonSerializationMode         = "json"
)

func main() {

	loadUserStorageFromFile()

	fmt.Println("Welcome To App")
	serializMode := flag.String("serialize-mode", ManDarAvordiSerializationMode, "save mandaravordi serialize data user ")

	command := flag.String("command", "no command", "run command")
	flag.Parse()

	switch *serializMode {
	case ManDarAvordiSerializationMode:
		serializationMode = ManDarAvordiSerializationMode
	default:
		serializationMode = JsonSerializationMode

	}

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter another Command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func loadUserStorageFromFile() {

	//load user storage from file

	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Println("cant open the file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)
	{
		if oErr != nil {
			fmt.Println("cant read the file", oErr)
		}
		//fmt.Println(data)
	}
	var datastr = string(data)
	userSlice := strings.Split(datastr, "\n")
	for _, u := range userSlice {
		if u == "" {
			continue
		}

		//fmt.Println("line of file", index, "user", u)
		var user = User{}
		userFields := strings.Split(u, ",")
		for _, field := range userFields {
			values := strings.Split(field, ": ")
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					fmt.Println("strconv err", err)

					return
				}
				user.ID = id
			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue
			}
		}
		fmt.Printf("user %+v\n", user)
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
		if listcat.UserID == authenticatedUser.ID {
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
		fmt.Printf("category-id is not int , %v\n", err)

		return
	}
	ifexist := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			ifexist = true

			break
		}
	}
	if !ifexist {
		fmt.Println("category id not valid")

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
	//fmt.Printf("userStorage: %+v\n", userStorage)
	writeUserToFile(user)
}

func writeUserToFile(user User) {

	//save user data in user.txt

	file, err := os.OpenFile(userStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file cant create or append", err)

		return
	}
	defer file.Close()

	var data []byte
	if serializationMode == ManDarAvordiSerializationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s", user.ID, user.Name,
			user.Email, user.Password))
	} else if serializationMode == JsonSerializationMode {
		//json
		var jErr error
		data, jErr = json.Marshal(user)
		if jErr != nil {
			fmt.Println("cant marshal user to struct", jErr)

			return
		}
	} else {
		fmt.Println("invalid serialization mode")

		return
	}

	var wErr error
	numberOfWrittenByte, wErr := file.Write(data)
	if wErr != nil {
		fmt.Printf("cant write in file ,  %v\n", wErr)
	}
	fmt.Println("number of written byte", numberOfWrittenByte)

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
