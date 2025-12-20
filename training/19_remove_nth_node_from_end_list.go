package training

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	if n > count {
		return nil
	}
	current = head
	for i := 0; i < count-n-1; i++ {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	if count == 1 && n == 1 {
		return nil
	}

	return head
}
