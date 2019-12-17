package path

// Path define the API for the problem that is there a path in the graph from a source.
type Path interface {

	// HasPathTo check if there is a path from source to v.
	HasPathTo(v int) bool

	// PathTo return the path from source to v, and nil if no such a path.
	PathTo(v int) []int
}


