package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	type testCase[T any] struct {
		name  string
		items []T
	}

	intTests := []testCase[int]{
		{name: "Add single int", items: []int{1}},
		{name: "Add multiple ints", items: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range intTests {
		q := NewQueue[int]()
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			for i, want := range tt.items {
				got := q.Dequeue()
				if got != want {
					t.Errorf("Dequeue() at index %d = %v, want %v", i, got, want)
				}
			}
		})
	}

	stringTests := []testCase[string]{
		{name: "Add single string", items: []string{"hello"}},
		{name: "Add multiple strings", items: []string{"hello", "world", "foo"}},
	}
	for _, tt := range stringTests {
		q := NewQueue[string]()
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			for i, want := range tt.items {
				got := q.Dequeue()
				if got != want {
					t.Errorf("Dequeue() at index %d = %v, want %v", i, got, want)
				}
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	type testCase[T any] struct {
		name  string
		items []T
		want  int
	}

	tests := []testCase[int]{
		{name: "Empty queue", items: []int{}, want: 0},
		{name: "Single element", items: []int{1}, want: 1},
		{name: "Multiple elements", items: []int{1, 2, 3, 4, 5}, want: 5},
	}
	for _, tt := range tests {
		q := NewQueue[int]()
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			if got := q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Len decreases after Dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		if got := q.Len(); got != 3 {
			t.Errorf("Len() after enqueue = %v, want 3", got)
		}

		q.Dequeue()
		if got := q.Len(); got != 2 {
			t.Errorf("Len() after 1 dequeue = %v, want 2", got)
		}

		q.Dequeue()
		if got := q.Len(); got != 1 {
			t.Errorf("Len() after 2 dequeues = %v, want 1", got)
		}

		q.Dequeue()
		if got := q.Len(); got != 0 {
			t.Errorf("Len() after 3 dequeues = %v, want 0", got)
		}
	})
}
