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
		"weightedGraph":    false,
	}
	sortingFlags := map[string]bool{
		"stupid":       false,
		"bubble":       false,
		"counting":     false,
		"selection":    false,
		"merge":        false,
		"insert":       false,
		"pairedInsert": false,
		"shell":        false,
		"coctail":      false,
		"tree":         false,
		"oddEven":      false,
		"comb":         false,
		"radix":        false,
		"pigeonhole":   false,
		"stooge":       false,
		"gnome":        false,
		"bucket":       false,
		"tim":          false,
		"cycle":        false,
		"quick":        false,
	}
	searchingFlags := map[string]bool{
		"binary":        false,
		"linear":        false,
		"jump":          false,
		"interpolation": false,
		"exponential":   false,
		"lcs":           true,
	}

	ds.RunDataStructuresTests(dataStructuresFlags)
	sort.RunSortTests(sortingFlags)
	search.RunSearchTests(searchingFlags)
}
