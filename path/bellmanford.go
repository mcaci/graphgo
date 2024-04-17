package path

import (
	"math"

	"github.com/mcaci/graphgo/graph"
)

type Weighter interface{ Weight() int }

func BellmanFordDist[T comparable](g graph.Graph[T], s *graph.Vertex[T]) map[*graph.Vertex[T]]*Distance[T] {
	d := make(map[*graph.Vertex[T]]*Distance[T])
	vs := g.Vertices()
	for i := range vs {
		switch v := vs[i]; v {
		case s:
			d[v] = &Distance[T]{v: s, u: v, d: 0}
		default:
			d[v] = &Distance[T]{v: s, u: v, d: math.MaxInt}
		}
	}
	canRelax := func(x, y *graph.Vertex[T], w Weighter) bool {
		return d[x].d+w.Weight() < d[y].d && d[x].d+w.Weight() > 0
	}
	relax := func(x, y *graph.Vertex[T], w Weighter) {
		d[y].setDistance(w.Weight() + d[x].d)
	}
	es := g.Edges()
	for range vs {
		for _, e := range es {
			w := e.P.(Weighter)
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
