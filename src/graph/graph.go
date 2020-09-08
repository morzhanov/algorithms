package graph

import (
	"github.com/morzhanov/algorithms/src/ds/directedgraph"
	"github.com/morzhanov/algorithms/src/graph/cycleindirectedgraph"
	"github.com/morzhanov/algorithms/src/graph/dijkstra"
	"github.com/morzhanov/algorithms/src/utils"
)

// RunGraphTests function runs graph tests
func RunGraphTests(flags utils.AlgorithmEnableFlags) {
	if flags["dijkstra"] == true {
		// Dijkstra
		dijkstra.Test()
	}
	if flags["cycleindirectedgraph"] == true {
		// CycleInDirectedGraph
		g := directedgraph.CreateGraph()
		cycleindirectedgraph.Test(g)
	}
}
