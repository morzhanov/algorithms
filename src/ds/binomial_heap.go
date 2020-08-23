package ds

// https://www.geeksforgeeks.org/binomial-heap-2/
// https://www.geeksforgeeks.org/memory-representation-of-binomial-heap/?ref=rp
// https://www.geeksforgeeks.org/implementation-binomial-heap/
// https://www.geeksforgeeks.org/implementation-binomial-heap-set-2/?ref=rp

// BinomialHeapNode is a Binomial Heap implementation
type BinomialHeapNode struct {
	data, degree           int
	child, parent, sibling *BinomialHeapNode
}

func (h *BinomialHeapNode) getLeftChildIndex(parentIdx int) int {
	return 2*parentIdx + 1
}

Node* newNode(int key) 
{ 
    Node *temp = new Node; 
    temp->data = key; 
    temp->degree = 0; 
    temp->child = temp->parent = temp->sibling = NULL; 
    return temp; 
} 

func (n* BinomialHeapNode) CreateNode (key int) *BinomialHeapNode {
	
}

// TestBinomialHeap method tests binomial heap
func TestBinomialHeap(heap BinomialHeap) {
	// elements := []int{25, 20, 6, 8, 7, 2, 1, 5, 2, 4, 67, 8, 20}
	// for _, el := range elements {
	// 	heap.Add(el)
	// 	fmt.Printf("Added element %v to min heap. Result MinHeap %v\n", el, heap)
	// }

	// peek, _ := heap.Peek()
	// fmt.Printf("MinHeap peek element: %v\n", peek)

	// pool, _ := heap.Pool()
	// fmt.Printf("MinHeap Pool called. Peek element: %v; MinHeap: %v\n", pool, heap)

	// fmt.Printf("Heap hight: %v\n", heap.GetHeight())
}
