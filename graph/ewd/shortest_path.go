package ewd

import "math"

// ShortestPath defines the API for shortest path problem from a source vertex s.
// The s is provided from the constructor.
type ShortestPath interface {
	// DistanceTo return the directed path from vertex s to v. If there is no such a path return math.MaxFloat64.
	DistanceTo(v int) float64

	// HasPathTo return if there is a path from vertex s to v.
	HasPathTo(v int) bool

	// PathTo returns the shortest path from vertex s to v or nil if there is no such a path.
	PathTo(v int) []DirectedEdge
}

type shortestPath struct {
	distTo []float64 // distTo[v] is the length of the shortest known path from s to v.
	edgeTo []*DirectedEdge // edgeTo[v] is the edge that connects v to its parent in the tree (the last edge from s to v)
}

// relax relax edge e is should not be in the shortest path.
func (impl *shortestPath) relax(e *DirectedEdge) {
	if impl.distTo[e.To] > impl.distTo[e.From] + e.Weight {
		impl.distTo[e.To] = impl.distTo[e.From] + e.Weight
		impl.edgeTo[e.To] = e
	}
}

func (impl *shortestPath) Relax(g Interface, v int) {
	for i := range g.Adjacent(v) {
		impl.relax(&g.Adjacent(v)[i])
	}
}

func (impl *shortestPath) DistanceTo(v int) float64 {
	return impl.distTo[v]
}

func (impl *shortestPath) HasPathTo(v int) bool {
	return impl.distTo[v] < math.MaxFloat64
}

func (impl *shortestPath) PathTo(v int) []*DirectedEdge {
	if !impl.HasPathTo(v) {
		return nil
	}
	path := make([]*DirectedEdge, 0)
	for e := impl.edgeTo[v]; e != nil; e = impl.edgeTo[e.From] {
		path = append([]*DirectedEdge{e}, path...)
	}
	return path
}

