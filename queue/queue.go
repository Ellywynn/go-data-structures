package queue

import "errors"

var (
	errEmptyQueue = errors.New("the queue is empty")
)

type Queue struct {
	size   int
	values []int
}

// Creates a new queue instance
func New() *Queue {
	return &Queue{
		size:   0,
		values: nil,
	}
}

// Returns the first emelent in the queue
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errEmptyQueue
	}

	return q.values[0], nil
}

// Returns the last emelent in the queue
func (q *Queue) Back() (int, error) {
	if q.IsEmpty() {
		return 0, errEmptyQueue
	}

	return q.values[q.size-1], nil
}

// Returns the first element and removes it from the queue
func (q *Queue) Pop() (int, error) {
	v, err := q.Front()
	if err != nil {
		return 0, err
	}

	q.values = q.values[1:]
	q.size--
	return v, nil
}

// Appends an element to the end of the queue
func (q *Queue) PushBack(values ...int) {
	q.values = append(q.values, values...)
	q.size += len(values)
}

// Appends an element to the end of the queue
func (q *Queue) PushFront(values ...int) {
	q.size += len(values)
	temp := make([]int, q.size)
	temp = append(temp, values...)
	temp = append(temp, (q.values)...)
	q.values = temp
}

// Return true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Returns the current size of the queue
func (q *Queue) Size() int {
	return q.size
}
