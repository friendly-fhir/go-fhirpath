package system

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	profile "github.com/friendly-fhir/go-fhir/r4/core/profiles"
)

// Integer is the Go-representation of the FHIRPath System.Integer type. This is
// a 32-bit signed integer value.
type Integer int32

// NewInteger constructs a new System.Integer object with the given value.
func NewInteger(value int32) Integer {
	return Integer(value)
}

// ParseInteger parses a string into the valid FHIRPath System.Integer type.
func ParseInteger(str string) (Integer, error) {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return Integer(0), newParseError[Integer](str, err)
	}

	return Integer(value), nil
}

// MustParseInteger parses an integer string, and panics if the value is invalid.
func MustParseInteger(str string) Integer {
	got, err := ParseInteger(str)
	if err != nil {
		panic(err)
	}
	return got
}

func (Integer) isAny() {}

// Negate returns the inverse polarity of this integer value.
func (i Integer) Negate() Integer {
	return -i
}

// Compare compares two integer values.
// This returns a negative value if this object is less than other,
// a positive number if other is greater than this, or equal if both values are
// the same.
//
// For example:
//
//	a, b := system.Integer(0), system.Integer(42)
//	assert.True(a.Compare(b) < 0)
//	assert.True(b.Compare(a) > 0)
//	assert.True(a.Compare(a) == 0)
func (i Integer) Compare(other Integer) int {
	return int(i - other)
}

// Int32 converts this system.Integer into an in32Go native type.
func (i Integer) Int32() int32 {
	return int32(i)
}

// Formatter

// String returns the string representation of the System.Integer.
func (i Integer) String() string {
	return strconv.Itoa(int(i))
}

// Format formats this System.Integer with the given format and arguments.
func (i Integer) Format(state fmt.State, verb rune) {
	fmt.Fprintf(state, "%"+string(verb), int32(i))
}

var (
	_ fmt.Stringer  = (*Integer)(nil)
	_ fmt.Formatter = (*Integer)(nil)
)

// R4 Conversions

// FromR4 converts a FHIR Integer type into a System.Integer type.
func (i *Integer) FromR4(in profile.Integer) {
	*i = Integer(in.GetValue())
}

// R4 converts this System.Integer into a FHIR Integer type.
func (i *Integer) R4() *fhir.Integer {
	return &fhir.Integer{Value: int32(*i)}
}

// JSON conversions

// MarshalJSON converts this Integer object into a JSON object.
func (i Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(i))
}

// UnmarshalJSON converts a JSON object into an Integer object.
func (i *Integer) UnmarshalJSON(data []byte) error {
	var value int32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Integer(value)
	return nil
}

var (
	_ json.Marshaler   = (*Integer)(nil)
	_ json.Unmarshaler = (*Integer)(nil)
)

// Text conversions

// MarshalText converts this Integer object into a text object.
func (i Integer) MarshalText() ([]byte, error) {
	return i.MarshalJSON()
}

// UnmarshalText converts a text object into an Integer object.
func (i *Integer) UnmarshalText(text []byte) error {
	return i.UnmarshalJSON(text)
}

var (
	_ encoding.TextMarshaler   = (*Integer)(nil)
	_ encoding.TextUnmarshaler = (*Integer)(nil)
)
