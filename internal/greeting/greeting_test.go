package greeting_test

import (
	"testing"

	"github.com/friendly-fhir/go-template/internal/greeting"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	want := "Hello, world!"

	got := greeting.New()

	if !cmp.Equal(got, want) {
		t.Errorf("greeting.New() = %q, want %q", got, want)
	}
}
