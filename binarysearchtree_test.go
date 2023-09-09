package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) String() string {
	return bst.Root.String()
}

func (bst *BinarySearchTree) Insert(n *Node) {
	node := &Node{Value: n.Value}

	if bst.Root == nil {
		bst.Root = node
		return
	}

	var insertHelper func(nodeToInsert *Node, current *Node)
	insertHelper = func(nodeToInsert *Node, current *Node) {
		if current == nil {
			return
		}

		if nodeToInsert.Value.(int) < current.Value.(int) {
			if current.Left != nil {
				insertHelper(nodeToInsert, current.Left)
			} else {
				current.Left = nodeToInsert
			}
		} else if nodeToInsert.Value.(int) > current.Value.(int) {
			if current.Right != nil {
				insertHelper(nodeToInsert, current.Right)
			} else {
				current.Right = nodeToInsert
			}
		}
	}

	insertHelper(node, bst.Root)
}

func (bst *BinarySearchTree) InOrder() []any {
	res := []any{}

	var inOrderHelper func(n *Node)
	inOrderHelper = func(n *Node) {
		if n == nil {
			return
		}

		inOrderHelper(n.Left)
		res = append(res, n.Value)
		inOrderHelper(n.Right)
	}

	inOrderHelper(bst.Root)
	return res
}

func TestBinarySearchTree(t *testing.T) {
	//       6
	//     /   \
	//    4     8
	//   / \   / \
	//  3   5 7   9

	bst := NewBinarySearchTree()

	bst.Insert(&Node{Value: 6})
	bst.Insert(&Node{Value: 4})
	bst.Insert(&Node{Value: 8})
	bst.Insert(&Node{Value: 3})
	bst.Insert(&Node{Value: 5})
	bst.Insert(&Node{Value: 9})
	bst.Insert(&Node{Value: 7})

	fmt.Println(bst.String())
	assert.Equal(t, "(((3) 4 (5)) 6 ((7) 8 (9)))", bst.String())

	assert.EqualValues(t, []any{3, 4, 5, 6, 7, 8, 9}, bst.InOrder())
}
