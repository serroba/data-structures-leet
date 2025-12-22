package leet

func CountHillValley(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	count := 0
	lp := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		rp := nums[i+1]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != nums[i] {
				rp = nums[j]
				break
			}
		}
		if nums[i] > lp && nums[i] > rp {
			count++
		} else if nums[i] < lp && nums[i] < rp {
			count++
		}
		if lp != nums[i] {
			lp = nums[i]
		}
	}

	return count
}
