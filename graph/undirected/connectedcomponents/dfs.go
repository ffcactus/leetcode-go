package connectedcomponents

import "leetcode-go/graph"

// dfsImpl implements the CC interface based on DFS algorithm.
type dfsImpl struct {
	marked []bool
	id     []int
	count  int
}

// NewDFSImpl returns a CC implementation based on DFS algorithm.
func NewDFSImpl(g graph.UndirectedGraph) CC {
	impl := dfsImpl{}
	impl.marked = make([]bool, g.Vertices())
	impl.id = make([]int, g.Vertices())
	for s := 0; s < g.Vertices(); s++ {
		// After a recursive call to dfs(), all the vertices that belongs to s will be marked,
		// So when we find another vertex that is not marked we find another component.
		if !impl.marked[s] {
			impl.dfs(g, s)
			impl.count++
		}
	}
	return &impl
}

func (impl *dfsImpl) dfs(g graph.UndirectedGraph, v int) {
	impl.marked[v] = true
	impl.id[v] = impl.count
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.dfs(g, w)
		}
	}
}

// Connected check if v and w are connected.
// Here we just check if they belong the same component.
func (impl *dfsImpl) Connected(v, w int) bool {
	return impl.id[v] == impl.id[w]
}

// Count return the number of connected components.
func (impl *dfsImpl) Count() int {
	return impl.count
}

// ID return the component identifier for vertex v, (between 0 and Count() - 1)
func (impl *dfsImpl) ID(v int) int {
	return impl.id[v]
}
