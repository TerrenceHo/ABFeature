package services

import "errors"

var (
	ErrProjectValidation = errors.New("Project model validation failed.")

	ErrExperimentValidation = errors.New("Experiment model validation failed.")

	ErrIdInvalid = errors.New("ID cannot be an empty string.")
)
