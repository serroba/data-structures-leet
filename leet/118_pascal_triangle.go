package leet

func Generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		level := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				level[j] = 1
			} else {
				level[j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}

		triangle = append(triangle, level)
	}

	return triangle
}
