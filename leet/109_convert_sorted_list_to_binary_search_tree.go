package leet

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	var prev, hare, tortoise *ListNode

	hare = head

	tortoise = head
	for hare != nil && hare.Next != nil {
		prev = tortoise
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	node := &TreeNode{
		Val:   tortoise.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(tortoise.Next),
	}

	return node
}
