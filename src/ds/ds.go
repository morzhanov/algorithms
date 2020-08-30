package ds

import (
	"github.com/morzhanov/algorithms/src/ds/binarysearchtree"
	"github.com/morzhanov/algorithms/src/ds/binarytree"
	"github.com/morzhanov/algorithms/src/ds/binomialheap"
	"github.com/morzhanov/algorithms/src/ds/doublylinkedlist"
	"github.com/morzhanov/algorithms/src/ds/dynamicarray"
	"github.com/morzhanov/algorithms/src/ds/fibonacciheap"
	"github.com/morzhanov/algorithms/src/ds/hashtable"
	"github.com/morzhanov/algorithms/src/ds/linkedlist"
	"github.com/morzhanov/algorithms/src/ds/maxheap"
	"github.com/morzhanov/algorithms/src/ds/minheap"
	"github.com/morzhanov/algorithms/src/ds/queue"
	"github.com/morzhanov/algorithms/src/ds/redblacktree"
	"github.com/morzhanov/algorithms/src/ds/set"
	"github.com/morzhanov/algorithms/src/ds/stack"
	"github.com/morzhanov/algorithms/src/ds/tree"
	"github.com/morzhanov/algorithms/src/ds/trie"
	"github.com/morzhanov/algorithms/src/utils"
)

// RunDataStructuresTests function runs data structures tests
func RunDataStructuresTests(flags utils.AlgorithmEnableFlags) {
	if flags["stack"] == true {
		// Stack
		var s stack.Stack
		stack.Test(s)
	}
	if flags["queue"] == true {
		// Queue
		var q queue.Queue
		queue.Test(q)
	}
	if flags["dynamicArray"] == true {
		// DynamicArray
		var da dynamicarray.DynamicArray
		dynamicarray.Test(da)
	}
	if flags["linkedList"] == true {
		// LinkedList
		var l linkedlist.LinkedList
		linkedlist.Test(l)
	}
	if flags["doublyLinkedList"] == true {
		// DoublyLinkedList
		var dl doublylinkedlist.DoublyLinkedList
		doublylinkedlist.Test(dl)
	}
	if flags["set"] == true {
		// Set
		var s set.Set
		set.Test(s)
	}
	if flags["hashTable"] == true {
		// HashTable
		ht := hashtable.NewHashTable()
		hashtable.Test(ht)
	}
	if flags["tree"] == true {
		// Tree
		var tn tree.Node
		tree.Test(tn)
	}
	if flags["binaryTree"] == true {
		// BinaryTree
		var tn binarytree.Node
		binarytree.Test(tn)
	}
	if flags["binarySearchTree"] == true {
		// BinarySearchTree
		var tn binarysearchtree.Node
		binarysearchtree.Test(tn)
	}
	if flags["minHeap"] == true {
		// MinHeap
		var h minheap.MinHeap
		minheap.Test(h)
	}
	if flags["maxHeap"] == true {
		// MaxHeap
		var h maxheap.MaxHeap
		maxheap.Test(h)
	}
	if flags["binomialHeap"] == true {
		// BinomialHeap
		var h binomialheap.BinomialHeap
		binomialheap.Test(h)
	}
	if flags["fibonacciHeap"] == true {
		// FibonacciHeap
		var h fibonacciheap.FibonacciHeap
		fibonacciheap.Test(h)
	}
	if flags["trie"] == true {
		// Trie
		t := trie.CreateTrie()
		trie.Test(t)
	}
	if flags["redBlackTree"] == true {
		// RedBlackTree
		t := redblacktree.CreateTree(10)
		redblacktree.Test(t)
	}
}
