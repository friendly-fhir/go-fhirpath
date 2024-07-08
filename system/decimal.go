package system

import (
	"encoding"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// Decimal is the Go-representation of the FHIRPath System.Decimal type. This is
// a fixed-point decimal value.
type Decimal decimal.Decimal

// NewDecimal constructs a Decimal object with the underlying value.
//
// This function primarily exists for symmetry with the other constructor
// functions.
func NewDecimal(d float64) Decimal {
	return Decimal(decimal.NewFromFloat(d))
}

// ParseDecimal parses a string into the valid FHIRPath System.Decimal type.
func ParseDecimal(str string) (Decimal, error) {
	value, err := decimal.NewFromString(str)
	if err != nil {
		return Decimal(decimal.New(0, 0)), newParseError[Decimal](str, err)
	}

	return Decimal(value), nil
}

func (Decimal) isAny() {}

// Comparisons

// Equal compares the other system.Decimal value to provide a total-ordering.
func (d Decimal) Equal(other Decimal) bool {
	return (*decimal.Decimal)(&d).Equal(decimal.Decimal(other))
}

// Equivalent compares two System.Decimal values for FHIRPath equivalence.
func (d Decimal) Equivalent(other Decimal) bool {
	// Clamp precision to [0, -N], and invert to be positive [N, 0].
	l := -min(0, (*decimal.Decimal)(&d).Exponent())
	r := -min(0, (*decimal.Decimal)(&other).Exponent())

	// Use the smallest precision to truncate both values.
	precision := min(l, r)
	lhs := (*decimal.Decimal)(&d).Truncate(precision)
	rhs := (*decimal.Decimal)(&other).Truncate(precision)
	return lhs.Equal(rhs)
}

// Compare compares two decimal values, returning a negative value if this
// object is less than other,
func (d Decimal) Compare(other Decimal) int {
	return (*decimal.Decimal)(&d).Cmp(decimal.Decimal(other))
}

// Conversions

// Integer converts this system.Decimal into a system.Integer type.
func (d Decimal) Integer() Integer {
	return Integer(decimal.Decimal(d).IntPart())
}

// Integer64 converts this system.Decimal into a system.Integer64 type.
func (d Decimal) Integer64() Integer64 {
	return Integer64(decimal.Decimal(d).IntPart())
}

// Float64 converts this system.Decimal into a float64 Go native type.
func (d Decimal) Float64() float64 {
	return decimal.Decimal(d).InexactFloat64()
}

// Formatter

// String returns the string representation of the System.Decimal.
func (d Decimal) String() string {
	return decimal.Decimal(d).String()
}

// Format implements the fmt.Formatter interface.
func (d Decimal) Format(state fmt.State, verb rune) {
	fmt.Fprintf(state, "%"+string(verb), decimal.Decimal(d).InexactFloat64())
}

var (
	_ fmt.Stringer  = (*Decimal)(nil)
	_ fmt.Formatter = (*Decimal)(nil)
)

// JSON conversions

// MarshalJSON converts this Decimal object into a JSON object.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalJSON converts a JSON object into an Decimal object.
func (d *Decimal) UnmarshalJSON(data []byte) error {
	value, err := ParseDecimal(string(data))
	if err != nil {
		return err
	}

	*d = value
	return nil
}

var (
	_ json.Marshaler   = (*Decimal)(nil)
	_ json.Unmarshaler = (*Decimal)(nil)
)

// Text conversions

// MarshalText converts this Decimal object into a text object.
func (d Decimal) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText converts a text object into an Decimal object.
func (d *Decimal) UnmarshalText(text []byte) error {
	value, err := ParseDecimal(string(text))
	*d = value
	return err
}

var (
	_ encoding.TextMarshaler   = (*Decimal)(nil)
	_ encoding.TextUnmarshaler = (*Decimal)(nil)
)
