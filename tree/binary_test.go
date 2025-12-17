package tree

import (
	"cmp"
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	type args[T cmp.Ordered] struct {
		node *Node[T]
		val  T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			args: args[int]{
				node: nil,
				val:  5,
			},
			want: false,
		},
		{
			name: "single node",
			args: args[int]{
				node: &Node[int]{val: 1},
				val:  1,
			},
			want: true,
		},
		{
			name: "multiple nodes",
			args: args[int]{
				node: &Node[int]{left: &Node[int]{val: 1}, right: &Node[int]{val: 10}, val: 5},
				val:  1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.node, tt.args.val); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllBFS(t *testing.T) {
	type args[T cmp.Ordered] struct {
		node *Node[T]
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			args: args[int]{
				node: nil,
			},
			want: []int{},
		},
		{
			name: "single node",
			args: args[int]{
				node: &Node[int]{val: 1},
			},
			want: []int{1},
		},
		{
			name: "multiple nodes",
			args: args[int]{
				node: &Node[int]{left: &Node[int]{val: 1}, right: &Node[int]{val: 10}, val: 5},
			},
			want: []int{5, 1, 10},
		},
		{
			name: "left-only child",
			args: args[int]{
				node: &Node[int]{val: 5, left: &Node[int]{val: 3}},
			},
			want: []int{5, 3},
		},
		{
			name: "right-only child",
			args: args[int]{
				node: &Node[int]{val: 5, right: &Node[int]{val: 7}},
			},
			want: []int{5, 7},
		},
		{
			name: "left-skewed tree",
			args: args[int]{
				node: &Node[int]{val: 5, left: &Node[int]{val: 3, left: &Node[int]{val: 1}}},
			},
			want: []int{5, 3, 1},
		},
		{
			name: "right-skewed tree",
			args: args[int]{
				node: &Node[int]{val: 5, right: &Node[int]{val: 7, right: &Node[int]{val: 9}}},
			},
			want: []int{5, 7, 9},
		},
		{
			name: "multi-level tree (BFS order)",
			args: args[int]{
				node: &Node[int]{
					val:   5,
					left:  &Node[int]{val: 3, left: &Node[int]{val: 1}, right: &Node[int]{val: 4}},
					right: &Node[int]{val: 7, right: &Node[int]{val: 9}},
				},
			},
			want: []int{5, 3, 7, 1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree[int]{
				root: tt.args.node,
			}
			if got := tree.FindAllBFS(); !slices.Equal(got, tt.want) {
				t.Errorf("FindAllBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllDFS(t *testing.T) {
	type args[T cmp.Ordered] struct {
		node *Node[T]
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Test on empty tree",
		},
		{
			name: "Test on single node",
			args: args[int]{
				node: &Node[int]{val: 1},
			},
			want: []int{1},
		},
		{
			name: "Test on two levels tree",
			args: args[int]{
				node: &Node[int]{left: &Node[int]{val: 1}, right: &Node[int]{val: 10}, val: 5},
			},
			want: []int{5, 1, 10},
		},
		{
			name: "Test on multiple levels tree",
			args: args[int]{
				node: &Node[int]{left: &Node[int]{val: 3, left: &Node[int]{val: 1}, right: &Node[int]{val: 2}}, right: &Node[int]{val: 10}, val: 5},
			},
			want: []int{5, 3, 1, 2, 10},
		},
		{
			name: "Test on left-only child",
			args: args[int]{
				node: &Node[int]{val: 5, left: &Node[int]{val: 3}},
			},
			want: []int{5, 3},
		},
		{
			name: "Test on right-only child",
			args: args[int]{
				node: &Node[int]{val: 5, right: &Node[int]{val: 7}},
			},
			want: []int{5, 7},
		},
		{
			name: "Test on left-skewed tree",
			args: args[int]{
				node: &Node[int]{val: 5, left: &Node[int]{val: 3, left: &Node[int]{val: 1}}},
			},
			want: []int{5, 3, 1},
		},
		{
			name: "Test on right-skewed tree",
			args: args[int]{
				node: &Node[int]{val: 5, right: &Node[int]{val: 7, right: &Node[int]{val: 9}}},
			},
			want: []int{5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree[int]{
				root: tt.args.node,
			}
			if got := tree.FindAllDFS(); !slices.Equal(got, tt.want) {
				t.Errorf("FindAllDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPathSum(t *testing.T) {
	type args struct {
		t        *Node[int]
		target   int
		sumSoFar int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty tree",
			args: args{
				t:        nil,
				target:   0,
				sumSoFar: 0,
			},
			want: false,
		},
		{
			name: "single node - matches",
			args: args{
				t:        &Node[int]{val: 5},
				target:   5,
				sumSoFar: 0,
			},
			want: true,
		},
		{
			name: "single node - does not match",
			args: args{
				t:        &Node[int]{val: 5},
				target:   6,
				sumSoFar: 0,
			},
			want: false,
		},
		{
			name: "two level - left leaf matches",
			args: args{
				t: &Node[int]{
					val:   5,
					left:  &Node[int]{val: 4},
					right: &Node[int]{val: 8},
				},
				target:   9, // 5 -> 4
				sumSoFar: 0,
			},
			want: true,
		},
		{
			name: "two level - internal sum matches but not a leaf path",
			args: args{
				t: &Node[int]{
					val:   5,
					left:  &Node[int]{val: 4, left: &Node[int]{val: 1}},
					right: &Node[int]{val: 8},
				},
				target:   9, // 5 -> 4 sums to 9, but 4 isn't a leaf
				sumSoFar: 0,
			},
			want: false,
		},
		{
			name: "multi-level - classic example matches (5-4-11-2)",
			args: args{
				t: &Node[int]{
					val: 5,
					left: &Node[int]{
						val: 4,
						left: &Node[int]{
							val:   11,
							left:  &Node[int]{val: 7},
							right: &Node[int]{val: 2},
						},
					},
					right: &Node[int]{
						val:   8,
						left:  &Node[int]{val: 13},
						right: &Node[int]{val: 4, right: &Node[int]{val: 1}},
					},
				},
				target:   22,
				sumSoFar: 0,
			},
			want: true,
		},
		{
			name: "multi-level - no matching path",
			args: args{
				t: &Node[int]{
					val: 1,
					left: &Node[int]{
						val:  2,
						left: &Node[int]{val: 3},
					},
					right: &Node[int]{
						val:   4,
						right: &Node[int]{val: 5},
					},
				},
				target:   100,
				sumSoFar: 0,
			},
			want: false,
		},
		{
			name: "handles negative values",
			args: args{
				t: &Node[int]{
					val: -2,
					right: &Node[int]{
						val: -3,
					},
				},
				target:   -5, // -2 -> -3
				sumSoFar: 0,
			},
			want: true,
		},
		{
			name: "non-zero sumSoFar (pre-accumulated) still works",
			args: args{
				t: &Node[int]{
					val:  5,
					left: &Node[int]{val: 4},
				},
				target:   12, // sumSoFar=3 + 5 + 4 = 12
				sumSoFar: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSum(tt.args.t, tt.args.target, tt.args.sumSoFar); got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
