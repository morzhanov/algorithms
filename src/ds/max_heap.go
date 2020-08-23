package ds

import (
	"errors"
	"fmt"
	"math"
)

// MaxHeap is a MaxHeap implementation
type MaxHeap []int

func (h *MaxHeap) getLeftChildIndex(parentIdx int) int {
	return 2*parentIdx + 1
}
func (h *MaxHeap) getRightChildIndex(parentIdx int) int {
	return 2*parentIdx + 2
}
func (h *MaxHeap) getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *MaxHeap) getLeftChild(idx int) int {
	return (*h)[h.getLeftChildIndex(idx)]
}
func (h *MaxHeap) getRightChild(idx int) int {
	return (*h)[h.getRightChildIndex(idx)]
}
func (h *MaxHeap) getParent(idx int) int {
	return (*h)[h.getParentIndex(idx)]
}

func (h *MaxHeap) hasLeftChild(idx int) bool {
	return h.getLeftChildIndex(idx) < len(*h)
}
func (h *MaxHeap) hasRightChild(idx int) bool {
	return h.getRightChildIndex(idx) < len(*h)
}
func (h *MaxHeap) hasParent(idx int) bool {
	return h.getParentIndex(idx) >= 0
}

func (h *MaxHeap) swap(firstIndex int, secondIndex int) {
	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}

// Peek method returns first heap element
func (h *MaxHeap) Peek() (int, error) {
	if len(*h) == 0 {
		return 0, errors.New("Heap is empty")
	}
	return (*h)[0], nil
}

func (h *MaxHeap) heapifyDown() {
	idx := 0
	for h.hasLeftChild(idx) {
		maxIdx := h.getLeftChildIndex(idx)
		if h.hasRightChild(idx) && h.getRightChild(idx) > h.getLeftChild(idx) {
			maxIdx = h.getRightChildIndex(idx)
		}

		if (*h)[idx] > (*h)[maxIdx] {
			break
		} else {
			h.swap(idx, maxIdx)
		}
		idx = maxIdx
	}
}

func (h *MaxHeap) heapifyUp() {
	idx := len(*h) - 1
	for h.hasParent(idx) && h.getParent(idx) < (*h)[idx] {
		h.swap(h.getParentIndex(idx), idx)
		idx = h.getParentIndex(idx)
	}
}

// Pool method returns first heap element
func (h *MaxHeap) Pool() (int, error) {
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
func (h *MaxHeap) Add(item int) {
	*h = append(*h, item)
	h.heapifyUp()
}

// GetHeight method returns max heap height
func (h *MaxHeap) GetHeight() float64 {
	return math.Ceil(math.Log2(float64(len(*h))))
}

// TestMaxHeap method tests heap
func TestMaxHeap(heap MaxHeap) {
	elements := []int{25, 20, 6, 8, 7, 2, 1, 5, 2, 4, 67, 8, 20}
	for _, el := range elements {
		heap.Add(el)
		fmt.Printf("Added element %v to max heap. Result MaxHeap %v\n", el, heap)
	}

	peek, _ := heap.Peek()
	fmt.Printf("MaxHeap peek element: %v\n", peek)

	pool, _ := heap.Pool()
	fmt.Printf("MaxHeap Pool called. Peek element: %v; MaxHeap: %v\n", pool, heap)

	fmt.Printf("Heap hight: %v\n", heap.GetHeight())
}
