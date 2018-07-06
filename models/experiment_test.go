package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExperimentValidate(t *testing.T) {
	assert := assert.New(t)

	experiment := &Experiment{
		ID:          "",
		Name:        "",
		Description: "",
		Percentage:  2.0,
		Enabled:     false,
		ProjectID:   "",
	}
	errs := experiment.Validate()
	assert.Equal(len(errs), 4, "Number of errors should be 4")
	assert.Equal(errs[0].Error(), "Experiments must have an ID")
	assert.Equal(errs[1].Error(), "Please provide a project name")
	assert.Equal(errs[2].Error(), "Please provide a percentage between 0 and 1")
	assert.Equal(errs[3].Error(), "Please select a project to associate the experiment with.")

	experiment.Percentage = -1
	errs = experiment.Validate()
	assert.Equal(len(errs), 4, "Number of errors should be 4")
	assert.Equal(errs[0].Error(), "Experiments must have an ID")
	assert.Equal(errs[1].Error(), "Please provide a project name")
	assert.Equal(errs[2].Error(), "Please provide a percentage between 0 and 1")
	assert.Equal(errs[3].Error(), "Please select a project to associate the experiment with.")
}

func TestExperimentUpdateFields(t *testing.T) {
	assert := assert.New(t)

	old_experiment := &Experiment{
		ID:          "ID",
		Name:        "Experiment Name",
		Description: "Description",
		Percentage:  0.4,
		Enabled:     false,
		ProjectID:   "ProjectID",
	}

	new_experiment := &Experiment{
		ID:          "ID",
		Name:        "New Experiment Name",
		Description: "",
		Percentage:  0.6,
		Enabled:     true,
		ProjectID:   "ProjectID",
	}

	errs := old_experiment.UpdateFields(new_experiment)
	assert.Equal(len(errs), 0, "There should be no errors updating experiment")
	assert.Equal(old_experiment.ID, "ID", "Experiment ID should be the same.")
	assert.Equal(old_experiment.Name, "New Experiment Name", "Experiment Name should be changed.")
	assert.Equal(old_experiment.Description, "Description", "Experiment Description should be unchanged.")
	assert.Equal(old_experiment.Percentage, 0.6, "Experiment Percentage should be changed.")
	assert.Equal(old_experiment.Enabled, true, "Experiment Enabled should be true.")
	assert.Equal(old_experiment.ProjectID, "ProjectID", "Experiment ProjectID should be unchanged.")
}
