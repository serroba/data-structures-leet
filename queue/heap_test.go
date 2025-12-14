package queue

import (
	"testing"
)

func TestMinHeap_Len(t *testing.T) {
	tests := []struct {
		name   string
		pushes []int
		want   int
	}{
		{name: "empty heap", pushes: []int{}, want: 0},
		{name: "single element", pushes: []int{5}, want: 1},
		{name: "multiple elements", pushes: []int{5, 3, 7, 1}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{}
			for _, v := range tt.pushes {
				h.Push(v)
			}
			if got := h.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Peek(t *testing.T) {
	tests := []struct {
		name    string
		pushes  []int
		want    int
		wantOk  bool
	}{
		{name: "empty heap", pushes: []int{}, want: 0, wantOk: false},
		{name: "single element", pushes: []int{5}, want: 5, wantOk: true},
		{name: "min is first pushed", pushes: []int{1, 3, 5}, want: 1, wantOk: true},
		{name: "min is last pushed", pushes: []int{5, 3, 1}, want: 1, wantOk: true},
		{name: "min is middle pushed", pushes: []int{5, 1, 3}, want: 1, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{}
			for _, v := range tt.pushes {
				h.Push(v)
			}
			got, ok := h.Peek()
			if ok != tt.wantOk {
				t.Errorf("Peek() ok = %v, wantOk %v", ok, tt.wantOk)
			}
			if got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
	t.Run("maintains min heap property", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(5)
		h.Push(3)
		h.Push(7)
		h.Push(1)

		min, _ := h.Peek()
		if min != 1 {
			t.Errorf("Peek() after pushes = %v, want 1", min)
		}
	})

	t.Run("push duplicate values", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(3)
		h.Push(3)
		h.Push(3)

		if h.Len() != 3 {
			t.Errorf("Len() = %v, want 3", h.Len())
		}
		min, _ := h.Peek()
		if min != 3 {
			t.Errorf("Peek() = %v, want 3", min)
		}
	})
}

func TestMinHeap_Pop(t *testing.T) {
	t.Run("pop from empty heap", func(t *testing.T) {
		h := &MinHeap{}
		_, ok := h.Pop()
		if ok {
			t.Errorf("Pop() on empty heap should return false")
		}
	})

	t.Run("pop single element", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(5)
		got, ok := h.Pop()
		if !ok {
			t.Errorf("Pop() ok = false, want true")
		}
		if got != 5 {
			t.Errorf("Pop() = %v, want 5", got)
		}
		if h.Len() != 0 {
			t.Errorf("Len() after pop = %v, want 0", h.Len())
		}
	})

	t.Run("pop returns elements in sorted order", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(5)
		h.Push(3)
		h.Push(7)
		h.Push(1)
		h.Push(4)

		expected := []int{1, 3, 4, 5, 7}
		for i, want := range expected {
			got, ok := h.Pop()
			if !ok {
				t.Errorf("Pop() #%d ok = false, want true", i)
			}
			if got != want {
				t.Errorf("Pop() #%d = %v, want %v", i, got, want)
			}
		}
	})

	t.Run("pop two elements", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(5)
		h.Push(3)

		got1, _ := h.Pop()
		got2, _ := h.Pop()

		if got1 != 3 {
			t.Errorf("First Pop() = %v, want 3", got1)
		}
		if got2 != 5 {
			t.Errorf("Second Pop() = %v, want 5", got2)
		}
	})

	t.Run("interleaved push and pop", func(t *testing.T) {
		h := &MinHeap{}
		h.Push(5)
		h.Push(3)

		got, _ := h.Pop() // should be 3
		if got != 3 {
			t.Errorf("Pop() = %v, want 3", got)
		}

		h.Push(1)
		got, _ = h.Pop() // should be 1
		if got != 1 {
			t.Errorf("Pop() = %v, want 1", got)
		}

		got, _ = h.Pop() // should be 5
		if got != 5 {
			t.Errorf("Pop() = %v, want 5", got)
		}
	})
}

func TestMinHeap_HeapProperty(t *testing.T) {
	t.Run("large random-ish sequence", func(t *testing.T) {
		h := &MinHeap{}
		values := []int{10, 4, 15, 20, 0, 8, 2, 14, 3, 9}
		for _, v := range values {
			h.Push(v)
		}

		prev := -1
		for h.Len() > 0 {
			got, _ := h.Pop()
			if got < prev {
				t.Errorf("Pop() = %v, but previous was %v (not sorted)", got, prev)
			}
			prev = got
		}
	})
}
