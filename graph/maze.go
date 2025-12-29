package graph

import "ds/stack"

type Coordinate struct {
	r, c int
}
type Maze struct {
	start, end Coordinate
	grid       [][]int
}

func (m Maze) HasPathFromStartToEnd() bool {
	s := stack.New(m.start)
	for !s.Empty() {
		vertex, _ := s.Pop()

		r, c := vertex.r, vertex.c
		if vertex == m.end && m.grid[r][c] == 1 {
			return true
		}

		if r < 0 || r >= len(m.grid) || c < 0 || c >= len(m.grid[0]) || m.grid[r][c] == 0 {
			continue
		}

		m.grid[r][c] = 0
		s.Push(Coordinate{r + 1, c})
		s.Push(Coordinate{r + 1, c})
		s.Push(Coordinate{r, c + 1})
		s.Push(Coordinate{r, c - 1})
	}

	return false
}
