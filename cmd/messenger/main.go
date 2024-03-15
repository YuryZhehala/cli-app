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
			user, err := user.RegisterUser()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("User data has been registered %s", user.Name)
		case 2:
			user, err := user.Login()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Welcome, %s \n", user.Name)
		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
