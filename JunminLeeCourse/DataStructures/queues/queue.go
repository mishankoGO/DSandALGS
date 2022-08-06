package queues

type Queue struct {
	items []int
}

// Enqueue method
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue method
func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
