package optimalbst

import (
	"fmt"
	"math"
)

// CalculateBST is a main function that calculates minimum cost of a Binary Search Tree
// returns cost and BST as an array
func CalculateBST(keys []int, freq []int) (int, []int) {
	n := len(keys)
	bst := make([]int, n)
	// create an auxiliary 2D matrix to store results of subproblems
	cost := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		cost[i] = make([]int, n+1)
	}

	// cost[i][j] = Optimal cost of binary search tree that
	// can be formed from keys[i] to keys[j]. cost[0][n-1]
	// will store the resultant cost

	// for a single key, cost is equal to frequency of the key
	for i := 0; i < n; i++ {
		cost[i][i] = freq[i]
	}

	// now we need to consider chains of length 2, 3, ... . l is chain length.
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l+1; i++ {
			// get column number j from row number i and chain length
			j := i + l - 1
			cost[i][j] = int(math.MaxInt32 >> 1)

			// try making all keys in interval keys[i..j] as root
			for r := i; r <= j; r++ {
				// c = cost when keys[r] becomes root of this subtree
				c := 0
				if r > i {
					c = cost[i][r-1]
				}
				if r < j {
					bst = append(bst, r)
					c += cost[r+1][j]
				}
				c += sumWithCheck(freq, i, j)
				if c < cost[i][j] {
					fmt.Printf("i: %v, j: %v\n", i, j)
					fmt.Printf("r: %v, val: %v\n", r, keys[r])
					cost[i][j] = c
				}
			}
		}
	}
	return cost[0][n-1], bst
}

// 2,0,1
// 1,0,2

// utility function to get sum of array elements freq[i] to freq[j]
func sumWithCheck(freq []int, i int, j int) int {
	s := 0
	for k := i; k <= j; k++ {
		if k >= len(freq) {
			continue
		}
		s += freq[k]
	}
	return s
}

// Test function tests OptimalBST
func Test(keys []int, freq []int) {
	cost, bst := CalculateBST(keys, freq)
	fmt.Printf("Cost of Optimal BST is %v\n", cost)
	fmt.Printf("Optimal BST is %v\n", bst)
}
