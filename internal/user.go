package user

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YuryZhehala/cli-app/internal/utils"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const filename = "users.json"

func Register() {
	name := utils.StringPrompt("What is your name?")
	password := utils.PasswordPrompt("What is your password?")
	email := utils.StringPrompt("What is your email?")

	users := readUsersFromFile()

	user := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	users = append(users, user)

	if err := writeUsersToFile(users); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User data has been written to users.json file successfully!")
}

func readUsersFromFile() []User {
	var users []User

	content, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error:", err)
		return users
	}

	if err := json.Unmarshal(content, &users); err != nil {
		fmt.Println("Error:", err)
		return users
	}

	return users
}

func writeUsersToFile(users []User) error {
	userData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Login() {
	name := utils.StringPrompt("Enter your name?")

	found, err := findUserByName(name)
	if err != nil {
		fmt.Println("Error:", err)
		return

	}

	if found == nil {
		fmt.Println("The user was not found")
		return
	}

	password := utils.PasswordPrompt("Enter your password?")

	if found.Password != password {
		fmt.Println("Wrong password")
		return
	}

	fmt.Printf("Welcome, %s \n", found.Name)
}

func findUserByName(name string) (*User, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Name == name {
			return &user, nil
		}
	}

	return nil, nil
}
