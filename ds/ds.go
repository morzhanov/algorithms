package ds

import (
	"fmt"

	"github.com/morzhanov/algorithms/ds/acyclicgraph"
	"github.com/morzhanov/algorithms/ds/avltree"
	"github.com/morzhanov/algorithms/ds/binarysearchtree"
	"github.com/morzhanov/algorithms/ds/binarytree"
	"github.com/morzhanov/algorithms/ds/binomialheap"
	"github.com/morzhanov/algorithms/ds/btree"
	"github.com/morzhanov/algorithms/ds/dag"
	"github.com/morzhanov/algorithms/ds/directedgraph"
	"github.com/morzhanov/algorithms/ds/doublylinkedlist"
	"github.com/morzhanov/algorithms/ds/dynamicarray"
	"github.com/morzhanov/algorithms/ds/fibonacciheap"
	"github.com/morzhanov/algorithms/ds/hashtable"
	"github.com/morzhanov/algorithms/ds/linkedlist"
	"github.com/morzhanov/algorithms/ds/maxheap"
	"github.com/morzhanov/algorithms/ds/minheap"
	"github.com/morzhanov/algorithms/ds/optimalbst"
	"github.com/morzhanov/algorithms/ds/quadtree"
	"github.com/morzhanov/algorithms/ds/queue"
	"github.com/morzhanov/algorithms/ds/redblacktree"
	"github.com/morzhanov/algorithms/ds/set"
	"github.com/morzhanov/algorithms/ds/simplegraph"
	"github.com/morzhanov/algorithms/ds/stack"
	"github.com/morzhanov/algorithms/ds/tree"
	"github.com/morzhanov/algorithms/ds/trie"
	"github.com/morzhanov/algorithms/ds/weightedgraph"
	"github.com/morzhanov/algorithms/utils"
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
	if flags["avlTree"] == true {
		// AVLTree
		var tn avltree.Node
		avltree.Test(tn)
	}
	if flags["optimalBSTCost"] == true {
		// OptimalBSTCost
		keys := []int{10, 12, 20}
		freq := []int{34, 8, 50}
		optimalbst.TestCost(keys, freq)
	}
	if flags["optimalBSTCost"] == true {
		// OptimalBSTCost
		keys := []int{10, 12, 20}
		freq := []int{34, 8, 50}
		optimalbst.TestCost(keys, freq)
	}
	if flags["optimalBST"] == true {
		// OptimalBST
		keys := []int{10, 12, 20}
		freq := []int{34, 8, 50}
		optimalbst.Test(keys, freq)
	}
	if flags["btree"] == true {
		// BTree
		btree := btree.CreateBTree(4, 9)
		arr := []int{4, 2, 7, 11, 1, 5, 3, 8}
		for _, key := range arr {
			fmt.Printf("Adding %v to BTree\n", key)
			btree.Add(key)
			fmt.Printf("%v\n", btree.Print())
		}
		for _, key := range arr {
			fmt.Printf("Removing %v from BTree\n", key)
			btree.Remove(key)
			fmt.Printf("%v\n", btree.Print())
		}
	}
	if flags["quadTree"] == true {
		// QuadTree
		qt := quadtree.Create(quadtree.CreateBounds(10, 20, 30, 40), 10, 4)
		quadtree.Test(qt)
	}
	if flags["graph"] == true {
		// Graph
		g := simplegraph.CreateGraph()
		simplegraph.Test(g)
	}
	if flags["directedGraph"] == true {
		// Directed Graph
		g := directedgraph.CreateGraph()
		directedgraph.Test(g)
	}
	if flags["acyclicGraph"] == true {
		// Acyclic Graph
		g := acyclicgraph.CreateGraph()
		acyclicgraph.Test(g)
	}
	if flags["weightedGraph"] == true {
		// Weighted Graph
		g := weightedgraph.CreateGraph()
		weightedgraph.Test(g)
	}
	if flags["dag"] == true {
		// DAG
		g := dag.CreateGraph()
		dag.Test(g)
	}
}
