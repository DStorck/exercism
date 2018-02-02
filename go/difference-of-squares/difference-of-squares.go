package diffsquares

func SquareOfSums(n int) int {
	var sum int
	for i := 1; i < n+1; i++ {
		sum += i
	}
	ans := sum * sum
	return ans
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i < n+1; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	ans := SquareOfSums(n) - SumOfSquares(n)
	return ans
}
