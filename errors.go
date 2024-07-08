package fhirpath

import (
	"errors"

	"github.com/friendly-fhir/go-fhirpath/collection"
)

var (
	// ErrUnimplemented is returned when a feature is not yet implemented.
	ErrUnimplemented = errors.New("unimplemented")

	// ErrNotSingleton is an error raised if a collection is not a singleton, but
	// one was expected.
	ErrNotSingleton = collection.ErrNotSingleton
)
