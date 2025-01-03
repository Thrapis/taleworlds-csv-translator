package utils

import "unicode"

func IsUpper(r rune) bool {
	return unicode.IsLetter(r) && unicode.IsUpper(r)
}

func IsLower(r rune) bool {
	return unicode.IsLetter(r) && unicode.IsLower(r)
}
