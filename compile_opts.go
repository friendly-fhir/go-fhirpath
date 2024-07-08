package fhirpath

type CompileOption interface {
	setCompile(*compileConfig) error
}

type compileConfig struct {
	Text string
}

func (c *compileConfig) apply(opts ...CompileOption) error {
	for _, opt := range opts {
		if err := opt.setCompile(c); err != nil {
			return err
		}
	}
	return nil
}

type compileOption func(cfg *compileConfig) error

func (f compileOption) setCompile(cfg *compileConfig) error {
	return f(cfg)
}

var _ CompileOption = (*compileOption)(nil)

// N1 returns a [CompileOption] that configures the compiler to use the FHIR N1
// version of the FHIRPath language.
//
// Only one of N1 or [N2] may be specified at a time.
func N1() CompileOption {
	return notImplemented{}
}

// N2 returns a [CompileOption] that configures the compiler to use the FHIR N2
// version of the FHIRPath language.
//
// Only one of [N1] or N2 may be specified at a time.
func N2() CompileOption {
	return notImplemented{}
}

// R4 returns a [CompileOption] that configures the compiler to use the FHIR R4
// version of the FHIRPath language.
func R4() CompileOption {
	return notImplemented{}
}

// SimpleR4 returns a [CompileOption] that configures the compiler to use the
// FHIR R4 version of the FHIRPath language without any additional features.
func SimpleR4() CompileOption {
	return notImplemented{}
}

// AddFunc returns a [CompileOption] that adds a custom function to the
// FHIRPath compiler.
//
// fn must be a valid function that takes a [collection.Collection], input
// parameters, and returns a [collection.Collection] and an [error]. Parameters
// must be a conforming FHIRPath type; either from the base FHIR instance, or
// from the [system] package.
func AddFunc(name string, fn any) CompileOption {
	return notImplemented{}
}

// AddFuncs returns a [CompileOption] that adds custom functions to the FHIRPath
// compiler.
//
// Supplied functions must be a valid function that takes a
// [collection.Collection], input parameters, and returns a
// [collection.Collection] and an [error]. Parameters must be a conforming
// FHIRPath type; either from the base FHIR instance, or from the [system] package.
func AddFuncs(funcs map[string]any) CompileOption {
	return notImplemented{}
}

type notImplemented struct{}

func (s notImplemented) setCompile(*compileConfig) error {
	return ErrUnimplemented
}

func (s notImplemented) setEvaluate(*evaluateConfig) error {
	return ErrUnimplemented
}

var (
	_ CompileOption = (*notImplemented)(nil)
	_ EvalOption    = (*notImplemented)(nil)
)
