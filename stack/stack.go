package stack

import "errors"

var (
	errStackEmpty = errors.New("the stack is empty")
)

type Stack struct {
	size   int
	values []int
}

func New() *Stack {
	return &Stack{
		size:   0,
		values: make([]int, 0, 5),
	}
}

func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errStackEmpty
	}

	v := s.values[s.size-1]
	s.values = s.values[:s.size-1]
	s.size--
	return v, nil
}

func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, errStackEmpty
	}

	return s.values[s.size-1], nil
}

func (s *Stack) Size() int {
	return s.size
}
