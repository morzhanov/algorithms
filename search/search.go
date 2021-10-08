package search

import (
	"github.com/morzhanov/algorithms/graph/dijkstra"
	"github.com/morzhanov/algorithms/search/binarysearch"
	"github.com/morzhanov/algorithms/search/exponentialsearch"
	"github.com/morzhanov/algorithms/search/interpolationserach"
	"github.com/morzhanov/algorithms/search/jumpsearch"
	"github.com/morzhanov/algorithms/search/lcs"
	"github.com/morzhanov/algorithms/search/levenshtein"
	"github.com/morzhanov/algorithms/search/linearsearch"
	"github.com/morzhanov/algorithms/search/substring"
	"github.com/morzhanov/algorithms/utils"
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
