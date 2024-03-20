package user

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/YuryZhehala/cli-app/internal/utils"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const filename = "users.json"

func RegisterUser() (User, error) {
	email := utils.StringPrompt("What is your email?")
	name := utils.StringPrompt("What is your name?")
	password := utils.PasswordPrompt("What is your password?")

	user, err := CreateUser(name, password, email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(name, password, email string) (User, error) {
	if email == "" {
		return User{}, fmt.Errorf("the email is empty")
	}

	if len(name) < 3 {
		return User{}, fmt.Errorf("the name length should be 3 or more characters")
	}

	if password == "" {
		return User{}, fmt.Errorf("the password is empty")
	}

	users := readUsersFromFile()

	idx := slices.IndexFunc(users, func(u User) bool { return u.Email == email })

	if idx != -1 {
		return User{}, fmt.Errorf("user with this email is already registered")
	}

	user := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	users = append(users, user)

	if err := writeUsersToFile(users); err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUser(email string) error {
	users := readUsersFromFile()
	idx := slices.IndexFunc(users, func(u User) bool { return u.Email == email })

	if idx == -1 {
		return fmt.Errorf("user with this email was not found")
	}

	users = slices.Delete(users, idx, idx+1)

	if err := writeUsersToFile(users); err != nil {
		return err
	}

	return nil
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

func Login() (User, error) {
	name := utils.StringPrompt("Enter your name?")

	if len(name) < 3 {
		return User{}, fmt.Errorf("the name length should be 3 or more characters")
	}

	found, err := findUserByName(name)
	if err != nil {
		return User{}, err
	}

	if found == nil {
		return User{}, fmt.Errorf("the user was not found")
	}

	password := utils.PasswordPrompt("Enter your password?")

	if found.Password != password {
		fmt.Println("Wrong password")
		return User{}, fmt.Errorf("wrong password")
	}

	return *found, nil
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
