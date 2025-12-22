package leet

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// [3,9,20,null,null,15,7] → [[3],[20,9],[15,7]]
		//       3          ← left to right
		//      / \
		//     9  20        ← right to left
		//        / \
		//       15  7      ← left to right
		{
			name: "example 1",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		// [1] → [[1]]
		{
			name: "example 2 - single node",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: [][]int{{1}},
		},
		// [] → []
		{
			name: "example 3 - empty tree",
			args: args{root: nil},
			want: nil,
		},
		// [1,2,3,4,null,null,5] → [[1],[3,2],[4,5]]
		//        1        ← left to right
		//       / \
		//      2   3      ← right to left
		//     /     \
		//    4       5    ← left to right
		{
			name: "example 4 - sparse tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
