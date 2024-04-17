package graph

import (
	"fmt"
	"strings"
)

type AdjacencyLists[T comparable] struct {
	l        []AdjacencyList[T]
	directed bool
}

func (g *AdjacencyLists[T]) AddVertex(v *Vertex[T]) {
	if g.ContainsVertex(v) {
		return
	}
	g.l = append(g.l, AdjacencyList[T]{v: v})
}

func (g *AdjacencyLists[T]) RemoveVertex(v *Vertex[T]) {
	if !g.ContainsVertex(v) {
		return
	}
	i := indexVertex(g.Vertices(), v)
	g.l = append(g.l[:i], g.l[i+1:]...)
}

func (g *AdjacencyLists[T]) ContainsVertex(v *Vertex[T]) bool {
	return indexVertex(g.Vertices(), v) >= 0
}

func (g *AdjacencyLists[T]) AddEdge(e *Edge[T]) {
	add := func(x, y *Vertex[T]) {
		i := indexVertex[T](g.Vertices(), x)
		if i < 0 {
			return
		}
		g.l[i].n = append(g.l[i].n, &Neighbour[T]{v: y, p: e.P})
	}
	x, y := e.X, e.Y
	vs := g.Vertices()
	iX := indexVertex[T](vs, x)
	if iX >= 0 {
		x = vs[iX]
	}
	iY := indexVertex[T](vs, y)
	if iY >= 0 {
		y = vs[iY]
	}
	add(x, y)
	if g.directed {
		return
	}
	add(y, x)
}

func (g *AdjacencyLists[T]) RemoveEdge(e *Edge[T]) {
	rm := func(x, y *Vertex[T]) {
		i := indexVertex[T](g.Vertices(), x)
		if i < 0 {
			return
		}
		for j, n := range g.l[i].n {
			if n.v.E != y.E {
				continue
			}
			g.l[i].n = append(g.l[i].n[:j], g.l[i].n[j+1:]...)
			return
		}
	}
	rm(e.X, e.Y)
	if g.directed {
		return
	}
	rm(e.Y, e.X)
}

func (g *AdjacencyLists[T]) ContainsEdge(e *Edge[T]) bool {
	return indexEdge[T](g, e) >= 0
}

func (g *AdjacencyLists[T]) AreAdjacent(a, b *Vertex[T]) bool {
	return g.ContainsEdge(&Edge[T]{X: a, Y: b})
}

func (g *AdjacencyLists[T]) Degree(v *Vertex[T]) int {
	i := indexVertex[T](g.Vertices(), v)
	if i < 0 {
		return 0
	}
	return len(g.l[i].n)
}

func (g *AdjacencyLists[T]) AdjacentNodes(v *Vertex[T]) []*Vertex[T] {
	i := indexVertex[T](g.Vertices(), v)
	if i < 0 {
		return nil
	}
	var l []*Vertex[T]
	for _, n := range g.l[i].n {
		l = append(l, n.v)
	}
	return l
}

func (g *AdjacencyLists[T]) IsDirected() bool { return g.directed }

func (g *AdjacencyLists[T]) Vertices() []*Vertex[T] {
	var vertices []*Vertex[T]
	for _, v := range g.l {
		vertices = append(vertices, v.v)
	}
	return vertices
}

func (g *AdjacencyLists[T]) Edges() []*Edge[T] {
	var edges []*Edge[T]
	for _, v := range g.l {
		for _, u := range v.n {
			edges = append(edges, &Edge[T]{X: v.v, Y: u.v, P: u.p})
		}
	}
	return edges
}

type AdjacencyList[T comparable] struct {
	v *Vertex[T]
	n []*Neighbour[T]
}

func (g AdjacencyList[T]) String() string {
	var s []string
	for _, v := range g.n {
		s = append(s, v.String())
	}
	return fmt.Sprintf("(v:%v, l:%v)", g.v, strings.Join(s, ","))
}

type Neighbour[T comparable] struct {
	v *Vertex[T]
	p EdgeProperty
}

func (n Neighbour[T]) String() string {
	return fmt.Sprintf("(v:%v, p:%v)", n.v, n.p)
}
