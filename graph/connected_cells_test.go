package graph

import "testing"

func TestConnectedCells(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		r    int
		c    int
		want int
	}{
		{
			name: "single cell with 1",
			grid: [][]int{
				{1},
			},
			r: 0, c: 0,
			want: 1,
		},
		{
			name: "single cell with 0",
			grid: [][]int{
				{0},
			},
			r: 0, c: 0,
			want: 0,
		},
		{
			name: "horizontal line of 1s",
			//  [1, 1, 1, 0]
			grid: [][]int{
				{1, 1, 1, 0},
			},
			r: 0, c: 0,
			want: 3,
		},
		{
			name: "vertical line of 1s",
			//  [1]
			//  [1]
			//  [1]
			//  [0]
			grid: [][]int{
				{1},
				{1},
				{1},
				{0},
			},
			r: 0, c: 0,
			want: 3,
		},
		{
			name: "L-shaped region",
			//  [1, 0]
			//  [1, 0]
			//  [1, 1]
			grid: [][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
			r: 0, c: 0,
			want: 4,
		},
		{
			name: "two separate regions start at first",
			//  [1, 1, 0, 1, 1]
			grid: [][]int{
				{1, 1, 0, 1, 1},
			},
			r: 0, c: 0,
			want: 2, // only counts connected region
		},
		{
			name: "start at 0 returns 0",
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			r: 1, c: 1,
			want: 0,
		},
		{
			name: "square block",
			//  [1, 1]
			//  [1, 1]
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			r: 0, c: 0,
			want: 4,
		},
		{
			name: "cross shape",
			//  [0, 1, 0]
			//  [1, 1, 1]
			//  [0, 1, 0]
			grid: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			r: 1, c: 1,
			want: 5,
		},
		{
			name: "out of bounds row",
			grid: [][]int{
				{1, 1},
			},
			r: 5, c: 0,
			want: 0,
		},
		{
			name: "out of bounds col",
			grid: [][]int{
				{1, 1},
			},
			r: 0, c: 5,
			want: 0,
		},
		{
			name: "4x4 grid with 3 islands - start at island A",
			//  [1, 1, 0, 1]   island A: 2 cells (top-left)
			//  [0, 0, 0, 1]   island B: 2 cells (right side)
			//  [1, 1, 0, 0]   island C: 3 cells (bottom-left)
			//  [1, 0, 0, 0]
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			r: 0, c: 0,
			want: 2, // island A only
		},
		{
			name: "4x4 grid with 3 islands - start at island B",
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			r: 0, c: 3,
			want: 2, // island B only
		},
		{
			name: "4x4 grid with 3 islands - start at island C",
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			r: 2, c: 0,
			want: 3, // island C only
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy grid to avoid mutation affecting other tests
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			if got := ConnectedCells(gridCopy, tt.r, tt.c); got != tt.want {
				t.Errorf("ConnectedCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountConnected(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		sr   int
		sc   int
		want int
	}{
		{
			name: "single cell with 1",
			grid: [][]int{
				{1},
			},
			sr: 0, sc: 0,
			want: 1,
		},
		{
			name: "single cell with 0",
			grid: [][]int{
				{0},
			},
			sr: 0, sc: 0,
			want: 0,
		},
		{
			name: "horizontal line of 1s",
			//  [1, 1, 1, 0]
			grid: [][]int{
				{1, 1, 1, 0},
			},
			sr: 0, sc: 0,
			want: 3,
		},
		{
			name: "vertical line of 1s",
			//  [1]
			//  [1]
			//  [1]
			//  [0]
			grid: [][]int{
				{1},
				{1},
				{1},
				{0},
			},
			sr: 0, sc: 0,
			want: 3,
		},
		{
			name: "L-shaped region",
			//  [1, 0]
			//  [1, 0]
			//  [1, 1]
			grid: [][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
			sr: 0, sc: 0,
			want: 4,
		},
		{
			name: "two separate regions start at first",
			//  [1, 1, 0, 1, 1]
			grid: [][]int{
				{1, 1, 0, 1, 1},
			},
			sr: 0, sc: 0,
			want: 2, // only counts connected region
		},
		{
			name: "start at 0 returns 0",
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			sr: 1, sc: 1,
			want: 0,
		},
		{
			name: "square block",
			//  [1, 1]
			//  [1, 1]
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			sr: 0, sc: 0,
			want: 4,
		},
		{
			name: "cross shape",
			//  [0, 1, 0]
			//  [1, 1, 1]
			//  [0, 1, 0]
			grid: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			sr: 1, sc: 1,
			want: 5,
		},
		{
			name: "out of bounds row",
			grid: [][]int{
				{1, 1},
			},
			sr: 5, sc: 0,
			want: 0,
		},
		{
			name: "out of bounds col",
			grid: [][]int{
				{1, 1},
			},
			sr: 0, sc: 5,
			want: 0,
		},
		{
			name: "4x4 grid with 3 islands - start at island A",
			//  [1, 1, 0, 1]   island A: 2 cells (top-left)
			//  [0, 0, 0, 1]   island B: 2 cells (right side)
			//  [1, 1, 0, 0]   island C: 3 cells (bottom-left)
			//  [1, 0, 0, 0]
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			sr: 0, sc: 0,
			want: 2, // island A only
		},
		{
			name: "4x4 grid with 3 islands - start at island B",
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			sr: 0, sc: 3,
			want: 2, // island B only
		},
		{
			name: "4x4 grid with 3 islands - start at island C",
			grid: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 1},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
			sr: 2, sc: 0,
			want: 3, // island C only
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy grid to avoid mutation affecting other tests
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			if got := CountConnected(gridCopy, tt.sr, tt.sc); got != tt.want {
				t.Errorf("CountConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
