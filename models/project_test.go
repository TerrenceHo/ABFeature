package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectValidate(t *testing.T) {
	assert := assert.New(t)

	project := &Project{
		ID:          "",
		Name:        "",
		Description: "",
	}
	errs := project.Validate()
	assert.Equal(len(errs), 2, "Number of errors should be 2")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")
	assert.Equal(errs[1].Error(), "Please provide a project name.")

	project.Name = "Name"
	errs = project.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "ID primary key cannot be empty")

	project.Name = ""
	project.ID = "ID"
	errs = project.Validate()
	assert.Equal(len(errs), 1, "Number of errors should be 1")
	assert.Equal(errs[0].Error(), "Please provide a project name.")
}

func TestProjectUpdateFields(t *testing.T) {
	assert := assert.New(t)
	old_project := &Project{
		ID:          "ID",
		Name:        "Project Name",
		Description: "Project Description",
	}

	new_project := &Project{
		ID:          "ID",
		Name:        "New Project Name",
		Description: "",
	}

	errs := old_project.UpdateFields(new_project)
	assert.Equal(len(errs), 0, "Number of errors should be zero.")
	assert.Equal(old_project.ID, "ID", "ID should be 'ID'")
	assert.Equal(old_project.Name, "New Project Name", "Project name should be updated.")
	assert.Equal(old_project.Description, "Project Description", "Description should be unchanged.")
}
