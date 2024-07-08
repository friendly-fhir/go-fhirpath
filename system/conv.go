package system

import (
	"fmt"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	profile "github.com/friendly-fhir/go-fhir/r4/core/profiles"
)

// FromR4 return a system type from a (valid) FHIR R4 element.
func FromR4(element fhir.Element) (Any, error) {
	switch e := element.(type) {
	case *fhir.Boolean:
		var b Boolean
		b.FromR4(e)
		return b, nil
	case profile.Integer:
		var i Integer
		i.FromR4(e)
		return i, nil
	case profile.String:
		var s String
		s.FromR4(e)
		return s, nil
	}
	return nil, fmt.Errorf("%w: %T is not a valid R4 type", ErrNotConvertible, element)
}

// Normalizes a FHIR R4 type into a system type, if able -- or just returns
// the input value if it's not a FHIR R4 type.
func Normalize(v any) any {
	element, ok := v.(fhir.Element)
	if !ok {
		return v
	}
	got, err := FromR4(element)
	if err != nil {
		return element
	}
	return got
}
