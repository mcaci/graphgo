package path

import (
	"math"

	"github.com/mcaci/graphgo/graph"
)

type EdgeWeight int

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
	canRelax := func(x, y *graph.Vertex[T], w EdgeWeight) bool { return d[x].d+int(w) < d[y].d && d[x].d+int(w) > 0 }
	relax := func(x, y *graph.Vertex[T], w EdgeWeight) { d[y].SetDist(int(w) + d[x].d) }
	es := g.Edges()
	for range vs {
		for _, e := range es {
			w := e.P.(EdgeWeight)
			switch {
			case canRelax(e.X, e.Y, w):
				relax(e.X, e.Y, w)
			case canRelax(e.Y, e.X, w):
				relax(e.Y, e.X, w)
			}
		}
	}
	return d
}
