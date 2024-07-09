package namespace

import (
	"fmt"
	"reflect"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhirpath/system"
)

var (
	// R4 is the FHIR namespace that contains all R4 FHIR definitions.
	R4 = New("FHIR", fhirNamer, r4Element, r4Resource, r4Domain, r4Backbone)

	// System is the system namespace that contains all system types.
	System = New("System", systemNamer, systemAny)
)

// Select returns the namespace that contains the specified type. If no
// namespace contains the type, then nil is returned.
func Select(t reflect.Type, namespaces ...*Namespace) *Namespace {
	for _, ns := range namespaces {
		if ns.Contains(t) {
			return ns
		}
	}
	return nil
}

// Namer is an interface that can be used to generate names for types.
type Namer interface {
	// Name returns the name of the type.
	Name(reflect.Type) string
}

// NamerFunc is a convenience type that implements the Namer interface. This
// simplifies constructing a function that acts as a name function.
type NamerFunc func(reflect.Type) string

// Name calls the underlying function to generate the name.
func (f NamerFunc) Name(t reflect.Type) string {
	return f(t)
}

var _ Namer = (*NamerFunc)(nil)

// Namespace represents a FHIR model namespace. This is used to group types
// together and provide a common name for them. This leverages interfaces for
// containment.
type Namespace struct {
	namer      Namer
	name       string
	interfaces []reflect.Type
}

// New constructs a new namespace with the specified name, namer, and interfaces.
func New(name string, namer Namer, ifaces ...reflect.Type) *Namespace {
	return &Namespace{
		name:       name,
		namer:      namer,
		interfaces: ifaces,
	}
}

// String returns the name of this namespace as a string.
func (n *Namespace) String() string {
	return n.name
}

// QualifiedName returns the qualified name of the type in this namespace.
// This is a namespace-prefixed type-name.
func (n *Namespace) QualifiedName(t reflect.Type) string {
	return fmt.Sprintf("%s.%s", n.name, n.namer.Name(t))
}

// Name returns the name of the type as it appears in the namespace. This will
// not include the namespace prefix.
func (n *Namespace) Name(t reflect.Type) string {
	return n.namer.Name(t)
}

// Contains returns true if the namespace contains the specified type.
func (n *Namespace) Contains(t reflect.Type) bool {
	for _, iface := range n.interfaces {
		if t.Implements(iface) {
			return true
		}
	}
	return false
}

var (
	systemAny  = reflect.TypeOf((*system.Any)(nil)).Elem()
	r4Element  = reflect.TypeOf((*fhir.Element)(nil)).Elem()
	r4Backbone = reflect.TypeOf((*fhir.BackboneElement)(nil)).Elem()
	r4Resource = reflect.TypeOf((*fhir.Resource)(nil)).Elem()
	r4Domain   = reflect.TypeOf((*fhir.DomainResource)(nil)).Elem()

	fhirNamer = NamerFunc(func(t reflect.Type) string {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		name := t.Name()
		// TODO: primitives should be in camelCase instead of PascalCase.
		return name
	})
	systemNamer = NamerFunc(func(t reflect.Type) string {
		return t.Name()
	})
)
