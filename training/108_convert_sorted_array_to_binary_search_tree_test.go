package training

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// [-10,-3,0,5,9] → [0,-3,9,-10,null,5]
		//        0
		//       / \
		//     -3   9
		//     /   /
		//   -10  5
		{
			name: "example 1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		// [1,3] → [3,1]
		//   3
		//  /
		// 1
		{
			name: "example 2",
			args: args{nums: []int{1, 3}},
			want: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
