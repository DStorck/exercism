package pangram

import (
	"regexp"
	"strings"
)

func isLetter(s string) bool {
	isAlpha := regexp.MustCompile(`[A-Za-z]`).MatchString
	return isAlpha(s)
}

func IsPangram(s string) bool {
	var letters = map[string]bool{}
	chars := strings.Split(s, "")
	for _, i := range chars {
		if isLetter(i) {
			i = strings.ToLower(i)
			letters[i] = true
		}
	}
	if len(letters) == 26 {
		return true
	}
	return false
}
