package ds

import (
	"errors"
	"fmt"
	"math"
)

// TreeNode is a Tree node implementation
type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	value  int
}

// GetLeftHeight method returns node left height
func (tn *TreeNode) GetLeftHeight() int {
	if tn.left == nil {
		return 0
	}
	return tn.left.GetHeight() + 1
}

// GetRightHeight method returns node right height
func (tn *TreeNode) GetRightHeight() int {
	if tn.right == nil {
		return 0
	}
	return tn.right.GetHeight() + 1
}

// GetHeight method returns node height
func (tn *TreeNode) GetHeight() int {
	return int(math.Max(float64(tn.GetLeftHeight()), float64(tn.GetRightHeight())))
}

// GetBalanceFactor returns node balance factor
func (tn *TreeNode) GetBalanceFactor() int {
	return tn.GetLeftHeight() - tn.GetRightHeight()
}

// Uncle method returns node uncle
func (tn *TreeNode) Uncle() *TreeNode {
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
func (tn *TreeNode) SetValue(value int) {
	tn.value = value
}

// SetLeft method sets left child
func (tn *TreeNode) SetLeft(node *TreeNode) {
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
func (tn *TreeNode) SetRight(node *TreeNode) {
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
func (tn *TreeNode) RemoveChild(nodeToRemove *TreeNode) error {
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
func (tn *TreeNode) ReplaceChild(nodeToReplace *TreeNode, replacementNode *TreeNode) error {
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
func (tn *TreeNode) Traverse() []*TreeNode {
	leftTraverse := make([]*TreeNode, 0)
	rightTraverse := make([]*TreeNode, 0)
	if tn.left != nil {
		leftTraverse = tn.left.Traverse()
	}
	if tn.right != nil {
		rightTraverse = tn.right.Traverse()
	}

	res := make([]*TreeNode, 1)
	res[0] = tn
	res = append(res, leftTraverse...)
	res = append(res, rightTraverse...)
	return res
}

// Insert inserts the Item t in the tree
func (tn *TreeNode) Insert(value int) {
	newNode := &TreeNode{nil, nil, nil, value}
	insertNode(tn, newNode)
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *TreeNode) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// TestTree method tests tree
func TestTree(tn TreeNode) {
	// add nodes
	values := [15]int{1, 5, 2, 7, 3, 9, 0, 5, 6, 2, 4, 1, 9, 11, 5}
	for i := 1; i < len(values); i++ {
		tn.Insert(values[i])
	}
	traverse := tn.Traverse()
	fmt.Printf("Traversing Tree...\n")
	for i := 0; i < len(traverse); i++ {
		fmt.Printf("Node #%v: Value = %v; Left = %v; Right = %v\n", i+1, traverse[i].value, traverse[i].left, traverse[i].right)
	}
	fmt.Printf("Traversing Tree completed\n")

	h := tn.GetHeight()
	fmt.Printf("Tree hight: %v\n", h)
}
