package graph

import (
	"slices"
)

type Graph[T comparable] interface {
	AreAdjacent(a, b *Vertex[T]) bool
	Degree(v *Vertex[T]) int
	AdjacentNodes(v *Vertex[T]) []*Vertex[T]
	IsDirected() bool

	Vertices() []*Vertex[T]
	Edges() []*Edge[T]
	AddVertex(v *Vertex[T])
	RemoveVertex(v *Vertex[T])
	ContainsVertex(v *Vertex[T]) bool
	AddEdge(e *Edge[T])
	RemoveEdge(e *Edge[T])
	ContainsEdge(e *Edge[T]) bool
}

const (
	ArcsListType = iota
	AdjacencyListType
	AdjacencyMatrixType
	// IncidenceListType
	// IncidenceMatrixType
)

func New[T comparable](graphType int, isDirected bool) Graph[T] {
	switch graphType {
	case ArcsListType:
		return &ArcsList[T]{directed: isDirected}
	case AdjacencyListType:
		return &AdjacencyLists[T]{directed: isDirected}
	case AdjacencyMatrixType:
		return &AdjacencyMatrix[T]{directed: isDirected}
	default:
		return &ArcsList[T]{}
	}
}

func Fill[T comparable](vs []*Vertex[T], es []*Edge[T], into Graph[T]) {
	for _, v := range vs {
		if into.ContainsVertex(v) {
			continue
		}
		into.AddVertex(v)
	}
	for _, e := range es {
		if into.ContainsEdge(e) {
			continue
		}
		v, u := vs[indexVertex(vs, e.X)], vs[indexVertex(vs, e.Y)]
		into.AddEdge(&Edge[T]{X: v, Y: u, P: e.P})
	}
}

func indexVertex[T comparable](vs []*Vertex[T], v *Vertex[T]) int {
	return slices.IndexFunc(vs, func(o *Vertex[T]) bool { return v.E == o.E })
}

func indexEdge[T comparable](g Graph[T], e *Edge[T]) int {
	edges := g.Edges()
	switch g.IsDirected() {
	case true:
		return slices.IndexFunc(edges, func(o *Edge[T]) bool {
			return e.X.E == o.X.E && e.Y.E == o.Y.E
		})
	case false:
		return slices.IndexFunc(edges, func(o *Edge[T]) bool {
			x2y := e.X.E == o.X.E && e.Y.E == o.Y.E
			y2x := e.X.E == o.Y.E && e.Y.E == o.X.E
			return x2y || y2x
		})
	default:
		return -1
	}
}

func Copy[T comparable](g Graph[T]) Graph[T] {
var o Graph[T]
switch g.(type) {
case *ArcsList:
		o = &ArcsList[T]{directed: g.isDirected}
	case *AdjacencyLists:
		o = &AdjacencyLists[T]{directed: g.isDirected}
	case *AdjacencyMatrix:
		o = &AdjacencyMatrix[T]{directed: g.isDirected}
default:
  o = &ArcsList[T]{directed: g.isDirected}
}
                for _, v := range g.Vertices() {
                        o.AddVertex(v)
                }
                for _, e := range g.Edges() {
                        o.AddEdge(e)
                }
return o
}