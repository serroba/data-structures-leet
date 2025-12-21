package training

func hasPathSum(root *TreeNode, targetSum int) bool {
	var addsTo func(*TreeNode, int, int) bool
	addsTo = func(node *TreeNode, targetSum, sumSoFar int) bool {
		if node == nil {
			return false
		}
		if node.Left == nil && node.Right == nil {
			return targetSum == sumSoFar+node.Val
		}
		return addsTo(node.Left, targetSum, sumSoFar+node.Val) || addsTo(node.Right, targetSum, sumSoFar+node.Val)
	}

	return addsTo(root, targetSum, 0)
}
