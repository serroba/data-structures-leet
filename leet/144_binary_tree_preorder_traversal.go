package leet

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		if node.Left != nil {
			traverse(node.Left)
		}
		if node.Right != nil {
			traverse(node.Right)
		}
	}
	traverse(root)
	return result
}
