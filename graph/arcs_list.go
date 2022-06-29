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

func (g *ArcsList[T]) ContainsVertex(v *Vertex[T]) bool {
	_, _, err := getVertex[T](g, v)
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

func (g *ArcsList[T]) ContainsEdge(e *Edge[T]) bool {
	_, _, err := getEdge[T](g, e)
	return err == nil
}

func (g *ArcsList[T]) AreAdjacent(a, b *Vertex[T]) bool {
	_, e, err := getEdge[T](g, &Edge[T]{X: a, Y: b})
	if err != nil {
		return false
	}
	return e != nil
}

func (g *ArcsList[T]) Degree(v *Vertex[T]) int {
	var d int
	for _, e := range g.e {
		switch v.E {
		case e.X.E, e.Y.E:
			d++
		}
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
