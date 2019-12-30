package ewd

import (
	"container/heap"
	"math"
)

// ShortestPath defines the API for shortest path problem from a source vertex s.
// The s is provided from the constructor.
type ShortestPath interface {
	// DistanceTo return the directed path from vertex s to v. If there is no such a path return math.MaxFloat64.
	DistanceTo(v int) float64

	// HasPathTo return if there is a path from vertex s to v.
	HasPathTo(v int) bool

	// PathTo returns the shortest path from vertex s to v or nil if there is no such a path.
	PathTo(v int) []*DirectedEdge
}

type dijkstraSP struct {
	distTo []float64 // distTo[v] is the length of the shortest known path from s to v.
	edgeTo []*DirectedEdge // edgeTo[v] is the edge that connects v to its parent in the tree (the last edge from s to v)
	pq priorityQueue
}

// relax relax edge e is should not be in the shortest path.
func (impl *dijkstraSP) relax(e *DirectedEdge) {
	if impl.distTo[e.To] > impl.distTo[e.From] + e.Weight {
		impl.distTo[e.To] = impl.distTo[e.From] + e.Weight
		impl.edgeTo[e.To] = e
		if i := impl.pq.Contains(e.To); i != nil {
			impl.pq.Update(i, i.Value, impl.distTo[e.To])
		} else {
			heap.Push(&impl.pq, item{
				Value: e.To,
				Priority: impl.distTo[e.To],
			})
		}
	}
}

func (impl *dijkstraSP) Relax(g Interface, v int) {
	for _, w := range g.Adjacent(v) {
		impl.relax(&w)
	}
}

func (impl *dijkstraSP) DistanceTo(v int) float64 {
	return impl.distTo[v]
}

func (impl *dijkstraSP) HasPathTo(v int) bool {
	return impl.distTo[v] < math.MaxFloat64
}

func (impl *dijkstraSP) PathTo(v int) []*DirectedEdge {
	if !impl.HasPathTo(v) {
		return nil
	}
	path := make([]*DirectedEdge, 0)
	for e := impl.edgeTo[v]; e != nil; e = impl.edgeTo[e.From] {
		path = append([]*DirectedEdge{e}, path...)
	}
	return path
}

// NewDijkstraShortestPath creates the shortest path solution based on Dijkstra's algorithm.
func NewDijkstraShortestPath(g Interface, s int) ShortestPath {
	impl := dijkstraSP{
		distTo: make([]float64, g.Vertices()),
		edgeTo: make([]*DirectedEdge, g.Vertices()),
		pq:     make([]item, 0),
	}
	for i := range impl.distTo {
		impl.distTo[i] = math.MaxFloat64
	}
	impl.distTo[s] = 0
	heap.Init(&impl.pq)
	heap.Push(&impl.pq, item{
		Value: s,
		Priority: 0.0,
		Index:    0,
	})
	heap.Fix(&impl.pq, s)
	for len(impl.pq) > 0 {
		i := heap.Pop(&impl.pq)
		impl.Relax(g, i.(item).Value)
		// It i is the only target vertex you are going to search the shortest path, you can break here.
	}
	return &impl
}

// An item is something we manage in a Priority queue.
type item struct {
	Value int
	Priority float64    // The Priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The Index of the item in the heap.
}

// A priorityQueue implements heap.Interface and holds Items.
type priorityQueue []item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Update modifies the Priority and Value of an item in the queue.
func (pq *priorityQueue) Update(item *item, value int, priority float64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func (pq *priorityQueue) Contains(value int) *item {
	for i := 0; i <len(*pq); i++ {
		if (*pq)[i].Value == value {
			return &(*pq)[i]
		}
	}
	return nil
}