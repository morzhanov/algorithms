package btree

import (
	"fmt"
	"math"
)

// NodeValue represents node boxed value
type NodeValue struct {
	value int
}

// Node is an implementation of BTree node
type Node struct {
	keysCount int
	keys      []*NodeValue
	children  []*Node
	tree      *BTree
}

// NodeSplitRes represents node.split function result
type NodeSplitRes struct {
	left  *Node
	key   *NodeValue
	right *Node
}

// CreateNode function creates BTree node
func CreateNode(t *BTree) *Node {
	node := Node{0, make([]*NodeValue, t.order), make([]*Node, t.order), t}
	return &node
}

func (n *Node) isLeaf() bool {
	return n.children[0] == nil
}

func (n *Node) isFull() bool {
	return n.keysCount == n.tree.order
}

func (n *Node) add(key int) *NodeSplitRes {
	if n.isLeaf() && n.isFull() {
		return n.split(key, nil)
	}
	if n.isLeaf() && !n.isFull() {
		n.insertKey(key)
		return nil
	}

	child := n.getChildContaining(key)
	split := child.add(key)
	if split == nil {
		return nil
	}
	if n.isFull() {
		return n.split(split.key.value, split.right)
	}
	n.insertSplit(split)
	return nil
}

func (n *Node) insertKey(key int) {
	pos := n.keysCount
	keys := n.keys
	for pos > 0 && keys[pos-1] != nil && keys[pos-1].value > key {
		keys[pos] = keys[pos-1]
		pos--
	}
	val := NodeValue{key}
	keys[pos] = &val
	n.keysCount++
}

func (n *Node) insertSplit(split *NodeSplitRes) {
	child := split.left

	// insert key with right child poped up from child node
	// case A: first child was split
	if child == n.children[0] {
		for i := n.keysCount; i > 0; i-- {
			n.keys[i] = n.keys[i-1]
		}
		n.keys[0] = split.key
		for i := n.keysCount + 1; i > 1; i-- {
			n.children[i] = n.children[i-1]
		}
		n.children[0] = child
		n.children[1] = split.right
		n.keysCount++
		return
	}

	// case B: [key][split-child] (split child is on the right)
	pos := n.keysCount
	for pos > 0 && n.children[pos] != child {
		n.keys[pos] = n.keys[pos-1]
		n.children[pos+1] = n.children[pos]
		pos--
	}
	n.keys[pos] = split.key
	n.children[pos+1] = split.right
	n.keysCount++
}

func (n *Node) getChildContaining(key int) *Node {
	for i := 0; i < n.keysCount; i++ {
		if key <= n.keys[i].value {
			return n.children[i]
		}
	}
	return n.children[n.keysCount]
}

func (n *Node) split(key int, keyRightChild *Node) *NodeSplitRes {
	left := n
	right := CreateNode(n.tree)

	keys := n.keys[:]
	keys = append(keys, nil)
	children := n.children[:]
	children = append(children, nil)

	// find new key position
	pos := len(keys) - 1
	for pos > 0 && keys[pos-1].value > key {
		keys[pos] = keys[pos-1]
		children[pos+1] = children[pos]
		pos--
	}
	val := NodeValue{key}
	keys[pos] = &val
	children[pos] = keyRightChild

	// split into two children and key
	medianIdx := int(math.Floor(float64(len(keys)) / 2))
	medianKey := n.keys[medianIdx]

	for i := 0; i < len(keys)-1; i++ {
		if i < medianIdx {
			left.children[i] = children[i]
			left.keys[i] = keys[i]
		} else if i == medianIdx {
			left.children[i] = children[i]
			left.keys[i] = nil
		}
		n.keys[i] = nil
		left.children[i] = nil
	}
	left.keysCount = medianIdx

	for i := 0; i < len(keys); i++ {
		if i > medianIdx {
			right.keys[i-medianIdx-1] = keys[i]
			right.children[i-medianIdx-1] = children[i]
			right.keysCount++
		}
	}
	right.children[len(keys)-medianIdx-1] = children[len(keys)-1]
	res := NodeSplitRes{left, medianKey, right}
	return &res
}

func (n *Node) remove(key int) bool {
	if n.isLeaf() {
		return n.removeKey(key)
	}
	keyIdx := n.indexOfKey(key)
	if keyIdx == -1 {
		child := n.getChildContaining(key)
		res := child.remove(key)
		n.rebalance(findIndexOfElement(n.children, child))
		return res
	}
	child := n.children[keyIdx]
	n.keys[keyIdx] = child.extractMax()
	n.rebalance(keyIdx)
	return true
}

func findIndexOfElement(arr []*Node, value *Node) int {
	for idx, el := range arr {
		if el == value {
			return idx
		}
	}
	return -1
}

