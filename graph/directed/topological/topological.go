// Package topological provides solutions for topological problem in directed graph.
package topological

type Topological interface {
	// IsDAG return if a directed graph is a directed acyclic graph (which has no directed cycle).
	IsDAG() bool

	// Order return the vertices in topological order.
	Order() []int
}
