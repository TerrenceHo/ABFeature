package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExperimentGroupValidate(t *testing.T) {
	assert := assert.New(t)
	eg := &ExperimentGroup{
		ID:           "",
		ExperimentID: "",
		GroupID:      "",
	}

	errs := eg.Validate()
	assert.Equal(len(errs), 3, "ExperimentGroup should have three errors.")
	assert.Equal(errs[0].Error(), "ExperimentGroup must have an ID", "ExperimentGroup ID cannot be empty")
	assert.Equal(errs[1].Error(), "ExperimentGroup must have an ExperimentID", "ExperimentGroup ExperimentID cannot be empty")
	assert.Equal(errs[2].Error(), "ExperimentGroup must have a GroupID", "ExperimentGroup GroupID cannot be empty")

}
