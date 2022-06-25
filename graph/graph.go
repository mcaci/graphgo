package graph

import "io"

type Graph[T comparable] interface {
	AreAdjacent(a, b *Vertex[T]) bool
	Degree(n *Vertex[T]) int
	AdjacentNodes(n *Vertex[T]) []*Vertex[T]

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

func New[T comparable](graphType int) Graph[T] {
	switch graphType {
	case ArcsListType:
		return &ArcsList[T]{}
	case AdjacencyListType:
		return &AdjacencyLists[T]{}
	case AdjacencyMatrixType:
		return &AdjacencyMatrix[T]{}
	default:
		return &ArcsList[T]{}
	}
}

func Create(graphType int, r io.Reader) Graph[string] {
	switch graphType {
	case ArcsListType:
		return NewArcsList(r)
	case AdjacencyListType:
		return NewAdjacencyList(r)
	case AdjacencyMatrixType:
		return NewAdjacencyMatrix(r)
	// case IncidenceListType:
	// 	return NewIncidenceList(r)
	// case IncidenceMatrixType:
	// 	return NewIncidenceMatrix(r)
	default:
		return NewArcsList(r)
	}
}
