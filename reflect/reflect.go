/*
Package reflect implements §10.2. Reflection from the FHIRPath specification.

The types defined in this package are a subset of the types defined in the
Clinical Query Language (CQL) specification, but are limited solely to the usage
within FHIRPath.

See https://hl7.org/fhirpath/N1/#reflection for more information.
*/
package reflect

// TypeSpecifier represents a reflection of a FHIRPath type specifier.
//
// Type specifier may name either FHIR or System-namespaced types, and may be
// either scalar or list values. In most FHIRPath applications, these will only
// ever be seen as scalar type values corresponding to the underlying resource or
// element. Lists are a byproduct that only present itself during reflection
// operations, but otherwise should never be seen in expressions.
type TypeSpecifier string

// Info represents a reflection of a FHIRPath type.
type Info interface {
	isInfo()
}

// InfoElement represents an element from a reflection of a FHIRPath type.
type InfoElement interface {
	isInfoElement()
}

// SimpleTypeInfo is a Go implementation of §10.2.1 "Primitive Types" from
// the FHIRPath specification.
//
// See: https://hl7.org/fhirpath/N1/#primitive-types
type SimpleTypeInfo struct {
	Namespace string        `json:"namespace" yaml:"namespace" fhirpath:"namespace"`
	Name      string        `json:"name" yaml:"name" fhirpath:"name"`
	BaseType  TypeSpecifier `json:"baseType" yaml:"baseType" fhirpath:"baseType"`
}

func (SimpleTypeInfo) isInfo() {}

var _ Info = (*SimpleTypeInfo)(nil)

// ClassInfoElement is a Go implementation of §10.2.2 "Class Types" from
// the FHIRPath specification.
//
// This is a compound type that is built into larger structures in the
// [ClassInfo] type.
//
// See https://hl7.org/fhirpath/N1/#class-types
type ClassInfoElement struct {
	Name       string        `json:"name" yaml:"name" fhirpath:"name"`
	Type       TypeSpecifier `json:"type" yaml:"type" fhirpath:"type"`
	IsOneBased bool          `json:"isOneBased" yaml:"isOneBased" fhirpath:"isOneBased"`
}

func (ClassInfoElement) isInfoElement() {}

// ClassInfo is a Go implementation of §10.2.2 "Class Types" from
// the FHIRPath specification.
//
// See https://hl7.org/fhirpath/N1/#class-types
type ClassInfo struct {
	Namespace string             `json:"namespace" yaml:"namespace" fhirpath:"namespace"`
	Name      string             `json:"name" yaml:"name" fhirpath:"name"`
	BaseType  TypeSpecifier      `json:"baseType" yaml:"baseType" fhirpath:"baseType"`
	Element   []ClassInfoElement `json:"element" yaml:"element" fhirpath:"element"`
}

func (ClassInfo) isInfo() {}

var _ Info = (*ClassInfo)(nil)

// ListTypeInfo is a Go implementation of §10.2.3 "Collection Types" from
// the FHIRPath specification.
//
// See https://hl7.org/fhirpath/N1/#collection-types
type ListTypeInfo struct {
	ElementType TypeSpecifier `json:"elementType" yaml:"elementType" fhirpath:"elementType"`
}

func (ListTypeInfo) isInfo() {}

var _ Info = (*ListTypeInfo)(nil)

// TupleTypeInfoElement is a Go implementation of §10.2.4 "Anonymous Types" from
// the FHIRPath specification.
//
// This is a compound type that is built into larger structures in the
// [TupleTypeInfo] type.
//
// See https://hl7.org/fhirpath/N1/#anonymous-types
type TupleTypeInfoElement struct {
	Name       string        `json:"name" yaml:"name" fhirpath:"name"`
	Type       TypeSpecifier `json:"type" yaml:"type" fhirpath:"type"`
	IsOneBased bool          `json:"isOneBased" yaml:"isOneBased" fhirpath:"isOneBased"`
}

func (TupleTypeInfoElement) isInfoElement() {}

var _ InfoElement = (*TupleTypeInfoElement)(nil)

// TupleTypeInfo is a Go implementation of §10.2.4 "Anonymous Types" from
// the FHIRPath specification.
//
// See https://hl7.org/fhirpath/N1/#anonymous-types
type TupleTypeInfo struct {
	Element []TupleTypeInfoElement `json:"element" yaml:"element" fhirpath:"element"`
}

func (TupleTypeInfo) isInfo() {}

var _ Info = (*TupleTypeInfo)(nil)
