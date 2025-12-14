package queue

type MinHeap struct {
	tree []int
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
	h.siftUp(len(h.tree) - 1)
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
		h.siftDown(0)
	}
	return root, true
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.tree[p] <= h.tree[i] {
			break
		}
		h.tree[p], h.tree[i] = h.tree[i], h.tree[p]
		i = p
	}
}

func (h *MinHeap) siftDown(i int) {
	n := h.Len()
	for {
		l := 2*i + 1
		r := 2*i + 2
		smallest := i

		if l < n && h.tree[l] < h.tree[smallest] {
			smallest = l
		}
		if r < n && h.tree[r] < h.tree[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.tree[i], h.tree[smallest] = h.tree[smallest], h.tree[i]
		i = smallest
	}
}
