package leet

import (
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		// [1,3,null,null,2] → [3,1,null,null,2]
		// Swapped: 1 and 3
		//   1          3
		//  /    →     /
		// 3          1
		//  \          \
		//   2          2
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 2},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
		},
		// [3,1,4,null,null,2] → [2,1,4,null,null,3]
		// Swapped: 2 and 3
		//     3            2
		//    / \    →     / \
		//   1   4        1   4
		//      /            /
		//     2            3
		{
			name: "example 2",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 2},
				},
			},
			want: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.root)
			if !reflect.DeepEqual(tt.root, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.root, tt.want)
			}
		})
	}
}
