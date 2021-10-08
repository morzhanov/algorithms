package btree

// BTree is a BTree implementation
type BTree struct {
	root  *Node
	order int
}

// CreateBTree function creates BTree
func CreateBTree(order int, initialKey int) *BTree {
	bt := BTree{nil, order}
	node := CreateNode(&bt)
	val := NodeValue{initialKey}
	node.keys[0] = &val
	node.keysCount++
	bt.root = node
	return &bt
}

// Add function adds key to BTree
func (t *BTree) Add(key int) {
	curr := t.root
	split := curr.add(key)
	if split == nil {
		return
	}
	t.root = fromSplit(t, split)
}

// Remove function removes key from BTree
func (t *BTree) Remove(key int) bool {
	removed := t.root.remove(key)
	if t.root.keysCount == 0 && t.root.children[0] != nil {
		t.root = t.root.children[0]
	}
	return removed
}

// Print function returns string representation of BTree
func (t *BTree) Print() string {
	return t.root.print()
}
