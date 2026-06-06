// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym
import "unicode"
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	tla := ""
    runes := []rune(s)
	flag := true
	for _, r := range runes {
		// If flag is true and it's a valid letter, grab it
		if flag && unicode.IsLetter(r) {
			tla += string(unicode.ToUpper(r))
			flag = false // Stop grabbing letters until the next word
		}

		// If we hit a space (or hyphen/punctuation), reset the flag for the next word
		if unicode.IsSpace(r) || r == '-' {
			flag = true
		}
	}
	return tla
}
