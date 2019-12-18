package directed

import "leetcode-go/graph"

// dfsCycleDetection implements the cycle interface based on DFS.
type dfsCycleDetection struct {
	marked  []bool
	edgeTo  []int
	cycle   []int  // vertices on the cycle.
	onStack []bool // vertices on the call stack.
}

func NewCycleDetection(g graph.DirectedGraph) graph.CycleDetector {
	impl := dfsCycleDetection{}
	impl.onStack = make([]bool, g.Vertices())
	impl.edgeTo = make([]int, g.Vertices())
	impl.marked = make([]bool, g.Vertices())
	for v := 0; v < g.Vertices(); v++ {
		if !impl.marked[v] {
			impl.dfs(g, v)
		}
	}
	return &impl
}

func (impl *dfsCycleDetection) dfs(g graph.DirectedGraph, v int) {
	impl.onStack[v] = true
	impl.marked[v] = true
	for _, w := range g.Adjacent(v) {
		if impl.HasCycle() {
			return
		}
		if !impl.marked[w] {
			impl.edgeTo[w] = v
			impl.dfs(g, w)
		} else if impl.onStack[w] { // records the vertices on the cycle.
			impl.cycle = make([]int, 0)
			for x := v; x != w; x = impl.edgeTo[x] {
				impl.cycle = append([]int{x}, impl.cycle...)
			}
			impl.cycle = append([]int{v, w}, impl.cycle...)
		}
	}
	impl.onStack[v] = false
}

// Cycle returns a cycle if there is. If there are multiple cycles, which on will be returned depends on the graph presentation.
// Since we already records all the vertices on the stack, so if the stack is empty, there is no cycle.
func (impl *dfsCycleDetection) HasCycle() bool {
	return len(impl.cycle) != 0
}

// Cycle returns a cycle if there is. If there are multiple cycles, which on will be returned depends on the graph presentation.
// Here, just return the records.
func (impl *dfsCycleDetection) Cycle() []int {
	return impl.cycle
}
