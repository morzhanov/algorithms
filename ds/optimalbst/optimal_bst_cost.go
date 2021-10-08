package optimalbst

import (
	"fmt"
	"math"
)

// Cost is a main function that calculates minimum cost of a Binary Search Tree
// This function only calculates cost, use OptimalBST to also test and return Tree
func Cost(keys []int, freq []int) int {
	return optCost(freq, 0, len(keys)-1)
}

// recursive function to calculate cost of optimal binary search tree
func optCost(freq []int, i int, j int) int {
	// no elements in subarray
	if j < i {
		return 0
	}
	// one element
	if j == i {
		return freq[i]
	}

	// get sum of freq[i], ... freq[j]
	fsum := sum(freq, i, j)
	min := int(math.MaxInt32 >> 1)

	// one by one consider all elements as root and
	// recursively find cost of the BST, compare the
	// cost with min and update min if needed
	for r := i; r <= j; r++ {
		cost := optCost(freq, i, r-1) + optCost(freq, r+1, j)
		if cost < min {
			min = cost
		}
	}
	return min + fsum
}

// utility function to get sum of array elements freq[i] to freq[j]
func sum(freq []int, i int, j int) int {
	s := 0
	for k := i; k <= j; k++ {
		s += freq[k]
	}
	return s
}

// TestCost function tests OptimalBST cost
func TestCost(keys []int, freq []int) {
	cost := Cost(keys, freq)
	fmt.Printf("Cost of Optimal BST is %v\n", cost)
}
