/*
Package envcontext provides definitions for Context objects that document
environment variables, as defined in FHIRPath.
*/
package envcontext

import "context"

type exprKey struct{}
type exprEntries map[string]any

// Lookup retrieves a value from the context by name.
func Lookup(ctx context.Context, name string) (any, bool) {
	if ctx == nil {
		return nil, false
	}
	value := ctx.Value(exprKey{})
	if value == nil {
		return nil, false
	}

	// This case should never logically occur, but is added as a precaution.
	entries := value.(exprEntries)
	if entries == nil {
		return nil, false
	}
	result, ok := entries[name]
	return result, ok
}

// Get retrieves a value from the context by name.
func Get(ctx context.Context, name string) any {
	value, _ := Lookup(ctx, name)
	return value
}

// GetOr retrieves a value from the context by name, or returns a default value.
func GetOr(ctx context.Context, name string, def any) any {
	value, ok := Lookup(ctx, name)
	if !ok {
		return def
	}
	return value
}

// WithEntry adds a single environment value by name to the context.
func WithEntry(ctx context.Context, name string, value any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	exprCtx := ctx.Value(exprKey{})
	if exprCtx == nil {
		exprCtx = exprEntries{}
	}
	entries := exprCtx.(exprEntries)
	entries[name] = value
	return context.WithValue(ctx, exprKey{}, entries)
}

// WithEntries adds multiple environment values by name to the context.
func WithEntries(ctx context.Context, values map[string]any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	exprCtx := ctx.Value(exprKey{})
	if exprCtx == nil {
		exprCtx = exprEntries{}
	}
	entries := exprCtx.(exprEntries)
	if entries == nil {
		entries = exprEntries{}
	}

	for key, value := range values {
		entries[key] = value
	}
	return context.WithValue(ctx, exprKey{}, entries)
}
