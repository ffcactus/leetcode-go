package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// undirectedGraphImpl represents a way to save the undirected graph information.
type undirectedGraphImpl struct {
	v   int     // number of vertices.
	e   int     // number of edges.
	adj [][]int // adjacency list.
}

// NewUndirectedGraph return a undirected graph that is represented by adjacency lists,
// v is the number of vertices in the graph.
func NewUndirectedGraph(v int) UndirectedGraph {
	impl := undirectedGraphImpl{
		v: v,
	}
	for i := 0; i < v; i++ {
		impl.adj = append(impl.adj, make([]int, 0))
	}
	return &impl
}

// NewUndirectedGraphFromReader return a undirected graph that is represented by adjacency lists,
// The graph will be generated from reader r.
// The first line is the number of vertices.
// The second line is the number of edges.
// The rest of lines are pair of vertex that is connected directly, separated by space.
func NewUndirectedGraphFromReader(r io.Reader) (UndirectedGraph, error) {
	impl := undirectedGraphImpl{}
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		v64, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		impl.v = int(v64)
	}
	if scanner.Scan() {
		v64, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		impl.e = int(v64)
	}
	for i := 0; i < impl.v; i++ {
		impl.adj = append(impl.adj, make([]int, 0))
	}
	for i := 0; i < impl.e; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		pair := strings.Split(line, " ")
		if len(pair) == 2 {
			v64, err := strconv.ParseInt(pair[0], 10, 64)
			if err != nil {
				return nil, err
			}
			v0 := int(v64)
			v64, err = strconv.ParseInt(pair[1], 10, 64)
			if err != nil {
				return nil, err
			}
			v1 := int(v64)
			impl.AddEdge(v0, v1)
		}
	}
	return &impl, nil
}

// AddEdge add an edge between vertex v0 and v1.
// For adjacency lists, we need record both from v0 to v1 and from v1 to v0.
func (impl *undirectedGraphImpl) AddEdge(v0, v1 int) {
	impl.adj[v0] = append(impl.adj[v0], v1)
	impl.adj[v1] = append(impl.adj[v1], v0)
	impl.e++
}

// Adjacent return the vertices that is adjacent to vertex v.
// For adjacency lists, we can simply return the lists at index of v.
func (impl *undirectedGraphImpl) Adjacent(v int) []int {
	return impl.adj[v]
}

// Vertices return the number of vertices.
func (impl *undirectedGraphImpl) Vertices() int {
	return impl.v
}

// Edges return the number of edges.
func (impl *undirectedGraphImpl) Edges() int {
	return impl.e
}
