package btree_test

import (
	"testing"

	"github.com/ellywynn/go-data-structures/btree"
	"github.com/stretchr/testify/assert"
)

func fillTree() *btree.Node {
	tree := btree.New(20)

	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(4)
	tree.Insert(30)
	tree.Insert(17)

	return tree
}

func TestBTree_Insert(t *testing.T) {
	tree := btree.New(10)
	assert.NotNil(t, tree)
	assert.Equal(t, 10, tree.Value)

	tree.Insert(5)
	assert.NotNil(t, tree.Left)
	assert.Equal(t, 5, tree.Left.Value)

	tree.Insert(15)
	assert.NotNil(t, tree.Right)
	assert.Equal(t, 15, tree.Right.Value)

	tree.Insert(3)
	assert.NotNil(t, tree.Left.Left)
	assert.Equal(t, 3, tree.Left.Left.Value)
}

func TestBTree_Search(t *testing.T) {
	tree := fillTree()

	found := tree.Search(10)
	assert.False(t, found)

	found = tree.Search(15)
	assert.True(t, found)

	found = tree.Search(20)
	assert.True(t, found)

	found = tree.Search(0)
	assert.False(t, found)
}
