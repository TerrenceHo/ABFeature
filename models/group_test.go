package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupValidate(t *testing.T) {
	assert := assert.New(t)

	group := &Group{
		ID:          "",
		Name:        "",
		Description: "",
	}
	errs := group.Validate()
	assert.Equal(len(errs), 2, "Number of errors should be 2")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")
	assert.Equal(errs[1].Error(), "Please provide a Group name.")

	group.Name = "Name"
	errs = group.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")

	group.Name = ""
	group.ID = "ID"
	errs = group.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "Please provide a Group name.")
}

func TestGroupUpdateFields(t *testing.T) {
	assert := assert.New(t)
	old_group := &Group{
		ID:          "ID",
		Name:        "Group Name",
		Description: "Group Description",
	}

	new_group := &Group{
		ID:          "ID",
		Name:        "New Group Name",
		Description: "",
	}

	errs := old_group.UpdateFields(new_group)
	assert.Equal(len(errs), 0, "Number of errors should be zero.")
	assert.Equal(old_group.ID, "ID", "ID should be 'ID'")
	assert.Equal(old_group.Name, "New Group Name", "Group name should be updated.")
	assert.Equal(old_group.Description, "Group Description", "Group Description should be unchanged.")
}
