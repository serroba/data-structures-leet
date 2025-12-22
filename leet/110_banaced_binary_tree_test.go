package leet

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// [3,9,20,null,null,15,7] → true
		//       3
		//      / \
		//     9  20
		//        / \
		//       15  7
		{
			name: "example 1 - balanced",
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
			want: true,
		},
		// [1,2,2,3,3,null,null,4,4] → false
		//          1
		//         / \
		//        2   2
		//       / \
		//      3   3
		//     / \
		//    4   4
		{
			name: "example 2 - unbalanced",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  &TreeNode{Val: 4},
							Right: &TreeNode{Val: 4},
						},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
		// [] → true (empty tree)
		{
			name: "example 3 - empty tree",
			args: args{root: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
