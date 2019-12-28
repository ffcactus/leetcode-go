// Package ewd includes the API and implementations for the problems in edge-weighted directed graph.
package ewd

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// DirectedEdge defines the edge in edge-weighted directed graph.
type DirectedEdge struct {
	From int
	To int
	Weight float64
}

// Interface defines the interface of edge-weighted directed graph.
type Interface interface {
	Vertices() int
	Edges() int
	AddEdge(e DirectedEdge)
	Adjacent(v int) []DirectedEdge
	AllEdges() []DirectedEdge
}


type implementation struct {
	v int
	e int
	adj [][]DirectedEdge
}

func (impl *implementation) Vertices() int {
	return impl.v
}

func (impl *implementation) Edges() int {
	return impl.e
}

func (impl *implementation) AddEdge(e DirectedEdge) {
	impl.adj[e.From] = append(impl.adj[e.From], e)
	impl.e++
}

func (impl *implementation) Adjacent(v int) []DirectedEdge {
	return impl.adj[v]
}

func (impl *implementation) AllEdges() []DirectedEdge {
	ret := make([]DirectedEdge, 0)
	for _, v := range impl.adj {
		for _, vv := range v {
			ret = append(ret, vv)
		}
	}
	return ret
}

func NewEdgeWeightedDirectedGraph(v int) Interface {
	impl := implementation{
		v:   v,
		e:   0,
		adj: make([][]DirectedEdge, v),
	}
	return &impl
}

func NewEdgeWeightedDirectedGraphFromStream(r io.Reader) (Interface, error) {
	impl := implementation{}
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
	impl.adj = make([][]DirectedEdge, impl.v)

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
			from := int(v64)
			v64, err = strconv.ParseInt(pair[1], 10, 64)
			if err != nil {
				return nil, err
			}
			to := int(v64)
			weight, err := strconv.ParseFloat(pair[2], 64)
			if err != nil {
				return nil, err
			}
			e := DirectedEdge{
				From:   from,
				To:     to,
				Weight: weight,
			}
			impl.adj[e.From] = append(impl.adj[e.From], e)
		}
	}
	return &impl, nil
}

