package avltree

import (
	"errors"
	"fmt"
	"math"

	"github.com/morzhanov/algorithms/utils"
)

// Node is AVL Tree node implementation
// AVL tree is a self-balancing Binary Search Tree (BST)
// where the difference between heights of left and right subtrees
// cannot be more than one for all nodes.
type Node struct {
	left   *Node
	right  *Node
	height int
	value  int
}

// A utility function to right rotate subtree rooted with y
func rightRotate(y *Node) *Node {
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2
	y.height = utils.IntMax(height(y.left), height(y.right)) + 1
	x.height = utils.IntMax(height(x.left), height(x.right)) + 1
	return x
}

// A utility function to left rotate subtree rooted with x
func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	x.height = utils.IntMax(height(x.left), height(x.right)) + 1
	y.height = utils.IntMax(height(y.left), height(y.right)) + 1
	return y
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
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

// GetBalanceFactor returns node balance factor
func (tn *Node) GetBalanceFactor() int {
	return tn.GetLeftHeight() - tn.GetRightHeight()
}

// SetValue method sets node value
func (tn *Node) SetValue(value int) {
	tn.value = value
}

// SetLeft method sets left child
func (tn *Node) SetLeft(node *Node) {
	// Attach new node to the left.
	tn.left = node
}

// SetRight method sets right child
func (tn *Node) SetRight(node *Node) {
	// Attach new node to the right.
	tn.right = node
}

func minValueNode(n *Node) *Node {
	curr := n
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

// RemoveNode method removes node from tree
// Should be called from root node
func (tn *Node) RemoveNode(value int) *Node {
	if tn == nil {
		return tn
	}
	if value < tn.value {
		tn.left = tn.left.RemoveNode(value)
	} else if value > tn.value {
		tn.right = tn.right.RemoveNode(value)
	} else {
		// node with only one child or no child
		if tn.left == nil || tn.right == nil {
			var temp *Node
			if temp == tn.left {
				temp = tn.right
			} else {
				temp = tn.left
			}
			if temp == nil {
				temp = tn
				tn = nil
			}
			tn = temp
		} else {
			// node with two children: Get the inorder successor (smallest in the right subtree)
			temp := minValueNode(tn.right)
			tn.value = temp.value
			tn.right = tn.right.RemoveNode(temp.value)
		}
	}

	// If the tree had only one node then return
	if tn == nil {
		return tn
	}

	// update height of the current node
	tn.height = utils.IntMax(height(tn.left), height(tn.right)) + 1
	balance := tn.GetBalanceFactor()
	// If this node becomes unbalanced, then there are 4 cases
	// Left Left Case
	if balance > 1 && tn.left.GetBalanceFactor() >= 0 {
		return rightRotate(tn)
	}
	// Left Right Case
	if balance > 1 && tn.left.GetBalanceFactor() < 0 {
		tn.left = leftRotate(tn.left)
		return rightRotate(tn)
	}
	// Right Right Case
	if balance < -1 && tn.right.GetBalanceFactor() <= 0 {
		return leftRotate(tn)
	}
	// Right Left Case
	if balance < -1 && tn.right.GetBalanceFactor() > 0 {
		tn.right = rightRotate(tn.right)
		return leftRotate(tn)
	}
	return tn
}

// ReplaceChild method replaces node child
func (tn *Node) ReplaceChild(nodeToReplace *Node, replacementNode *Node) error {
	if tn.left != nil && tn.left == nodeToReplace {
		tn.left = replacementNode
		return nil
	}
	if tn.right != nil && tn.right == nodeToReplace {
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
func (tn *Node) Insert(value int) *Node {
	newNode := &Node{nil, nil, 1, value}
	return insert(tn, newNode)
}

// internal function to find the correct place for a node in a tree
func insert(node, newNode *Node) *Node {
	if node == nil {
		return newNode
	}
	if newNode.value < node.value {
		node.left = insert(node.left, newNode)
	} else if newNode.value > node.value {
		node.right = insert(node.right, newNode)
	} else {
		return node
	}

	// Update height of this ancestor node
	node.height = 1 + utils.IntMax(height(node.left), height(node.right))
	// Get the balance factor of this ancestor
	balance := node.GetBalanceFactor()

	// If this node becomes unbalanced, then there are 4 cases
	// Left Left Case
	if balance > 1 && newNode.value < node.left.value {
		return rightRotate(node)
	}
	// Right Right Case
	if balance < -1 && newNode.value > node.right.value {
		return leftRotate(node)
	}
	// Left Right Case
	if balance > 1 && newNode.value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	// Right Left Case
	if balance < -1 && newNode.value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}
	return node
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

// Test method tests AVL tree
func Test(tn Node) {
	// add nodes
	values := [15]int{1, 5, 2, 7, 3, 9, 0, 5, 6, 2, 4, 1, 9, 11, 5}
	for i := 1; i < len(values); i++ {
		tn.Insert(values[i])
	}
	traverse := tn.Traverse()
	fmt.Printf("Traversing AVLTree...\n")
	for i := 0; i < len(traverse); i++ {
		fmt.Printf("Node #%v: Value = %v; Left = %v; Right = %v\n", i+1, traverse[i].value, traverse[i].left, traverse[i].right)
	}
	fmt.Printf("Traversing AVLTree completed\n")

	h := tn.GetHeight()
	fmt.Printf("AVLTree hight: %v\n", h)

	val := 7
	node := tn.Search(val)
	fmt.Printf("AVLTree node with value %v found: %v\n", val, node)

	val = 100
	node = tn.Search(val)
	fmt.Printf("AVLTree node with value %v not found: %v\n", val, node)
}
