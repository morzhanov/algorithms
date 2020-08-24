package main

import (
	"local/algorithms/src/ds"
	"local/algorithms/src/search"
	"local/algorithms/src/sort"
	"math/rand"
)

func prepareValuesForSeach(performSorting bool) ([]int, int) {
	values := make([]int, 100)
	value := 112
	for i := range values {
		values[i] = rand.Intn(500)
	}
	values[9] = value

	if performSorting {
		sort.MergeSort(values[:], 0, len(values)-1)
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
		var stack ds.Stack
		ds.TestStack(stack)
	}
	if flags["queue"] == true {
		// Queue
		var queue ds.Queue
		ds.TestQueue(queue)
	}
	if flags["dynamicArray"] == true {
		// DynamicArray
		var da ds.DynamicArray
		ds.TestDynamicArray(da)
	}
	if flags["linkedList"] == true {
		// LinkedList
		var l ds.LinkedList
		ds.TestLinkedList(l)
	}
	if flags["doublyLinkedList"] == true {
		// DoublyLinkedList
		var dl ds.DoublyLinkedList
		ds.TestDoublyLinkedList(dl)
	}
	if flags["set"] == true {
		// Set
		var s ds.Set
		ds.TestSet(s)
	}
	if flags["hashTable"] == true {
		// HashTable
		ht := ds.NewHashTable()
		ds.TestHashTable(ht)
	}
	if flags["tree"] == true {
		// Tree
		var tn ds.TreeNode
		ds.TestTree(tn)
	}
	if flags["binaryTree"] == true {
		// BinaryTree
		var tn ds.BinaryTreeNode
		ds.TestBinaryTree(tn)
	}
	if flags["binarySearchTree"] == true {
		// BinarySearchTree
		var tn ds.BinarySearchTreeNode
		ds.TestBinarySearchTree(tn)
	}
	if flags["minHeap"] == true {
		// MinHeap
		var h ds.MinHeap
		ds.TestMinHeap(h)
	}
	if flags["maxHeap"] == true {
		// MaxHeap
		var h ds.MaxHeap
		ds.TestMaxHeap(h)
	}
	if flags["binomialHeap"] == true {
		// BinomialHeap
		var h ds.BinomialHeap
		ds.TestBinomialHeap(h)
	}
	if flags["fibonacciHeap"] == true {
		// FibonacciHeap
		var h ds.FibonacciHeap
		ds.TestFibonacciHeap(h)
	}
	if flags["trie"] == true {
		// Trie
		t := ds.CreateTrie()
		ds.TestTrie(t)
	}
}

func sorting(flags map[string]bool) {
	if flags["stupid"] == true {
		// StupidSorting
		// extremely stupid sorting,
		// do not provide more than 5-10 elements or app could stuck
		values := prepareValuesForSort(5)
		sort.TestStupidSort(values[:])
	}
	if flags["bubble"] == true {
		// BubbleSorting
		values := prepareValuesForSort(100)
		sort.TestBubbleSort(values[:])
	}
	if flags["counting"] == true {
		// CountingSorting
		values := prepareValuesForSort(100)
		sort.TestCountingSort(values[:])
	}
	if flags["selection"] == true {
		// SelectionSorting
		values := prepareValuesForSort(100)
		sort.TestSelectionSort(values[:])
	}
	if flags["merge"] == true {
		// MergeSorting
		values := prepareValuesForSort(100)
		sort.TestMergeSort(values[:])
	}
	if flags["insert"] == true {
		// InsertionSorting
		values := prepareValuesForSort(100)
		sort.TestInsertionSort(values[:])
	}
}

func searching(flags map[string]bool) {
	if flags["binary"] == true {
		// BinarySearch
		// NOTE: array should be sorted
		values, value := prepareValuesForSeach(true)
		search.TestBinarySearch(values[:], value)
	}
	if flags["linear"] == true {
		// LinearSearch
		values, value := prepareValuesForSeach(false)
		search.TestLinearSearch(values[:], value)
	}
	if flags["jump"] == true {
		// JumpSearch
		// NOTE: array should be sorted
		values, value := prepareValuesForSeach(false)
		search.TestJumpSearch(values[:], value)
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
