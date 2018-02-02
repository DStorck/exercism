package bob

import (
	"regexp"
	"strings"
)

func hasLetters(s string) bool {
	isAlpha := regexp.MustCompile(`[A-Za-z]+`).MatchString
	return isAlpha(s)
}

func isCaps(s string) bool {
	if strings.ToUpper(s) == s && hasLetters(s) {
		return true
	} else {
		return false
	}
}

func isQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[len(s)-1:] == "?" {
		return true
	} else {
		return false
	}
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isQuestion(remark) {
		if isCaps(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else if isCaps(remark) {
		return "Whoa, chill out!"
	} else if len(remark) == 0 {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
