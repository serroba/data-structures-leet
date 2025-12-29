package leet

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	for i := range nums {
		if nums[i] == target {
			return i
		}

		if i < len(nums)-1 {
			if nums[i] < target && target < nums[i+1] {
				return i + 1
			}
		} else {
			return i + 1
		}
	}

	return 0
}
