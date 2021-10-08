package topologicalsorting

import (
	"fmt"

	"github.com/morzhanov/algorithms/ds/dag"
	"github.com/morzhanov/algorithms/ds/stack"
)

// TopologicalSorting function implement graph topoligocal sorting algorithm
func TopologicalSorting(g *dag.Graph) {
	stack := stack.Stack{}
	// mark all the vertices as not visited
	visited := make([]bool, g.GetVertexCount()+1)
	// call the recursive helper function to store
	// Topological Sort starting from all vertices one by one
	for i := 0; i < g.GetVertexCount(); i++ {
		if visited[i] == false {
			topologicalSortUtil(g, i, visited, &stack)
		}
	}
	// print contents of stack
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		fmt.Printf("%v\n", int(val))
	}
}

// a recursive function used by topologicalSort
func topologicalSortUtil(g *dag.Graph, v int, visited []bool, stack *stack.Stack) {
	// mark the current node as visited.
	visited[v] = true
	// recur for all the vertices adjacent to this vertex
	for _, val := range (*g.GetAdjList())[v] {
		if !visited[val] {
			topologicalSortUtil(g, val, visited, stack)
		}
	}
	// push current vertex to stack which stores result
	stack.Push(float64(v))
}

// Test function tests TopologicalSorting
func Test(g *dag.Graph) {
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

	TopologicalSorting(g)
}
