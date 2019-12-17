package path

import "leetcode-go/graph"

// NewDFSPath return DFS implementation of interface Path.
func NewDFSPath(g graph.UndirectedGraph, s int) Path {
	impl := dfsImpl{}
	impl.marked = make([]bool, g.Vertices())
	impl.edgeTo = make([]int, g.Vertices())
	impl.s = s
	impl.dfs(g, s)
	return &impl
}

type dfsImpl struct {
	marked []bool // Has dfs been called for this vertex?
	edgeTo []int  // last vertex on known path to this vertex.
	s      int    // source.
}

func (impl *dfsImpl) dfs(g graph.UndirectedGraph, v int) {
	impl.marked[v] = true
	for _, w := range g.Adjacent(v) {
		if !impl.marked[w] {
			impl.edgeTo[w] = v
			impl.dfs(g, w)
		}
	}
}

// HasPathTo check if there is a path from source to v.
// For DFS implementation it just check the marker initialized during the constructor.
func (impl *dfsImpl) HasPathTo(v int) bool {
	return impl.marked[v]
}

// PathTo return the path from source to v, and nil if no such a path.
// In the DFS implementation, since we recorded all the path from any vertex to the source,
// we can simply use this record.
func (impl *dfsImpl) PathTo(v int) []int {
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
