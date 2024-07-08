package system

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
)

// Boolean is the FHIRPath system-type representation of the "boolean" value.
type Boolean bool

// NewBoolean constructs a Boolean object with the underlying value.
//
// This function primarily exists for symmetry with the other constructor
// functions.
func NewBoolean(b bool) Boolean {
	return Boolean(b)
}

// ParseBoolean parses a string into the valid FHIRPath System.Boolean type.
func ParseBoolean(str string) (Boolean, error) {
	switch str {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, newParseError[Boolean](str, nil)
}

// MustParseBoolean parses a boolean string, and panics if the value is invalid.
func MustParseBoolean(str string) Boolean {
	got, err := ParseBoolean(str)
	if err != nil {
		panic(err)
	}
	return got
}

func (Boolean) isAny() {}

// Negate returns the inverse polarity of this boolean value.
func (b Boolean) Negate() Boolean {
	return !b
}

// Bool returns the Go boolean representation of the System.Boolean.
func (b Boolean) Bool() bool {
	return bool(b)
}

// Formatting

// String returns the string representation of the System.Boolean.
func (b Boolean) String() string {
	return strconv.FormatBool(bool(b))
}

// Format implements the fmt.Formatter interface.
func (b Boolean) Format(s fmt.State, verb rune) {
	fmt.Fprintf(s, "%"+string(verb), bool(b))
}

var (
	_ fmt.Stringer  = (*Boolean)(nil)
	_ fmt.Formatter = (*Boolean)(nil)
)

// R4 conversions

// FromR4 converts a FHIR Boolean type into a System.Boolean type.
func (b *Boolean) FromR4(r *fhir.Boolean) {
	*b = Boolean(r.Value)
}

// R4 converts this System.Boolean into a FHIR Boolean type.
func (b Boolean) R4() *fhir.Boolean {
	return &fhir.Boolean{Value: bool(b)}
}

// JSON conversions

// MarshalJSON converts this Boolean object into a JSON object.
func (b Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

// UnmarshalJSON converts a JSON object into a Boolean object.
func (b *Boolean) UnmarshalJSON(data []byte) error {
	var value bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = Boolean(value)
	return nil
}

var (
	_ json.Marshaler   = (*Boolean)(nil)
	_ json.Unmarshaler = (*Boolean)(nil)
)

// Text conversions

// MarshalText converts this Boolean object into a text object.
func (b Boolean) MarshalText() ([]byte, error) {
	return b.MarshalJSON()
}

// UnmarshalText converts a text object into a Boolean object.
func (b *Boolean) UnmarshalText(text []byte) error {
	return b.UnmarshalJSON(text)
}

var (
	_ encoding.TextMarshaler   = (*Boolean)(nil)
	_ encoding.TextUnmarshaler = (*Boolean)(nil)
)
