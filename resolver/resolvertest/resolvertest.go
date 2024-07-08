/*
Package resolvertest provides test-doubles for working with the FHIRPath
resolver package.
*/
package resolvertest

import (
	"context"

	"github.com/friendly-fhir/go-fhirpath/resolver"
)

// Return returns a [resolver.Resolver] that always resolves to the specified
// value.
func Return(v any) resolver.Resolver {
	return &fixedResolver{value: v}
}

// Error returns a [resolver.Resolver] that always returns the specified error.
func Error(err error) resolver.Resolver {
	return &fixedResolver{err: err}
}

type fixedResolver struct {
	value any
	err   error
	resolver.BaseResolver
}

func (r *fixedResolver) Resolve(_ context.Context, ref string) (any, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.value, nil
}
