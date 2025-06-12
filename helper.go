package main

import (
	"strings"
	"unicode"
)

func matchCase(original, replacement string) string {
	if isUpper(original) {
		return strings.ToUpper(replacement)
	}
	if isCapitalized(original) {
		return strings.Title(replacement)
	}
	return replacement
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func isCapitalized(word string) bool {
	if len(word) == 0 {
		return false
	}
	return unicode.IsUpper(rune(word[0]))
}

func stripPunctuation(word string) (prefix, base, suffix string) {
	start := 0
	end := len(word)
	for start < len(word) && !unicode.IsLetter(rune(word[start])) {
		start++
	}
	for end > start && !unicode.IsLetter(rune(word[end-1])) {
		end--
	}
	return word[:start], word[start:end], word[end:]
}
