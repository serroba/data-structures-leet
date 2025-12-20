package training

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if p.Left == nil && q.Left != nil || p.Left != nil && q.Left == nil {
		return false
	}
	if p.Right == nil && q.Right != nil || p.Right != nil && q.Right == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
