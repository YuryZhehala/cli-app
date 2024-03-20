package user_test

import (
	"testing"

	user "github.com/YuryZhehala/cli-app/internal"

	"github.com/stretchr/testify/assert"
)

func Test_RegisterUser(t *testing.T) {
	res, err := user.RegisterUser()

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}

func Test_CreateUser(t *testing.T) {
	expected := user.User{
		Name:     "John",
		Password: "123",
		Email:    "john@gmail.com",
	}

	res, err := user.CreateUser("John", "123", "john@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func Test_CreateUserWithEmptyName(t *testing.T) {
	res, err := user.CreateUser("", "123", "john@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}

func Test_CreateUserWithEmptyEmail(t *testing.T) {
	res, err := user.CreateUser("John", "123", "")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}

func Test_CreateUserWithEmptyPassword(t *testing.T) {
	res, err := user.CreateUser("John", "", "john@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}

func Test_CreateUserWithTwoCharsInName(t *testing.T) {
	res, err := user.CreateUser("Jo", "123", "john@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}

func Test_CreateUserWithTheSameEmail(t *testing.T) {
	expected := user.User{
		Name:     "John",
		Password: "123",
		Email:    "johnTest@gmail.com",
	}

	_ = user.DeleteUser("johnTest@gmail.com")

	res, err := user.CreateUser("John", "123", "johnTest@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, expected, res)

	res, err = user.CreateUser("John", "123", "johnTest@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}
