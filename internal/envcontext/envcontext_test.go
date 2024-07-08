package envcontext_test

import (
	"context"
	"testing"

	"github.com/friendly-fhir/go-fhirpath/internal/envcontext"
)

func TestLookup(t *testing.T) {
	testCases := []struct {
		name   string
		ctx    context.Context
		key    string
		want   any
		wantOK bool
	}{
		{
			name:   "Nil context returns nil",
			ctx:    nil,
			key:    "key",
			want:   nil,
			wantOK: false,
		}, {
			name:   "Empty context returns nil",
			ctx:    context.Background(),
			key:    "key",
			want:   nil,
			wantOK: false,
		}, {
			name:   "Context with no entries returns nil",
			ctx:    envcontext.WithEntries(context.Background(), nil),
			key:    "key",
			want:   nil,
			wantOK: false,
		}, {
			name:   "Context with entry returns value",
			ctx:    envcontext.WithEntry(context.Background(), "key", "value"),
			key:    "key",
			want:   "value",
			wantOK: true,
		}, {
			name:   "Context with non-matching entries returns nil",
			ctx:    envcontext.WithEntry(context.Background(), "other-key", "value"),
			key:    "key",
			want:   nil,
			wantOK: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := envcontext.Lookup(tc.ctx, tc.key)

			if got, want := ok, tc.wantOK; got != want {
				t.Fatalf("Lookup() ok = %v; want %v", got, want)
			}

			if got, want := got, tc.want; got != want {
				t.Errorf("Lookup() = %v; want %v", got, want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name string
		ctx  context.Context
		key  string
		want any
	}{
		{
			name: "Nil context returns nil",
			ctx:  nil,
			key:  "key",
			want: nil,
		}, {
			name: "Empty context returns nil",
			ctx:  context.Background(),
			key:  "key",
			want: nil,
		}, {
			name: "Context with entry returns value",
			ctx:  envcontext.WithEntry(context.Background(), "key", "value"),
			key:  "key",
			want: "value",
		}, {
			name: "Context with non-matching entries returns nil",
			ctx:  envcontext.WithEntry(context.Background(), "other-key", "value"),
			key:  "key",
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := envcontext.Get(tc.ctx, tc.key)

			if got, want := got, tc.want; got != want {
				t.Errorf("Get() = %v; want %v", got, want)
			}
		})
	}
}

func TestGetOr(t *testing.T) {
	testCases := []struct {
		name string
		ctx  context.Context
		key  string
		def  any
		want any
	}{
		{
			name: "Nil context returns default",
			ctx:  nil,
			key:  "key",
			def:  "default",
			want: "default",
		}, {
			name: "Empty context returns default",
			ctx:  context.Background(),
			key:  "key",
			def:  "default",
			want: "default",
		}, {
			name: "Context with entry returns value",
			ctx:  envcontext.WithEntry(context.Background(), "key", "value"),
			key:  "key",
			def:  "default",
			want: "value",
		}, {
			name: "Context with non-matching entries returns default",
			ctx:  envcontext.WithEntry(context.Background(), "other-key", "value"),
			key:  "key",
			def:  "default",
			want: "default",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := envcontext.GetOr(tc.ctx, tc.key, tc.def)

			if got, want := got, tc.want; got != want {
				t.Errorf("GetOr() = %v; want %v", got, want)
			}
		})
	}
}

func TestWithEntry(t *testing.T) {
	testCases := []struct {
		name string
		ctx  context.Context
		key  string
		val  any
	}{
		{
			name: "Nil context adds entry",
			ctx:  nil,
			key:  "key",
			val:  "value",
		}, {
			name: "Empty context adds entry",
			ctx:  context.Background(),
			key:  "key",
			val:  "value",
		}, {
			name: "Context with entry adds entry",
			ctx:  envcontext.WithEntry(context.Background(), "other-key", "value"),
			key:  "key",
			val:  "value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := envcontext.WithEntry(tc.ctx, tc.key, tc.val)

			got := envcontext.Get(ctx, tc.key)

			if got, want := got, tc.val; got != want {
				t.Errorf("WithEntry() = %v; want %v", got, want)
			}
		})
	}
}

func TestWithEntries(t *testing.T) {
	testCases := []struct {
		name   string
		ctx    context.Context
		values map[string]any
	}{
		{
			name:   "Nil context adds entries",
			ctx:    nil,
			values: map[string]any{"key": "value"},
		}, {
			name:   "Empty context adds entries",
			ctx:    context.Background(),
			values: map[string]any{"key": "value"},
		}, {
			name:   "Context with entry adds entries",
			ctx:    envcontext.WithEntry(context.Background(), "other-key", "value"),
			values: map[string]any{"key": "value"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := envcontext.WithEntries(tc.ctx, tc.values)

			for key, val := range tc.values {
				got := envcontext.Get(ctx, key)

				if got, want := got, val; got != want {
					t.Errorf("WithEntries() = %v; want %v", got, want)
				}
			}
		})
	}
}
