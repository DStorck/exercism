package grains

import (
	"errors"
	"math"
)

var points = map[string]int{}

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("error message")
	} else {
		power := float64(n - 1)
		exp := math.Pow(float64(2), power)
		ans := uint64(exp)
		return ans, nil
	}
}

func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		temp, _ := Square(i)
		total += uint64(temp)
	}
	return total
}
