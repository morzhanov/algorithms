package dijkstra

import (
	"fmt"
	"math"
)

const maxInt = int(math.MaxInt16 >> 1)

func minDistance(dist []int, sptSet []bool, V int) int {
	// initialize min value
	min := maxInt
	minIndex := -1

	for v := 0; v < V; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}

// a utility function to print the constructed distance array
func printSolution(V int, dist []int) {
	fmt.Printf("Dijkstra algorithm results: \n")
	fmt.Printf("Vertex \t\t Distance from Source\n")
	for i := 0; i < V; i++ {
		fmt.Printf("%v \t\t %v\n", i, dist[i])
	}
}

// Dijkstra function calulates distances of shortest paths from src to all vertices. It is a O(ELogV) function
func Dijkstra(graph [][]int, src int) {
	V := len(graph)
	// the output array. dist[i] will hold the shortest distance from src to i
	dist := make([]int, V)
	// sptSet[i] will true if vertex i is included in shortest
	// path tree or shortest distance from src to i is finalized
	sptSet := make([]bool, V)
	// initialize all distances as INFINITE and stpSet[] as false
	for i := 0; i < V; i++ {
		dist[i] = maxInt
		sptSet[i] = false
	}
	// distance of source vertex from itself is always 0
	dist[src] = 0

	// find shortest path for all vertices
	for count := 0; count < V-1; count++ {
		// pick the minimum distance vertex from the set of vertices
		// not yet processed. u is always equal to src in first iteration.
		u := minDistance(dist, sptSet, V)
		// Mark the picked vertex as processed
		sptSet[u] = true
		// Update dist value of the adjacent vertices of the
		// picked vertex.
		for v := 0; v < V; v++ {
			// update dist[v] only if is not in sptSet, there is an
			// edge from u to v, and total weight of path from src to
			// v through u is smaller than current value of dist[v]
			if !sptSet[v] && graph[u][v] != 0 && dist[u] != maxInt && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}
	// print the constructed distance array
	printSolution(V, dist)
}

// Test function tests Dijkstra algorithm
func Test() {
	/* Let us create the example graph discussed above */
	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}
	Dijkstra(graph, 0)
}
