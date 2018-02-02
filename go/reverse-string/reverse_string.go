package reverse

import "strings"

func String(s string) string {
	var str string
	letters := strings.Split(s, "")
	for i := len(letters) - 1; i >= 0; i-- {
		str += letters[i]
	}
	return str
}
