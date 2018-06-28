package models

import (
	"errors"
	"time"
)

// Each Project represents a software project that you want to run experiments
// on. A Project is associated with many experiments, representing all the
// feature tests/experiments that are being conducted in the software project.
type Project struct {
	// ID is programmatically generated
	ID string `json:"ID" db:"id"`

	// Name of Project is typically the name of the software project being
	// tested, provided by the user to be human readable.
	Name string `json:"Name" db:"name"`

	// Description of a project. For human purposes only
	Description string `json:"Description" db:"description"`

	// Attributes allows users to store custom data
	// Attributes map[string]interface{} `json:"Attributes"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
	// DeletedAt time.Time `json:"DeletedAt"`
}

// Runs data validation on a Project object. It checks that data provided is of
// proper values, and returns the list of all errors.
// Validations:
// (1) Projects must have a name
// (2) Projects must have an ID
func (p *Project) Validate() []error {
	var errs []error
	if p.ID == "" {
		errs = append(errs, errors.New("ID primary key cannot be empty"))
	}
	if p.Name == "" {
		errs = append(errs, errors.New("Please provide a project name."))
	}

	return errs
}
