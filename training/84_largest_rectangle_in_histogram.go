package training

func largestRectangleArea(heights []int) int {
	maxArea := 0
	s := stack{-1}
	for i := 0; i <= len(heights); i++ {
		var currentHeight int
		if i < len(heights) {
			currentHeight = heights[i]
		}
		for s.Len() > 1 && currentHeight <= heights[s.Peek()] {
			h := heights[s.Pop()]
			width := i - s.Peek() - 1
			if h*width > maxArea {
				maxArea = h * width
			}
		}

		s.Push(i)
	}
	return maxArea
}
