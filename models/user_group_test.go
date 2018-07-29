package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGroupValidate(t *testing.T) {
	assert := assert.New(t)

	ug := &UserGroup{
		ID:      "",
		UserID:  "",
		GroupID: "",
	}

	errs := ug.Validate()

	assert.Equal(len(errs), 3, "UserGroup should have three errors.")
	assert.Equal(errs[0].Error(), "UserGroup must have an ID", "UserGroup ID cannot be empty")
	assert.Equal(errs[1].Error(), "UserGroup must have an UserID", "UserGroup UserID cannot be empty")
	assert.Equal(errs[2].Error(), "UserGroup must have a GroupID", "UserGroup GroupID cannot be empty")
}
