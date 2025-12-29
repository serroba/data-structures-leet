package leet

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	hare := head
	tortoise := head

	for hare != nil && hare.Next != nil && hare.Next.Next != nil {
		hare = hare.Next.Next

		tortoise = tortoise.Next
		if hare == tortoise {
			return true
		}
	}

	return false
}
