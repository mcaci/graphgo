package path

import (
	"math"

	"github.com/mcaci/graphgo/graph"
)

func BellmanFordDist[T comparable](g graph.Graph[T], s *graph.Vertex[T]) map[*graph.Vertex[T]]*Distance[T] {
	d := make(map[*graph.Vertex[T]]*Distance[T])
	vs := g.Vertices()
	for i := range vs {
		v := vs[i]
		var dist int
		if v != s {
			dist = math.MaxInt
		}
		d[v] = &Distance[T]{v: s, u: v, d: dist}
	}
	canRelax := func(x, y *graph.Vertex[T], w int) bool { return d[x].d+w < d[y].d && d[x].d+w > 0 }
	relax := func(x, y *graph.Vertex[T], w int) { d[y].SetDist(w + d[x].d) }
	es := g.Edges()
	for range vs {
		for _, e := range es {
			switch {
			case canRelax(e.X, e.Y, e.P.W):
				relax(e.X, e.Y, e.P.W)
			case canRelax(e.Y, e.X, e.P.W):
				relax(e.Y, e.X, e.P.W)
			}
		}
	}
	return d
}
