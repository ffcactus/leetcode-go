// Package cycle provide an implementation to check if there is a cycle in the graph.
// The implementation use DFS to solve the problem.
package cycle

import "leetcode-go/graph"

// DFSImpl use DFS to check if there is a cycle in the graph.
type DFSImpl struct {
	marked   []bool
	hasCycle bool
}

// NewDFSImpl generate the DFSImpl from a graph.
func NewDFSImpl(g graph.UndirectedGraph) *DFSImpl {
	impl := DFSImpl{
		marked:   make([]bool, g.Vertices()),
		hasCycle: false,
	}
	for i := 0; i < g.Vertices(); i++ {
		if !impl.marked[i] {
			impl.dfs(g, i, i)
		}
	}
	return &impl
}

func (impl *DFSImpl) dfs(g graph.UndirectedGraph, next, previous int) {
	if impl.hasCycle {
		return
	}
	impl.marked[next] = true
	for _, w := range g.Adjacent(next) {
		if !impl.marked[w] {
			impl.dfs(g, w, next)
		} else if w != previous { // if your next's next vertex is marked and not equals to your previous vertex, it means there is a cycle.
			impl.hasCycle = true
			return
		}
	}
}

// HasCycle return if there is a cycle in the graph.
func (impl *DFSImpl) HasCycle() bool {
	return impl.hasCycle
}
