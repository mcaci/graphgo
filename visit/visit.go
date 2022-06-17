package visit

import (
	"github.com/mcaci/graphgo/graph"
)

func Generic[T comparable](g graph.Graph[T], s *graph.Vertex[T]) *Tree[T] {
	if !g.ContainsVertex(s) {
		return nil
	}
	t := &Tree[T]{element: &s.E}
	vert := g.Vertices()
	for i := range vert {
		if vert[i].E != s.E {
			continue
		}
		vert[i].SetVisited()
		break
	}
	var f []graph.Vertex[T]
	f = append(f, *s)
	for len(f) != 0 {
		var u *graph.Vertex[T]
		u, f = &f[0], f[1:]
		for _, v := range g.AdjacentNodes(u) {
			if v.IsVisited() {
				continue
			}
			v.SetVisited()
			f = append(f, *v)
			tree := t.Find(&u.E)
			if tree != nil {
				tree.children = append(tree.children, &Tree[T]{element: &v.E})
			}
		}
	}
	return t
}
