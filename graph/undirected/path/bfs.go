package path

import "leetcode-go/graph"

// NewBFSPath return BFS implementation of interface Path.
// BFS can find out the shortest path.
func NewBFSPath(g graph.UndirectedGraph, s int) Path {
	impl := bfsImpl{}
	impl.marked = make([]bool, g.Vertices())
	impl.edgeTo = make([]int, g.Vertices())
	impl.s = s
	impl.bfs(g, s)
	return &impl
}

// bfsImpl is the BFS implementation of interface Path.
type bfsImpl struct {
	marked []bool // Has dfs been called for this vertex?
	edgeTo []int  // last vertex on known path to this vertex.
	s      int    // source.
}

func (impl *bfsImpl) bfs(g graph.UndirectedGraph, s int) {
	queue := make([]int, 0)
	impl.marked[s] = true
	queue = append([]int{s}, queue...)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range g.Adjacent(v) {
			if !impl.marked[w] {
				impl.edgeTo[w] = v
				impl.marked[w] = true
				queue = append([]int{w}, queue...)
			}
		}
	}
}

// HasPathTo check if there is a path from source to v.
// For BFS implementation it just check the marker initialized during the constructor.
func (impl *bfsImpl) HasPathTo(v int) bool {
	return impl.marked[v]
}

// PathTo return the path from source to v, and nil if no such a path.
// In the BFS implementation, since we recorded all the path from any vertex to the source,
// we can simply use this record.
func (impl *bfsImpl) PathTo(v int) []int {
	if !impl.HasPathTo(v) {
		return nil
	}
	ret := make([]int, 0)
	for x := v; x != impl.s; x = impl.edgeTo[x] {
		ret = append([]int{x}, ret...)
	}
	ret = append([]int{impl.s}, ret...)
	return ret
}
