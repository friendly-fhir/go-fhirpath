package fhirpath

import (
	"context"
	"fmt"

	"github.com/friendly-fhir/go-fhirpath/collection"
	"github.com/friendly-fhir/go-fhirpath/system"
)

// Collection is a collection of FHIRPath values.
type Collection = collection.Collection

// Path represents a compiled FHIRPath expression.
type Path struct {
	path string
}

// Compile compiles the FHIRPath expression and returns a Path object. If the
// expression is invalid, an error is returned.
//
// Compilation will always use the latest version of the FHIRPath language,
// but may be configured with options to enable other language features.
func Compile(path string, opts ...CompileOption) (*Path, error) {
	var cfg compileConfig
	if err := cfg.apply(opts...); err != nil {
		return nil, err
	}
	return nil, ErrUnimplemented
}

// MustCompile is a convenience function that compiles the FHIRPath expression
// and panics if the expression is invalid.
func MustCompile(path string, opts ...CompileOption) *Path {
	result, err := Compile(path, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// Eval evaluates the FHIRPath expression and returns the result as a
// collection of values. If the expression is invalid, an error is returned.
// The resource argument is the FHIR resource to evaluate the expression against.
func (p *Path) Eval(ctx context.Context, resource any, opts ...EvalOption) (Collection, error) {
	var cfg evaluateConfig
	if err := cfg.apply(opts...); err != nil {
		return nil, err
	}

	return nil, ErrUnimplemented
}

// MustEval is a convenience function that evaluates the FHIRPath expression
// and panics if the expression is invalid.
func (p *Path) MustEval(ctx context.Context, resource any, opts ...EvalOption) Collection {
	result, err := p.Eval(ctx, resource, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// EvalBool evaluates the FHIRPath expression and returns the result as a
// boolean value. If the result is a singleton, it will be converted to a
// truthy boolean value. If
func (p *Path) EvalBool(ctx context.Context, resource any, opts ...EvalOption) (bool, error) {
	result, err := p.Eval(ctx, resource, opts...)
	if err != nil {
		return false, err
	}

	if result.IsEmpty() {
		return false, nil
	}
	singleton, err := result.Singleton()
	if err != nil {
		return false, err
	}

	if val, ok := system.Normalize(singleton).(system.Boolean); ok {
		return val.Bool(), nil
	}
	return true, nil
}

// MustEvalBool is a convenience function that evaluates the FHIRPath expression
// and panics if the expression is invalid.
func (p *Path) MustEvalBool(ctx context.Context, resource any, opts ...EvalOption) bool {
	result, err := p.EvalBool(ctx, resource, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// EvalString evaluates the FHIRPath expression and returns the result as a
// string value. If the result is a singleton, it will be converted to a string.
// If the result is not a string, an error is returned.
func (p *Path) EvalString(ctx context.Context, resource any, opts ...EvalOption) (string, error) {
	result, err := p.Eval(ctx, resource, opts...)
	if err != nil {
		return "", err
	}

	singleton, err := result.Singleton()
	if err != nil {
		return "", err
	}

	if val, ok := system.Normalize(singleton).(system.String); ok {
		return val.String(), nil
	}
	return "", fmt.Errorf("expected string result, got %T", result[0])
}

// MustEvalString is a convenience function that evaluates the FHIRPath expression
// and panics if the expression is invalid.
func (p *Path) MustEvalString(ctx context.Context, resource any, opts ...EvalOption) string {
	result, err := p.EvalString(ctx, resource, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// EvalFloat64 evaluates the FHIRPath expression and returns the result as a
// float64 value. If the result is a singleton, it will be converted to a float64.
func (p *Path) EvalFloat64(ctx context.Context, resource any, opts ...EvalOption) (float64, error) {
	result, err := p.Eval(ctx, resource, opts...)
	if err != nil {
		return 0, err
	}

	singleton, err := result.Singleton()
	if err != nil {
		return 0, err
	}

	if val, ok := system.Normalize(singleton).(system.Decimal); ok {
		return val.Float64(), nil
	}
	return 0, fmt.Errorf("expected number result, got %T", result[0])
}

// MustEvalFloat64 is a convenience function that evaluates the FHIRPath expression
// and panics if the expression is invalid.
func (p *Path) MustEvalFloat64(ctx context.Context, resource any, opts ...EvalOption) float64 {
	result, err := p.EvalFloat64(ctx, resource, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// String returns the FHIRPath expression as a string.
func (p *Path) String() string {
	if p == nil {
		return ""
	}
	return p.path
}

var _ fmt.Stringer = (*Path)(nil)

func (p *Path) Equal(other *Path) bool {
	if p == nil || other == nil {
		return p == other
	}
	return p.path == other.path
}
