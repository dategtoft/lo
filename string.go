package lo

import (
	"strings"
	"unicode"
)

// Capitalize returns the string with the first letter uppercased.
func Capitalize(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// CamelCase converts a string to camelCase.
func CamelCase(str string) string {
	words := splitWords(str)
	if len(words) == 0 {
		return ""
	}
	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		result += Capitalize(strings.ToLower(word))
	}
	return result
}

// SnakeCase converts a string to snake_case.
func SnakeCase(str string) string {
	words := splitWords(str)
	lower := make([]string, len(words))
	for i, w := range words {
		lower[i] = strings.ToLower(w)
	}
	return strings.Join(lower, "_")
}

// KebabCase converts a string to kebab-case.
func KebabCase(str string) string {
	words := splitWords(str)
	lower := make([]string, len(words))
	for i, w := range words {
		lower[i] = strings.ToLower(w)
	}
	return strings.Join(lower, "-")
}

// PascalCase converts a string to PascalCase.
func PascalCase(str string) string {
	words := splitWords(str)
	result := ""
	for _, word := range words {
		result += Capitalize(strings.ToLower(word))
	}
	return result
}

// Truncate shortens a string to maxLen characters, appending suffix if truncated.
func Truncate(str string, maxLen int, suffix string) string {
	runes := []rune(str)
	if len(runes) <= maxLen {
		return str
	}
	suffixRunes := []rune(suffix)
	cut := maxLen - len(suffixRunes)
	if cut < 0 {
		cut = 0
	}
	return string(runes[:cut]) + suffix
}

// splitWords splits a string into words by spaces, underscores, hyphens, or camelCase boundaries.
func splitWords(str string) []string {
	var words []string
	var current strings.Builder
	runes := []rune(str)
	for i, r := range runes {
		if r == ' ' || r == '_' || r == '-' {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
		} else if i > 0 && unicode.IsUpper(r) && unicode.IsLower(runes[i-1]) {
			words = append(words, current.String())
			current.Reset()
			current.WriteRune(r)
		} else {
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}
	return words
}
