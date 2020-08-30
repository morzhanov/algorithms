package main

import (
	"math/rand"

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
	"github.com/morzhanov/algorithms/src/ds/set"
	"github.com/morzhanov/algorithms/src/ds/stack"
	"github.com/morzhanov/algorithms/src/ds/tree"
	"github.com/morzhanov/algorithms/src/ds/trie"
	"github.com/morzhanov/algorithms/src/search/binarysearch"
	"github.com/morzhanov/algorithms/src/search/jumpsearch"
	"github.com/morzhanov/algorithms/src/search/linearsearch"
	"github.com/morzhanov/algorithms/src/sort/bubblesort"
	"github.com/morzhanov/algorithms/src/sort/countingsort"
	"github.com/morzhanov/algorithms/src/sort/insertionsort"
	"github.com/morzhanov/algorithms/src/sort/mergesort"
	"github.com/morzhanov/algorithms/src/sort/selectionsort"
	"github.com/morzhanov/algorithms/src/sort/stupidsort"
)

func prepareValuesForSeach(performSorting bool) ([]int, int) {
	values := make([]int, 100)
	value := 112
	for i := range values {
		values[i] = rand.Intn(500)
	}
	values[9] = value

	if performSorting {
		mergesort.MergeSort(values[:], 0, len(values)-1)
	}
	return values, value
}

func prepareValuesForSort(count int) []int {
	values := make([]int, count)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func dataStructures(flags map[string]bool) {
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
}

func sorting(flags map[string]bool) {
	if flags["stupid"] == true {
		// StupidSorting
		// extremely stupid sorting,
		// do not provide more than 5-10 elements or app could stuck
		values := prepareValuesForSort(5)
		stupidsort.Test(values[:])
	}
	if flags["bubble"] == true {
		// BubbleSorting
		values := prepareValuesForSort(100)
		bubblesort.Test(values[:])
	}
	if flags["counting"] == true {
		// CountingSorting
		values := prepareValuesForSort(100)
		countingsort.Test(values[:])
	}
	if flags["selection"] == true {
		// SelectionSorting
		values := prepareValuesForSort(100)
		selectionsort.Test(values[:])
	}
	if flags["merge"] == true {
		// MergeSorting
		values := prepareValuesForSort(100)
		mergesort.Test(values[:])
	}
	if flags["insert"] == true {
		// InsertionSorting
		values := prepareValuesForSort(100)
		insertionsort.Test(values[:])
	}
}

func searching(flags map[string]bool) {
	if flags["binary"] == true {
		// BinarySearch
		// NOTE: array should be sorted
		values, value := prepareValuesForSeach(true)
		binarysearch.Test(values[:], value)
	}
	if flags["linear"] == true {
		// LinearSearch
		values, value := prepareValuesForSeach(false)
		linearsearch.Test(values[:], value)
	}
	if flags["jump"] == true {
		// JumpSearch
		// NOTE: array should be sorted
		values, value := prepareValuesForSeach(false)
		jumpsearch.Test(values[:], value)
	}
}

func main() {
	// Use those maps to enable/disable algorithms
	dataStructuresFlags := map[string]bool{
		"stack":            false,
		"queue":            false,
		"dynamicArray":     false,
		"linkedList":       false,
		"doublyLinkedList": false,
		"set":              false,
		"hashTable":        false,
		"tree":             false,
		"binaryTree":       false,
		"binarySearchTree": false,
		"minHeap":          false,
		"maxHeap":          false,
		"binomialHeap":     false,
		"fibonacciHeap":    false,
		"trie":             true,
	}
	sortingFlags := map[string]bool{
		"stupid":    false,
		"bubble":    false,
		"counting":  false,
		"selection": false,
		"merge":     false,
		"insert":    false,
	}
	searchingFlags := map[string]bool{
		"binary": false,
		"linear": false,
		"jump":   false,
	}

	dataStructures(dataStructuresFlags)
	sorting(sortingFlags)
	searching(searchingFlags)
}
