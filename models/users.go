package models

import (
	"errors"
	"time"
)

// A user is an indivudual unit that a test/experiment can be applied to. Users
// can belong to groups, in which case all group rules apply to that user. They
// can also belong to an experiment, in which case the experiment applies to the
// user, regardless of the percentage set in the experiment.
type User struct {
	// ID is programmatically genereated
	ID string `json:ID" db:"id"`

	// Human readable Name
	Name string `json:"Name" db:"name"`

	// Description of user. User details. Human purposes only.
	Description string `json:"Description" db:"description"`

	Active bool `json:"Active" db:"active"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

// Runs data validation on a User object. It checks that data provided is of
// proper values, and returns a list of all errors.
// Validations:
// (1) Users must have a name
// (2) Users must have an ID
func (u *User) Validate() []error {
	var errs []error
	if u.ID == "" {
		errs = append(errs, errors.New("ID primary key cannot be empty"))
	}
	if u.Name == "" {
		errs = append(errs, errors.New("Please provide a user name."))
	}

	return errs
}

// Updates User fields, runs validation
func (u *User) UpdateFields(user *User) []error {
	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Description != "" {
		u.Description = user.Description
	}
	errs := u.Validate()
	return errs
}
