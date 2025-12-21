package training

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// [1,null,2,3] → [1,2,3]
		// 1
		//  \
		//   2
		//  /
		// 3
		{
			name: "example 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: []int{1, 2, 3},
		},
		// [1] → [1]
		{
			name: "example 2 - single node",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: []int{1},
		},
		// [] → []
		{
			name: "example 3 - empty tree",
			args: args{root: nil},
			want: nil,
		},
		// [1,2,3] → [1,2,3]
		//     1
		//    / \
		//   2   3
		{
			name: "example 4 - balanced",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
