package lcs

import (
	"fmt"

	"github.com/morzhanov/algorithms/src/utils"
)

// LCS function implements Longest common subseauence algorithm
func LCS(A string, B string) {
	X := []rune(A)
	Y := []rune(B)
	m := len(X)
	n := len(Y)
	L := make([][]rune, m+1)
	for i := 0; i < len(L); i++ {
		L[i] = make([]rune, n+1)
	}

	// following steps build L[m+1][n+1] in bottom up fashion.
	// note that L[i][j] contains length of LCS of X[0..i-1] and Y[0..j-1]
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0
			} else if X[i-1] == Y[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = utils.RuneMax(L[i-1][j], L[i][j-1])
			}
		}
	}

	// following code is used to print LCS
	index := L[m][n]
	// create a character array to store the lcs string
	lcs := make([]rune, index+1)
	// set the terminating character
	lcs[index] = '\u0000'
	// start from the right-most-bottom-most corner and one by one store characters in lcs[]
	i := m
	j := n

	for i > 0 && j > 0 {
		// if current character in X[] and Y are same, then current character is part of LCS
		if X[i-1] == Y[j-1] {
			// put current character in result
			lcs[index-1] = X[i-1]
			// reduce values of i, j and index
			i--
			j--
			index--
		} else if L[i-1][j] > L[i][j-1] {
			// if not same, then find the larger of two and go in the direction of larger value
			i--
		} else {
			j--
		}
	}

	// print the lcs
	fmt.Printf("LCS of " + A + " and " + B + " is ")
	for k := 0; k <= len(lcs)-1; k++ {
		fmt.Printf("%v", string(lcs[k]))
	}
	fmt.Printf("\n")
}

// Test function tests LCS
func Test(A string, B string) {
	fmt.Printf("LCS started\n")
	LCS(A, B)
}
