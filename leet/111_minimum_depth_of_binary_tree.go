package leet

import "ds/queue"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	q := queue.New(root)
	depth := 1
	for !q.Empty() {
		levelSize := q.Len()
		for i := 0; i < levelSize; i++ {
			n := q.Dequeue()
			if n.Left == nil && n.Right == nil {
				return depth
			}
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
		depth++
	}
	return depth
}
