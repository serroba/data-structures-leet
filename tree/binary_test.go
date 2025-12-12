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
