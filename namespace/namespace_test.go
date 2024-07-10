package namespace_test

import (
	"reflect"
	"testing"

	fhir "github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/resources/patient"
	"github.com/friendly-fhir/go-fhirpath/namespace"
	fpreflect "github.com/friendly-fhir/go-fhirpath/reflect"
	"github.com/friendly-fhir/go-fhirpath/system"
)

func TestSelect(t *testing.T) {
	testCases := []struct {
		name  string
		input reflect.Type
		want  *namespace.Namespace
	}{
		{
			name:  "FHIR resource",
			input: reflect.TypeOf((*patient.Patient)(nil)),
			want:  namespace.R4,
		}, {
			name:  "FHIR complex type",
			input: reflect.TypeOf((*fhir.Quantity)(nil)),
			want:  namespace.R4,
		}, {
			name:  "FHIR abstract type",
			input: reflect.TypeOf((*fhir.Element)(nil)).Elem(),
			want:  namespace.R4,
		}, {
			name:  "System type",
			input: reflect.TypeOf((*system.String)(nil)).Elem(),
			want:  namespace.System,
		}, {
			name:  "Unknown type",
			input: reflect.TypeOf((*testing.T)(nil)),
			want:  nil,
		}, {
			name:  "Reflect type",
			input: reflect.TypeOf((*fpreflect.ClassInfo)(nil)),
			want:  namespace.Reflect,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := namespace.Select(tc.input, namespace.R4, namespace.System, namespace.Reflect)

			if got != tc.want {
				t.Errorf("namespace.Select() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestNamespaceName(t *testing.T) {
	testCases := []struct {
		name      string
		input     reflect.Type
		namespace *namespace.Namespace
		want      fpreflect.TypeSpecifier
	}{
		{
			name:      "FHIR resource",
			input:     reflect.TypeOf((*patient.Patient)(nil)),
			namespace: namespace.R4,
			want:      "Patient",
		}, {
			name:      "FHIR complex type",
			input:     reflect.TypeOf((*fhir.Quantity)(nil)),
			namespace: namespace.R4,
			want:      "Quantity",
		}, {
			name:      "FHIR abstract type",
			input:     reflect.TypeOf((*fhir.Element)(nil)).Elem(),
			namespace: namespace.R4,
			want:      "Element",
		}, {
			name:      "System type",
			input:     reflect.TypeOf((*system.String)(nil)).Elem(),
			namespace: namespace.System,
			want:      "String",
		}, {
			name:      "Reflect type",
			input:     reflect.TypeOf((*fpreflect.ClassInfo)(nil)),
			namespace: namespace.Reflect,
			want:      "ClassInfo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.namespace.Name(tc.input)

			if got != tc.want {
				t.Errorf("Namespace.Name() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestNamespaceString(t *testing.T) {
	testCases := []struct {
		name      string
		namespace *namespace.Namespace
		want      string
	}{
		{
			name:      "FHIR namespace",
			namespace: namespace.R4,
			want:      "FHIR",
		}, {
			name:      "System namespace",
			namespace: namespace.System,
			want:      "System",
		}, {
			name:      "Reflect namespace",
			namespace: namespace.Reflect,
			want:      "Reflect",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.namespace.String()

			if got != tc.want {
				t.Errorf("Namespace.String() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestNamespaceQualifiedName(t *testing.T) {
	testCases := []struct {
		name      string
		input     reflect.Type
		namespace *namespace.Namespace
		want      fpreflect.TypeSpecifier
	}{
		{
			name:      "FHIR resource",
			input:     reflect.TypeOf((*patient.Patient)(nil)),
			namespace: namespace.R4,
			want:      "FHIR.Patient",
		}, {
			name:      "FHIR complex type",
			input:     reflect.TypeOf((*fhir.Quantity)(nil)),
			namespace: namespace.R4,
			want:      "FHIR.Quantity",
		}, {
			name:      "FHIR abstract type",
			input:     reflect.TypeOf((*fhir.Element)(nil)).Elem(),
			namespace: namespace.R4,
			want:      "FHIR.Element",
		}, {
			name:      "System type",
			input:     reflect.TypeOf((*system.String)(nil)).Elem(),
			namespace: namespace.System,
			want:      "System.String",
		}, {
			name:      "Reflect type",
			input:     reflect.TypeOf((*fpreflect.ClassInfo)(nil)),
			namespace: namespace.Reflect,
			want:      "Reflect.ClassInfo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.namespace.QualifiedName(tc.input)

			if got != tc.want {
				t.Errorf("Namespace.QualifiedName() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestNamespaceContains(t *testing.T) {
	testCases := []struct {
		name      string
		input     reflect.Type
		namespace *namespace.Namespace
		want      bool
	}{
		{
			name:      "FHIR resource in FHIR namespace",
			input:     reflect.TypeOf((*patient.Patient)(nil)),
			namespace: namespace.R4,
			want:      true,
		}, {
			name:      "FHIR complex type in FHIR namespace",
			input:     reflect.TypeOf((*fhir.Quantity)(nil)),
			namespace: namespace.R4,
			want:      true,
		}, {
			name:      "System type in System namespace",
			input:     reflect.TypeOf((*system.String)(nil)).Elem(),
			namespace: namespace.System,
			want:      true,
		}, {
			name:      "FHIR resource in System namespace",
			input:     reflect.TypeOf((*patient.Patient)(nil)),
			namespace: namespace.System,
			want:      false,
		}, {
			name:      "System type in FHIR namespace",
			input:     reflect.TypeOf((*system.String)(nil)).Elem(),
			namespace: namespace.R4,
			want:      false,
		}, {
			name:      "Reflect type in Reflect namespace",
			input:     reflect.TypeOf((*fpreflect.ClassInfo)(nil)),
			namespace: namespace.Reflect,
			want:      true,
		}, {
			name:      "FHIR resource in Reflect namespace",
			input:     reflect.TypeOf((*patient.Patient)(nil)),
			namespace: namespace.Reflect,
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.namespace.Contains(tc.input)

			if got != tc.want {
				t.Errorf("Namespace.Contains() = %v; want %v", got, tc.want)
			}
		})
	}
}
