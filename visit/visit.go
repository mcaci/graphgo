package visit

import (
	"github.com/mcaci/graphgo/graph"
)

func Generic[T comparable](g graph.Graph[T], v graph.Vertex[T]) *Tree[T] {
	return &Tree[T]{}
}
