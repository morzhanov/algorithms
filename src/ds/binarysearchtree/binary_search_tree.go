package binarysearchtree

import (
	"errors"
	"fmt"
	"math"
)

// Node is a BinarySearchTree node implementation
type Node struct {
	left   *Node
	right  *Node
	parent *Node
	value  int
}

// GetLeftHeight method returns node left height
func (tn *Node) GetLeftHeight() int {
	if tn.left == nil {
		return 0
	}
	return tn.left.GetHeight() + 1
}

// GetRightHeight method returns node right height
func (tn *Node) GetRightHeight() int {
	if tn.right == nil {
		return 0
	}
	return tn.right.GetHeight() + 1
}

// GetHeight method returns node height
func (tn *Node) GetHeight() int {
	return int(math.Max(float64(tn.GetLeftHeight()), float64(tn.GetRightHeight())))
}

// GetValue method returns node value
func (tn *Node) GetValue() int {
	return tn.value
}

// GetBalanceFactor returns node balance factor
func (tn *Node) GetBalanceFactor() int {
	return tn.GetLeftHeight() - tn.GetRightHeight()
}

// Uncle method returns node uncle
func (tn *Node) Uncle() *Node {
	if tn.parent == nil || tn.parent.parent == nil {
		return nil
	}

	// Check if grand-parent has two children.
	if tn.parent.parent.left == nil || tn.parent.parent.right == nil {
		return nil
	}

	// So for now we know that current node has grand-parent and this
	// grand-parent has two children. Let's find out who is the uncle.
	if tn.parent == tn.parent.parent.left {
		return tn.parent.parent.right
	}
	return tn.parent.parent.left
}

// SetValue method sets node value
func (tn *Node) SetValue(value int) {
	tn.value = value
}

// SetLeft method sets left child
func (tn *Node) SetLeft(node *Node) {
	// Reset parent for left node since it is going to be detached.
	if tn.left != nil {
		tn.left.parent = nil
	}

	// Attach new node to the left.
	tn.left = node

	// Make current node to be a parent for new left one.
	if tn.left != nil {
		tn.left.parent = tn
	}
}

// SetRight method sets right child
func (tn *Node) SetRight(node *Node) {
	// Reset parent for right node since it is going to be detached.
	if tn.right != nil {
		tn.right.parent = nil
	}

	// Attach new node to the right.
	tn.right = node

	// Make current node to be a parent for new right one.
	if tn.right != nil {
		tn.right.parent = tn
	}
}

// RemoveChild method removes node child
func (tn *Node) RemoveChild(nodeToRemove *Node) error {
	if tn.left != nil && tn.left == nodeToRemove {
		tn.left.parent = nil
		tn.left = nil
		return nil
	}
	if tn.right != nil && tn.right == nodeToRemove {
		tn.right.parent = nil
		tn.right = nil
		return nil
	}
	return errors.New(("Node has no children"))
}

// ReplaceChild method replaces node child
func (tn *Node) ReplaceChild(nodeToReplace *Node, replacementNode *Node) error {
	if tn.left != nil && tn.left == nodeToReplace {
		tn.left.parent = nil
		tn.left = replacementNode
		return nil
	}
	if tn.right != nil && tn.right == nodeToReplace {
		tn.right.parent = nil
		tn.right = replacementNode
		return nil
	}
	return errors.New(("Node has no children"))
}

// Traverse tree
func (tn *Node) Traverse() []*Node {
	leftTraverse := make([]*Node, 0)
	rightTraverse := make([]*Node, 0)
	if tn.left != nil {
		leftTraverse = tn.left.Traverse()
	}
	if tn.right != nil {
		rightTraverse = tn.right.Traverse()
	}

	res := make([]*Node, 1)
	res[0] = tn
	res = append(res, leftTraverse...)
	res = append(res, rightTraverse...)
	return res
}

// Insert inserts the Item t in the tree
func (tn *Node) Insert(value int) {
	newNode := &Node{nil, nil, nil, value}
	insert(tn, newNode)
}

// internal function to find the correct place for a node in a tree
func insert(node, newNode *Node) {
	if newNode.value < node.value {
		if node.left == nil {
			newNode.parent = node
			node.left = newNode
		} else {
			insert(node.left, newNode)
		}
	} else {
		if node.right == nil {
			newNode.parent = node
			node.right = newNode
		} else {
			insert(node.right, newNode)
		}
	}
}

// Search function searches node by value
func (tn *Node) Search(value int) *Node {
	if tn.value == value {
		return tn
	}
	if tn.value > value && tn.left != nil {
		return tn.left.Search(value)
	}
	if tn.value < value && tn.right != nil {
		return tn.right.Search(value)
	}
	return nil
}

// InorderTraversal is a function to do inorder traversal of BST
func (tn *Node) InorderTraversal() []int {
	res := make([]int, 0)
	if tn != nil {
		res = append(res, tn.left.InorderTraversal()...)
		res = append(res, tn.value)
		res = append(res, tn.right.InorderTraversal()...)
	}
	return res
}

// Test method tests binary tree
func Test(tn Node) {
	// add nodes
	values := [15]int{1, 5, 2, 7, 3, 9, 0, 5, 6, 2, 4, 1, 9, 11, 5}
	for i := 1; i < len(values); i++ {
		tn.Insert(values[i])
	}
	traverse := tn.Traverse()
	fmt.Printf("Traversing BinarySearchTree...\n")
	for i := 0; i < len(traverse); i++ {
		fmt.Printf("Node #%v: Value = %v; Left = %v; Right = %v\n", i+1, traverse[i].value, traverse[i].left, traverse[i].right)
	}
	fmt.Printf("Traversing BinarySearchTree completed\n")

	h := tn.GetHeight()
	fmt.Printf("BinarySearchTree hight: %v\n", h)

	val := 7
	node := tn.Search(val)
	fmt.Printf("BinarySearchTree node with value %v found: %v\n", val, node)

	val = 100
	node = tn.Search(val)
	fmt.Printf("BinarySearchTree node with value %v not found: %v\n", val, node)
}
