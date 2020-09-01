package acyclicgraph

import (
	"errors"
	"fmt"

	"github.com/morzhanov/algorithms/src/ds/queue"
)

// AdjList represents graph adjacency list
type AdjList map[int][]int

// Graph represents undirected acyclic graph
type Graph struct {
	vertexCount int
	adjList     AdjList
}

// CreateGraph creates graph
func CreateGraph() *Graph {
	return &Graph{0, make(AdjList, 0)}
}

// AddVertex function adds vertex to the graph
func (g *Graph) AddVertex(v int) {
	g.adjList[v] = make([]int, 0)
	g.vertexCount++
}

// AddEdge function adds edge to the graph
func (g *Graph) AddEdge(v int, w int) error {
	// dont allow cycles
	if v == w {
		return errors.New("Cycles are not allowed")
	}
	// get the list for vertex v and put the
	// vertex w denoting edge between v and w
	g.adjList[v] = append(g.adjList[v], w)
	// since graph is undirected,
	// add an edge from w to v also
	g.adjList[w] = append(g.adjList[w], v)
	return nil
}

// Print function prints the vertex and adjacency list
func (g *Graph) Print() {
	for k, v := range g.adjList {
		conc := ""
		for _, val := range v {
			conc += fmt.Sprintf("%v  ", val)
		}
		fmt.Printf("%v -> %v\n", k, conc)
	}
}

// BFS function implements graph BFS
func (g *Graph) BFS(node int) (res []int) {
	res = make([]int, 0)
	// create a visited array
	visited := make(map[int]bool, 0)
	for i := 0; i < g.vertexCount; i++ {
		visited[i] = false
	}

	// create an object for queue
	var q queue.Queue

	// add the starting node to the queue
	visited[node] = true
	q.Enqueue(float64(node))

	// loop until queue is element
	for !q.IsEmpty() {
		// get the element from the queue
		getQueueElement, _ := q.Dequeue()

		// saving the current vertex
		res = append(res, int(getQueueElement))

		// get the adjacent list for current vertex
		list := g.adjList[int(getQueueElement)]

		// loop through the list and add the element to the
		// queue if it is not processed yet
		for _, val := range list {
			if !visited[val] {
				visited[val] = true
				q.Enqueue(float64(val))
			}
		}
	}
	return res
}

// DFS function implements graph DFS
func (g *Graph) DFS(node int) (res []int) {
	// create a visited array
	visited := make(map[int]bool, 0)
	for i := 0; i < g.vertexCount; i++ {
		visited[i] = false
	}

	return g.dfsUtil(node, visited)
}

// DFS utility function
func (g *Graph) dfsUtil(node int, visited map[int]bool) (res []int) {
	res = []int{node}
	visited[node] = true
	neighbours := g.adjList[node]

	for _, value := range neighbours {
		if !visited[value] {
			res = append(res, g.dfsUtil(value, visited)...)
		}
	}
	return res
}

// Test function tests graph
func Test(g *Graph) {
	vertices := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(vertices); i++ {
		g.AddVertex(vertices[i])
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(5, 3)
	g.AddEdge(3, 6)
	// cycle, get and show error
	err := g.AddEdge(2, 2)

	if err != nil {
		fmt.Printf("Error when trying to add cycle edge: %v\n", err)
	}

	g.Print()

	fmt.Printf("BFS from 1: %v\n", g.BFS(1))
	fmt.Printf("BFS from 6: %v\n", g.BFS(6))
	fmt.Printf("DFS from 1: %v\n", g.DFS(1))
	fmt.Printf("DFS from 6: %v\n", g.DFS(6))
}
