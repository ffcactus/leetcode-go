// Package graph includes various implementations related to algorithms in graph problems.
package graph

// Graph provide the common interface for both directed and undirected graph.
type Graph interface {
	// Vertices return the number of vertices.
	Vertices() int

	// Edges return the number of edges.
	Edges() int

	// AddEdge add an edge between vertex v0 and v1.
	AddEdge(v0, v1 int)

	// Adjacent return the vertices that is adjacent to vertex v.
	Adjacent(v int) []int
}

// UndirectedGraph defines the interface for undirected graph.
type UndirectedGraph interface {
	Graph
}

// DirectedGraph defines the interface for directed graph.
type DirectedGraph interface {
	Graph

	// Reverse reverse this graph.
	Reverse() DirectedGraph
}
