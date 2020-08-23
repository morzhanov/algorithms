package ds

import (
	"fmt"
	"math/rand"
	"time"
)

// BinomialHeapNode is a Binomial Heap Node implementation
type BinomialHeapNode struct {
	data, degree           int
	child, parent, sibling *BinomialHeapNode
}

// BinomialHeap is a Binomial Heap implementation
type BinomialHeap struct {
	head      *BinomialHeapNode
	nodeCount int
}

// Clear function clears the heap's data, making it an empty heap.
func (h *BinomialHeap) Clear() {
	h.head = nil
	h.nodeCount = 0
}

// ExtractMin function extracts and returns the minimum node from the heap
func (h *BinomialHeap) ExtractMin() *BinomialHeapNode {
	if h.head == nil {
		return nil
	}

	min := h.head
	var minPrev *BinomialHeapNode
	next := min.sibling
	nextPrev := min

	for next != nil {
		if compareBinomialHeapNodes(next, min) < 0 {
			min = next
			minPrev = nextPrev
		}
		nextPrev = next
		next = next.sibling
	}

	removeBinomialTreeRoot(h, min, minPrev)
	h.nodeCount--
	return min
}

// GetMin function returns the minimum node from the heap
func (h *BinomialHeap) GetMin() *BinomialHeapNode {
	if h.head == nil {
		return nil
	}

	min := h.head
	next := min.sibling
	for next != nil {
		if compareBinomialHeapNodes(next, min) < 0 {
			min = next
		}
		next = next.sibling
	}
	return min
}

// Insert function inserts a new data node into the heap
func (h *BinomialHeap) Insert(data int) *BinomialHeapNode {
	var tempHeap BinomialHeap
	var newNode BinomialHeapNode
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

	var prev *BinomialHeapNode
	curr := newHead
	next := newHead.sibling
	for next != nil {
		if curr.degree != next.degree ||
			(next.sibling != nil && next.sibling.degree == curr.degree) {
			prev = curr
			curr = next
		} else if compareBinomialHeapNodes(curr, next) < 0 {
			curr.sibling = next.sibling
			linkBinomialTrees(curr, next)
		} else {
			if prev == nil {
				newHead = next
			} else {
				prev.sibling = next
			}
			linkBinomialTrees(next, curr)
			curr = next
		}
		next = curr.sibling
	}
	h.head = newHead
}

// MergeHeap function merges another heap with the current
// returns the head of the merged heap.
func (h *BinomialHeap) MergeHeap(other *BinomialHeap) *BinomialHeapNode {
	if h.head == nil {
		return other.head
	}
	if other.head == nil {
		return h.head
	}

	var head *BinomialHeapNode
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
func removeBinomialTreeRoot(heap *BinomialHeap, root *BinomialHeapNode, prev *BinomialHeapNode) {
	// remove root from the heap
	if root == heap.head {
		heap.head = root.sibling
	} else {
		prev.sibling = root.sibling
	}

	// reverse the order of root's children and make a new heap
	var newHead BinomialHeapNode
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

// linkBinomialTrees function links two binomial trees of the same order.
// minNodeTree - the head of the first tree to link.
// other - the head of the second tree to link.
func linkBinomialTrees(minNodeTree *BinomialHeapNode, other *BinomialHeapNode) {
	other.parent = minNodeTree
	other.sibling = minNodeTree.child
	minNodeTree.child = other
	minNodeTree.degree++
}

// compareBinomialHeapNodes function compares two nodes with each other
// returns -1, 0 or 1 if a < b, a == b or a > b respectively
func compareBinomialHeapNodes(a *BinomialHeapNode, b *BinomialHeapNode) int {
	if a.data > b.data {
		return 1
	}
	if a.data < b.data {
		return -1
	}
	return 0
}

func traverseBinomialHeap(h *BinomialHeap) [][]int {
	res := make([][]int, 0)
	curr := h.head
	for curr != nil {
		res = append(res, traverseBinomialTree(curr))
		curr = curr.sibling
	}
	return res
}

func traverseBinomialTree(node *BinomialHeapNode) []int {
	nodes := make([]int, 0)
	nodes = append(nodes, node.data)
	if node.child != nil {
		nodes = append(nodes, traverseBinomialTree(node.child)...)
	}
	if node.sibling != nil {
		nodes = append(nodes, traverseBinomialTree(node.sibling)...)
	}
	return nodes
}

func printTraverse(heap *BinomialHeap, data [][]int) {
	fmt.Printf("Traversing binominal heap after insert. Heap size = %v\n", heap.Size())
	for i, el := range data {
		fmt.Printf("№%v tree = %v\n", i, el)
	}
}

// TestBinomialHeap method tests binomial heap
func TestBinomialHeap(heap BinomialHeap) {
	heap.Insert(3)
	printTraverse(&heap, traverseBinomialHeap(&heap))
	heap.Insert(7)
	heap.Insert(8)
	heap.Insert(1)
	heap.Insert(2)
	printTraverse(&heap, traverseBinomialHeap(&heap))
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(9)
	printTraverse(&heap, traverseBinomialHeap(&heap))
	heap.Insert(11)
	heap.Insert(81)
	heap.Insert(2)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(1)
	printTraverse(&heap, traverseBinomialHeap(&heap))

	heap.Clear()
	fmt.Printf("Cleared heap. Heap size = %v\n", heap.Size())

	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn((i + 1) * 10)
		heap.Insert(r)
		printTraverse(&heap, traverseBinomialHeap(&heap))
	}
}
