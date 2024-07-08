package system

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

// Integer64 is the Go-representation of the FHIRPath System.Integer type. This
// is a 64-bit signed integer value.
type Integer64 int64

// NewInteger64 constructs a new System.Integer object with the given value.
func NewInteger64(value int64) Integer64 {
	return Integer64(value)
}

// ParseInteger64 parses a string into the valid FHIRPath System.Integer type.
func ParseInteger64(str string) (Integer64, error) {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Integer64(0), newParseError[Integer64](str, err)
	}

	return Integer64(value), nil
}

// MustParseInteger64 parses an integer string, and panics if the value is invalid.
func MustParseInteger64(str string) Integer64 {
	got, err := ParseInteger64(str)
	if err != nil {
		panic(err)
	}
	return got
}

func (Integer64) isAny() {}

// Negate returns the inverse polarity of this integer value.
func (i Integer64) Negate() Integer64 {
	return -i
}

// Compare compares two integer values.
// This returns a negative value if this object is less than other,
// a positive number if other is greater than this, or equal if both values are
// the same.
//
// For example:
//
//	a, b := system.Integer64(0), system.Integer64(42)
//	assert.True(a.Compare(b) < 0)
//	assert.True(b.Compare(a) > 0)
//	assert.True(a.Compare(a) == 0)
func (i Integer64) Compare(other Integer64) int {
	return int(i - other)
}

// Int64 converts this system.Integer into an in64Go native type.
func (i Integer64) Int64() int64 {
	return int64(i)
}

// Formatting

// String returns the string representation of the System.Integer.
func (i Integer64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Format implements the fmt.Formatter interface.
func (i Integer64) Format(state fmt.State, verb rune) {
	fmt.Fprintf(state, "%"+string(verb), int64(i))
}

var (
	_ fmt.Stringer  = (*Integer64)(nil)
	_ fmt.Formatter = (*Integer64)(nil)
)

// JSON conversions

// MarshalJSON converts this Integer object into a JSON object.
func (i Integer64) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalJSON converts a JSON object into an Integer object.
func (i *Integer64) UnmarshalJSON(data []byte) error {
	value, err := ParseInteger64(string(data))
	if err != nil {
		return err
	}

	*i = value
	return nil
}

var (
	_ json.Marshaler   = (*Integer64)(nil)
	_ json.Unmarshaler = (*Integer64)(nil)
)

// Text conversions

// MarshalText converts this Integer object into a text object.
func (i Integer64) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts a text object into an Integer object.
func (i *Integer64) UnmarshalText(text []byte) error {
	value, err := ParseInteger64(string(text))
	if err != nil {
		return err
	}

	*i = value
	return nil
}

var (
	_ encoding.TextMarshaler   = (*Integer64)(nil)
	_ encoding.TextUnmarshaler = (*Integer64)(nil)
)
