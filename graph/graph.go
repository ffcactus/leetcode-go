// Package graph includes various implementations related to algorithms in graph problems.
package graph

import (
	"errors"
	"fmt"
)

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

// Edge defines the edge used in edge weighted graph.
type Edge struct {
	weight float64
	v      int
	w      int
}

func (e Edge) Weight() float64 {
	return e.weight
}

func (e Edge) Either() int {
	return e.v
}

func (e Edge) Other(v int) (int, error) {
	if v == e.v {
		return e.w, nil
	} else if v == e.w {
		return e.v, nil
	}
	return 0, errors.New("inconsistent edge")
}

type EdgeWeightedGraph interface {
	// Vertices return the number of vertices.
	Vertices() int

	// Edges return the number of edges.
	Edges() int

	// AddEdge adds edge to this graph.
	AddEdge(e Edge)

	// Adjacent return all the edges incident to v.
	Adjacent(v int) []Edge

	// AllEdges return all the edges of this graph.
	AllEdges() []Edge
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}
