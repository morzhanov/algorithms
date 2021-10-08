package redblacktree

import (
	"fmt"

	"github.com/morzhanov/algorithms/utils"
)

// Node represents RedBlackTree node implementation
type Node struct {
	left   *Node
	right  *Node
	value  int
	height int
}

// CreateNode function returns new RedBlackTree node
func CreateNode(value int) Node {
	return Node{value: value, height: 1}
}

// Add function adds new node to the current node
func (n *Node) Add(value int) []*Node {
	var newNode *Node
	relation := utils.IntComparator(value, n.value)
	addRes := make([]*Node, 0)

	if relation == 0 {
		return []*Node{n, n}
	}
	if relation < 0 {
		if n.left != nil {
			addRes = n.left.Add(value)
			n.left = addRes[0]
			newNode = addRes[1]
		} else {
			node := CreateNode(value)
			n.left = &node
			newNode = n.left
		}
	} else {
		if n.right != nil {
			addRes = n.right.Add(value)
			n.right = addRes[0]
			newNode = addRes[1]
		} else {
			node := CreateNode(value)
			n.right = &node
			newNode = n.right
		}
	}
	return []*Node{n.BalanceTree(), newNode}
}

// BalanceTree function balances rbtree node
func (n *Node) BalanceTree() *Node {
	var leftH, rightH int
	var res *Node
	if n.left != nil {
		leftH = n.left.height
	} else {
		leftH = 0
	}
	if n.right != nil {
		rightH = n.right.height
	} else {
		rightH = 0
	}

	if leftH > rightH+1 {
		res = n.SwingRight()
	} else if rightH > leftH+1 {
		res = n.SwingLeft()
	} else {
		n.SetHeight()
		res = n
	}
	return res
}

// Join function joins rbtree node
func (n *Node) Join(node *Node) *Node {
	if node == nil {
		return n
	}
	top := n
	if n.height > node.height {
		top = n
		top.right = node.Join(top.right)
	} else {
		top = node
		top.left = n.Join(top.left)
	}
	return top.BalanceTree()
}

// MoveLeft function moves node to the left
func (n *Node) MoveLeft() *Node {
	right := n.right
	if right == nil {
		return nil
	}
	rightLeft := right.left

	n.right = rightLeft
	right.left = n
	n.SetHeight()
	right.SetHeight()
	return right
}

// MoveRight function moves node to the right
func (n *Node) MoveRight() *Node {
	left := n.left
	leftRight := left.right

	n.left = leftRight
	left.right = n
	n.SetHeight()
	left.SetHeight()
	return left
}

// Remove function removes node
func (n *Node) Remove(value int) []*Node {
	relation := utils.IntComparator(value, n.value)
	var remNode *Node
	var remRes []*Node
	var res *Node

	if relation != 0 {
		if relation < 0 {
			remNode = n
			if n.left == nil {
				remRes = n.left.Remove(value)
				n.left = remRes[0]
				remNode = remRes[1]
			} else {
				remNode = nil
			}
		} else {
			if n.right != nil {
				remRes = n.right.Remove(value)
				n.right = remRes[0]
				remNode = remRes[1]
			} else {
				remNode = nil
			}
		}
		res = n
	} else {
		remNode = n
		if n.left == nil {
			res = n.right
		} else if n.right == nil {
			res = n.left
		} else {
			res = n.left.Join(n.right)
			n.left = nil
			n.right = nil
		}
	}

	if remNode == nil {
		return []*Node{n, nil}
	}
	if res != nil {
		return []*Node{res.BalanceTree(), remNode}
	}
	return []*Node{res, remNode}
}

// SetHeight function calculates and sets node height
func (n *Node) SetHeight() {
	var leftH, rightH int
	if n.left != nil {
		leftH = n.left.height
	} else {
		leftH = 0
	}
	if n.right != nil {
		rightH = n.right.height
	} else {
		rightH = 0
	}

	if leftH < rightH {
		n.height = rightH + 1
	} else {
		n.height = leftH + 1
	}
}

// SwingLeft function checks node right children and moves left
func (n *Node) SwingLeft() *Node {
	right := n.right
	rightLeft := n.right.left
	rightRight := n.right.right
	var rightLeftH, rightRightH int

	if rightLeft != nil {
		rightLeftH = rightLeft.height
	} else {
		rightLeftH = 0
	}
	if rightRight != nil {
		rightRightH = rightRight.height
	} else {
		rightRightH = 0
	}

	if rightLeftH > rightRightH {
		n.right = right.MoveRight()
	}
	return n.MoveLeft()
}

// SwingRight function checks node left children and moves right
func (n *Node) SwingRight() *Node {
	left := n.left
	leftRight := n.left.right
	leftLeft := n.left.left
	var leftRightH, leftLeftH int

	if leftRight != nil {
		leftRightH = leftRight.height
	} else {
		leftRightH = 0
	}
	if leftLeft != nil {
		leftLeftH = leftLeft.height
	} else {
		leftLeftH = 0
	}

	if leftRightH > leftLeftH {
		n.left = left.MoveLeft()
	}
	return n.MoveLeft()
}

// Traverse function performs node traverse
func (n *Node) Traverse() {
	if n.left != nil {
		n.left.Traverse()
	}
	fmt.Printf("value: %v | height: %v\n", n.value, n.height)
	if n.right != nil {
		n.right.Traverse()
	}
}
