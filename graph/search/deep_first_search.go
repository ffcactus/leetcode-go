package search

import "leetcode-go/graph"

// NewDepthFirstSearch return a deep first search.
func NewDepthFirstSearch(g graph.Graph, s int) Search {
	impl := depthFirstSearch{}
	impl.marked = make([]bool, g.Vertices())
	impl.dfs(g, s) // mark all the vertices that is connected to the source and get the count.
	return &impl
}

// depthFirstSearch provide the depth first search of Search interface.
type depthFirstSearch struct {
	marked []bool
	count int
}

func (impl *depthFirstSearch) dfs(g graph.Graph, v int) {
	impl.marked[v] = true
	impl.count++
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.dfs(g, w)
		}
	}
}

// Marked return if v is connected to the source in the search.
func (impl *depthFirstSearch) Marked(v int) bool {
	return impl.marked[v]
}

// Count return the number of vertices that is connected to source.
func (impl *depthFirstSearch) Count() int {
	return impl.count
}
