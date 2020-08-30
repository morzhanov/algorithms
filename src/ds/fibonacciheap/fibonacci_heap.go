package fibonacciheap

import (
	"fmt"
)

// Node is a Fibonacci Heap node implementation
type Node struct {
	key, degree               int
	isMarked                  bool
	prev, next, child, parent *Node
}

// FibonacciHeap is a Fibonacci Heap implementation
type FibonacciHeap struct {
	minNode   *Node
	nodeCount int
}

// NodeListIterator creates an Iterator used to simplify the consolidate() method. It works by
// making a shallow copy of the nodes in the root list and iterating over the
// shallow copy instead of the source as the source will be modified
type NodeListIterator struct {
	index int
	items []*Node
}

// CreateNodeListIterator function creates iterator
// start - a node from the root list
func CreateNodeListIterator(start *Node) *NodeListIterator {
	var iterator NodeListIterator
	iterator.index = -1
	iterator.items = make([]*Node, 0)
	current := start
	for {
		iterator.items = append(iterator.items, current)
		current = current.next
		if start == current {
			break
		}
	}
	return &iterator
}

// HasNext function returns whether there is a next node in the iterator
func (it *NodeListIterator) HasNext() bool {
	return it.index < len(it.items)-1
}

// Next function returns next node
func (it *NodeListIterator) Next() *Node {
	it.index++
	return it.items[it.index]
}

// Clear function clears the heap's key, making it an empty heap.
func (h *FibonacciHeap) Clear() {
	h.minNode = nil
	h.nodeCount = 0
}

// DecreaseKey function decreases a key of a node
func (h *FibonacciHeap) DecreaseKey(node *Node, newKey int) {
	if node == nil {
		fmt.Printf("Cannot decrease key of non-existent node")
		return
	}
	var tempNode *Node
	tempNode.key = newKey
	if compareNodes(node, tempNode) < 0 {
		fmt.Printf("New key is larger than old key")
		return
	}

	node.key = newKey
	parent := node.parent
	if parent != nil && compareNodes(node, parent) < 0 {
		cut(node, parent, h.minNode)
		cascadingCut(parent, h.minNode)
	}
	if compareNodes(node, h.minNode) < 0 {
		h.minNode = node
	}
}

// Delete function deletes a node
func (h *FibonacciHeap) Delete(node *Node) {
	parent := node.parent
	if parent != nil {
		cut(node, parent, h.minNode)
		cascadingCut(parent, h.minNode)
	}
	h.minNode = node
	h.ExtractMin()
}

// ExtractMin function extracts and returns the minimum node from the heap
func (h *FibonacciHeap) ExtractMin() *Node {
	extractedMin := h.minNode
	if extractedMin != nil {
		if extractedMin.child != nil {
			child := extractedMin.child
			for {
				child.parent = nil
				child = child.next
				if child == extractedMin.child {
					break
				}
			}
		}

		var nextInRootList *Node
		if extractedMin.next != extractedMin {
			nextInRootList = extractedMin.next
		}
		removeNodeFromList(extractedMin)
		h.nodeCount--

		// merge the children of the minimum node with the root list
		h.minNode = mergeLists(nextInRootList, extractedMin.child)
		if h.minNode != nil {
			h.minNode = consolidate(h.minNode)
		}
	}
	return extractedMin
}

// FindMin function returns the minimum node from the heap
// node - the heap's minimum node or undefined if the heap is empty
func (h *FibonacciHeap) FindMin() *Node {
	return h.minNode
}

// Insert function inserts a new key into the heap
func (h *FibonacciHeap) Insert(key int) *Node {
	var node Node
	node.key = key
	node.prev = &node
	node.next = &node
	node.degree = 0
	h.minNode = mergeLists(h.minNode, &node)
	h.nodeCount++
	return &node
}

// IsEmpty function returns is heap empty
func (h *FibonacciHeap) IsEmpty() bool {
	return h.minNode == nil
}

// Size function returns heap size
func (h *FibonacciHeap) Size() int {
	if h.IsEmpty() {
		return 0
	}
	return h.GetNodeListSize(h.minNode)
}

// GetNodeListSize function returns the size of a node list
func (h *FibonacciHeap) GetNodeListSize(node *Node) int {
	count := 0
	current := node
	for {
		count++
		if current.child != nil {
			count += h.GetNodeListSize(current.child)
		}
		current = current.next
		if current == node {
			break
		}
	}
	return count
}

// Union function joins another heap to this heap
func (h *FibonacciHeap) Union(other *FibonacciHeap) {
	h.minNode = mergeLists(h.minNode, other.minNode)
	h.nodeCount += other.nodeCount
}

