package leet

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// [1,2,3] vs [1,2,3] → true
		{
			name: "example 1",
			args: args{
				p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			want: true,
		},
		// [1,2] vs [1,null,2] → false
		{
			name: "example 2",
			args: args{
				p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
				q: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			},
			want: false,
		},
		// [1,2,1] vs [1,1,2] → false
		{
			name: "example 3",
			args: args{
				p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
				q: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
