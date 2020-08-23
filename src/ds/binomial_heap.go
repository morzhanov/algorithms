package ds

import (
	"errors"
	"fmt"
	"math"
)

// https://www.geeksforgeeks.org/binomial-heap-2/
// https://www.geeksforgeeks.org/memory-representation-of-binomial-heap/?ref=rp
// https://www.geeksforgeeks.org/implementation-binomial-heap/
// https://www.geeksforgeeks.org/implementation-binomial-heap-set-2/?ref=rp

// MinHeap is a Min Binary Heap implementation
type MinHeap []int

func (h *MinHeap) getLeftChildIndex(parentIdx int) int {
	return 2*parentIdx + 1
}
func (h *MinHeap) getRightChildIndex(parentIdx int) int {
	return 2*parentIdx + 2
}
func (h *MinHeap) getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *MinHeap) getLeftChild(idx int) int {
	return (*h)[h.getLeftChildIndex(idx)]
}
func (h *MinHeap) getRightChild(idx int) int {
	return (*h)[h.getRightChildIndex(idx)]
}
func (h *MinHeap) getParent(idx int) int {
	return (*h)[h.getParentIndex(idx)]
}

func (h *MinHeap) hasLeftChild(idx int) bool {
	return h.getLeftChildIndex(idx) < len(*h)
}
func (h *MinHeap) hasRightChild(idx int) bool {
	return h.getRightChildIndex(idx) < len(*h)
}
func (h *MinHeap) hasParent(idx int) bool {
	return h.getParentIndex(idx) >= 0
}

func (h *MinHeap) swap(firstIndex int, secondIndex int) {
	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}

// Peek method returns first heap element
func (h *MinHeap) Peek() (int, error) {
	if len(*h) == 0 {
		return 0, errors.New("Heap is empty")
	}
	return (*h)[0], nil
}

func (h *MinHeap) heapifyDown() {
	idx := 0
	for h.hasLeftChild(idx) {
		minIdx := h.getLeftChildIndex(idx)
		if h.hasRightChild(idx) && h.getRightChild(idx) < h.getLeftChild(idx) {
			minIdx = h.getRightChildIndex(idx)
		}

		if (*h)[idx] < (*h)[minIdx] {
			break
		} else {
			h.swap(idx, minIdx)
		}
		idx = minIdx
	}
}

func (h *MinHeap) heapifyUp() {
	idx := len(*h) - 1
	for h.hasParent(idx) && h.getParent(idx) > (*h)[idx] {
		h.swap(h.getParentIndex(idx), idx)
		idx = h.getParentIndex(idx)
	}
}

// Pool method returns first heap element
func (h *MinHeap) Pool() (int, error) {
	if len(*h) == 0 {
		return 0, errors.New("Heap is empty")
	}

	item := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapifyDown()

	return item, nil
}

// Add method adds element to Heap
func (h *MinHeap) Add(item int) {
	*h = append(*h, item)
	h.heapifyUp()
}

// GetHeight method returns min heap height
func (h *MinHeap) GetHeight() float64 {
	return math.Ceil(math.Log2(float64(len(*h))))
}

// TestMinHeap method tests heap
func TestMinHeap(heap MinHeap) {
	elements := []int{25, 20, 6, 8, 7, 2, 1, 5, 2, 4, 67, 8, 20}
	for _, el := range elements {
		heap.Add(el)
		fmt.Printf("Added element %v to min heap. Result MinHeap %v\n", el, heap)
	}

	peek, _ := heap.Peek()
	fmt.Printf("MinHeap peek element: %v\n", peek)

	pool, _ := heap.Pool()
	fmt.Printf("MinHeap Pool called. Peek element: %v; MinHeap: %v\n", pool, heap)

	fmt.Printf("Heap hight: %v\n", heap.GetHeight())
}
