package main

import (
	"github.com/morzhanov/algorithms/src/ds"
	"github.com/morzhanov/algorithms/src/search"
	"github.com/morzhanov/algorithms/src/sort"
)

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
		"trie":             false,
		"redBlackTree":     false,
		"avlTree":          false,
		"optimalBSTCost":   false,
		"optimalBST":       false,
		"btree":            false,
		"quadTree":         false,
		"graph":            false,
		"directedGraph":    false,
		"acyclicGraph":     false,
		"weightedGraph":    true,
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

	ds.RunDataStructuresTests(dataStructuresFlags)
	sort.RunSortTests(sortingFlags)
	search.RunSearchTests(searchingFlags)
}
