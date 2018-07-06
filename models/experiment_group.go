package models

import (
	"errors"
	"time"
)

// Joining table, to establish many-to-many relationship for Experiments and
// Groups.
type ExperimentGroup struct {
	// Unique ID for identification
	ID string `json:"ID" db:"id"`

	// Experiment ID to reference experiment
	ExperimentID string `json:"ExperimentID" db:"experiment_id"`

	// Group ID to reference group
	GroupID string `json:"GroupID" db:"group_id"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

// Runs data validation on ExperimentGroup object.  It checks that the data
// provided is of proper values. None of the ID values may be empty.
func (eg *ExperimentGroup) Validate() (errs []error) {
	if eg.ID == "" {
		errs = append(errs, errors.New("ExperimentGroup must have an ID"))
	}
	if eg.ExperimentID == "" {
		errs = append(errs, errors.New("ExperimentGroup must have an ExperimentID"))
	}
	if eg.GroupID == "" {
		errs = append(errs, errors.New("ExperimentGroup must have a GroupID"))
	}

	return errs
}
