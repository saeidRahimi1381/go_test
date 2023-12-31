package function

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var UserStorage []User

func Runcommand(command string) {
	switch command {
	case "creat-task":
		CreatTask()
	case "creat-catgory":
		createCategory()
	case "register-user":
		RegisterUser()
	case "login":
		Login()
	case "Exit":
		os.Exit(0)
	default:
		fmt.Println("command us not valid", command)
	}
}

func CreatTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, duedate, catgory string
	fmt.Println("plase enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("plase enter the task catgory")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("plase enter the task due date")
	scanner.Scan()
	catgory = scanner.Text()

	fmt.Println("task:", name, duedate, catgory)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("plase enter the catgory title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("plase enter the catgory color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("catgory:", title, color)
}

func RegisterUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	id = email
	fmt.Println("register-user", id, name, email, password)

	user := User{
		ID:       len(UserStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	UserStorage = append(UserStorage, user)
}

func Login() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("plase enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("plase enter the password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", id, email, password)
}
