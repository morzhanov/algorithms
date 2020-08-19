package ds

import (
	"errors"
	"fmt"
)

// TreeNode is a Tree node implementation
type TreeNode struct {
	children []*TreeNode
	parent   *TreeNode
	value    int
}

// GetHeight method returns node height
func (tn *TreeNode) GetHeight() int {
	if len(tn.children) == 0 {
		return 1
	}
	max := 0
	for _, child := range tn.children {
		curr := child.GetHeight()
		if curr > max {
			max = curr
		}
	}
	return max + 1
}

// SetValue method sets node value
func (tn *TreeNode) SetValue(value int) {
	tn.value = value
}

// RemoveChild method removes node child
func (tn *TreeNode) RemoveChild(nodeToRemove *TreeNode) error {
	for idx, child := range tn.children {
		if child == nodeToRemove {
			tn.children = append(tn.children[:idx], tn.children[idx+1:]...)
			return nil
		}
	}
	return errors.New(("Node has no children"))
}

// ReplaceChild method replaces node child
func (tn *TreeNode) ReplaceChild(nodeToReplace *TreeNode, replacementNode *TreeNode) error {
	for idx, child := range tn.children {
		if child == nodeToReplace {
			tn.children[idx] = replacementNode
			return nil
		}
	}
	return errors.New(("Node not found"))
}

// Traverse tree
func (tn *TreeNode) Traverse() []*TreeNode {
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, tn)
	for _, child := range tn.children {
		nodes = append(nodes, child.Traverse()...)
	}
	return nodes
}

// AddChild inserts new node to the node children array
func (tn *TreeNode) AddChild(n *TreeNode) {
	n.parent = tn
	tn.children = append(tn.children, n)
}

// TestTree method tests tree
func TestTree(tn TreeNode) {
	// add 5 children nodes to root and 5 nodes to each root child
	for i := 0; i < 5; i++ {
		node := TreeNode{value: i}
		tn.AddChild(&node)
	}
	for _, child := range tn.children {
		for i := 0; i < 5; i++ {
			node := TreeNode{value: i}
			child.AddChild(&node)
		}
	}

	traverse := tn.Traverse()
	fmt.Printf("Traversing Tree...\n")
	for i := 0; i < len(traverse); i++ {
		fmt.Printf("Node #%v: Value = %v; Parent = %v; Children count = %v\n", i+1, traverse[i].value, traverse[i].parent, len(traverse[i].children))
	}
	fmt.Printf("Traversing Tree completed\n")

	h := tn.GetHeight()
	fmt.Printf("Tree hight: %v\n", h)
}
