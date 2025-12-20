package training

func maxAreaFun(height []int) int {
	area := 0
	lp := 0
	rp := len(height) - 1
	for lp < rp {
		currentArea := min(height[lp], height[rp]) * (rp - lp)
		if currentArea > area {
			area = currentArea
		}
		if height[lp] < height[rp] {
			lp++
		} else {
			rp--
		}
	}
	return area
}
