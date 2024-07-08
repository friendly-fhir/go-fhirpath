package system_test

import (
	"errors"
	"testing"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhirpath/system"
	"github.com/google/go-cmp/cmp"
)

func TestParseBoolean(t *testing.T) {
	testCases := []struct {
		input string
		want  system.Boolean
	}{
		{"false", false},
		{"true", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := system.ParseBoolean(tc.input)
			if err != nil {
				t.Fatalf("ParseBoolean() = %v; want nil", err)
			}

			if got, want := got, tc.want; got != want {
				t.Errorf("ParseBoolean() = %v; want %v", got, want)
			}

		})
	}
}

func TestParseBoolean_InvalidString_ReturnsParseError(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"bad value"},
		{"b"},
		{"TRUE"},
		{"1"},
		{"FALSE"},
		{"0"},
		{"False"},
		{"True"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			_, err := system.ParseBoolean(tc.input)

			var parseErr *system.ParseError
			ok := errors.As(err, &parseErr)

			if got, want := ok, true; got != want {
				t.Errorf("ParseBoolean() = %v; want %v", got, want)
			}
		})
	}
}

func TestMustParseBoolean(t *testing.T) {
	testCases := []struct {
		input string
		want  system.Boolean
	}{
		{"false", false},
		{"true", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := system.MustParseBoolean(tc.input)

			if got, want := got, tc.want; got != want {
				t.Errorf("ParseBoolean() = %v; want %v", got, want)
			}
		})
	}
}

func TestMustParseBoolean_InvalidString_Panics(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"bad value"},
		{"b"},
		{"TRUE"},
		{"1"},
		{"FALSE"},
		{"0"},
		{"False"},
		{"True"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			defer func() { _ = recover() }()

			system.MustParseBoolean(tc.input)

			t.Errorf("MustParseBoolean() = want panic")
		})
	}
}

func TestBooleanBool(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"false", false},
		{"true", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			v := system.MustParseBoolean(tc.input)

			got := v.Bool()

			if got, want := got, tc.want; got != want {
				t.Errorf("Boolean.Bool() = %v; want %v", got, want)
			}
		})
	}
}

func TestBooleanR4(t *testing.T) {
	testCases := []struct {
		input string
		want  *fhir.Boolean
	}{
		{"false", &fhir.Boolean{Value: false}},
		{"true", &fhir.Boolean{Value: true}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			v := system.MustParseBoolean(tc.input)

			got := v.R4()

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("Boolean.ToBoolean() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
