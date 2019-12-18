package graph

// Search defines the interface for a search operation in a graph concerned to a source.
type Search interface {
	// Marked return if v is connected to the source in the search.
	Marked(v int) bool

	// Count return the number of vertices that is connected to source.
	Count() int
}