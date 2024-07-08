package system

// Any is the top-level system type for FHIRPath.
type Any interface {
	isAny()
}

// IsType checks whether a given type string is a valid FHIRPath System type
// name value. This function is case-sensitive.
func IsType(ty string) bool {
	switch ty {
	case "Boolean", "Integer", "Any", "Date", "DateTime", "Decimal", "Quantity",
		"String", "Time":
		return true
	}
	return false
}
