package path

import "github.com/mcaci/graphgo/graph"

func Shortest[T comparable](g graph.Graph[T], d map[*graph.Vertex[T]]*Distance[T], x, y *graph.Vertex[T]) []*graph.Vertex[T] {
	if len(g.Vertices()) < 2 {
		return nil
	}
	path := []*graph.Vertex[T]{y}
	v := y
	isShortestDist := func(u, v *graph.Vertex[T], w Weighter) bool { return d[u].d+w.Weight() == d[v].d }
	es := g.Edges()
	for v != x {
	edges:
		for _, u := range g.AdjacentNodes(v) {
			for _, edge := range es {
				switch {
				case edge.X == u && edge.Y == v,
					edge.X == v && edge.Y == u:

					if isShortestDist(u, v, edge.P.(Weighter)) {
						path = append([]*graph.Vertex[T]{u}, path...)
						v = u
						break edges
					}

				}
			}
		}
	}
	return path
}
