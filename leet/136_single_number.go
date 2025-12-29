package leet

func singleNumber(nums []int) int {
	seen := map[int]int{}
	for i := range nums {
		seen[nums[i]]++
	}

	for i, v := range seen {
		if v == 1 {
			return i
		}
	}

	return -1
}
