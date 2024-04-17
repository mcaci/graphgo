package path

import "github.com/mcaci/graphgo/graph"

func Shortest[T comparable](g graph.Graph[T], d map[*graph.Vertex[T]]*Distance[T], x, y *graph.Vertex[T]) []*graph.Vertex[T] {
	if len(g.Vertices()) < 2 {
		return nil
	}
	path := []*graph.Vertex[T]{y}
	v := y
	isShortestDist := func(u, v *graph.Vertex[T], w Weighter) bool { return d[u].d+w.Weight() == d[v].d }
	isConnectingEdge := func(u, v *graph.Vertex[T], e *graph.Edge[T]) bool {
		return (e.X == u && e.Y == v) || (e.X == v && e.Y == u)
	}
	es := g.Edges()
	for v != x {
	neighbourSearch:
		for _, u := range g.AdjacentNodes(v) {
			for _, edge := range es {
				if !isConnectingEdge(u, v, edge) {
					continue
				}
				if !isShortestDist(u, v, edge.P.(Weighter)) {
					continue
				}
				path = append([]*graph.Vertex[T]{u}, path...)
				v = u
				break neighbourSearch
			}
		}
	}
	return path
}
