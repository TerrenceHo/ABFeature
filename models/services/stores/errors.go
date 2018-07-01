package stores

import "errors"

var (
	// Inidcates that the model object failed validation
	ErrInvalidEntry = errors.New("invalid project entry")

	// Indicates the the model object could not be found
	ErrNoEntryFound = errors.New("no project found")

	ErrNoRows = errors.New("no rows found")
)
