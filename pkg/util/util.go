package util

import (
	"regexp"
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
		return r == '.' || r == '_' || r == '-'
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

func OtelStringToCamelCaseField(input string) string {
	str := OtelStringToCamelCase(input)
	if len(str) == 0 {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func OtelStringToPromLabel(input string) string {
	var b strings.Builder
	for _, r := range input {
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
		} else if r == '.' || r == '-' {
			b.WriteRune('_')
		} else if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		}
	}
	label := b.String()
	label = strings.Trim(label, "_")
	return label
}

func ValueTypeToAttributeConstructor(input string) string {
	switch input {
	case "int":
		return "Int"
	case "int64":
		return "Int64"
	case "bool":
		return "Bool"
	case "float64":
		return "Float64"
	case "string":
		return "String"
	case "[]int":
		return "IntSlice"
	case "[]int64":
		return "Int64Slice"
	case "[]float64":
		return "Float64Slice"
	case "[]bool":
		return "BoolSlice"
	case "[]string":
		return "StringSlice"
	}
	panic("unkown value type")
}

func HasDuplicateStrings(slice []string) bool {
	seen := make(map[string]struct{})
	for _, s := range slice {
		if _, exists := seen[s]; exists {
			return true
		}
		seen[s] = struct{}{}
	}
	return false
}

func MarkdownLinkAnchor(header string) string {
	anchor := strings.ToLower(header)
	// Remove all non-alphanumeric characters except hyphens and spaces
	re := regexp.MustCompile(`[^a-z0-9 -]`)
	anchor = re.ReplaceAllString(anchor, "")
	// Replace spaces with hyphens
	anchor = strings.ReplaceAll(anchor, " ", "-")
	return anchor
}
