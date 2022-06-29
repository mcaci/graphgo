package graph

import (
	"fmt"
)

type AdjacencyMatrix[T comparable] struct {
	m [][]*EdgeProperty
	v []*Vertex[T]
}

func (g *AdjacencyMatrix[T]) AddVertex(v *Vertex[T]) {
	g.v = append(g.v, v)
	g.m = append(g.m, make([]*EdgeProperty, len(g.m)))
	for i := range g.m {
		g.m[i] = append(g.m[i], nil)
	}
}

func (g *AdjacencyMatrix[T]) RemoveVertex(v *Vertex[T]) {
	iRm, _, err := getVertex[T](g, v)
	if err != nil {
		return
	}
	g.v = append(g.v[:iRm], g.v[iRm+1:]...)
	m := make([][]*EdgeProperty, len(g.m)-1)
	reduce := func(i, iRm int) int {
		if i > iRm {
			return i - 1
		}
		return i
	}
	for i := range m {
		m[i] = append(m[i], make([]*EdgeProperty, len(g.m)-1)...)
	}
	for i := range g.m {
		if i == iRm {
			continue
		}
		mi := reduce(i, iRm)
		for j := range g.m[i] {
			if j == iRm {
				continue
			}
			mj := reduce(j, iRm)
			m[mi][mj] = g.m[i][j]
		}
	}
	g.m = m
}

func (g *AdjacencyMatrix[T]) ContainsVertex(v *Vertex[T]) bool {
	_, _, err := getVertex[T](g, v)
	return err == nil
}

func (g *AdjacencyMatrix[T]) AddEdge(e *Edge[T]) {
	i, _, err := getVertex[T](g, e.X)
	if err != nil {
		return
	}
	j, _, err := getVertex[T](g, e.Y)
	if err != nil {
		return
	}
	g.m[i][j] = &e.P
	g.m[j][i] = &e.P
}

func (g *AdjacencyMatrix[T]) RemoveEdge(e *Edge[T]) {
	i, _, err := getVertex[T](g, e.X)
	if err != nil {
		return
	}
	j, _, err := getVertex[T](g, e.Y)
	if err != nil {
		return
	}
	g.m[i][j] = nil
	g.m[j][i] = nil
}

func (g *AdjacencyMatrix[T]) ContainsEdge(e *Edge[T]) bool {
	i, _, err := getVertex[T](g, e.X)
	if err != nil {
		return false
	}
	j, _, err := getVertex[T](g, e.Y)
	if err != nil {
		return false
	}
	return g.m[i][j] != nil && g.m[j][i] != nil
}

func (g *AdjacencyMatrix[T]) AreAdjacent(a, b *Vertex[T]) bool {
	_, e, err := getEdge[T](g, &Edge[T]{X: a, Y: b})
	if err != nil {
		return false
	}
	return e != nil
}

func (g *AdjacencyMatrix[T]) Degree(v *Vertex[T]) int {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return 0
	}
	var d int
	for j := range g.m[i] {
		if g.m[i][j] == nil {
			continue
		}
		d++
	}
	return d
}

func (g *AdjacencyMatrix[T]) AdjacentNodes(v *Vertex[T]) []*Vertex[T] {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return nil
	}
	var nodes []*Vertex[T]
	for j := range g.m[i] {
		if g.m[i][j] == nil {
			continue
		}
		nodes = append(nodes, g.v[j])
	}
	return nodes
}

func (g *AdjacencyMatrix[T]) Vertices() []*Vertex[T] { return g.v }
func (g *AdjacencyMatrix[T]) Edges() []*Edge[T] {
	var edges []*Edge[T]
	for i := range g.m {
		for j := range g.m[i] {
			if g.m[i][j] == nil {
				continue
			}
			edges = append(edges, &Edge[T]{X: g.v[i], Y: g.v[j], P: *g.m[i][j]})
		}
	}
	return edges
}

func (g AdjacencyMatrix[T]) String() string {
	return fmt.Sprintf("(m:%v, v:%v)", g.m, g.v)
}
