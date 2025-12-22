package leet

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum=22 → true
		//         5
		//        / \
		//       4   8
		//      /   / \
		//     11  13  4
		//    /  \      \
		//   7    2      1
		// Path: 5→4→11→2 = 22
		{
			name: "example 1 - has path",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		// [1,2,3], targetSum=5 → false
		//     1
		//    / \
		//   2   3
		// Paths: 1→2=3, 1→3=4
		{
			name: "example 2 - no path",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				targetSum: 5,
			},
			want: false,
		},
		// [], targetSum=0 → false (empty tree)
		{
			name: "example 3 - empty tree",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
		// [1,2], targetSum=0 → false (single child case)
		//   1
		//  /
		// 2
		{
			name: "example 4 - single child",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
				targetSum: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
