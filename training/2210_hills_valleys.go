package training

func CountHillValley(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	hillsAndValleys := 0
	var nonRepeated []int
	nonRepeated = append(nonRepeated, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		nonRepeated = append(nonRepeated, nums[i])
	}
	for i := 1; i < len(nonRepeated)-1; i++ {
		if nonRepeated[i] > nonRepeated[i-1] && nonRepeated[i] > nonRepeated[i+1] {
			hillsAndValleys++
		} else if nonRepeated[i] < nonRepeated[i-1] && nonRepeated[i] < nonRepeated[i+1] {
			hillsAndValleys++
		}
	}
	return hillsAndValleys
}
