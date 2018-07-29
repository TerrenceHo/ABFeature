package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserExperimentValidate(t *testing.T) {
	assert := assert.New(t)

	ue := &UserExperiment{
		ID:           "",
		UserID:       "",
		ExperimentID: "",
	}

	errs := ue.Validate()

	assert.Equal(len(errs), 3, "UserExperiment should have three errors.")
	assert.Equal(errs[0].Error(), "UserExperiment must have an ID", "UserExperiment ID cannot be empty")
	assert.Equal(errs[1].Error(), "UserExperiment must have an UserID", "UserExperiment UserID cannot be empty")
	assert.Equal(errs[2].Error(), "UserExperiment must have a ExperimentID", "UserExperiment ExperimentID cannot be empty")
}
