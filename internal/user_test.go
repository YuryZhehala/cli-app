package user_test

import (
	"testing"

	user "github.com/YuryZhehala/cli-app/internal"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	res, err := user.CreateUser("", "123", "john@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, user.User{}, res)
}
