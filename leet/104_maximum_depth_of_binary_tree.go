package leet

import "ds/stack"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	type item struct {
		node  *TreeNode
		depth int
	}
	s := stack.New(item{root, 1})
	for !s.Empty() {
		i, _ := s.Pop()
		if depth < i.depth {
			depth = i.depth
		}
		if i.node.Left != nil {
			s.Push(item{i.node.Left, i.depth + 1})
		}
		if i.node.Right != nil {
			s.Push(item{i.node.Right, i.depth + 1})
		}
	}
	return depth
}
