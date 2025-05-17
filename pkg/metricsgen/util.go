package metricsgen

import (
	"strings"
	"unicode"
)

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func OtelStringToCamelCase(input string) string {
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '.' || r == '_' || r == '-' // split on common delimiters
	})

	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}

	return strings.Join(parts, "")
}
