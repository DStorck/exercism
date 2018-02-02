package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	words := regexp.MustCompile("[- ]").Split(s, -1)
	var acr string = ""
	for _, word := range words {
		acr += word[0:1]
	}
	acr = strings.ToUpper(acr)
	return acr
}
