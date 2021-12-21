package queue_test

import (
	"testing"

	"github.com/ellywynn/go-data-structures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue_New(t *testing.T) {
	q := queue.New()
	assert.Equal(t, 0, q.Size())
}

func TestQueue_Push(t *testing.T) {
	q := queue.New()
	q.PushBack(10, 15, 20)

	assert.Equal(t, 3, q.Size())

	v, err := q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 10, v)

	v, err = q.Back()
	assert.NoError(t, err)
	assert.Equal(t, 20, v)

	q.PushFront(50, 30)
	assert.Equal(t, 5, q.Size())

	v, err = q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 50, v)

	v, err = q.Back()
	assert.NoError(t, err)
	assert.Equal(t, 20, v)
}

func TestQueue_Pop(t *testing.T) {
	q := queue.New()
	v, err := q.Pop()
	assert.Error(t, err)
	assert.Equal(t, 0, v)

	q.PushBack(10, 20, 30)

	v, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, v)
	assert.Equal(t, 2, q.Size())

	q.Pop()
	q.Pop()

	assert.True(t, q.IsEmpty())
}

func TestQueue_IsEmpty(t *testing.T) {
	q := queue.New()
	assert.True(t, q.IsEmpty())

	q.PushFront(10, 20)
	assert.False(t, q.IsEmpty())

	q.Pop()
	q.Pop()
	assert.True(t, q.IsEmpty())
}
