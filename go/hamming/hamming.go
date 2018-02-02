package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Not the same distance")
	}
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")
	var count int
	for i := range aArr {
		if aArr[i] != bArr[i] {
			count += 1
		}
	}
	return count, nil
}
