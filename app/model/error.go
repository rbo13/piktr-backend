package model

import "errors"

var (
	// ErrNotInserted ...
	ErrNotInserted = errors.New("User not inserted")
	// ErrNotUpdated ...
	ErrNotUpdated = errors.New("User not updated")
	// ErrIDIsRequired ...
	ErrIDIsRequired = errors.New("ID is required")
	// ErrNotFound ...
	ErrNotFound = errors.New("User not found")
)
