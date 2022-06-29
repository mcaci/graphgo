package graph

import (
	"fmt"
	"strings"
)

type AdjacencyLists[T comparable] []AdjacencyList[T]

func (g *AdjacencyLists[T]) AddVertex(v *Vertex[T]) {
	*g = append(*g, AdjacencyList[T]{v: v})
}

func (g *AdjacencyLists[T]) RemoveVertex(v *Vertex[T]) {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return
	}
	*g = append((*g)[:i], (*g)[i+1:]...)
}

func (g *AdjacencyLists[T]) ContainsVertex(v *Vertex[T]) bool {
	_, _, err := getVertex[T](g, v)
	return err == nil
}

func (g *AdjacencyLists[T]) AddEdge(e *Edge[T]) {
	add := func(x, y *Vertex[T]) {
		i, _, err := getVertex[T](g, x)
		if err != nil {
			return
		}
		(*g)[i].n = append((*g)[i].n, &Neighbour[T]{v: y, p: e.P})
	}
	add(e.X, e.Y)
	add(e.Y, e.X)
}

func (g *AdjacencyLists[T]) RemoveEdge(e *Edge[T]) {
	rm := func(x, y *Vertex[T]) {
		i, _, err := getVertex[T](g, x)
		if err != nil {
			return
		}
		for j, n := range (*g)[i].n {
			if n.v.E != y.E {
				continue
			}
			(*g)[i].n = append((*g)[i].n[:j], (*g)[i].n[j+1:]...)
			return
		}
	}
	rm(e.X, e.Y)
	rm(e.Y, e.X)
}

func (g *AdjacencyLists[T]) ContainsEdge(e *Edge[T]) bool {
	_, _, err := getEdge[T](g, e)
	return err == nil
}

func (g *AdjacencyLists[T]) AreAdjacent(a, b *Vertex[T]) bool {
	_, e, err := getEdge[T](g, &Edge[T]{X: a, Y: b})
	if err != nil {
		return false
	}
	return e != nil
}

func (g *AdjacencyLists[T]) Degree(v *Vertex[T]) int {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return 0
	}
	return len((*g)[i].n)
}

func (g *AdjacencyLists[T]) AdjacentNodes(v *Vertex[T]) []*Vertex[T] {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return nil
	}
	var l []*Vertex[T]
	for _, n := range (*g)[i].n {
		l = append(l, n.v)
	}
	return l
}

func (g *AdjacencyLists[T]) Vertices() []*Vertex[T] {
	var vertices []*Vertex[T]
	for _, v := range *g {
		vertices = append(vertices, v.v)
	}
	return vertices
}

func (g *AdjacencyLists[T]) Edges() []*Edge[T] {
	var edges []*Edge[T]
	for _, v := range *g {
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
