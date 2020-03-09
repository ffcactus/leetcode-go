package finish

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := NewDirectedGraphFromArray(numCourses, prerequisites)
	detector := NewCycleDetection(g)
	return !detector.HasCycle()
}

// dfsCycleDetection implements the cycle interface based on DFS.
type dfsCycleDetection struct {
	marked  []bool
	edgeTo  []int
	cycle   []int  // vertices on the cycle.
	onStack []bool // vertices on the call stack.
}

// NewCycleDetection return a CycleDetector from directed graph.
func NewCycleDetection(g DirectedGraph) CycleDetector {
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

func (impl *dfsCycleDetection) dfs(g DirectedGraph, v int) {
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

// CycleDetector provides the interface for cycle detection in graph.
type CycleDetector interface {

	// HasCycle check if there is any cycle in the graph. For directed graph, it check directed cycle.
	HasCycle() bool

	// Cycle returns a cycle if there is. If there are multiple cycles, which on will be returned depends on the graph presentation.
	Cycle() []int
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

// directedGraphImpl represents a way to save the directed graph information.
type directedGraphImpl struct {
	v   int     // number of vertices.
	e   int     // number of edges.
	adj [][]int // adjacency list.
}

// NewDirectedGraph return a undirected graph that is represented by adjacency lists,
// v is the number of vertices in the graph.
func NewDirectedGraph(v int) DirectedGraph {
	impl := directedGraphImpl{
		v: v,
	}
	for i := 0; i < v; i++ {
		impl.adj = append(impl.adj, make([]int, 0))
	}
	return &impl
}

// NewDirectedGraphFromArray construct the directed graph from array.
func NewDirectedGraphFromArray(v int, edges [][]int) DirectedGraph {
	impl := directedGraphImpl{
		v:   v,
		e:   len(edges),
		adj: make([][]int, v),
	}
	for _, edge := range edges {
		impl.AddEdge(edge[0], edge[1])
	}
	return &impl
}

// DirectedGraph defines the interface for directed graph.
type DirectedGraph interface {
	Graph

	// Reverse reverse this graph.
	Reverse() DirectedGraph
}

// Graph provide the common interface for both directed and undirected graph.
type Graph interface {
	// Vertices return the number of vertices.
	Vertices() int

	// Edges return the number of edges.
	Edges() int

	// AddEdge add an edge between vertex v0 and v1.
	AddEdge(v0, v1 int)

	// Adjacent return the vertices that is adjacent to vertex v.
	Adjacent(v int) []int
}

// AddEdge add an edge v0->v1.
func (impl *directedGraphImpl) AddEdge(v0, v1 int) {
	impl.adj[v0] = append(impl.adj[v0], v1)
	impl.e++
}

// Adjacent return the vertices that is adjacent to vertex v.
// For adjacency lists, we can simply return the lists at index of v.
func (impl *directedGraphImpl) Adjacent(v int) []int {
	return impl.adj[v]
}

// Vertices return the number of vertices.
func (impl *directedGraphImpl) Vertices() int {
	return impl.v
}

// Edges return the number of edges.
func (impl *directedGraphImpl) Edges() int {
	return impl.e
}

// Reverse reverse this graph.
func (impl *directedGraphImpl) Reverse() DirectedGraph {
	newOne := NewDirectedGraph(impl.v)
	for v := 0; v < impl.v; v++ {
		for _, w := range impl.Adjacent(v) {
			newOne.AddEdge(w, v)
		}
	}
	return newOne
}
