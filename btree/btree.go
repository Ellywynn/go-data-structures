package btree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// Creates a new Binary Tree instance
func New(value int) *Node {
	return &Node{Value: value}
}

func (t *Node) Insert(value int) {
	if t.Value < value {
		// Move right
		if t.Right == nil {
			t.Right = &Node{Value: value}
		} else {
			t.Right.Insert(value)
		}
	} else {
		// Move left
		if t.Left == nil {
			t.Left = &Node{Value: value}
		} else {
			t.Left.Insert(value)
		}
	}
}

func (t *Node) Search(value int) bool {
	if t == nil {
		return false
	}

	if value < t.Value {
		return t.Left.Search(value)
	} else if value > t.Value {
		return t.Right.Search(value)
	} else {
		return true
	}
}