func (n *Node) rebalance(childIdx int) {
	minKeys := n.tree.order / 2
	child := n.children[childIdx]
	if child.keysCount >= minKeys {
		return
	}

	// borrow from left child
	if childIdx != 0 {
		leftChild := n.children[childIdx-1]
		if leftChild.keysCount > minKeys {
			lastKey := leftChild.keys[leftChild.keysCount-1]
			lastChild := leftChild.children[leftChild.keysCount]
			leftChild.keysCount--

			key := n.keys[childIdx-1]
			n.keys[childIdx-1] = lastKey

			for i := child.keysCount - 1; i >= 0; i-- {
				child.keys[i+1] = child.keys[i]
			}
			child.keys[0] = key
			for i := child.keysCount; i >= 0; i-- {
				child.children[i+1] = child.children[i]
			}
			child.children[0] = lastChild
			child.keysCount++
			return
		}
	}

	// borrow from right child
	if childIdx < n.keysCount {
		rightChild := n.children[childIdx+1]
		if rightChild.keysCount > minKeys {
			firstKey := rightChild.keys[0]
			firstChild := rightChild.children[0]
			for i := 0; i < rightChild.keysCount-1; i++ {
				rightChild.keys[i] = rightChild.keys[i+1]
			}
			for i := 0; i < rightChild.keysCount; i++ {
				rightChild.children[i] = rightChild.children[i+1]
			}
			rightChild.keysCount--
			child.keys[child.keysCount] = n.keys[childIdx]
			n.keys[childIdx] = firstKey
			child.children[child.keysCount+1] = firstChild
			child.keysCount++
			return
		}
	}

	// merge
	if childIdx != 0 {
		// merge left and current
		childIdx--
	}

	// childIndex will point to the *left* node of two merged nodes
	merged := n.mergeChildren(childIdx)
	for i := childIdx; i < n.keysCount-1; i++ {
		n.keys[i] = n.keys[i+1]
	}
	for i := childIdx; i < n.keysCount; i++ {
		n.children[i] = n.children[i+1]
	}
	n.keysCount--
	n.children[childIdx] = merged
}

func (n *Node) mergeChildren(leftIdx int) *Node {
	key := n.keys[leftIdx]
	left := n.children[leftIdx]
	right := n.children[leftIdx+1]
	left.keys[left.keysCount] = key
	left.keysCount++

	// copy right keys and childs into left
	for i := 0; i < right.keysCount; i++ {
		left.children[left.keysCount] = right.children[i]
		left.keys[left.keysCount] = right.keys[i]
		left.keysCount++
	}
	left.children[left.keysCount] = right.children[right.keysCount]
	return left
}

func (n *Node) extractMax() *NodeValue {
	var key *NodeValue
	if n.isLeaf() {
		key = n.keys[n.keysCount-1]
		n.keysCount--
	} else {
		child := n.children[n.keysCount]
		key = child.extractMax()
		n.rebalance(n.keysCount)
	}
	return key
}

func (n *Node) indexOfKey(key int) int {
	for i := 0; i < n.keysCount; i++ {
		if n.keys[i].value == key {
			return i
		}
	}
	return -1
}

func (n *Node) removeKey(key int) bool {
	keyIdx := n.indexOfKey(key)
	if keyIdx == -1 {
		return false
	}
	// delete key
	for i := keyIdx + 1; i < n.keysCount; i++ {
		n.keys[i-1] = n.keys[i]
	}
	n.keysCount--
	return true
}

func (n *Node) print() string {
	const indent = " "
	const indentOpt = ""

	if n.isLeaf() {
		curr := ""
		keys := n.keys[0:n.keysCount]
		for idx, key := range keys {
			if key == nil {
				continue
			}
			if idx == 0 {
				curr += fmt.Sprintf("%v", key.value)
				continue
			}
			curr += fmt.Sprintf(", %v", key.value)
		}
		if curr == "" {
			return ""
		}
		return indentOpt + "[" + curr + "]"
	}

	childIndent := indentOpt + indent
	childStrings := make([]string, 0)
	temp := n.children[0 : n.keysCount+1]
	for _, child := range temp {
		childStrings = append(childStrings, child.print())
	}

	str := indentOpt + "[\n" + childStrings[0] + "\n"
	for i := 1; i < len(childStrings); i++ {
		str += fmt.Sprintf("%v %v \n %v \n", childIndent, n.keys[i-1].value, childStrings[i])
	}
	str += indentOpt + "]"
	return str
}

func fromSplit(t *BTree, split *NodeSplitRes) *Node {
	node := CreateNode(t)
	node.keysCount = 1
	node.keys[0] = split.key
	node.children[0] = split.left
	node.children[1] = split.right
	return node
}
