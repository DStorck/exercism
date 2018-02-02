package strand

import "strings"

var conversions = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	convertedString := ""
	letters := strings.Split(dna, "")
	for _, letter := range letters {
		convertedString += conversions[letter]
	}
	return convertedString
}
