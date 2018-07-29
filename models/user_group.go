package models

import (
	"errors"
	"time"
)

// Joining table to establish many-to-many relationship between Users and Groups
type UserGroup struct {
	// Unique ID for identification
	ID string `json:"ID" db:"id"`

	// User ID to reference experiment
	UserID string `json:"UserID" db:"user_id"`

	// Group ID to reference group
	GroupID string `json:"GroupID" db:"group_id"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

// Runs data validation on UserGroup object. Checks that the data is of proper
// values. None of the ID values may be empty.
func (ug *UserGroup) Validate() (errs []error) {
	if ug.ID == "" {
		errs = append(errs, errors.New("UserGroup must have an ID"))
	}
	if ug.UserID == "" {
		errs = append(errs, errors.New("UserGroup must have an UserID"))
	}
	if ug.GroupID == "" {
		errs = append(errs, errors.New("UserGroup must have a GroupID"))
	}

	return errs
}
