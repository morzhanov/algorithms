package main

import "local/algorithms/src/sort"

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
	values := [8]int{102, 4, 2, 212, 3, 4, 5, 22}
	sort.TestStupidSort(values[:])
}
