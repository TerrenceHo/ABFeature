package models

import (
	"errors"
	"time"
)

// An Experiment is a feature test that is being conducted on a project. This
// can be anything from changing the color of the button, to showing different
// static web pages.
type Experiment struct {
	// Unique ID used for identification
	ID string `json:"ID" db:"id"`

	// Human readable name
	Name string `json:"Name" db:"name"`

	// Description of the experiment.  For human purposes only
	Description string `json:"Description" db:"description"`

	// Sets what percentage of users see the new feature.  Given as a float,
	// with values between 0.0 - 1.0 inclusive.
	Percentage float64 `json:"Percentage" db:"percentage"`

	// Sets whether the experiment is active or not.
	Enabled bool `json:"Enabled" db:"enabled"`

	// User set key-value stores about the Experiment, to store custom values.
	// Attributes map[string]interface{} `json:"Attributes"`

	// ID of Project associated with Experiment
	ProjectID string `json:"ProjectID" db:"project_id"`

	// Metadata
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
	// DeletedAt time.Time `json:"DeletedAt"`
}

// Runs data validation on an Experiment object. It checks that data provided
// is of proper values, and returns the list of all errors.
// Validations:
// (1) Experiments must have a name.
// (2) Percentage must between be between 0 and 1.
// (3) Experiments must be associated with a project
// (4) Experiments must have an ID
func (e *Experiment) Validate() (errs []error) {
	if e.ID == "" {
		errs = append(errs, errors.New("Experiments must have an ID"))
	}

	if e.Name == "" {
		errs = append(errs, errors.New("Please provide a project name"))
	}

	if e.Percentage > 1.0 || e.Percentage < 0.0 {
		errs = append(errs, errors.New("Please provide a percentage between 0 and 1"))
	}

	if e.ProjectID == "" {
		errs = append(errs,
			errors.New("Please select a project to associate the experiment with."))
	}

	return errs
}

// Updates Experiment's fields, and runs validation on data
func (e *Experiment) UpdateFields(experiment *Experiment) []error {
	if experiment.Name != "" {
		e.Name = experiment.Name
	}
	if experiment.Description != "" {
		e.Description = experiment.Description
	}
	if experiment.Percentage != -1 {
		e.Percentage = experiment.Percentage
	}
	e.Enabled = experiment.Enabled

	errs := e.Validate()
	return errs
}
