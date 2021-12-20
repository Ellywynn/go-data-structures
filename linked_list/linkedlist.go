package linkedlist

import "errors"

var (
	errEmptyList = errors.New("the linked list is empty")
)

// Linked list node
type Node struct {
	Value int
	Next  *Node
}

// Linked list structure
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Creates a new list instance
func New() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Appends a value to the linked list
func (list *LinkedList) Add(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	// Initialize the head
	if list.IsEmpty() {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.Next = newNode
		list.head = list.head.Next // Update the head
	}

	list.size++
}

// Returns the size of the linked list
func (list *LinkedList) Size() int {
	return list.size
}

// Removes the last node from list and returns its value
func (list *LinkedList) Remove() (int, error) {
	if list.IsEmpty() {
		return 0, errEmptyList
	}

	var prev, node *Node = nil, list.tail

	for node != list.head {
		if node.Next == list.head {
			prev = node
		}
		node = node.Next
	}

	v := list.head.Value
	list.head = prev
	list.head.Next = nil
	list.size--

	if list.IsEmpty() {
		list.tail = nil
		list.head = nil
	}

	return v, nil
}

// Returns the value of the last element or error if the list is empty
func (list *LinkedList) Last() (int, error) {
	if list.IsEmpty() {
		return 0, errEmptyList
	}

	return list.head.Value, nil
}

// Check if list is empty
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}
