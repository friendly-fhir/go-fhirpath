/*
Package collection provides definitions for FHIRPath Collection objects, for
semantically clear behavior of FHIRPath operations involving collections.

This includes general equality, equivalence, comparison, filtering, etc.
*/
package collection

import (
	"fmt"

	r4 "github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhirpath/system"
)

// Collection represents a FHIRPath Collection, which is both the input set and
// output result of any FHIRPath expression.
type Collection []any

var (
	// True is a Collection containing only a single Boolean "true" value.
	True = Collection{system.Boolean(true)}

	// False is a Collection containing only a single Boolean "false" value.
	False = Collection{system.Boolean(false)}

	// Empty is a Collection containing no values.
	Empty Collection = nil
)

// Of creates a fhirpath Collection of the input arguments.
//
// This function will panic if any of the inputs do not satisfy the following
// valid types:
// - fhir.Resource
// - fhir.Element
// - system.Any
// - reflect.TypeSpecifier
func Of(args ...any) Collection {
	for _, arg := range args {
		if !isValidType(arg) {
			panic(fmt.Sprintf("invalid input type %T: inputs must be valid collection type (fhir.Element, fhir.Resource, system.Any, or reflect.TypeSpecifier)", arg))
		}
	}
	return Collection(args)
}

func isValidType(a any) bool {
	switch a.(type) {
	case r4.Resource, r4.Element, system.Any:
		return true
	}
	return false
}

// FromSlice converts a slice of an underlying type into a Collection object.
//
// This function will panic if T does not satisfy one of the valid FHIR Collection
// types:
// - fhir.Resource
// - fhir.Element
// - system.Any
// - reflect.TypeSpecifier
func FromSlice[S ~[]T, T any](s S) Collection {
	slices := make([]any, len(s))
	for i, v := range s {
		slices[i] = v
	}
	return Of(slices...)
}

// Singleton creates a singleton FHIRPath Collection containing only the one
// input argument. This function mostly exists for self-documenting purposes.
//
// This function will panic if the input does not satisfy one of the following
// valid types:
// - fhir.Resource
// - fhir.Element
// - system.Any
// - reflect.TypeSpecifier
func Singleton(v any) Collection {
	return Of(v)
}

// Join returns a Collection that contains all the elements of the collections
// being joined together.
func Join(collections ...Collection) Collection {
	length := 0
	for _, collection := range collections {
		length += len(collection)
	}

	result := make(Collection, 0, length)
	for _, collection := range collections {
		result = append(result, collection...)
	}

	return result
}

// IsEmpty queries whether this collection is empty.
func (c Collection) IsEmpty() bool {
	return len(c) == 0
}

// IsSingleton queries whether this collection contains exactly 1 value.
func (c Collection) IsSingleton() bool {
	return len(c) == 1
}

// Len returns the number of elements in the collection.
func (c Collection) Len() int {
	return len(c)
}

// Singleton converts a Collection into a singleton value of an unspecified type.
// If the value is not a singleton, an ErrNotSingleton is returned instead.
func (c Collection) Singleton() (any, error) {
	if !c.IsSingleton() {
		return nil, ErrNotSingleton
	}
	return c[0], nil
}

// SingletonBoolean converts a Collection into a boolean, following the
// FHIRPath Singleton Evaluation of Booleans. This behaves in the following way:
//
//   - If the underly collection is empty, this returns empty
//   - If the collection is a singleton of boolean, it returns that boolean
//   - If the collection is a singleton containing a value, this returns true
//   - Otherwise, this raises an error for non-singleton values
//
// There error raised when not a singleton is ErrNotSingleton.
//
// See: https://hl7.org/fhirpath/N1/#singleton-evaluation-of-collections
func (c Collection) SingletonBoolean() (Collection, error) {
	if c.IsEmpty() {
		return Empty, nil
	}

	if !c.IsSingleton() {
		return Empty, ErrNotSingleton
	}

	if val, ok := system.Normalize(c[0]).(system.Boolean); ok {
		return Singleton(val), nil
	}

	// Single-value nodes are implicitly converted to true if it's not a bool
	return True, nil
}

// Bool converts this collection into a single boolean value.
//
// This function will error in the following circumstances:
//
//   - If this collection is not a singleton, an ErrNotSingleton is returned.
func (c Collection) Bool() (bool, error) {
	if !c.IsSingleton() {
		return false, ErrNotSingleton
	}

	if val, ok := system.Normalize(c[0]).(system.Boolean); ok {
		return val.Bool(), nil
	}
	return true, nil
}

// String converts this collection into a single string value. This function
// will return an error if the collection is not a singleton (e.g. contains only
// one value), or if the entry is not either a System.Boolean or derived FHIR.String
// type.
//
// This function will error in the following circumstances:
//
//   - If this collection is not a singleton, an ErrNotSingleton is returned.
//   - If the underlying type of the singleton is not convertible to string, an
//     ErrNotConvertible is returned.
func (c Collection) String() (string, error) {
	if !c.IsSingleton() {
		return "", ErrNotSingleton
	}

	if val, ok := system.Normalize(c[0]).(system.String); ok {
		return val.String(), nil
	}

	return "", c.convertErr(c[0], "string")
}

// Int32 converts this collection into an int32 string value. This function
// will return an error if the collection is not a singleton (e.g. contains only
// one value), or if the entry is not either a System.Integer type.
//
// This function will error in the following circumstances:
//
//   - If this collection is not a singleton, an ErrNotSingleton is returned.
//   - If the underlying type of the singleton is not convertible to int32, an
//     ErrNotConvertible is returned.
func (c Collection) Int32() (int32, error) {
	if !c.IsSingleton() {
		return 0, ErrNotSingleton
	}

	if val, ok := system.Normalize(c[0]).(system.Integer); ok {
		return int32(val), nil
	}

	return 0, c.convertErr(c[0], "int32")
}

// Normalize returns a new collection containing all the elements from this
// collection normalized into FHIRPath system types, where possible.
func (c Collection) Normalize() Collection {
	result := make(Collection, 0, len(c))
	for _, v := range c {
		result = append(result, system.Normalize(v))
	}
	return result
}

func (c Collection) Equal(other Collection) bool {
	if len(c) != len(other) {
		return false
	}

	for i := range c {
		// TODO: implement proper comparison
		if c[i] != other[i] {
			return false
		}
	}

	return true
}

func (c Collection) convertErr(got any, want string) error {
	return fmt.Errorf("type %T %w to %v", got, ErrNotConvertible, want)
}
