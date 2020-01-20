package edgeweighted

import (
	"container/heap"

	"leetcode-go/graph"
)

// MST provide the interface for minimal spinning tree.
type MST interface {
	// Edges returns all of the MST edges.
	Edges() []graph.Edge

	// Weight return the weight of MST.
	Weight() float64
}

// lazyPrimMST implements the MST interface with Prim algorithm.
type lazyPrimMST struct {
	marked   []bool
	mstQueue []graph.Edge
	pq       edgePQ
}

// NewLazyPrimMST constructs the MST implementation.
func NewLazyPrimMST(g graph.EdgeWeightedGraph) MST {
	impl := lazyPrimMST{
		marked:   make([]bool, g.Vertices()),
		mstQueue: make([]graph.Edge, 0),
		pq:       make(edgePQ, 0),
	}
	heap.Init(&impl.pq)
	// assumes G is connected.
	impl.visit(g, 0)
	for impl.pq.Len() > 0 {
		e := heap.Pop(&impl.pq).(graph.Edge) // get lowest-weight edge from pq.
		v := e.Either()
		w, _ := e.Other(v)
		if impl.marked[v] && impl.marked[w] { // skip if ineligible.
			continue
		}
		impl.mstQueue = append(impl.mstQueue, e) // add edge to tree.
		if !impl.marked[v] {
			impl.visit(g, v)
		}
		if !impl.marked[w] {
			impl.visit(g, w)
		}
	}
	return &impl
}

func (impl *lazyPrimMST) visit(g graph.EdgeWeightedGraph, v int) {
	// mark v and add to pq all edges from v to unmarked vertices.
	impl.marked[v] = true
	for _, e := range g.Adjacent(v) {
		other, _ := e.Other(v)
		if !impl.marked[other] {
			heap.Push(&impl.pq, e)
		}
	}
}

// Edges returns all of the MST edges.
func (impl *lazyPrimMST) Edges() []graph.Edge {
	return impl.mstQueue
}

// Weight return the weight of MST.
func (impl *lazyPrimMST) Weight() float64 {
	return 0
}

type edgePQ []graph.Edge

func (pq edgePQ) Len() int { return len(pq) }

func (pq edgePQ) Less(i, j int) bool {
	return pq[i].Weight() < pq[j].Weight()
}

func (pq edgePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *edgePQ) Push(x interface{}) {
	item := x.(graph.Edge)
	*pq = append(*pq, item)
}

func (pq *edgePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
