package graph

// CycleDetector provides the interface for cycle detection in graph.
type CycleDetector interface {

	// HasCycle check if there is any cycle in the graph. For directed graph, it check directed cycle.
	HasCycle() bool

	// Cycle returns a cycle if there is. If there are multiple cycles, which on will be returned depends on the graph presentation.
	Cycle() []int
}
