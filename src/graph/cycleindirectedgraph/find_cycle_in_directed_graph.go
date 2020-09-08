package cycleindirectedgraph

import (
	"fmt"

	"github.com/morzhanov/algorithms/src/ds/directedgraph"
)

// FindCycleInDirectedGraph finds cycle in directed graph
// returns true if the graph contains a cycle, else false
func FindCycleInDirectedGraph(g *directedgraph.Graph) bool {
	// mark all the vertices as not visited and
	// not part of recursion stack
	visited := make([]bool, g.GetVertexCount())
	recStack := make([]bool, g.GetVertexCount())
	// call the recursive helper function to detect cycle in different DFS trees
	for i := 0; i < g.GetVertexCount(); i++ {
		if isCyclicUtil(g, i, visited, recStack) {
			return true
		}
	}
	return false
}

// mark the current node as visited and
// part of recursion stack
func isCyclicUtil(g *directedgraph.Graph, i int, visited []bool, recStack []bool) bool {
	// mark the current node as visited and
	// part of recursion stack
	if recStack[i] {
		return true
	}
	if visited[i] {
		return false
	}

	visited[i] = true
	recStack[i] = true
	children := (*g.GetAdjList())[i]
	for c := range children {
		if isCyclicUtil(g, c, visited, recStack) {
			return true
		}
	}
	recStack[i] = false
	return false
}

// Test function tests FindCycleInDirectedGraph
func Test(g *directedgraph.Graph) {
	vertices := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(vertices); i++ {
		g.AddVertex(vertices[i])
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(2, 3)

	g.AddEdge(3, 6)

	g.AddEdge(4, 5)

	g.AddEdge(5, 6)
	g.AddEdge(5, 3)

	g.AddEdge(6, 1)

	isCycle := FindCycleInDirectedGraph(g)

	if isCycle {
		fmt.Printf("Graph is cyclic:\n")
		g.Print()
	} else {
		fmt.Printf("Graph is not cyclic:\n")
		g.Print()
	}

	g = directedgraph.CreateGraph()
	for i := 0; i < len(vertices); i++ {
		g.AddVertex(vertices[i])
	}
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(5, 5)
	g.AddEdge(5, 6)
	isCycle = FindCycleInDirectedGraph(g)
	if isCycle {
		fmt.Printf("Graph is cyclic:\n")
		g.Print()
	} else {
		fmt.Printf("Graph is not cyclic:\n")
		g.Print()
	}

	g = directedgraph.CreateGraph()
	for i := 0; i < len(vertices); i++ {
		g.AddVertex(vertices[i])
	}
	g.AddEdge(1, 2)
	isCycle = FindCycleInDirectedGraph(g)
	if isCycle {
		fmt.Printf("Graph is cyclic:\n")
		g.Print()
	} else {
		fmt.Printf("Graph is not cyclic:\n")
		g.Print()
	}
}
