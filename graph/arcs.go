package graph

import (
	"fmt"
	"strings"
)

type ArcsList[T comparable] struct {
	v []*Vertex[T]
	e []*Edge[T]
}

func (g *ArcsList[T]) AddVertex(v *Vertex[T]) {
	g.v = append(g.v, v)
}

func (g *ArcsList[T]) RemoveVertex(v *Vertex[T]) {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return
	}
	g.v = append(g.v[:i], g.v[i+1:]...)
}

func (g ArcsList[T]) ContainsVertex(v *Vertex[T]) bool {
	_, _, err := getVertex[T](&g, v)
	return err == nil
}

func (g *ArcsList[T]) AddEdge(e *Edge[T]) {
	g.e = append(g.e, e)
}

func (g *ArcsList[T]) RemoveEdge(e *Edge[T]) {
	i, _, err := getEdge[T](g, e)
	if err != nil {
		return
	}
	g.e = append(g.e[:i], g.e[i+1:]...)
}

func (g ArcsList[T]) ContainsEdge(e *Edge[T]) bool {
	_, _, err := getEdge[T](&g, e)
	return err == nil
}

func (g ArcsList[T]) AreAdjacent(a, b *Vertex[T]) bool {
	for i := range g.e {
		e := g.e[i]
		switch {
		case e.X.E == a.E && e.Y.E == b.E,
			e.X.E == b.E && e.Y.E == a.E:
			return true
		}
	}
	return false
}

func (g ArcsList[T]) Degree(n *Vertex[T]) int {
	var d int
	for _, a := range g.e {
		switch n.E {
		case a.X.E, a.Y.E:
			d++
		}
	}
	return d
}

func (g ArcsList[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	var nodes []*Vertex[T]
	for _, a := range g.e {
		switch n.E {
		case a.X.E:
			nodes = append(nodes, a.Y)
		case a.Y.E:
			nodes = append(nodes, a.X)
		}
	}
	return nodes
}

func (g ArcsList[T]) Vertices() []*Vertex[T] { return g.v }
func (g ArcsList[T]) Edges() []*Edge[T]      { return g.e }

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
