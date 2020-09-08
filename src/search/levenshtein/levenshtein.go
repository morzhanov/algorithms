package levenshtein

import (
	"fmt"

	"github.com/morzhanov/algorithms/src/utils"
)

// Levenshtein function implements Levenshtein distance algorithm
func Levenshtein(A string, B string) {
	X := []rune(A)
	Y := []rune(B)
	m := len(X)
	n := len(Y)
	L := make([][]int, m+1)
	for i := 0; i < len(L); i++ {
		L[i] = make([]int, n+1)
	}

	// fill L[][] in bottom up manner
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// if first string is empty, only option is to
			// insert all characters of second string
			if i == 0 {
				// min operations = j
				L[i][j] = j
			} else if j == 0 {
				// if second string is empty, only option is to
				// remove all characters of second string
				// min operations = i
				L[i][j] = i
			} else if X[i-1] == Y[j-1] {
				// if last characters are same, ignore last char
				// and recur for remaining string
				L[i][j] = L[i-1][j-1]
			} else {
				// if the last character is different, consider all possibilities and find the minimum
				L[i][j] = 1 + utils.IntMinArr(
					[]int{L[i][j-1], // insert
						L[i-1][j],    // remove
						L[i-1][j-1]}, // replace
				)
			}
		}
	}

	fmt.Printf("Levenshtein distance of %v and %v is %v\n", A, B, L[m][n])
}

// Test function tests Levenshtein distance
func Test(A string, B string) {
	fmt.Printf("Levenshtein distance started\n")
	Levenshtein(A, B)
}
