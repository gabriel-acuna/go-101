package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"./Models"
)

var reader *bufio.Reader
var users map[string]Models.User

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createUser() {
	clearConsole()
	var user Models.User

	fmt.Println("Enter user's first name: ")
	user.SetFristname(readLine())

	fmt.Println("Enter user's last name: ")
	user.SetLastname(readLine())

	fmt.Println("Enter user's email: ")
	user.SetEmail(readLine())

	fmt.Println("Enter user's age: ")
	if age, err := strconv.Atoi(readLine()); err == nil {
		user.SetAge(age)
	}

	fmt.Print(">>> User created successfully!\n\n")
	users[user.GetEmail()] = user
}
func listUser() {
	clearConsole()

	if len(users) > 0 {
		fmt.Println("********** Users ************ ")
		for _, user := range users {
			fmt.Println("-----------------------")
			fmt.Printf(" Full name: %s %s \n Email: %s \n Age: %d\n", user.GetFristname(), user.GetFristname(), user.GetEmail(), user.GetAge())
			fmt.Println("-----------------------")
		}
	} else {
		fmt.Print(">>> There aren't users registered yet!\n\n")
	}

}

func updateUser() {
	fmt.Print(">>> User updated successfully!\n\n")
}

func deleteUser() {
	clearConsole()
	fmt.Println("Enter user's email:")
	email := readLine()

	if _, ok := users[email]; ok {
		delete(users, email)
		fmt.Print(">>> User deleted successfully!\n\n")
	} else {
		fmt.Print(">>> User not found!\n\n")
	}

}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("Couldn't be read the value")
	} else {
		return strings.TrimSuffix(option, "\n")
	}

}

func main() {

	reader = bufio.NewReader(os.Stdin)

	var option string

	users = make(map[string]Models.User)

	for {
		fmt.Println("A) Create")
		fmt.Println("B) List")
		fmt.Println("C) Update")
		fmt.Println("D) Delete")

		fmt.Println("Choose an option:")
		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "create":
			createUser()
		case "b", "list":
			listUser()
		case "c", "update":
			updateUser()
		case "d", "delete":
			deleteUser()
		default:
			fmt.Println("Option not valid!")

		}

	}
	fmt.Println("Bye! ;)")

}
