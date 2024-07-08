package system_test

import (
	"errors"
	"testing"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhirpath/system"
	"github.com/google/go-cmp/cmp"
)

func TestParseInteger(t *testing.T) {
	testCases := []struct {
		input string
		want  system.Integer
	}{
		{"0", 0},
		{"1234", 1234},
		{"-74", -74},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := system.ParseInteger(tc.input)

			if err != nil {
				t.Fatalf("ParseInteger() = %v; want nil", err)
			}

			if got, want := got, tc.want; got != want {
				t.Errorf("ParseInteger() = %v; want %v", got, want)
			}
		})
	}
}

func TestParseInteger_InvalidString_ReturnsParseError(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"bad value"},
		{"0x1234"},
		{"v12345v"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			_, err := system.ParseInteger(tc.input)

			var parseErr *system.ParseError
			ok := errors.As(err, &parseErr)

			if got, want := ok, true; got != want {
				t.Errorf("ParseInteger() = %v; want %v", got, want)
			}
		})
	}
}

func TestMustParseInteger(t *testing.T) {
	testCases := []struct {
		input string
		want  system.Integer
	}{
		{"0", 0},
		{"1234", 1234},
		{"-74", -74},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := system.MustParseInteger(tc.input)

			if got, want := got, tc.want; got != want {
				t.Errorf("MustParseInteger() = %v; want %v", got, want)
			}
		})
	}
}

func TestMustParseInteger_InvalidString_Panics(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"bad value"},
		{"0x1234"},
		{"v12345v"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			defer func() { _ = recover() }()

			system.MustParseInteger(tc.input)

			t.Errorf("MustParseInteger() did not panic")
		})
	}
}

func TestIntegerInt32(t *testing.T) {
	testCases := []struct {
		input string
		want  int32
	}{
		{"0", 0},
		{"1234", 1234},
		{"-74", -74},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			v := system.MustParseInteger(tc.input)

			got := v.Int32()

			if got, want := got, tc.want; got != want {
				t.Errorf("Integer.Int32() = %v; want %v", got, want)
			}
		})
	}
}

func TestIntegerR4(t *testing.T) {
	testCases := []struct {
		input string
		want  *fhir.Integer
	}{
		{"0", &fhir.Integer{Value: 0}},
		{"1234", &fhir.Integer{Value: 1234}},
		{"-74", &fhir.Integer{Value: -74}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			v := system.MustParseInteger(tc.input)

			got := v.R4()

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("Integer.R4() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
