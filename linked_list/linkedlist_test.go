package linkedlist_test

import (
	"testing"

	linkedlist "github.com/ellywynn/go-data-structures/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	l := linkedlist.New()
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.IsEmpty())

	// Try to remove value from empty list
	v, err := l.Remove()
	assert.Equal(t, 0, v)
	assert.Error(t, err)

	// Get the last element from empty list
	v, err = l.Last()
	assert.Equal(t, 0, v)
	assert.Error(t, err)

	// Add elements to the list
	for i := 1; i <= 5; i++ {
		l.Add(i)
	}

	// Check IsEmpty and Size functions
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 5, l.Size())

	// Check the last element
	v, err = l.Last()
	assert.NoError(t, err)
	assert.Equal(t, 5, v)

	// Try to remove the last value and check list for update
	v, err = l.Remove()
	assert.NoError(t, err)
	assert.Equal(t, 5, v)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 4, l.Size())

	// Check new last value after remove
	v, err = l.Last()
	assert.NoError(t, err)
	assert.Equal(t, 4, v)
}
