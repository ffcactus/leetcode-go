// Package twocolor provide the solution to check if the graph's vertex can be colored in two color that no edge connected to
// the vertices with the same color.
package twocolor

import "leetcode-go/graph"

// DFSImpl provider the DFS solution to solve the two color problem.
type DFSImpl struct {
	marked         []bool
	color          []bool
	isTwoColorable bool
}

// NewDFSImpl generate the DFSImpl from graph.
func NewDFSImpl(g graph.UndirectedGraph) *DFSImpl {
	impl := DFSImpl{
		marked:         make([]bool, g.Vertices()),
		color:          make([]bool, g.Vertices()),
		isTwoColorable: true,
	}
	return &impl
}

func (impl *DFSImpl) dfs(g graph.UndirectedGraph, v int) {
	if impl.isTwoColorable == false {
		return
	}
	impl.marked[v] = true
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.color[w] = !impl.color[v]
			impl.dfs(g, w)
		} else if impl.color[w] == impl.color[v] {
			impl.isTwoColorable = false
			return
		}
	}
}

func (impl *DFSImpl) IsTwoColorable() bool {
	return impl.isTwoColorable
}
