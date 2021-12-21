package btree

type node struct {
	Left  *node
	Right *node
	Value int
}

func newNode(value int) *node {
	return &node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

type BTree struct {
	root *node
}

// Creates a new Binary Tree instance
func New() *BTree {
	return &BTree{
		root: nil,
	}
}

func (t *BTree) Add(value int) {
	t.root = add(t.root, value)
}

func (t *BTree) FindNode(value int) *node {
	return findNode(t.root, value)
}

func (t *BTree) DeleteNode(value int) {
	node := t.FindNode(value)
	if node == nil {
		return
	}

	node = nil
}

func (t *BTree) Root() *node {
	return t.root
}

func add(current *node, value int) *node {
	if current == nil {
		current = newNode(value)
		return current
	}

	if value < current.Value {
		current.Left = add(current.Left, value)
	} else if value > current.Value {
		current.Right = add(current.Right, value)
	}

	return current
}

func findNode(current *node, value int) *node {
	if current == nil {
		return nil
	}

	if value == current.Value {
		return current
	}

	if value < current.Value {
		return findNode(current.Left, value)
	} else {
		return findNode(current.Right, value)
	}
}
