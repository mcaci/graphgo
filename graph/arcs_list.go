package graph

import (
	"fmt"
	"strings"
)

type ArcsList[T comparable] struct {
	v        []*Vertex[T]
	e        []*Edge[T]
	directed bool
}

func (g *ArcsList[T]) AddVertex(v *Vertex[T]) {
	if g.ContainsVertex(v) {
		return
	}
	g.v = append(g.v, v)
}

func (g *ArcsList[T]) RemoveVertex(v *Vertex[T]) {
	i := indexVertex(g.Vertices(), v)
	if i < 0 {
		return
	}
	g.v = append(g.v[:i], g.v[i+1:]...)
}

func (g *ArcsList[T]) ContainsVertex(v *Vertex[T]) bool {
	return indexVertex(g.Vertices(), v) >= 0
}

func (g *ArcsList[T]) AddEdge(e *Edge[T]) {
	if g.ContainsEdge(e) {
		return
	}
	x, y, p := e.X, e.Y, e.P
	vs := g.Vertices()
	iX := indexVertex[T](vs, x)
	if iX >= 0 {
		x = vs[iX]
	}
	iY := indexVertex[T](vs, y)
	if iY >= 0 {
		y = vs[iY]
	}
	g.e = append(g.e, &Edge[T]{X: x, Y: y, P: p})
	if g.directed {
		return
	}
	g.e = append(g.e, &Edge[T]{X: y, Y: x, P: p})
}

func (g *ArcsList[T]) RemoveEdge(e *Edge[T]) {
	i := indexEdge[T](g, e)
	if i < 0 {
		return
	}
	g.e = append(g.e[:i], g.e[i+1:]...)

	if g.directed {
		return
	}

	j := indexEdge[T](g, &Edge[T]{X: e.Y, Y: e.X})
	if j < 0 {
		return
	}
	g.e = append(g.e[:j], g.e[j+1:]...)
}
func (g *ArcsList[T]) ContainsEdge(e *Edge[T]) bool {
	return indexEdge[T](g, e) >= 0
}

func (g *ArcsList[T]) AreAdjacent(a, b *Vertex[T]) bool {
	return g.ContainsEdge(&Edge[T]{X: a, Y: b})
}

func (g *ArcsList[T]) Degree(v *Vertex[T]) int {
	var d int
	for _, e := range g.e {
		if e.X.E != v.E {
			continue
		}
		d++
	}
	return d
}

func (g *ArcsList[T]) AdjacentNodes(v *Vertex[T]) []*Vertex[T] {
	var nodes []*Vertex[T]
	for _, e := range g.e {
		switch v.E {
		case e.X.E:
			nodes = append(nodes, e.Y)
		case e.Y.E:
			nodes = append(nodes, e.X)
		}
	}
	return nodes
}

func (g *ArcsList[T]) IsDirected() bool       { return g.directed }
func (g *ArcsList[T]) Vertices() []*Vertex[T] { return g.v }
func (g *ArcsList[T]) Edges() []*Edge[T]      { return g.e }

func (g ArcsList[T]) String() string {
	var e []string
	for i := range g.e {
		e = append(e, g.e[i].String())
	}
	var v []string
	for i := range g.v {
		v = append(v, g.v[i].String())
	}
	return fmt.Sprintf("(v:%v,e:%v)", strings.Join(v, ","), strings.Join(e, ","))
}
