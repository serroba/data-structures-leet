package leet

func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode

	var findSwapped func(node *TreeNode)
	findSwapped = func(node *TreeNode) {
		if node == nil {
			return
		}

		findSwapped(node.Left)

		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node

		findSwapped(node.Right)
	}

	findSwapped(root)
	first.Val, second.Val = second.Val, first.Val
}
