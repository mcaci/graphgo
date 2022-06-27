package graph

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type AdjacencyLists[T comparable] []AdjacencyList[T]

func NewAdjacencyList(r io.Reader) *AdjacencyLists[string] {
	var g AdjacencyLists[string]
	s := bufio.NewScanner(r)
	lMap := make(map[string]*AdjacencyList[string])
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		if _, ok := lMap[f[0]]; !ok {
			lMap[f[0]] = &AdjacencyList[string]{v: &Vertex[string]{E: f[0]}}
		}
		if _, ok := lMap[f[1]]; !ok {
			lMap[f[1]] = &AdjacencyList[string]{v: &Vertex[string]{E: f[1]}}
		}
		n0 := &Neighbour[string]{v: lMap[f[1]].v}
		if len(f) >= 3 {
			w, err := strconv.Atoi(f[2])
			if err != nil {
				log.Panic(err)
			}
			n0.p = EdgeProperty{W: w}
		}
		lMap[f[0]].n = append(lMap[f[0]].n, n0)
		n1 := &Neighbour[string]{v: lMap[f[0]].v}
		if len(f) >= 3 {
			w, err := strconv.Atoi(f[2])
			if err != nil {
				log.Panic(err)
			}
			n0.p = EdgeProperty{W: w}
		}
		lMap[f[1]].n = append(lMap[f[1]].n, n1)
	}
	for _, v := range lMap {
		g = append(g, *v)
	}
	return &g
}

func (g *AdjacencyLists[T]) AddVertex(v *Vertex[T]) {
	*g = append(*g, AdjacencyList[T]{v: v})
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
		for i, gv := range *g {
			if gv.v.E != x.E {
				continue
			}
			for j, lv := range gv.n {
				if lv.v.E != y.E {
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
	has := func(x, y *Vertex[T]) bool {
		if !g.ContainsVertex(x) {
			return false
		}
		for _, gv := range g {
			if gv.v.E != x.E {
				continue
			}
			for _, lv := range gv.n {
				if lv.v.E != y.E {
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
