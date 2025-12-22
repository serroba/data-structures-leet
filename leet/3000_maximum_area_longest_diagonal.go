package leet

import "math"

func AreaOfMaxDiagonal(dimensions [][]int) int {
	maxDiagonal := 0.0
	ma := 0

	for i := range dimensions {
		diagonal := math.Sqrt(math.Pow(float64(dimensions[i][0]), 2) + math.Pow(float64(dimensions[i][1]), 2))
		if maxDiagonal < diagonal {
			maxDiagonal = diagonal
			ma = dimensions[i][0] * dimensions[i][1]
		}
		if maxDiagonal == diagonal && ma < dimensions[i][0]*dimensions[i][1] {
			ma = dimensions[i][0] * dimensions[i][1]
		}
	}
	return ma
}
