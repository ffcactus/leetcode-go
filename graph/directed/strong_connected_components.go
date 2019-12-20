package directed

import (
	"leetcode-go/graph"
	order "leetcode-go/graph/directed/order"
)

// SCCChecker provide the interface to check strong connected components.
type SCCChecker interface {

	// IsStrongConnected check if v and w are strongly connected.
	IsStrongConnected(v, w int) bool

	// Count return the number of strong components.
	Count() int

	// ID returns the strong component identifier for vertex v (between 0 and Count() -1)
	ID(v int) int
}

type kosaraju struct {
	marked []bool
	id []int
	count int
}

// NewKosarajuSSCChecker return a SCCChecker based on Kosaraju's algorithm.
func NewKosarajuSSCChecker(g graph.DirectedGraph) SCCChecker {
	impl := kosaraju{
		marked: make([]bool, g.Vertices()),
		id:     make([]int, g.Vertices()),
		count:  0,
	}
	dfsOrder := order.NewDFSOrder(g.Reverse())
	for _, w := range dfsOrder.ReversePost() {
		if !impl.marked[w] {
			impl.dfs(g, w)
			impl.count++
		}
	}
	return &impl
}

func (impl *kosaraju) dfs(g graph.DirectedGraph, v int) {
	impl.marked[v] = true
	impl.id[v] = impl.count
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.dfs(g, w)
		}
	}
}

// IsStrongConnected check if v and w are strongly connected.
// Here we simply check if they belongs to the same strong component.
func (impl *kosaraju) IsStrongConnected(v, w int) bool {
	return impl.id[v] == impl.id[w]
}

// Count return the number of strong components.
func (impl *kosaraju) Count() int {
	return impl.count
}

// ID returns the strong component identifier for vertex v (between 0 and Count() -1)
func (impl *kosaraju) ID(v int) int {
	return impl.id[v]
}
