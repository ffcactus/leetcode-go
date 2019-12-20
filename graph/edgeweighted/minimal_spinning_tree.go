package edgeweighted

import "leetcode-go/graph"

// MST provide the interface for minimal spinning tree.
type MST interface {
	// Edges returns all of the MST edges.
	Edges() []graph.Edge

	// Weight return the weight of MST.
	Weight() float64
}

