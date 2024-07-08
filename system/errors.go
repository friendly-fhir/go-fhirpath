package system

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// ErrNotConvertible is an error raised when attempting to call Collection.To*
	// to a type that is not convertible.
	ErrNotConvertible = fmt.Errorf("not convertible")
)

// ParseError is returned from objects during failed parsing.
// This error optionally contains an underlying error reason that may propagate
// failures caused by internal system APIs; the type and value of which is not
// guaranteed to be stable, but may be used for logging purposes.
type ParseError struct {
	Type   reflect.Type
	Input  string
	Reason error
}

func (e *ParseError) Error() string {
	name := e.Type.Name()
	input := e.Input
	return fmt.Sprintf(
		"%v parse '%v': %v",
		strings.ToLower(name),
		input,
		e.Reason,
	)
}

func newParseError[T Any](input string, reason error) error {
	var v T
	return &ParseError{
		Type:   reflect.TypeOf(v),
		Input:  input,
		Reason: reason,
	}
}
