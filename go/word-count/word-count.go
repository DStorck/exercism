package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func returnValidWord(word string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9']+")
	valid_word := reg.ReplaceAllString(word, "")
	return valid_word
}

func WordCount(str string) Frequency {
	freq := make(Frequency)
	str = strings.TrimSpace(str)
	words := regexp.MustCompile("[#;, ]+").Split(str, -1)
	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.TrimSpace(word)
		word = returnValidWord(word)
		freq[word] += 1
	}
	return freq
}
