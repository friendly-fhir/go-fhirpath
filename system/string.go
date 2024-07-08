package system

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strings"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	profile "github.com/friendly-fhir/go-fhir/r4/core/profiles"
	"github.com/friendly-fhir/go-fhirpath/internal/esc"
)

type String string

// NewString constructs a new System.String object with the given value.
func NewString(format string, args ...any) String {
	return String(fmt.Sprintf(format, args...))
}

// ParseString parses a string from a FHIRPath string literal. This will translate
// any embedded escapes into the associated values.
func ParseString(str string) (String, error) {
	result, ok := strings.CutPrefix(str, "'")
	if !ok {
		return "", newParseError[String](str, fmt.Errorf("missing prefix quote"))
	}

	result, ok = strings.CutSuffix(result, "'")
	if !ok {
		return "", newParseError[String](str, fmt.Errorf("missing suffix quote"))
	}

	result, err := esc.Parse(result)
	if err != nil {
		return "", newParseError[String](str, err)
	}

	return String(result), nil
}

// MustParseString parses a string from a FHIRPath string literal, and panics if
// the value is invalid.
func MustParseString(str string) String {
	result, err := ParseString(str)
	if err != nil {
		panic(err)
	}
	return result
}

func (String) isAny() {}

// Compares compares the other system.String value to provide a total-ordering.
func (s String) Compare(other String) int {
	return strings.Compare(string(s), string(other))
}

// Equivalent compares two System.String values for FHIRPath equivalence.
//
// This will ignore casing and normalize whitespace.
func (s String) Equivalent(other String) bool {
	return s.normalize() == other.normalize()
}

func (s String) normalize() string {
	result := strings.ToLower(string(s))
	result = whitespaceNormalizer.Replace(result)
	return result
}

// Formatting

// String returns the Go string representation of this System.String.
func (s String) String() string {
	return string(s)
}

// Format formats this System.String with the given format and arguments.
func (s String) Format(state fmt.State, verb rune) {
	fmt.Fprintf(state, "%"+string(verb), string(s))
}

var (
	_ fmt.Stringer  = (*String)(nil)
	_ fmt.Formatter = (*String)(nil)
)

// R4 conversions

// R4 converts this System.String into a FHIR String type.
func (s String) R4() *fhir.String {
	return &fhir.String{Value: string(s)}
}

// FromR4 converts a FHIR String type into a System.String type.
func (s *String) FromR4(r profile.String) {
	*s = String(r.GetValue())
}

// Marshal conversions

// MarshalJSON converts this String object into a JSON object.
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON converts a JSON object into a String object.
func (s *String) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = String(value)
	return nil
}

var (
	_ json.Marshaler   = (*String)(nil)
	_ json.Unmarshaler = (*String)(nil)
)

// Text conversions

// MarshalText converts this String object into a text object.
func (s String) MarshalText() ([]byte, error) {
	return s.MarshalJSON()
}

// UnmarshalText converts a text object into a String object.
func (s *String) UnmarshalText(text []byte) error {
	return s.UnmarshalJSON(text)
}

var (
	_ encoding.TextMarshaler   = (*String)(nil)
	_ encoding.TextUnmarshaler = (*String)(nil)
)

var whitespaceNormalizer *strings.Replacer

func init() {
	whitespaceNormalizer = strings.NewReplacer(
		"\t", " ",
		"\n", " ",
		"\r", " ",
	)
}
