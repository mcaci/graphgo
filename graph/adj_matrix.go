package graph

import (
	"fmt"
	"io"
)

type AdjacencyMatrix[T comparable] struct {
	m [][]bool
	v []*Vertex[T]
}

func NewAdjacencyMatrix(r io.Reader) *AdjacencyMatrix[string] {
	edges := NewArcsList(r)
	nMap := make(map[string]struct{})
	for _, e := range edges.e {
		nMap[e.X.E] = struct{}{}
		nMap[e.Y.E] = struct{}{}
	}
	var mat [][]bool
	var nodes []*Vertex[string]
	for n := range nMap {
		mat = append(mat, make([]bool, len(nMap)))
		nodes = append(nodes, &Vertex[string]{E: n})
	}
	for _, e := range edges.e {
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
		mat[i][j] = true
		mat[j][i] = true
	}
	return &AdjacencyMatrix[string]{
		m: mat,
		v: nodes,
	}
}

func (g *AdjacencyMatrix[T]) AddVertex(v *Vertex[T]) {
	g.v = append(g.v, v)
	g.m = append(g.m, make([]bool, len(g.m)))
	for i := range g.m {
		g.m[i] = append(g.m[i], false)
	}
}

func (g *AdjacencyMatrix[T]) RemoveVertex(v *Vertex[T]) {
	if !g.ContainsVertex(v) {
		return
	}
	var id int
	for i, gv := range g.v {
		if gv.E != v.E {
			continue
		}
		id = i
		break
	}
	g.v = append(g.v[:id], g.v[id+1:]...)
	mat := make([][]bool, len(g.m)-1)
	for i := range mat {
		mat[i] = append(mat[i], make([]bool, len(g.m)-1)...)
	}
	for i := range g.m {
		if i == id {
			continue
		}
		iMat := i
		if i > id {
			iMat = i - 1
		}
		for j := range g.m[i] {
			if j == id {
				continue
			}
			jMat := j
			if j > id {
				jMat = j - 1
			}
			mat[iMat][jMat] = g.m[i][j]
		}
	}
	g.m = mat
}

func (g AdjacencyMatrix[T]) ContainsVertex(v *Vertex[T]) bool {
	for _, gv := range g.v {
		if gv.E != v.E {
			continue
		}
		return true
	}
	return false
}

func (g *AdjacencyMatrix[T]) AddEdge(e *Edge[T]) {
	if !g.ContainsVertex(e.X) {
		g.AddVertex(e.X)
	}
	if !g.ContainsVertex(e.Y) {
		g.AddVertex(e.Y)
	}
	x, y := e.X, e.Y
	var ix, iy int
	for i, gv := range g.v {
		switch gv.E {
		case x.E:
			ix = i
		case y.E:
			iy = i
		}
	}
	g.m[ix][iy] = true
}

func (g *AdjacencyMatrix[T]) RemoveEdge(e *Edge[T]) {
	if !g.ContainsEdge(e) {
		return
	}
	x, y := e.X, e.Y
	var ix, iy int
	for i, gv := range g.v {
		switch gv.E {
		case x.E:
			ix = i
		case y.E:
			iy = i
		}
	}
	g.m[ix][iy] = false
}

func (g AdjacencyMatrix[T]) ContainsEdge(e *Edge[T]) bool {
	if !g.ContainsVertex(e.X) || !g.ContainsVertex(e.Y) {
		return false
	}
	x, y := e.X, e.Y
	var ix, iy int
	for i, gv := range g.v {
		switch gv.E {
		case x.E:
			ix = i
		case y.E:
			iy = i
		}
	}
	return g.m[ix][iy]
}

func (g *AdjacencyMatrix[T]) AreAdjacent(a, b *Vertex[T]) bool {
	var i, j int
	for k := range g.v {
		if g.v[k].E != a.E {
			continue
		}
		i = k
		break
	}
	for k := range g.v {
		if g.v[k].E != b.E {
			continue
		}
		j = k
		break
	}
	return g.m[i][j]
}

func (g *AdjacencyMatrix[T]) Degree(n *Vertex[T]) int {
	var i int
	for index, node := range g.v {
		if node.E != n.E {
			continue
		}
		i = index
	}
	var d int
	for j := range g.m[i] {
		if !g.m[i][j] {
			continue
		}
		d++
	}
	return d
}

func (g *AdjacencyMatrix[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	var idx int
	for i, node := range g.v {
		if node.E != n.E {
			continue
		}
		idx = i
		break
	}
	var nodes []*Vertex[T]
	for j, arc := range g.m[idx] {
		if arc {
			nodes = append(nodes, g.v[j])
		}
	}
	return nodes
}

func (g AdjacencyMatrix[T]) String() string {
	return fmt.Sprintf("(m:%v, v:%v)", g.m, g.v)
}
