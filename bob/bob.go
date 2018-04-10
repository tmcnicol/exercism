package bob

import (
	// 	"fmt"
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	// Test if empty.
	isEmpty := len(remark) == 0

	// Test if has letters
	reg, _ := regexp.Compile("[a-zA-Z]")
	removedLetters := reg.ReplaceAllString(remark, "")
	hasLetters := remark != removedLetters

	isQuestion := strings.HasSuffix(remark, "?")
	isForceful := strings.ToUpper(remark) == remark

	if isEmpty {
		return "Fine. Be that way!"
	}
	if isQuestion && isForceful && hasLetters {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}

	if isForceful && hasLetters {
		return "Whoa, chill out!"
	}
	if !hasLetters {
		return "Whatever."
	}

	return "Whatever."
}
