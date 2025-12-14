package queue

type Queue[T any] struct {
	queue []T
}

func New[T any](items ...T) *Queue[T] {
	queue := make([]T, 0)
	for _, item := range items {
		queue = append(queue, item)
	}
	return &Queue[T]{queue: queue}
}

func (q *Queue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *Queue[T]) Dequeue() T {
	top := q.queue[0]
	q.queue = q.queue[1:]
	return top
}

func (q *Queue[T]) Len() int {
	return len(q.queue)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}
