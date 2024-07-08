package system_test

import (
	"errors"
	"testing"

	"github.com/friendly-fhir/go-fhirpath/system"
)

type cmpResult func(int) bool

var (
	less cmpResult = func(v int) bool {
		return v < 0
	}
	greater cmpResult = func(v int) bool {
		return v > 0
	}
	equal cmpResult = func(v int) bool {
		return v == 0
	}
)

func TestParseString(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  system.String
	}{
		{"ASCII", "'hello world'", "hello world"},
		{"Form Feed", `'\f'`, "\f"},
		{"Newline", `'\n'`, "\n"},
		{"Carriage Return", `'\r'`, "\r"},
		{"Single-quote", `'\''`, `'`},
		{"Double-quote", `'\"'`, `"`},
		{"Tab", `'\t'`, "\t"},
		{"Backtick", "'\\`'", "`"},
		{"Backslash", `'\\'`, `\`},
		{"Double-escape Form Feed", `'\\f'`, `\f`},
		{"Double-escape Newline", `'\\n'`, `\n`},
		{"Double-escape Carriage Return", `'\\r'`, `\r`},
		{"Double-escape Single-quote", `'\\''`, `\'`},
		{"Double-escape Double-quote", `'\\"'`, `\"`},
		{"Double-escape Tab", `'\\t'`, `\t`},
		{"Double-escape Backtick", "'\\\\`'", "\\`"},
		{"Double-escape Backslash", `'\\\\'`, `\\`},
		{"UTF-8-1", `'\u044d\u0442\u043e'`, `это`},
		{"UTF-8-2", `'\u0442\u0435\u0441\u0442'`, `тест`},
		{"UTF-8-3", `'\u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435'`, `сообщение`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := system.ParseString(tc.input)
			if err != nil {
				t.Fatalf("ParseString() = %v; want nil", err)
			}

			if got != tc.want {
				t.Errorf("ParseString() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestParseString_BadInput_ReturnsParseError(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Missing suffix quote", "'hello"},
		{"Missing prefix quote", "world'"},
		{"Missing both quotes", "hello world"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := system.ParseString(tc.input)

			var parseErr *system.ParseError
			ok := errors.As(err, &parseErr)

			if got, want := ok, true; got != want {
				t.Errorf("ParseString() = %v; want %v", got, want)
			}
		})
	}
}

func TestStringCompare(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs system.String
		cmp      cmpResult
	}{
		{"Equal", "hello", "hello", equal},
		{"Less", "123", "987", less},
		{"Greater", "987", "123", greater},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Compare(tc.rhs)

			ok := tc.cmp(got)

			if got, want := ok, true; got != want {
				t.Errorf("String.Compare() = %v; want %v", got, want)
			}
		})
	}
}

func TestStringEquivalent(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs system.String
	}{
		{"Equal", "hello", "hello"},
		{"Different Case", "HELLO", "hello"},
		{"Normalized Whitespace", "\t\r\nHELLO ", "   hello "},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Equivalent(tc.rhs)

			if got, want := got, true; got != want {
				t.Errorf("String.Equivalent() = %v; want %v", got, want)
			}
		})
	}
}
