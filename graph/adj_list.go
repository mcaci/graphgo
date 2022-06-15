package graph

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type AdjacencyList[T comparable] struct {
	v Vertex[T]
	l []*Vertex[T]
}

func (g AdjacencyList[T]) String() string {
	var s []string
	for _, a := range g.l {
		s = append(s, a.String())
	}
	return fmt.Sprintf("(v:%v, l:%v)", g.v, strings.Join(s, ","))
}

type AdjacencyLists[T comparable] []AdjacencyList[T]

func NewAdjacencyList(r io.Reader) *AdjacencyLists[string] {
	var g AdjacencyLists[string]
	s := bufio.NewScanner(r)
	lMap := make(map[string]*AdjacencyList[string])
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		if _, ok := lMap[f[0]]; !ok {
			lMap[f[0]] = &AdjacencyList[string]{v: Vertex[string]{E: f[0]}}
		}
		if _, ok := lMap[f[1]]; !ok {
			lMap[f[1]] = &AdjacencyList[string]{v: Vertex[string]{E: f[1]}}
		}
		lMap[f[0]].l = append(lMap[f[0]].l, &lMap[f[1]].v)
		lMap[f[1]].l = append(lMap[f[1]].l, &lMap[f[0]].v)
	}
	for _, v := range lMap {
		g = append(g, *v)
	}
	return &g
}

func (g *AdjacencyLists[T]) AddVertex(v *Vertex[T]) {
	*g = append(*g, AdjacencyList[T]{v: *v})
}

func (g *AdjacencyLists[T]) RemoveVertex(v *Vertex[T]) {
	id := -1
	for i, gv := range *g {
		if gv.v.E != v.E {
			continue
		}
		id = i
	}
	if id < 0 {
		return
	}
	*g = append((*g)[:id], (*g)[id+1:]...)
}

func (g AdjacencyLists[T]) ContainsVertex(v *Vertex[T]) bool {
	for _, gv := range g {
		if gv.v.E != v.E {
			continue
		}
		return true
	}
	return false
}

func (g *AdjacencyLists[T]) AddEdge(e *Edge[T]) {
	add := func(x, y *Vertex[T]) {
		if !g.ContainsVertex(x) {
			g.AddVertex(x)
			return
		}
		for i, gv := range *g {
			if gv.v.E != x.E {
				continue
			}
			(*g)[i].l = append((*g)[i].l, y)
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
		for i, gv := range *g {
			if gv.v.E != x.E {
				continue
			}
			for j, lv := range gv.l {
				if lv.E != y.E {
					continue
				}
				(*g)[i].l = append((*g)[i].l[:j], (*g)[i].l[j+1:]...)
				return
			}
		}
	}
	rm(e.X, e.Y)
	rm(e.Y, e.X)
}

func (g AdjacencyLists[T]) ContainsEdge(e *Edge[T]) bool {
	has := func(x, y *Vertex[T]) bool {
		if !g.ContainsVertex(x) {
			return false
		}
		for _, gv := range g {
			if gv.v.E != x.E {
				continue
			}
			for _, lv := range gv.l {
				if lv.E != y.E {
					continue
				}
				return true
			}
			return false
		}
		return false
	}
	return has(e.X, e.Y) //&& has(e.Y, e.X)
}

func (g *AdjacencyLists[T]) AreAdjacent(a, b *Vertex[T]) bool {
	for _, l := range *g {
		if l.v.E == a.E {
			for j := range l.l {
				if l.l[j].E == b.E {
					return true
				}
			}
		}
		if l.v.E == b.E {
			for j := range l.l {
				if l.l[j].E == a.E {
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
		return len(node.l)
	}
	return 0
}

func (g *AdjacencyLists[T]) AdjacentNodes(n *Vertex[T]) []*Vertex[T] {
	for _, list := range *g {
		if list.v.E != n.E {
			continue
		}
		return list.l
	}
	return nil
}
