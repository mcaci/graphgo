package graph

import (
	"fmt"
	"io"
)

type IncidenceList[T comparable] struct {
	n Vertex[T]
	l []int
}

func (g IncidenceList[T]) String() string {
	return fmt.Sprintf("(%v, %v)", g.n, g.l)
}

type IncidenceLists[T comparable] struct {
	l []IncidenceList[T]
	a ArcsList[T]
}

func NewIncidenceList(r io.Reader) *IncidenceLists[string] {
	var g IncidenceLists[string]
	g.a = *NewArcsList(r)
	nMap := make(map[string]IncidenceList[string])
	for i, a := range g.a.e {
		nx := nMap[a.X.E]
		ny := nMap[a.Y.E]
		nx.n = *a.X
		nx.l = append(nx.l, i)
		ny.n = *a.Y
		ny.l = append(ny.l, i)
		nMap[a.X.E] = nx
		nMap[a.Y.E] = ny
	}
	for _, v := range nMap {
		g.l = append(g.l, v)
	}
	return &g
}

func (g *IncidenceLists[T]) AreAdjacent(a, b *Vertex[T]) bool {
	var aList, bList IncidenceList[T]
	for _, l := range g.l {
		switch l.n.E {
		case a.E:
			aList = l
		case b.E:
			bList = l
		default:
			continue
		}
	}
	for _, i := range aList.l {
		if g.a.e[i].X.E == b.E || g.a.e[i].Y.E == b.E {
			return true
		}
	}
	for _, i := range bList.l {
		if g.a.e[i].X.E == a.E || g.a.e[i].Y.E == a.E {
			return true
		}
	}
	return false
}

func (g *IncidenceLists[T]) Degree(n *Vertex[T]) int {
	for _, node := range g.l {
		if node.n.E != n.E {
			continue
		}
		return len(node.l)
	}
	return 0
}

func (g *IncidenceLists[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	var l IncidenceList[T]
	for _, list := range g.l {
		if list.n.E != n.E {
			continue
		}
		l = list
		break
	}
	var nodes []*Vertex[T]
	for _, i := range l.l {
		arc := g.a.e[i]
		switch n.E {
		case arc.X.E:
			nodes = append(nodes, arc.Y)
		case arc.Y.E:
			nodes = append(nodes, arc.X)
		}
	}
	return nodes
}
