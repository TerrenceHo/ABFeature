package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	assert := assert.New(t)

	user := &User{
		ID:          "",
		Name:        "",
		Description: "",
	}

	errs := user.Validate()
	assert.Equal(len(errs), 2, "Number of errors should be 2")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")
	assert.Equal(errs[1].Error(), "Please provide a user name.")

	user.Name = "Name"
	errs = user.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")

	user.Name = ""
	user.ID = "ID"
	errs = user.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "Please provide a user name.")
}

func TestUserUpdateFields(t *testing.T) {
	assert := assert.New(t)
	old_user := &User{
		ID:          "ID",
		Name:        "User Name",
		Description: "User Description",
	}

	new_user := &User{
		ID:          "ID",
		Name:        "New User Name",
		Description: "",
	}

	errs := old_user.UpdateFields(new_user)
	assert.Equal(len(errs), 0, "Number of errors should be zero.")
	assert.Equal(old_user.ID, "ID", "ID should be 'ID'")
	assert.Equal(old_user.Name, "New User Name", "user name should be updated.")
	assert.Equal(old_user.Description, "User Description", "user Description should be unchanged.")
}
