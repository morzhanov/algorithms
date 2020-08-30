package redblacktree

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/morzhanov/algorithms/src/utils"
)

// RedBlackTree represents Red-Black tree implementation
// 1) Every node has a color either red or black.
// 2) Root of tree is always black.
// 3) There are no two adjacent red nodes (A red node cannot have a red parent or red child).
// 4) Every path from a node (including root) to any of its descendant NULL node has the same number of black nodes.
type RedBlackTree struct {
	root      *Node
	cursor    *Node
	ancestors []*Node
}

// CreateTree function creates and returns red-black tree
func CreateTree(rootValue int) *RedBlackTree {
	root := CreateNode(rootValue)
	return &RedBlackTree{root: &root, cursor: &root, ancestors: make([]*Node, 0)}
}

func (t *RedBlackTree) findNode(value int, saveAncestors bool) *Node {
	res := t.root
	if saveAncestors {
		t.ancestors = make([]*Node, 0)
	}

	for res != nil {
		relation := utils.IntComparator(value, res.value)
		if relation == 0 {
			break
		}
		if saveAncestors {
			t.ancestors = append(t.ancestors, res)
		}
		if relation < 0 {
			return res.left
		}
		return res.right
	}
	return res
}

func (t *RedBlackTree) maxNode(node *Node, saveAncestors bool) *Node {
	if node == nil {
		node = t.root
	}
	if node == nil {
		return node
	}
	for node.right != nil {
		if saveAncestors {
			t.ancestors = append(t.ancestors, node)
		}
		node = node.right
	}
	return node
}

func (t *RedBlackTree) minNode(node *Node, saveAncestors bool) *Node {
	if node == nil {
		node = t.root
	}
	if node == nil {
		return node
	}
	for node.left != nil {
		if saveAncestors {
			t.ancestors = append(t.ancestors, node)
		}
		node = node.left
	}
	return node
}

func (t *RedBlackTree) nextNode(node *Node) *Node {
	if node == nil {
		t.ancestors = make([]*Node, 0)
		return t.minNode(t.root, true)
	}
	if node.right != nil {
		t.ancestors = append(t.ancestors, node)
		return t.minNode(node.right, true)
	}
	ancestors := t.ancestors
	parent, ancestors := ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1]
	for parent != nil && parent.right == node {
		node = parent
		parent, ancestors = ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1]
	}
	return parent
}

func (t *RedBlackTree) prevNode(node *Node) *Node {
	if node == nil {
		t.ancestors = make([]*Node, 0)
		return t.maxNode(t.root, true)
	}
	if node.left != nil {
		t.ancestors = append(t.ancestors, node)
		return t.maxNode(node.left, true)
	}
	ancestors := t.ancestors
	parent, ancestors := ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1]
	for parent != nil && parent.left == node {
		node = parent
		parent, ancestors = ancestors[len(ancestors)-1], ancestors[:len(ancestors)-1]
	}
	return parent
}

// Add function adds node to the rbtree
func (t *RedBlackTree) Add(value int) *Node {
	if t.root == nil {
		root := CreateNode(value)
		t.root = &root
		return t.root
	}
	addRes := t.root.Add(value)
	t.root = addRes[0]
	return addRes[1]
}

// Find function returns node from rbtree
func (t *RedBlackTree) Find(value int) *Node {
	return t.findNode(value, false)
}

// FindNext function returns next node from rbtree
func (t *RedBlackTree) FindNext(value int) *Node {
	curr := t.findNode(value, true)
	return t.nextNode(curr)
}

// FindPrev function returns previous node from rbtree
func (t *RedBlackTree) FindPrev(value int) *Node {
	curr := t.findNode(value, true)
	return t.prevNode(curr)
}

// Max function returns node with max value from rbtree
func (t *RedBlackTree) Max(value int) *Node {
	return t.maxNode(nil, false)
}

// Min function returns node with min value from rbtree
func (t *RedBlackTree) Min(value int) *Node {
	return t.minNode(nil, false)
}

// Next function returns next node via cursor
func (t *RedBlackTree) Next(value int) *Node {
	return t.nextNode(t.cursor)
}

// Prev function returns previous node via cursor
func (t *RedBlackTree) Prev(value int) *Node {
	return t.prevNode(t.cursor)
}

// Remove function removes node from rbtree
func (t *RedBlackTree) Remove(value int) *Node {
	if t.root == nil {
		return nil
	}
	remRes := t.root.Remove(value)
	t.root = remRes[0]
	return remRes[1]
}

// Traverse function traverses rbtree
func (t *RedBlackTree) Traverse() {
	if t.root != nil {
		fmt.Printf("Traversing Red-Black Tree\n")
		t.root.Traverse()
	}
}

// Test function tests Red-Black tree
func Test(t *RedBlackTree) {
	addedValues := make([]int, 11)
	fmt.Printf("Creating Red-Black tree\n")
	for i := 0; i < 11; i++ {
		num := 1
		rand.Seed(time.Now().UnixNano())
		num = rand.Intn((i + 1) * 10)
		fmt.Printf("%v\n", num)
		addedValues[i] = num
		t.Add(num)
	}
	fmt.Printf("Created Red-Black Tree\n")
	t.Traverse()

	fmt.Printf("Removed %v value from tree\n", addedValues[0])
	t.Traverse()
	fmt.Printf("Removed %v value from tree\n", addedValues[10])
	t.Traverse()
	fmt.Printf("Value %v found in tree in node: %v\n", addedValues[8], t.Find(addedValues[8]))
}
