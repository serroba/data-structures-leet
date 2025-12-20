package training

func removeElement(nums []int, val int) int {
	left := 0
	for len(nums) > 0 && left < len(nums) {
		if nums[left] == val {
			nums[0], nums[left] = nums[left], nums[0]
			nums = nums[1:]
			left--
		}
		left++
	}
	return len(nums)
}
