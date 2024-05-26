package utils

import "unicode"

func CamelToPascal(s string) string {
	if s == "" {
		return ""
	}

	asRunes := []rune(s)
	asRunes[0] = unicode.ToUpper(asRunes[0])
	return string(asRunes)
}
