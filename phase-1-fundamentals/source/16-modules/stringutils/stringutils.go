package stringutils

/*
	Package stringutils: exported string utility functions.

	EXPORTED vs UNEXPORTED — the Go visibility rule:
	- Identifier starts with UPPERCASE => exported => accessible from any other package
	- Identifier starts with lowercase => unexported => private to this package only

	This is Go's entire visibility system. No private/public/protected keywords.
	Just capitalisation. This feels strange coming from TS (private/public) but
	becomes natural quickly — you can tell visibility at a glance from any identifier.

	In this package:
	- Reverse()     => exported, callable from main
	- CountVowels() => exported, callable from main
	- isVowel()     => unexported, internal helper — invisible outside this package
*/

// Reverse returns the string s reversed character by character.
// Exported — uppercase R means any package importing stringutils can call this.
func Reverse(s string) string {
	// Convert to rune slice to handle multi-byte Unicode characters correctly
	// string indexing in Go gives bytes, not characters — rune handles Unicode safely
	runes := []rune(s)

	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

// CountVowels returns the number of vowels (a, e, i, o, u) in string s.
// Exported — uses the unexported isVowel() helper internally.
func CountVowels(s string) int {
	count := 0
	for _, ch := range s {
		if isVowel(ch) { // calling unexported helper — fine because we're in the same package
			count++
		}
	}
	return count
}

// isVowel reports whether rune ch is an English vowel (case-insensitive).
// Unexported — lowercase i means this is an internal helper.
// main package CANNOT call stringutils.isVowel() — compile error if tried.
func isVowel(ch rune) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
