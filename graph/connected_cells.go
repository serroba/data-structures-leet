package graph

import "ds/stack"

func ConnectedCells(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0
	count := 1
	count += ConnectedCells(grid, r+1, c)
	count += ConnectedCells(grid, r-1, c)
	count += ConnectedCells(grid, r, c+1)
	count += ConnectedCells(grid, r, c-1)
	return count
}

type Cell struct {
	r, c int
}

func CountConnected(grid [][]int, sr, sc int) int {
	if sr < 0 || sr >= len(grid) || sc < 0 || sc >= len(grid[0]) {
		return 0
	}
	if grid[sr][sc] == 0 {
		return 0
	}

	s := stack.New(Cell{sr, sc})
	count := 0

	for !s.Empty() {
		cell, _ := s.Pop()
		r, c := cell.r, cell.c
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
			continue
		}
		grid[r][c] = 0
		count++
		s.Push(Cell{r + 1, c})
		s.Push(Cell{r - 1, c})
		s.Push(Cell{r, c + 1})
		s.Push(Cell{r, c - 1})

	}
	return count
}
