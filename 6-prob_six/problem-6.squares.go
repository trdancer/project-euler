package probsix

import (
	"math"
)

func SumSquareDifference(n int) int {

	sum := 0
	sumOfSquares := 0
	for i := range n {
		sum += i + 1
		sumOfSquares += int(math.Pow(float64(i+1), 2))
	}
	squareOfSums := int(math.Pow(float64(sum), 2))

	return squareOfSums - sumOfSquares
}
