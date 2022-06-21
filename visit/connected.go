package visit

import "github.com/mcaci/graphgo/graph"

func Connected[T comparable](g graph.Graph[T]) bool {
	return len(g.Vertices()) == Generic(g, g.Vertices()[0]).Size()
}
