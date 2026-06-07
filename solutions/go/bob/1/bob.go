// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey determines Bob's response based on the formatting of the input remark.
func Hey(remark string) string {
	// Remove all leading and trailing whitespace (spaces, tabs, newlines)
	remark = strings.TrimSpace(remark)

	// Rule 1: Silence
	if remark == "" {
		return "Fine. Be that way!"
	}

	// Check if the remark ends with a question mark
	isQuestion := strings.HasSuffix(remark, "?")

	// Check if the user is yelling (contains letters, and all letters are uppercase)
	hasLetters := strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	isYelling := hasLetters && remark == strings.ToUpper(remark)

	// Rule 2: Yelled Question
	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	// Rule 3: Just Yelling
	if isYelling {
		return "Whoa, chill out!"
	}

	// Rule 4: Just a Question
	if isQuestion {
		return "Sure."
	}

	// Rule 5: Anything else
	return "Whatever."
}
