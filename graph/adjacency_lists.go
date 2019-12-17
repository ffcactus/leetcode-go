package graph

import (
	"bufio"
	"io"
	"strconv"
)

// adjacencyLists represents a way to save the graph information.
type adjacencyLists struct {
	v int // number of vertices.
	e int // number of edges.
	adj [][]int // adjacency list.
}

// NewAdjacencyLists return a graph that is represented by adjacency lists,
// v is the number of vertices in the graph.
func NewAdjacencyLists(v int) Graph {
	impl := adjacencyLists{
		v:   v,
	}
	for i:=0; i < v; i++ {
		impl.adj = append(impl.adj, make([]int, 0))
	}
	return &impl
}

// NewAdjacencyListsFromReader return a graph that is represented by adjacency lists,
// The graph will be generated from reader r.
// The first line is the number of vertices.
// The second line is the number of edges.
// The rest of lines are pair of vertex that is connected directly, separated by space.
func NewAdjacencyListsFromReader(r io.Reader) Graph, error{
	impl := adjacencyLists{}
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		if v64, err := strconv.ParseInt(scanner.Text(), 10, 64); err != nil {
			return nil
		}
		impl.v = int(v64)
	}

}

// AddEdge add an edge between vertex v0 and v1.
// For adjacency lists, we need record both from v0 to v1 and from v1 to v0.
func (impl *adjacencyLists) AddEdge(v0, v1 int) {
	impl.adj[v0] = append(impl.adj[v0], v1)
	impl.adj[v1] = append(impl.adj[v1], v0)
	impl.e++
}

// Adjacent return the vertices that is adjacent to vertex v.
// For adjacency lists, we can simply return the lists at index of v.
func (impl *adjacencyLists) Adjacent(v int) []int {
	return impl.adj[v]
}

// Vertices return the number of vertices.
func (impl *adjacencyLists) Vertices() int {
	return impl.v
}

// Edges return the number of edges.
func (impl *adjacencyLists) Edges() int {
	return impl.e
}