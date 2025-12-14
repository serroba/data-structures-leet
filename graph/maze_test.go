package graph

import "testing"

func TestMaze_HasPathFromStartToEnd(t *testing.T) {
	tests := []struct {
		name  string
		start Coordinate
		end   Coordinate
		grid  [][]int
		want  bool
	}{
		{
			name:  "simple straight path",
			//  S = start, E = end, 1 = path, 0 = wall
			//  [S, 1, 1, E]
			grid: [][]int{
				{1, 1, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 3},
			want:  true,
		},
		{
			name: "simple vertical path",
			//  [S]
			//  [1]
			//  [1]
			//  [E]
			grid: [][]int{
				{1},
				{1},
				{1},
				{1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{3, 0},
			want:  true,
		},
		{
			name: "L-shaped path",
			//  [S, 0, 0]
			//  [1, 0, 0]
			//  [1, 1, E]
			grid: [][]int{
				{1, 0, 0},
				{1, 0, 0},
				{1, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{2, 2},
			want:  true,
		},
		{
			name: "blocked by wall",
			//  [S, 0, E]
			grid: [][]int{
				{1, 0, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 2},
			want:  false,
		},
		{
			name: "start is a wall",
			grid: [][]int{
				{0, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 2},
			want:  false,
		},
		{
			name: "end is a wall",
			grid: [][]int{
				{1, 1, 0},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 2},
			want:  false,
		},
		{
			name: "start equals end",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 0},
			want:  true,
		},
		{
			name: "maze with dead ends - has path",
			//  [S, 1, 0, 0, 0]
			//  [0, 1, 0, 1, 1]
			//  [0, 1, 1, 1, 0]
			//  [0, 0, 0, 1, E]
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 0, 1, 1},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{3, 4},
			want:  true,
		},
		{
			name: "maze with dead ends - no path",
			//  [S, 1, 0, 0, 0]
			//  [0, 1, 0, 1, 1]
			//  [0, 1, 0, 1, 0]
			//  [0, 0, 0, 1, E]
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 0, 1, 1},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{3, 4},
			want:  false,
		},
		{
			name: "spiral maze",
			//  [S, 1, 1, 1, 1]
			//  [0, 0, 0, 0, 1]
			//  [1, 1, 1, 0, 1]
			//  [1, 0, 0, 0, 1]
			//  [1, 1, 1, 1, E]
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1},
				{1, 1, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{4, 4},
			want:  true,
		},
		{
			name: "isolated islands",
			//  [S, 1, 0, 1, E]
			//  [1, 1, 0, 1, 1]
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1},
			},
			start: Coordinate{0, 0},
			end:   Coordinate{0, 4},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{
				start: tt.start,
				end:   tt.end,
				grid:  tt.grid,
			}
			if got := m.HasPathFromStartToEnd(); got != tt.want {
				t.Errorf("HasPathFromStartToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
