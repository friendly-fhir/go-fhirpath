/*
Package resolver provides an interface for resolving FHIRPath references.
*/
package resolver

import (
	"context"
	"fmt"
)

// Resolver is an interface for resolving FHIRPath references.
type Resolver interface {
	// Resolver a reference to a FHIR resource.
	Resolve(ctx context.Context, reference string) (any, error)

	isResolver()
}

// NoopResolver is a Resolver that does nothing.
type NoopResolver struct {
	BaseResolver
}

// BaseResolver is an embeddable type that implements the [Resolver] interface.
//
// This type must be embedded into a type that wishes to implement the [Resolver]
// interface, as it will enable forward-compatibility, and implement unexported
// functions that are required in the interface.
type BaseResolver struct{}

func (BaseResolver) Resolve(ctx context.Context, reference string) (any, error) {
	return nil, fmt.Errorf("no resolver configured")
}

func (BaseResolver) isResolver() {}
