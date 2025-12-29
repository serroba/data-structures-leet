package tree

import (
	"cmp"

	"ds/queue"
	"ds/stack"
)

type Node[T cmp.Ordered] struct {
	left  *Node[T]
	right *Node[T]
	val   T
}

type Tree[T cmp.Ordered] struct {
	root *Node[T]
}

func Insert[T cmp.Ordered](node *Node[T], val T) {
	if node == nil {
		node = &Node[T]{val: val}
	}

	if val < node.left.val {
		Insert(node.left, val)
	} else {
		Insert(node.right, val)
	}
}

func Find[T cmp.Ordered](node *Node[T], val T) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	}

	if val < node.val {
		return Find(node.left, val)
	}

	return Find(node.right, val)
}

func (t *Tree[T]) FindAllBFS() []T {
	var output []T
	if t.root == nil {
		return output
	}

	q := queue.New[*Node[T]]()
	q.Enqueue(t.root)

	for !q.Empty() {
		item := q.Dequeue()

		output = append(output, item.val)
		if item.left != nil {
			q.Enqueue(item.left)
		}

		if item.right != nil {
			q.Enqueue(item.right)
		}
	}

	return output
}

func (t *Tree[T]) FindAllDFS() []T {
	var output []T
	if t.root == nil {
		return output
	}

	s := stack.New(t.root)
	for !s.Empty() {
		item, _ := s.Pop()

		output = append(output, item.val)
		if item.right != nil {
			s.Push(item.right)
		}

		if item.left != nil {
			s.Push(item.left)
		}
	}

	return output
}

func HasPathSum(t *Node[int], target, sumSoFar int) bool {
	if t == nil {
		return false
	}

	sumSoFar += t.val
	if t.left == nil && t.right == nil {
		return target-sumSoFar == 0
	}

	return HasPathSum(t.left, target, sumSoFar) || HasPathSum(t.right, target, sumSoFar)
}
