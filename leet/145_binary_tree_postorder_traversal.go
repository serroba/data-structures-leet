package leet

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			traverse(node.Left)
		}
		if node.Right != nil {
			traverse(node.Right)
		}
		result = append(result, node.Val)
	}
	traverse(root)
	return result
}
