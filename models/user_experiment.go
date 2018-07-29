package models

import (
	"errors"
	"time"
)

// Joining table to establish many-to-many relationshiip betwen Users and
// Experiments
type UserExperiment struct {
	// Unique ID for identification
	ID string `json:"ID" db:"id"`

	// User ID to reference experiment
	UserID string `json:"UserID" db:"user_id"`

	// Experiment ID to reference experiment
	ExperimentID string `json:"ExperimentID" db:"experiment_id"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

// Runs data validation on UserExperiment object. Checks that the data is of
// proper values. None of the ID values may be empty.
func (ue *UserExperiment) Validate() (errs []error) {
	if ue.ID == "" {
		errs = append(errs, errors.New("UserExperiment must have an ID"))
	}
	if ue.UserID == "" {
		errs = append(errs, errors.New("UserExperiment must have an UserID"))
	}
	if ue.ExperimentID == "" {
		errs = append(errs, errors.New("UserExperiment must have a ExperimentID"))
	}

	return errs
}