// compareNodes function compares two nodes with each other
// returns -1, 0 or 1 if a < b, a == b or a > b respectively
func compareNodes(a *Node, b *Node) int {
	if a.key > b.key {
		return 1
	}
	if a.key < b.key {
		return -1
	}
	return 0
}

// cut function cuts the link between a node and its parent, moving the node to the root list.
// node - the node being cut.
// parent - the parent of the node being cut.
// minNode - the minimum node in the root list.
// compare - the node comparison function to use.
// returns the heap's new minimum node.
func cut(node *Node, parent *Node, minNode *Node) *Node {
	node.parent = nil
	parent.degree--
	if node.next == node {
		parent.child = nil
	} else {
		parent.child = node.next
	}
	removeNodeFromList(node)
	minNode = mergeLists(minNode, node)
	node.isMarked = false
	return minNode
}

// cascadingCut function performs a cascading cut on a node; mark the node if it is not marked,
// otherwise cut the node and perform a cascading cut on its parent.
// node - the node being considered to be cut.
// minNode - the minimum node in the root list.
// compare - the node comparison function to use.
// returns the heap's new minimum node.
func cascadingCut(node *Node, minNode *Node) *Node {
	parent := node.parent
	if parent != nil {
		if node.isMarked {
			minNode = cut(node, parent, minNode)
			minNode = cascadingCut(parent, minNode)
		} else {
			node.isMarked = true
		}
	}
	return minNode
}

// consolidate function merges all trees of the same order together until there are no two trees of the same order
// minNode - the current minimum node.
// returns the new minimum node
func consolidate(minNode *Node) *Node {
	aux := map[int]*Node{}
	it := CreateNodeListIterator(minNode)

	for it.HasNext() {
		curr := it.Next()
		// if there exists another node with the same degree, merge them
		for aux[curr.degree] != nil {
			if compareNodes(curr, aux[curr.degree]) > 0 {
				temp := curr
				curr = aux[curr.degree]
				aux[curr.degree] = temp
			}
			link(aux[curr.degree], curr)
			aux[curr.degree] = nil
			curr.degree++
		}
		aux[curr.degree] = curr
	}

	minNode = nil
	for i := 0; i < len(aux); i++ {
		if aux[i] != nil {
			// remove siblings before merging
			aux[i].next = aux[i]
			aux[i].prev = aux[i]
			minNode = mergeLists(minNode, aux[i])
		}
	}
	return minNode
}

// removeNodeFromList function removes a node from a node list
func removeNodeFromList(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.next = node
	node.prev = node
}

// link function links two heaps of the same order together
// max - the heap with the larger root.
// min - the heap with the smaller root.
func link(max *Node, min *Node) {
	removeNodeFromList(max)
	min.child = mergeLists(max, min.child)
	max.parent = min
	max.isMarked = false
}

// mergeLists function merges two lists of nodes together
// returns the new minimum node from the two lists
func mergeLists(a *Node, b *Node) *Node {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	temp := a.next
	a.next = b.next
	a.next.prev = a
	b.next = temp
	b.next.prev = b
	if compareNodes(a, b) < 0 {
		return a
	}
	return b
}

func traverse(h *FibonacciHeap) [][]int {
	res := make([][]int, 0)
	curr := h.minNode
	for {
		res = append(res, traverseTree(curr))
		if curr != nil {
			curr = curr.next
			if curr == h.minNode {
				break
			}
		}
	}
	return res
}

func traverseTree(node *Node) []int {
	if node == nil {
		return []int{}
	}
	nodes := make([]int, 0)
	nodes = append(nodes, node.key)
	if node.child != nil {
		nodes = append(nodes, traverseTree(node.child)...)
	}
	return nodes
}

func printTraverse(heap *FibonacciHeap, key [][]int) {
	fmt.Printf("Traversing fibonacci heap after insert. Heap size = %v\n", heap.Size())
	for i, el := range key {
		fmt.Printf("№%v tree = %v\n", i, el)
	}
}

// Test method tests fibonacci heap
func Test(heap FibonacciHeap) {
	heap.Insert(3)
	heap.Insert(3)
	heap.Insert(3)
	heap.Insert(3)
	heap.ExtractMin()
	printTraverse(&heap, traverse(&heap))
	heap.Insert(7)
	heap.Insert(8)
	heap.Insert(1)
	heap.Insert(2)
	heap.ExtractMin()
	printTraverse(&heap, traverse(&heap))
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(9)
	heap.ExtractMin()
	printTraverse(&heap, traverse(&heap))
	heap.Insert(11)
	heap.Insert(81)
	heap.Insert(2)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(1)
	heap.ExtractMin()
	printTraverse(&heap, traverse(&heap))
}
