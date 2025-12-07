package list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var output []string
	for l != nil {
		output = append(output, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(output, "->")
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	dummy := node
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
	} else {
		node.Next = list2
	}

	return dummy.Next
}
