package collection_test

import (
	"testing"

	"github.com/friendly-fhir/go-fhirpath/collection"
	"github.com/friendly-fhir/go-fhirpath/system"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestOf(t *testing.T) {
	sut := collection.Of(system.String("Hello"), system.Integer(42))

	if got, want := len(sut), 2; got != want {
		t.Errorf("Of() = %d; want %d", got, want)
	}
}

func TestOf_BadInput_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Of() did not panic")
		}
	}()

	collection.Of("go string")
}

func TestFromSlice(t *testing.T) {
	want := []any{
		system.String("Hello"),
		system.Integer(42),
	}
	got := collection.FromSlice(want)

	if diff := cmp.Diff(got, collection.Collection(want)); diff != "" {
		t.Errorf("FromSlice() mismatch (-got +want):\n%s", diff)
	}
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		other      collection.Collection
		want       collection.Collection
	}{
		{
			name:       "Joining two empty collections",
			collection: collection.Empty,
			other:      collection.Empty,
			want:       collection.Empty,
		}, {
			name:       "Joining empty collection with non-empty collection",
			collection: collection.Empty,
			other:      collection.Of(system.String("Hello")),
			want:       collection.Of(system.String("Hello")),
		}, {
			name:       "Joining non-empty collection with empty collection",
			collection: collection.Of(system.String("Hello")),
			other:      collection.Empty,
			want:       collection.Of(system.String("Hello")),
		}, {
			name:       "Joining two non-empty collections",
			collection: collection.Of(system.String("Hello")),
			other:      collection.Of(system.Integer(42)),
			want:       collection.Of(system.String("Hello"), system.Integer(42)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Join(tc.collection, tc.other)

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("Join() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSingleton(t *testing.T) {
	testCases := []struct {
		name string
		item any
		want collection.Collection
	}{
		{"Boolean", system.Boolean(true), collection.Of(system.Boolean(true))},
		{"String", system.String("Hello"), collection.Of(system.String("Hello"))},
		{"Integer", system.Integer(42), collection.Of(system.Integer(42))},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Singleton(tc.item)

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("Singleton() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestCollectionIsEmpty(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       bool
	}{
		{"Is not empty", collection.Of(system.Boolean(true)), false},
		{"Is empty", collection.Empty, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.collection.IsEmpty()

			if got != tc.want {
				t.Errorf("Collection.IsEmpty() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestCollectionIsSingleton(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       bool
	}{
		{"Is too large", collection.Of(system.Boolean(true), system.Integer(42)), false},
		{"Is singleton", collection.Of(system.Boolean(true)), true},
		{"Is empty", collection.Empty, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.collection.IsSingleton()

			if got != tc.want {
				t.Errorf("Collection.IsSingleton() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestCollectionLen(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       int
	}{
		{"Empty collection", collection.Empty, 0},
		{"Singleton collection", collection.Of(system.String("Hello")), 1},
		{"Non-empty collection", collection.Of(system.String("Hello"), system.Integer(42)), 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.collection.Len()

			if got != tc.want {
				t.Errorf("Collection.Len() = %d; want %d", got, tc.want)
			}
		})
	}
}

func TestCollectionSingleton(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       any
		wantErr    error
	}{
		{
			name:       "Singleton collection",
			collection: collection.Of(system.String("Hello")),
			want:       system.String("Hello"),
			wantErr:    nil,
		}, {
			name:       "Empty collection",
			collection: collection.Empty,
			want:       nil,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-singleton collection",
			collection: collection.Of(system.String("Hello"), system.Integer(42)),
			want:       nil,
			wantErr:    collection.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.collection.Singleton()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Singleton() error = %v; want %v", got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.Singleton() = %v; want %v", got, want)
			}
		})
	}
}

func TestCollectionSingletonBoolean(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       collection.Collection
		wantErr    error
	}{
		{
			name:       "Singleton boolean collection",
			collection: collection.Of(system.Boolean(true)),
			want:       collection.True,
			wantErr:    nil,
		}, {
			name:       "Empty collection",
			collection: collection.Empty,
			want:       collection.Empty,
			wantErr:    nil,
		}, {
			name:       "Non-singleton collection",
			collection: collection.Of(system.Boolean(true), system.Boolean(false)),
			want:       collection.Empty,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-boolean singleton collection",
			collection: collection.Of(system.Integer(42)),
			want:       collection.True,
			wantErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.collection.SingletonBoolean()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonBoolean() error = %v; want %v", got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonBoolean() = %v; want %v", got, want)
			}
		})
	}
}

func TestCollectionBool(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       bool
		wantErr    error
	}{
		{
			name:       "Singleton boolean collection",
			collection: collection.Of(system.Boolean(true)),
			want:       true,
			wantErr:    nil,
		}, {
			name:       "Empty collection",
			collection: collection.Empty,
			want:       false,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-singleton collection",
			collection: collection.Of(system.Boolean(true), system.Boolean(false)),
			want:       false,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-boolean singleton collection",
			collection: collection.Of(system.Integer(42)),
			want:       true,
			wantErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.collection.Bool()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Bool() error = %v; want %v", got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.Bool() = %v; want %v", got, want)
			}
		})
	}
}

func TestCollectionString(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       string
		wantErr    error
	}{
		{
			name:       "Singleton string collection",
			collection: collection.Of(system.String("Hello")),
			want:       "Hello",
			wantErr:    nil,
		}, {
			name:       "Empty collection",
			collection: collection.Empty,
			want:       "",
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-singleton collection",
			collection: collection.Of(system.String("Hello"), system.String("world")),
			want:       "",
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-string singleton collection",
			collection: collection.Of(system.Integer(42)),
			want:       "",
			wantErr:    collection.ErrNotConvertible,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.collection.String()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.String() error = %v; want %v", got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.String() = %v; want %v", got, want)
			}
		})
	}
}

func TestCollectionInt32(t *testing.T) {
	testCases := []struct {
		name       string
		collection collection.Collection
		want       int32
		wantErr    error
	}{
		{
			name:       "Singleton integer collection",
			collection: collection.Of(system.Integer(42)),
			want:       42,
			wantErr:    nil,
		}, {
			name:       "Empty collection",
			collection: collection.Empty,
			want:       0,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-singleton collection",
			collection: collection.Of(system.Integer(42), system.Integer(43)),
			want:       0,
			wantErr:    collection.ErrNotSingleton,
		}, {
			name:       "Non-integer singleton collection",
			collection: collection.Of(system.String("Hello")),
			want:       0,
			wantErr:    collection.ErrNotConvertible,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.collection.Int32()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Int32() error = %v; want %v", got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.Int32() = %v; want %v", got, want)
			}
		})
	}
}
