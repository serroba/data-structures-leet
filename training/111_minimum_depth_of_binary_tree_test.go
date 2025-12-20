package training

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// [3,9,20,null,null,15,7] → 2
		//       3
		//      / \
		//     9  20
		//        / \
		//       15  7
		// Min path: 3→9
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
			want: 2,
		},
		// [2,null,3,null,4,null,5,null,6] → 5
		// 2
		//  \
		//   3
		//    \
		//     4
		//      \
		//       5
		//        \
		//         6
		{
			name: "example 2 - skewed tree",
			args: args{
				root: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val:   5,
								Right: &TreeNode{Val: 6},
							},
						},
					},
				},
			},
			want: 5,
		},
		// [] → 0
		{
			name: "example 3 - empty tree",
			args: args{root: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
