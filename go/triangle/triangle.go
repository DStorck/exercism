package triangle

import "math"

type Kind int

const (
	NaT = iota
	Equ = iota
	Iso = iota
	Sca = iota
)

func KindFromSides(a, b, c float64) Kind {
	if !isTri(a, b, c) {
		return NaT
	}
	if a == b && a == c && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func isTri(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return (a+b >= c) && (a+c >= b) && (b+c >= a) // this really bothers me.  The sum of any two sides should be greater than the third.
}
