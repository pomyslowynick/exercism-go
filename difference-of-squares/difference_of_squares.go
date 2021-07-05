// diffsquares is package for calculating the sequence of square sum differences
package diffsquares

import "math"

// makeRange is a helper function returning a slice of integers between two arguments
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// SquareOfSum returns a result of squaring the sum of natural numbers up to a given number
func SquareOfSum(input int) int {
	sum := 0
	for _, number := range makeRange(0, input) {
		sum += number
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns the sum of squaring the natural numbers up to a given number
func SumOfSquares(input int) int {
	sum := 0.0
	for _, number := range makeRange(0, input) {
		sum += math.Pow(float64(number), 2)
	}
	return int(sum)
}

// Difference returns the difference (duh) between the SumOfSquares and SquareOfSum for a given
// natural numbers sequence.
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
