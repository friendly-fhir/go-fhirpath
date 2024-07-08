package collection

import "errors"

var (
	// ErrNotSingleton is an error raised if a collection is not a singleton, but
	// one was expected.
	ErrNotSingleton = errors.New("collection is not singleton")

	// ErrNotConvertible is an error raised when attempting to call Collection.To*
	// to a type that is not convertible.
	ErrNotConvertible = errors.New("not convertible")
)
