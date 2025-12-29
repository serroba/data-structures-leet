package queue

type MinHeap struct {
	tree []int
}

// Item represents an element with a priority for the priority queue.
type Item struct {
	Priority int
	Value    int
}

// PriorityQueue is a min-heap based priority queue for (priority, value) pairs.
type PriorityQueue struct {
	items []Item
}

func (h *MinHeap) Len() int {
	return len(h.tree)
}

func (h *MinHeap) Empty() bool {
	return h.Len() == 0
}

func (h *MinHeap) Peek() (int, bool) {
	if len(h.tree) == 0 {
		return 0, false
	}

	return h.tree[0], true
}

func (h *MinHeap) Push(i int) {
	h.tree = append(h.tree, i)
	h.siftUp()
}

func (h *MinHeap) Pop() (int, bool) {
	if len(h.tree) == 0 {
		return 0, false
	}

	root := h.tree[0]
	last := h.tree[len(h.tree)-1]

	h.tree = h.tree[:len(h.tree)-1]
	if !h.Empty() {
		h.tree[0] = last
		h.siftDown()
	}

	return root, true
}

// siftUp restores the min-heap property after inserting a new element at the end.
// It "bubbles up" the newly inserted element until it's in the correct position.
//
// Heap as array: for node at index i, parent is at (i-1)/2
//
//	    0
//	   / \
//	  1   2      indices
//	 / \
//	3   4
//
// Example: heap [1, 3, 5] after Push(2) becomes [1, 3, 5, 2]
//
//	Step 1: i=3, parent p=1, tree[1]=3 > tree[3]=2 → swap → [1, 2, 5, 3]
//	Step 2: i=1, parent p=0, tree[0]=1 <= tree[1]=2 → stop (heap property satisfied)
func (h *MinHeap) siftUp() {
	i := len(h.tree) - 1 // start at the last element (just inserted)
	for i > 0 {
		p := (i - 1) / 2 // parent index formula for array-based heap
		if h.tree[p] <= h.tree[i] {
			break // parent is smaller or equal, heap property satisfied
		}

		h.tree[p], h.tree[i] = h.tree[i], h.tree[p] // swap with parent
		i = p                                       // move up to parent's position
	}
}

// siftDown restores the min-heap property after replacing the root with the last element.
// It "bubbles down" the root element until it's in the correct position.
//
// Heap as array: for node at index i, left child is at 2i+1, right child is at 2i+2
//
//	    0
//	   / \
//	  1   2      indices
//	 / \
//	3   4
//
// Example: after Pop from [1, 3, 5, 7], we move last (7) to root → [7, 3, 5]
//
//	Step 1: i=0, left=1 (val=3), right=2 (val=5)
//	        smallest = 1 (3 < 7 and 3 < 5) → swap → [3, 7, 5]
//	Step 2: i=1, left=3 (out of bounds), right=4 (out of bounds)
//	        smallest = i → stop (no children, heap property satisfied)
func (h *MinHeap) siftDown() {
	n := h.Len()

	i := 0 // start at the root
	for {
		l := 2*i + 1 // left child index
		r := 2*i + 2 // right child index
		smallest := i

		// find the smallest among current node and its children
		if l < n && h.tree[l] < h.tree[smallest] {
			smallest = l
		}

		if r < n && h.tree[r] < h.tree[smallest] {
			smallest = r
		}

		if smallest == i {
			break // current node is smaller than both children, heap property satisfied
		}

		h.tree[i], h.tree[smallest] = h.tree[smallest], h.tree[i] // swap with smallest child
		i = smallest                                              // move down to child's position
	}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Empty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue) Push(priority, value int) {
	pq.items = append(pq.items, Item{Priority: priority, Value: value})
	pq.siftUp()
}

func (pq *PriorityQueue) Pop() (Item, bool) {
	if len(pq.items) == 0 {
		return Item{}, false
	}

	root := pq.items[0]
	last := pq.items[len(pq.items)-1]

	pq.items = pq.items[:len(pq.items)-1]
	if !pq.Empty() {
		pq.items[0] = last
		pq.siftDown()
	}

	return root, true
}

func (pq *PriorityQueue) siftUp() {
	i := len(pq.items) - 1
	for i > 0 {
		p := (i - 1) / 2
		if pq.items[p].Priority <= pq.items[i].Priority {
			break
		}

		pq.items[p], pq.items[i] = pq.items[i], pq.items[p]
		i = p
	}
}

func (pq *PriorityQueue) siftDown() {
	n := pq.Len()

	i := 0
	for {
		l := 2*i + 1
		r := 2*i + 2
		smallest := i

		if l < n && pq.items[l].Priority < pq.items[smallest].Priority {
			smallest = l
		}

		if r < n && pq.items[r].Priority < pq.items[smallest].Priority {
			smallest = r
		}

		if smallest == i {
			break
		}

		pq.items[i], pq.items[smallest] = pq.items[smallest], pq.items[i]
		i = smallest
	}
}
