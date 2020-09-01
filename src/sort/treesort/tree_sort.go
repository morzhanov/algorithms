package treesort

import (
	"fmt"

	"github.com/morzhanov/algorithms/src/ds/binarysearchtree"
)

// TreeSort function implements tree sort algorithm using binary search tree
func TreeSort(arr []int) []int {
	bst := binarysearchtree.Node{}
	bst.SetValue(arr[0])
	for _, el := range arr {
		bst.Insert(el)
	}
	return bst.InorderTraversal()
}

// Test function tests TreeSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Tree sorting started\n")
	res := TreeSort(values[:])
	fmt.Printf("Tree sorting finished. Elements count = %v\n", len(res))
	fmt.Printf("Tree sorting result = %v\n", res)
}
