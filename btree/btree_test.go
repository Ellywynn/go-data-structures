package btree_test

import (
	"testing"

	"github.com/ellywynn/go-data-structures/btree"
	"github.com/stretchr/testify/assert"
)

func fillTree() *btree.BTree {
	tree := btree.New()

	tree.Add(20)
	tree.Add(14)
	tree.Add(15)
	tree.Add(25)
	tree.Add(4)
	tree.Add(30)
	tree.Add(17)

	return tree
}

func TestBTree_Add(t *testing.T) {
	tree := btree.New()
	assert.Nil(t, tree.Root())

	tree.Add(10)
	assert.NotNil(t, tree.Root())
	assert.Equal(t, 10, tree.Root().Value)

	tree.Add(5)
	assert.NotNil(t, tree.Root().Left)
	assert.Equal(t, 5, tree.Root().Left.Value)

	tree.Add(15)
	assert.NotNil(t, tree.Root().Right)
	assert.Equal(t, 15, tree.Root().Right.Value)

	tree.Add(3)
	assert.NotNil(t, tree.Root().Left.Left)
	assert.Equal(t, 3, tree.Root().Left.Left.Value)
}

func TestBTree_FindNode(t *testing.T) {
	tree := fillTree()

	node := tree.FindNode(10)
	assert.Nil(t, node)

	node = tree.FindNode(15)
	assert.NotNil(t, node)

	node = tree.FindNode(20)
	assert.NotNil(t, node)
}
