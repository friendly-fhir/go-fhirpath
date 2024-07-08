package fhirpath

import (
	"time"

	"github.com/friendly-fhir/go-fhirpath/resolver"
	"github.com/friendly-fhir/go-fhirpath/tracer"
)

type Tracer = tracer.Tracer
type Resolver = resolver.Resolver

type EvalOption interface {
	setEvaluate(*evaluateConfig) error
}

type evaluateConfig struct {
	Time     time.Time
	Tracer   Tracer
	Resolver Resolver
}

func (c *evaluateConfig) apply(opts ...EvalOption) error {
	for _, opt := range opts {
		if err := opt.setEvaluate(c); err != nil {
			return err
		}
	}
	return nil
}

type evaluateOption func(cfg *evaluateConfig) error

func (f evaluateOption) setEvaluate(cfg *evaluateConfig) error {
	return f(cfg)
}

var _ EvalOption = (*evaluateOption)(nil)

// WithTime returns an [EvalOption] that configures the evaluator to use the
// specified time for evaluating time-dependent FHIRPath expressions.
//
// By default, without this specified, the current time is used.
func WithTime(time time.Time) EvalOption {
	return evaluateOption(func(cfg *evaluateConfig) error {
		cfg.Time = time
		return nil
	})
}

// WithTracer returns an [EvalOption] that configures the evaluator to use the
// specified tracer for tracing FHIRPath evaluation.
func WithTracer(tracer tracer.Tracer) EvalOption {
	return evaluateOption(func(cfg *evaluateConfig) error {
		cfg.Tracer = tracer
		return nil
	})
}

// WithResolver returns an [EvalOption] that configures the evaluator to
// use the specified resolver for resolving FHIRPath references.
func WithResolver(resolver resolver.Resolver) EvalOption {
	return evaluateOption(func(cfg *evaluateConfig) error {
		cfg.Resolver = resolver
		return nil
	})
}
