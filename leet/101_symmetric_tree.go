package leet

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}

	var isMirror func(*TreeNode, *TreeNode) bool

	isMirror = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}
