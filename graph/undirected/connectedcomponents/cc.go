package connectedcomponents

// CC defines the interface for connected components in graph.
type CC interface {

	// Connected check if v and w are connected.
	Connected(v, w int) bool

	// Count return the number of connected components.
	Count() int

	// ID return the component identifier for vertex v, (between 0 and Count() - 1)
	ID(v int) int
}
