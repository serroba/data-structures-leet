package leet

import "math"

func MaxSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxNonPositive := math.MinInt
	seen := map[int]bool{}

	var hasNegative, hasPositive bool

	total := 0

	for _, num := range nums {
		if num < 0 {
			if num > maxNonPositive {
				maxNonPositive = num
				hasNegative = true
			}

			continue
		} else {
			if !seen[num] {
				seen[num] = true
				total += num
			}

			hasPositive = true
		}

		if num == 0 {
			continue
		}

		seen[num] = true
	}

	if hasNegative && !hasPositive {
		return total + maxNonPositive
	}

	return total
}
