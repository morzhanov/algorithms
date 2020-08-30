package redblacktree

// Node represents Red-Black tree node implementation
type Node struct {
}

// RedBlackTree represents Red-Black tree implementation
// 1) Every node has a color either red or black.
// 2) Root of tree is always black.
// 3) There are no two adjacent red nodes (A red node cannot have a red parent or red child).
// 4) Every path from a node (including root) to any of its descendant NULL node has the same number of black nodes.
type RedBlackTree struct {
	root *Node
}
