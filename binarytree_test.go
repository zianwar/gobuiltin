package main

import (
	"container/list"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(0)
}

type Node struct {
	Left  *Node
	Value any
	Right *Node
}

func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprint(n.Value)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Depth(target *Node) int {
	depth := 0

	stack := list.New()
	stack.PushBack(bt.Root)

	// Apply DFS until we reach the target node or end of tree.
	for stack.Len() > 0 {
		// Pop from the stack
		node, _ := stack.Back().Value.(*Node), stack.Remove(stack.Back())

		// If we reached the target, or end of the tree, return the depth.
		if (target != nil && node.Value == target.Value) ||
			node.Left == nil || node.Right == nil {
			return depth
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		depth++
	}

	return depth
}

func (bt *BinaryTree) String() string {
	return bt.Root.String()
}

func (bt *BinaryTree) PreOrder() []any {
	result := []any{}
	stack := list.New()
	stack.PushBack(bt.Root)

	for stack.Len() > 0 {
		node, _ := stack.Back().Value.(*Node), stack.Remove(stack.Back())
		result = append(result, node.Value)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return result
}

func (bt *BinaryTree) InOrder() []any {
	result := []any{}
	stack := list.New()
	node := bt.Root

	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		// Pop from the stack
		node, _ = stack.Back().Value.(*Node), stack.Remove(stack.Back())
		result = append(result, node.Value)
		node = node.Right
	}

	return result
}

func (bt *BinaryTree) PostOrder() []any {
	result := []any{}
	stack := list.New()
	stack.PushBack(bt.Root)

	for stack.Len() > 0 {
		node, _ := stack.Back().Value.(*Node), stack.Remove(stack.Back())

		// Prepend to result array
		result = append([]any{node.Value}, result...)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	return result
}

func (bt *BinaryTree) PreOrderRecursive() []any {
	result := []any{}

	var helper func(n *Node)
	helper = func(n *Node) {
		if n == nil {
			return
		}
		result = append(result, n.Value)
		helper(n.Left)
		helper(n.Right)
	}

	helper(bt.Root)
	return result
}

func (bt *BinaryTree) InOrderRecursive() []any {
	result := []any{}

	var helper func(n *Node)
	helper = func(n *Node) {
		if n == nil {
			return
		}
		helper(n.Left)
		result = append(result, n.Value)
		helper(n.Right)
	}

	helper(bt.Root)
	return result
}

func (bt *BinaryTree) PostOrderRecursive() []any {
	result := []any{}

	var helper func(n *Node)
	helper = func(n *Node) {
		if n == nil {
			return
		}
		helper(n.Left)
		helper(n.Right)
		result = append(result, n.Value)
	}

	helper(bt.Root)
	return result
}

func (bt *BinaryTree) LevelOrder() [][]any {
	result := [][]any{}

	q := list.New()
	q.PushBack(bt.Root)

	for q.Len() > 0 {
		level := []any{}
		size := q.Len()

		for i := 0; i < size; i++ {
			node, _ := q.Front().Value.(*Node), q.Remove(q.Front())

			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
			level = append(level, node.Value)
		}
		result = append(result, level)
	}

	return result
}

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree()

	//       0
	//     /   \
	//    3     6
	//   / \   / \
	//  4   5 7   8

	node1 := &Node{Value: 3}
	node1.Left = &Node{Value: 4}
	node1.Right = &Node{Value: 5}

	node2 := &Node{Value: 6}
	node2.Left = &Node{Value: 7}
	node2.Right = &Node{Value: 8}

	root := &Node{Value: 0}
	root.Left = node1
	root.Right = node2

	bt.Root = root

	log.Println(bt.String())

	log.Printf("Depth -> %d\n", bt.Depth(nil))
	log.Printf("Depth of node (6) -> %d\n", bt.Depth(node2))

	//
	// Iterative traversals
	//
	preOrderItems := bt.PreOrder()
	log.Printf("PreOrder -> %v\n", preOrderItems)
	assert.EqualValues(t, []any{0, 3, 4, 5, 6, 7, 8}, preOrderItems)

	inOrderItems := bt.InOrder()
	log.Printf("InOrder -> %v\n", inOrderItems)
	assert.EqualValues(t, []any{4, 3, 5, 0, 7, 6, 8}, inOrderItems)

	postOrderItems := bt.PostOrder()
	log.Printf("PostOrder -> %v\n", postOrderItems)
	assert.EqualValues(t, []any{4, 5, 3, 7, 8, 6, 0}, postOrderItems)

	levelOrderItems := bt.LevelOrder()
	log.Printf("LevelOrder -> %v\n", levelOrderItems)
	assert.EqualValues(t, [][]any{{0}, {3, 6}, {4, 5, 7, 8}}, levelOrderItems)

	//
	// Recursive traversals
	//
	preOrderItemsRec := bt.PreOrderRecursive()
	log.Printf("PreOrderRecursive -> %v\n", preOrderItemsRec)
	assert.EqualValues(t, []any{0, 3, 4, 5, 6, 7, 8}, preOrderItemsRec)

	inOrderItemsRec := bt.InOrderRecursive()
	log.Printf("InOrderRecursive -> %v\n", inOrderItemsRec)
	assert.EqualValues(t, []any{4, 3, 5, 0, 7, 6, 8}, inOrderItemsRec)

	postOrderItemsRec := bt.PostOrderRecursive()
	log.Printf("PostOrderRecursive -> %v\n", postOrderItemsRec)
	assert.EqualValues(t, []any{4, 5, 3, 7, 8, 6, 0}, postOrderItemsRec)
}
