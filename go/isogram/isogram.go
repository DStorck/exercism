package isogram

import (
	"regexp"
	"strings"
)

func isLetter(s string) bool {
	isAlpha := regexp.MustCompile(`[A-Za-z]`).MatchString
	return isAlpha(s)
}

func IsIsogram(s string) bool {
	if s == "" {
		return true
	}
	letters := make(map[string]bool)
	word := strings.ToLower(s)
	chars := strings.Split(word, "")
	for _, letter := range chars {
		if isLetter(letter) {
			if letters[letter] != false {
				return false
			} else {
				letters[letter] = true
			}
		}
	}

	return true
}
