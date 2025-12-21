package training

import "ds/queue"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var levels [][]int
	q := queue.New(root)
	for !q.Empty() {
		levelSize := q.Len()
		var level []int
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue()
			level = append(level, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		levels = append(levels, level)
	}
	return levels
}
