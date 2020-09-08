package search

import (
	"github.com/morzhanov/algorithms/src/graph/dijkstra"
	"github.com/morzhanov/algorithms/src/search/binarysearch"
	"github.com/morzhanov/algorithms/src/search/exponentialsearch"
	"github.com/morzhanov/algorithms/src/search/interpolationserach"
	"github.com/morzhanov/algorithms/src/search/jumpsearch"
	"github.com/morzhanov/algorithms/src/search/lcs"
	"github.com/morzhanov/algorithms/src/search/levenshtein"
	"github.com/morzhanov/algorithms/src/search/linearsearch"
	"github.com/morzhanov/algorithms/src/search/substring"
	"github.com/morzhanov/algorithms/src/utils"
)

// RunSearchTests function runs search tests
func RunSearchTests(flags utils.AlgorithmEnableFlags) {
	if flags["binary"] == true {
		// BinarySearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(true)
		binarysearch.Test(values[:], value)
	}
	if flags["linear"] == true {
		// LinearSearch
		values, value := utils.PrepareValuesForSeach(false)
		linearsearch.Test(values[:], value)
	}
	if flags["jump"] == true {
		// JumpSearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(true)
		jumpsearch.Test(values[:], value)
	}
	if flags["interpolation"] == true {
		// InterpolationSearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(true)
		interpolationserach.Test(values[:], value)
	}
	if flags["exponential"] == true {
		// ExponentialSearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(true)
		exponentialsearch.Test(values[:], value)
	}
	if flags["lcs"] == true {
		// LCS
		lcs.Test("AGGTAB", "GXTXAYB")
	}
	if flags["levenshtein"] == true {
		// Levenshtein distance
		levenshtein.Test("AGGTAB", "GXTXAYB")
		levenshtein.Test("food", "money")
		levenshtein.Test("sunday", "saturday")
	}
	if flags["substring"] == true {
		// Substring
		substring.Test("AGG", "AGGTAB")
		substring.Test("money", "foodmoney")
		substring.Test("sunday", "saturday")
		substring.Test("day", "saturday")
	}
	if flags["dijkstra"] == true {
		// Dijkstra
		dijkstra.Test()
	}
}
