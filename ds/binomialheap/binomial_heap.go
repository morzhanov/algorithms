package binomialheap

import (
	"fmt"
	"math/rand"
	"time"
)

// Node is a Binomial Heap Node implementation
type Node struct {
	data, degree           int
	child, parent, sibling *Node
}

// BinomialHeap is a Binomial Heap implementation
type BinomialHeap struct {
	head      *Node
	nodeCount int
}

// Clear function clears the heap's data, making it an empty heap.
func (h *BinomialHeap) Clear() {
	h.head = nil
	h.nodeCount = 0
}

// ExtractMin function extracts and returns the minimum node from the heap
func (h *BinomialHeap) ExtractMin() *Node {
	if h.head == nil {
		return nil
	}

	min := h.head
	var minPrev *Node
	next := min.sibling
	nextPrev := min

	for next != nil {
		if compareNodes(next, min) < 0 {
			min = next
			minPrev = nextPrev
		}
		nextPrev = next
		next = next.sibling
	}

	removeRoot(h, min, minPrev)
	h.nodeCount--
	return min
}

// GetMin function returns the minimum node from the heap
func (h *BinomialHeap) GetMin() *Node {
	if h.head == nil {
		return nil
	}

	min := h.head
	next := min.sibling
	for next != nil {
		if compareNodes(next, min) < 0 {
			min = next
		}
		next = next.sibling
	}
	return min
}

// Insert function inserts a new data node into the heap
func (h *BinomialHeap) Insert(data int) *Node {
	var tempHeap BinomialHeap
	var newNode Node
	newNode.data = data
	tempHeap.head = &newNode
	tempHeap.nodeCount++
	h.Union(&tempHeap)
	return &newNode
}

// IsEmpty function returns true if heap is empty
func (h *BinomialHeap) IsEmpty() bool {
	return h.head == nil
}

// Size function returns heap size
func (h *BinomialHeap) Size() int {
	return h.nodeCount
}

// Union function joins another heap to the current
func (h *BinomialHeap) Union(other *BinomialHeap) {
	h.nodeCount += other.nodeCount
	newHead := h.MergeHeap(other)
	h.head = nil
	other.head = nil
	if newHead == nil {
		return
	}

	var prev *Node
	curr := newHead
	next := newHead.sibling
	for next != nil {
		if curr.degree != next.degree ||
			(next.sibling != nil && next.sibling.degree == curr.degree) {
			prev = curr
			curr = next
		} else if compareNodes(curr, next) < 0 {
			curr.sibling = next.sibling
			linkTrees(curr, next)
		} else {
			if prev == nil {
				newHead = next
			} else {
				prev.sibling = next
			}
			linkTrees(next, curr)
			curr = next
		}
		next = curr.sibling
	}
	h.head = newHead
}

// MergeHeap function merges another heap with the current
// returns the head of the merged heap.
func (h *BinomialHeap) MergeHeap(other *BinomialHeap) *Node {
	if h.head == nil {
		return other.head
	}
	if other.head == nil {
		return h.head
	}

	var head *Node
	hNext := h.head
	otherNext := other.head
	if h.head.degree <= other.head.degree {
		head = h.head
		hNext = hNext.sibling
	} else {
		head = other.head
		otherNext = otherNext.sibling
	}

	tail := head
	for hNext != nil && otherNext != nil {
		if hNext.degree <= otherNext.degree {
			tail.sibling = hNext
			hNext = hNext.sibling
		} else {
			tail.sibling = otherNext
			otherNext = otherNext.sibling
		}
		tail = tail.sibling
	}

	if hNext != nil {
		tail.sibling = hNext
	} else {
		tail.sibling = otherNext
	}
	return head
}

// removeBinomialTreeRoot function links two binomial trees of the same order
// minNodeTree - the head of the first tree to link.
// other - the head of the second tree to link.
func removeRoot(heap *BinomialHeap, root *Node, prev *Node) {
	// remove root from the heap
	if root == heap.head {
		heap.head = root.sibling
	} else {
		prev.sibling = root.sibling
	}

	// reverse the order of root's children and make a new heap
	var newHead Node
	child := root.child
	for child != nil {
		next := child.sibling
		child.sibling = &newHead
		child.parent = nil
		newHead = *child
		child = next
	}
	var newHeap BinomialHeap
	newHeap.head = &newHead
	heap.Union(&newHeap)
}

// linkTrees function links two binomial trees of the same order.
// minNodeTree - the head of the first tree to link.
// other - the head of the second tree to link.
func linkTrees(minNodeTree *Node, other *Node) {
	other.parent = minNodeTree
	other.sibling = minNodeTree.child
	minNodeTree.child = other
	minNodeTree.degree++
}

// compareNodes function compares two nodes with each other
// returns -1, 0 or 1 if a < b, a == b or a > b respectively
func compareNodes(a *Node, b *Node) int {
	if a.data > b.data {
		return 1
	}
	if a.data < b.data {
		return -1
	}
	return 0
}

func traverse(h *BinomialHeap) [][]int {
	res := make([][]int, 0)
	curr := h.head
	for curr != nil {
		res = append(res, traverseTree(curr))
		curr = curr.sibling
	}
	return res
}

func traverseTree(node *Node) []int {
	nodes := make([]int, 0)
	nodes = append(nodes, node.data)
	if node.child != nil {
		nodes = append(nodes, traverseTree(node.child)...)
	}
	if node.sibling != nil {
		nodes = append(nodes, traverseTree(node.sibling)...)
	}
	return nodes
}

func printTraverse(heap *BinomialHeap, data [][]int) {
	fmt.Printf("Traversing binominal heap after insert. Heap size = %v\n", heap.Size())
	for i, el := range data {
		fmt.Printf("№%v tree = %v\n", i, el)
	}
}

// Test method tests binomial heap
func Test(heap BinomialHeap) {
	heap.Insert(3)
	printTraverse(&heap, traverse(&heap))
	heap.Insert(7)
	heap.Insert(8)
	heap.Insert(1)
	heap.Insert(2)
	printTraverse(&heap, traverse(&heap))
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(9)
	printTraverse(&heap, traverse(&heap))
	heap.Insert(11)
	heap.Insert(81)
	heap.Insert(2)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(1)
	printTraverse(&heap, traverse(&heap))

	heap.Clear()
	fmt.Printf("Cleared heap. Heap size = %v\n", heap.Size())

	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn((i + 1) * 10)
		heap.Insert(r)
		printTraverse(&heap, traverse(&heap))
	}
}
