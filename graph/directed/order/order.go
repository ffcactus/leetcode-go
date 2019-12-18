// Package order provider interface and implementation that output the vertices in order.
package order

import (
	"leetcode-go/graph"
)

// DFSOrder provider various ordered output.
type DFSOrder interface {

	// Pre return the vertices in pre-order.
	Pre() []int

	// Post return the vertices in post-order.
	Post() []int

	// ReversePost return the vertices in reverse post-order.
	ReversePost() []int
}

// dfsOrderImpl provider the implementation of interface DFSOrder.
type dfsOrderImpl struct {
	marked      []bool
	pre         []int
	post        []int
	reversePost []int
}

// NewDFSOrder creates a DFSOrder from a graph.
func NewDFSOrder(g graph.DirectedGraph) DFSOrder {
	impl := dfsOrderImpl{}
	impl.pre = make([]int, 0)
	impl.post = make([]int, 0)
	impl.reversePost = make([]int, 0)
	impl.marked = make([]bool, g.Vertices())
	for v := 0; v < g.Vertices(); v++ {
		if !impl.marked[v] {
			impl.dfs(g, v)
		}
	}
	return &impl
}

func (impl *dfsOrderImpl) dfs(g graph.DirectedGraph, v int) {
	impl.pre = append(impl.pre, v)
	impl.marked[v] = true
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.dfs(g, w)
		}
	}
	impl.post = append(impl.post, v)
	impl.reversePost = append([]int{v}, impl.reversePost...)
}

// Pre return the vertices in pre-order.
func (impl *dfsOrderImpl) Pre() []int {
	return impl.pre
}

// Post return the vertices in post-order.
func (impl *dfsOrderImpl) Post() []int {
	return impl.post
}

// ReversePost return the vertices in reverse post-order.
func (impl *dfsOrderImpl) ReversePost() []int {
	return impl.reversePost
}
