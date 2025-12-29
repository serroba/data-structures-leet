package leet

import "math"

func FindClosest(x int, y int, z int) int {
	if x == y || math.Abs(float64(z-x)) == math.Abs(float64(z-y)) {
		return 0
	}

	if math.Abs(float64(z-x)) < math.Abs(float64(z-y)) {
		return 1
	}

	return 2
}
