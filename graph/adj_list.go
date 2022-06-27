package graph

import (
	"fmt"
	"strings"
)

type AdjacencyLists[T comparable] []AdjacencyList[T]

func (g *AdjacencyLists[T]) AddVertex(v *Vertex[T]) {
	if g.ContainsVertex(v) {
		return
	}
	*g = append(*g, AdjacencyList[T]{v: v})
}

func (g *AdjacencyLists[T]) RemoveVertex(v *Vertex[T]) {
	i, _, err := getVertex[T](g, v)
	if err != nil {
		return
	}
	*g = append((*g)[:i], (*g)[i+1:]...)
}

func (g AdjacencyLists[T]) ContainsVertex(v *Vertex[T]) bool {
	_, _, err := getVertex[T](&g, v)
	return err == nil
}

func (g *AdjacencyLists[T]) AddEdge(e *Edge[T]) {
	if g.ContainsEdge(e) {
		return
	}
	add := func(x, y *Vertex[T]) {
		for i, gv := range *g {
			if gv.v.E != x.E {
				continue
			}
			(*g)[i].n = append((*g)[i].n, &Neighbour[T]{v: y, p: e.P})
			return
		}
	}
	add(e.X, e.Y)
	add(e.Y, e.X)
}

func (g *AdjacencyLists[T]) RemoveEdge(e *Edge[T]) {
	rm := func(x, y *Vertex[T]) {
		if !g.ContainsVertex(x) {
			return
		}
		for i, l := range *g {
			if l.v.E != x.E {
				continue
			}
			for j, nv := range l.n {
				if nv.v.E != y.E {
					continue
				}
				(*g)[i].n = append((*g)[i].n[:j], (*g)[i].n[j+1:]...)
				return
			}
		}
	}
	rm(e.X, e.Y)
	rm(e.Y, e.X)
}

func (g AdjacencyLists[T]) ContainsEdge(e *Edge[T]) bool {
	_, _, err := getEdge[T](&g, e)
	// _, _, othererr := getEdge[T](&g, e.Y.E, e.X.E)
	return err == nil // && othererr != nil
}

func (g *AdjacencyLists[T]) AreAdjacent(a, b *Vertex[T]) bool {
	for _, l := range *g {
		if l.v.E == a.E {
			for j := range l.n {
				if l.n[j].v.E == b.E {
					return true
				}
			}
		}
		if l.v.E == b.E {
			for j := range l.n {
				if l.n[j].v.E == a.E {
					return true
				}
			}
		}
	}
	return false
}

func (g *AdjacencyLists[T]) Degree(n *Vertex[T]) int {
	for _, node := range *g {
		if node.v.E != n.E {
			continue
		}
		return len(node.n)
	}
	return 0
}

func (g *AdjacencyLists[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	var nghrs []*Neighbour[T]
	for _, list := range *g {
		if list.v.E != n.E {
			continue
		}
		nghrs = list.n
		break
	}
	var l []*Vertex[T]
	for _, n := range nghrs {
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
