package stack_test

import (
	"testing"

	"github.com/ellywynn/go-data-structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// Create the stack
	s := stack.New()
	assert.Equal(t, 0, s.Size)

	// Pop and Peek empty stack
	_, err := s.Pop()
	assert.Error(t, err)
	_, err = s.Peek()
	assert.Error(t, err)

	// Push elements
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}

	// Peek
	peek, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 5, peek)

	// Pops
	pop1, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 5, pop1)

	pop2, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 4, pop2)

	// Size
	assert.Equal(t, 3, s.Size)
}
