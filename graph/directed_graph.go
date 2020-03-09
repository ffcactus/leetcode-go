package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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

// NewDirectedGraphFromReader return a directed graph that is represented by adjacency lists,
// The graph will be generated from reader r.
// The first line is the number of vertices.
// The second line is the number of edges.
// The rest of lines are pair of vertex that is connected directly, separated by space.
func NewDirectedGraphFromReader(r io.Reader) (DirectedGraph, error) {
	impl := directedGraphImpl{}
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
