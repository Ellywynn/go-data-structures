package stack

import "errors"

var (
	errStackEmpty = errors.New("the stack is empty")
)

type Stack struct {
	Size   int
	values []int
}

func New() *Stack {
	return &Stack{
		Size:   0,
		values: make([]int, 0, 5),
	}
}

func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
	s.Size++
}

func (s *Stack) Pop() (int, error) {
	if s.Size == 0 {
		return 0, errStackEmpty
	}

	v := s.values[s.Size-1]
	s.values = s.values[:s.Size-1]
	s.Size--
	return v, nil
}

func (s *Stack) Peek() (int, error) {
	if s.Size == 0 {
		return 0, errStackEmpty
	}

	return s.values[s.Size-1], nil
}
