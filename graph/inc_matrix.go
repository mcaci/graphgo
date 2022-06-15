package graph

import (
	"fmt"
	"io"
)

type IncidenceMatrix[T comparable] struct {
	mat   [][]bool
	nodes []*Vertex[T]
	a     ArcsList[T]
}

func NewIncidenceMatrix(r io.Reader) *IncidenceMatrix[string] {
	edges := NewArcsList(r)
	nMap := make(map[string]struct{})
	for _, e := range edges.e {
		nMap[e.X.E] = struct{}{}
		nMap[e.Y.E] = struct{}{}
	}
	var mat [][]bool
	var nodes []*Vertex[string]
	for n := range nMap {
		mat = append(mat, make([]bool, len(edges.e)))
		nodes = append(nodes, &Vertex[string]{E: n})
	}
	for eid, e := range edges.e {
		var i, j int
		for k := range nodes {
			if nodes[k].E != e.X.E {
				continue
			}
			i = k
			break
		}
		for k := range nodes {
			if nodes[k].E != e.Y.E {
				continue
			}
			j = k
			break
		}
		mat[i][eid] = true
		mat[j][eid] = true
	}
	return &IncidenceMatrix[string]{
		mat:   mat,
		nodes: nodes,
		a:     *edges,
	}
}

func (g *IncidenceMatrix[T]) AreAdjacent(a, b *Vertex[T]) bool {
	return g.a.AreAdjacent(a, b)
}

func (g *IncidenceMatrix[T]) Degree(n *Vertex[T]) int {
	var i int
	for index, node := range g.nodes {
		if node.E != n.E {
			continue
		}
		i = index
	}
	var d int
	for j := range g.mat[i] {
		if !g.mat[i][j] {
			continue
		}
		d++
	}
	return d
}

func (g *IncidenceMatrix[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	var i int
	for idx, node := range g.nodes {
		if node.E != n.E {
			continue
		}
		i = idx
		break
	}
	var nodes []*Vertex[T]
	for j := range g.mat[i] {
		if !g.mat[i][j] {
			continue
		}
		arc := g.a.e[j]
		switch n.E {
		case arc.X.E:
			nodes = append(nodes, arc.Y)
		case arc.Y.E:
			nodes = append(nodes, arc.X)
		}
	}
	return nodes
}

func (g IncidenceMatrix[T]) String() string {
	return fmt.Sprintf("(%v, %v)", g.mat, g.nodes)
}
