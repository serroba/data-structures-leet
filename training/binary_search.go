package training

func binarySearch(nums []int, target int) int {
	lowEnd := 0
	upperEnd := len(nums) - 1
	for lowEnd < upperEnd {
		middle := lowEnd + (upperEnd-lowEnd)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			upperEnd = middle - 1
		} else {
			lowEnd = middle + 1
		}
	}
	return -1
}
