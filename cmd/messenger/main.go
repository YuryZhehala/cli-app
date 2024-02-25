package main

import (
	"fmt"

	user "github.com/YuryZhehala/cli-app/internal"
)

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Register a new user")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			user.Register()
		case 2:
			user.Login()
		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
