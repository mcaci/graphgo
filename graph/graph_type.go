package graph

type Type[T comparable] interface {
	AreAdjacent(a, b *Vertex[T]) bool
	Degree(n *Vertex[T]) int
	AdjacentNodes(n *Vertex[T]) []*Vertex[T]

	AddVertex(v *Vertex[T])
	RemoveVertex(v *Vertex[T])
	ContainsVertex(v *Vertex[T]) bool
	AddEdge(e *Edge[T])
	RemoveEdge(e *Edge[T])
	ContainsEdge(e *Edge[T]) bool
}
