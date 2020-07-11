package main

import (
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

func prepareValuesForSort() []int {
	values := make([]int, 1000)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func main() {

	/*
	* Data Structures
	 */

	// Stack
	// var stack ds.Stack
	// ds.TestStack(stack)

	// Queue
	// var queue ds.Queue
	// ds.TestQueue(queue)

	// DynamicArray
	// var da ds.DynamicArray
	// ds.TestDynamicArray(da)

	// LinkedList
	// var l ds.LinkedList
	// ds.TestLinkedList(l)

	// DoublyLinkedList
	// var dl ds.DoublyLinkedList
	// ds.TestDoublyLinkedList(dl)

	// Set
	// var s ds.Set
	// ds.TestSet(s)

	// Set
	// ht := ds.NewHashTable()
	// ds.TestHashTable(ht)

	// Tree
	// var tn ds.TreeNode
	// ds.TestTree(tn)

	/*
	* Sorting
	 */

	// StupidSorting
	// extremely stupid sorting, do not provide more than 5-10 elements or app could stuck
	// values := prepareValuesForSort()
	// sort.TestStupidSort(values[:])

	// BubbleSorting
	// values := prepareValuesForSort()
	// sort.TestBubbleSort(values[:])

	// CountingSorting
	// values := prepareValuesForSort()
	// sort.TestCountingSort(values[:])

	// SelectionSorting
	// values := prepareValuesForSort()
	// sort.TestSelectionSort(values[:])

	// MergeSorting
	// values := prepareValuesForSort()
	// sort.TestMergeSort(values[:])

	// InsertionSorting
	// values := prepareValuesForSort()
	// sort.TestInsertionSort(values[:])

	/*
	* Search
	 */

	// BinarySearch
	// NOTE: array should be sorted
	// values, value := prepareValuesForSeach(true)
	// search.TestBinarySearch(values[:], value)

	// BinarySearch
	values, value := prepareValuesForSeach(false)
	search.TestLinearSearch(values[:], value)
}
