package leet

import "testing"

func Test_hasCycle(t *testing.T) {
	// [3,2,0,-4], pos=1 → cycle at node 1
	// 3 → 2 → 0 → -4
	//     ↑_________|
	node1_3 := &ListNode{Val: 3}
	node1_2 := &ListNode{Val: 2}
	node1_0 := &ListNode{Val: 0}
	node1_4 := &ListNode{Val: -4}
	node1_3.Next = node1_2
	node1_2.Next = node1_0
	node1_0.Next = node1_4
	node1_4.Next = node1_2 // cycle back to pos 1

	// [1,2], pos=0 → cycle at node 0
	// 1 → 2
	// ↑___|
	node2_1 := &ListNode{Val: 1}
	node2_2 := &ListNode{Val: 2}
	node2_1.Next = node2_2
	node2_2.Next = node2_1 // cycle back to pos 0

	// [1], pos=-1 → no cycle
	node3_1 := &ListNode{Val: 1}

	// [1,2], pos=-1 → no cycle
	node4_1 := &ListNode{Val: 1}
	node4_2 := &ListNode{Val: 2}
	node4_1.Next = node4_2

	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "example 1 - cycle at pos 1",
			head: node1_3,
			want: true,
		},
		{
			name: "example 2 - cycle at pos 0",
			head: node2_1,
			want: true,
		},
		{
			name: "example 3 - no cycle",
			head: node3_1,
			want: false,
		},
		{
			name: "example 4 - empty list",
			head: nil,
			want: false,
		},
		{
			name: "example 5 - two nodes no cycle",
			head: node4_1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
