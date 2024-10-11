package visit

import "github.com/mcaci/graphgo/graph"

func Connected[T comparable](g graph.Graph[T]) bool {
	if len(g.Vertices()) == 0 {
		return true
	}
	return len(g.Vertices()) == Generic(g, g.Vertices()[0]).Size()
}

func ExistsPath[T comparable](g graph.Graph[T], x, y *graph.Vertex[T]) bool {
	return Generic(g, x).Contains(&y.E)
}
