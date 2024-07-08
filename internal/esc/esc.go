/*
Package esc provides basic functionality for handling the ESC (escape) sequences
in the FHIRPath grammar.
*/
package esc

import (
	"fmt"
	"strconv"
	"strings"
)

var escaper *strings.Replacer

func init() {
	escaper = strings.NewReplacer(
		`\'`, `\u0027`,
		`\"`, `\u0022`,
		"\\`", `\u0060`,
		"\\r", `\u000d`,
		"\\n", `\u000a`,
		"\\t", `\u0009`,
		"\\f", `\u000c`,
		"\\\\", `\u005c`,
	)
}

func Parse(input string) (string, error) {
	result := input
	result = escaper.Replace(result)
	// Re-escape any remaining quotes, so that Unquote won't fail.
	result = strings.ReplaceAll(result, `"`, `\u0022`)
	result, err := strconv.Unquote(fmt.Sprintf(`"%v"`, result))
	if err != nil {
		return "", err
	}
	return result, nil
}
