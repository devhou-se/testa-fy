package main

import (
	"strings"
)

var ignoreWords = map[string]struct{}{
	"i": {}, "am": {}, "i'm": {}, "me": {}, "be": {}, "see": {}, "run": {},
	"and": {}, "or": {}, "the": {}, "a": {}, "an": {}, "to": {}, "in": {}, "on": {}, "also": {},
	"that": {}, "is": {}, "it": {}, "of": {}, "by": {}, "we": {}, "are": {}, "were": {},
	"he": {}, "she": {}, "at": {}, "if": {}, "so": {}, "do": {}, "but": {}, "you": {}, "my": {},
}

// only true exceptions left here
var specialCases = map[string]string{
	"authenticity":  "authentestity",
	"action":        "action",
	"tickle":        "test-tickle",
	"demonstration": "damianstration",
	"Damian Testa":  "Damian Testa AKA Pooplord 5000",
}

func Testafy(input string) string {
	// phrase-level override
	input = strings.Replace(
		input,
		"Damian Testa",
		"Damian Testa AKA Pooplord 5000",
		-1,
	)

	words := strings.Fields(input)

	for i, word := range words {
		prefix, base, suffix := stripPunctuation(word)
		lower := strings.ToLower(base)

		if _, skip := ignoreWords[lower]; skip || len(lower) < 4 {
			continue
		}

		// explicit overrides first
		if rep, ok := specialCases[lower]; ok {
			words[i] = prefix + matchCase(base, rep) + suffix
			continue
		}

		// pattern-based fallbacks
		modified := applyPatterns(lower)
		if modified != lower {
			words[i] = prefix + matchCase(base, modified) + suffix
		}
	}

	return strings.Join(words, " ")
}

func applyPatterns(word string) string {
	// always map "trans..." → "test..."
	if strings.HasPrefix(word, "trans") {
		return "test" + word[len("trans"):]
	}
	switch {
	case strings.HasSuffix(word, "ction"):
		return strings.TrimSuffix(word, "ction") + "ctestation"
	case strings.HasSuffix(word, "ation"):
		return strings.TrimSuffix(word, "ation") + "estation"
	case strings.HasSuffix(word, "ment"):
		return strings.TrimSuffix(word, "ment") + "testment"
	case strings.HasSuffix(word, "east"):
		return strings.TrimSuffix(word, "east") + "easta"
	case strings.HasSuffix(word, "est"):
		return word + "a"
	case strings.HasSuffix(word, "ane"):
		return strings.TrimSuffix(word, "ane") + "anian"
	case strings.HasSuffix(word, "ain"):
		return strings.TrimSuffix(word, "ain") + "ainian"
	case strings.HasSuffix(word, "ame"):
		return strings.TrimSuffix(word, "ame") + "amian"
	case strings.HasSuffix(word, "ess"):
		return strings.TrimSuffix(word, "ess") + "essta"
	case strings.HasSuffix(word, "ster"):
		return strings.TrimSuffix(word, "ster") + "sta"
	}

	return word
}
