/*
Package tracer provides FHIRPath tracer implementations that may be used
for the "trace" FHIRPath function.
*/
package tracer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/friendly-fhir/go-fhirpath/collection"
)

// Tracer is an interface for tracing FHIRPath evaluation using the 'trace'
// function.
type Tracer interface {
	// Trace is called when the 'trace' function is invoked in a FHIRPath
	// expression.
	Trace(name string, collection collection.Collection) error
}

// NoopTracer is a Tracer that does nothing.
type NoopTracer struct{}

func (NoopTracer) Trace(string, collection.Collection) error {
	return nil
}

var _ Tracer = (*NoopTracer)(nil)

// JSONTracer is a Tracer that writes the trace output as JSON to an io.Writer.
type JSONTracer struct {
	// Writer is the io.Writer to write the JSON output to.
	Writer io.Writer

	// Indent is a string that is used to indent the JSON output.
	Indent string
}

// NewJSONTracer creates a new JSONTracer that writes to the specified io.Writer.
func (t *JSONTracer) Trace(name string, collection collection.Collection) error {
	bytes, err := json.MarshalIndent(collection, name, t.Indent)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(t.out(), "%s\n", string(bytes))
	return err
}

func (t *JSONTracer) out() io.Writer {
	if t.Writer == nil {
		return os.Stdout
	}
	return t.Writer
}

var _ Tracer = (*JSONTracer)(nil)
