package models

import (
	"errors"
	"time"
)

type Group struct {
	// ID is programmatically generated
	ID string `json:"ID" db:"id"`

	// Name of Group is typically the name of the software Group being
	// tested, provided by the user to be human readable.
	Name string `json:"Name" db:"name"`

	// Description of a Group. For human purposes only
	Description string `json:"Description" db:"description"`

	// Attributes allows users to store custom data
	// Attributes map[string]interface{} `json:"Attributes"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
	// DeletedAt time.Time `json:"DeletedAt"`
}

// Runs data validation on a Group object. It checks that data provided is of
// proper values, and returns the list of all errors.
// Validations:
// (1) Groups must have a name
// (2) Groups must have an ID
func (g *Group) Validate() []error {
	var errs []error
	if p.ID == "" {
		errs = append(errs, errors.New("ID primary key cannot be empty"))
	}
	if p.Name == "" {
		errs = append(errs, errors.New("Please provide a Group name."))
	}

	return errs
}
