package leet

import "ds/queue"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := queue.New(root)

	var levels [][]int

	leftToRight := true

	for !q.Empty() {
		levelSize := q.Len()

		level := make([]int, levelSize)
		for i := range levelSize {
			node := q.Dequeue()
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		leftToRight = !leftToRight

		levels = append(levels, level)
	}

	return levels
}
