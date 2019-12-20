package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// NewEdgeWeightedGraph generate an edge weighted graph that have a specific number of vertices.
func NewEdgeWeightedGraph(numOfVertices int) EdgeWeightedGraph {
	impl := adjacentListImplOfEWG{
		v: numOfVertices,
		adj: make([][]Edge, numOfVertices),
	}
	return &impl
}

// NewEdgeWeightedGraphFromReader generated an edge weighted graph from a reader.
// The first line is the number of vertices.
// The second line is the number of edges.
// The rest of lines, each includes a pair of vertex and the weight of edge, separated by space.
func NewEdgeWeightedGraphFromReader(r io.Reader) (EdgeWeightedGraph, error) {
	impl := adjacentListImplOfEWG{}
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
	impl.adj = make([][]Edge, impl.v)
	for i := 0; i < impl.e; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		pair := strings.Split(line, " ")
		if len(pair) == 3 {
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
			w64, err := strconv.ParseFloat(pair[2], 64)
			if err != nil {
				return nil, err
			}
			impl.AddEdge(Edge{
				v:v0,
				w:v1,
				weight:w64,
			})
		}
	}
	return &impl, nil
}

// adjacentListImplOfEWG implements the EdgeWeightedGraph based on adjacent lists.
type adjacentListImplOfEWG struct {
	v int
	e int
	adj [][]Edge
}

// Vertices return the number of vertices.
func (impl *adjacentListImplOfEWG) Vertices() int {
	return impl.v
}

// Edges return the number of edges.
func (impl *adjacentListImplOfEWG) Edges() int {
	return impl.e
}

// AddEdge adds edge to this graph.
func (impl *adjacentListImplOfEWG) AddEdge(e Edge) {
	v := e.Either()
	w, _ := e.Other(v)
	impl.adj[v] = append(impl.adj[v], e)
	impl.adj[w] = append(impl.adj[w], e)
	impl.e++
}

// Adjacent return all the edges incident to v.
func (impl *adjacentListImplOfEWG) Adjacent(v int) []Edge {
	return impl.adj[v]
}

// AllEdges return all the edges of this graph.
func (impl *adjacentListImplOfEWG) AllEdges() []Edge {
	ret := make([]Edge, 0)
	for v := 0; v < impl.v; v++ {
		for _, e := range impl.Adjacent(v) {
			if w, _ := e.Other(v); w > v {
				ret = append(ret, e)
			}
		}
	}
	return ret
}