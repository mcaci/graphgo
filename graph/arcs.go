package graph

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type ArcsList[T comparable] struct {
	v []*Vertex[T]
	e []*Edge[T]
}

func NewArcsList(r io.Reader) *ArcsList[string] {
	var g ArcsList[string]
	s := bufio.NewScanner(r)
	vMap := make(map[string]*Vertex[string])
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		if _, ok := vMap[f[0]]; !ok {
			vMap[f[0]] = &Vertex[string]{E: f[0]}
			g.AddVertex(vMap[f[0]])
		}
		if _, ok := vMap[f[1]]; !ok {
			vMap[f[1]] = &Vertex[string]{E: f[1]}
			g.AddVertex(vMap[f[1]])
		}
		d, err := strconv.Atoi(f[2])
		if err != nil {
			log.Fatal(err)
		}
		g.AddEdge(&Edge[string]{X: vMap[f[0]], Y: vMap[f[1]], P: EdgeProperty{W: d}})
	}
	return &g
}

func (g *ArcsList[T]) AddVertex(v *Vertex[T]) {
	g.v = append(g.v, v)
}

func (g *ArcsList[T]) RemoveVertex(v *Vertex[T]) {
	id := -1
	for i, gv := range g.v {
		if gv.E != v.E {
			continue
		}
		id = i
	}
	if id < 0 {
		return
	}
	g.v = append(g.v[:id], g.v[id+1:]...)
}

func (g ArcsList[T]) ContainsVertex(v *Vertex[T]) bool {
	for _, gv := range g.v {
		if gv.E != v.E {
			continue
		}
		return true
	}
	return false
}

func (g *ArcsList[T]) AddEdge(e *Edge[T]) {
	g.e = append(g.e, e)
}

func (g *ArcsList[T]) RemoveEdge(e *Edge[T]) {
	id := -1
	for i, ge := range g.e {
		if ge.X.E != e.X.E && ge.X.E != e.Y.E {
			continue
		}
		if ge.Y.E != e.X.E && ge.Y.E != e.Y.E {
			continue
		}
		id = i
		break
	}
	if id < 0 {
		return
	}
	g.e = append(g.e[:id], g.e[id+1:]...)
}

func (g ArcsList[T]) ContainsEdge(e *Edge[T]) bool {
	for _, ge := range g.e {
		if ge.X.E != e.X.E && ge.X.E != e.Y.E {
			continue
		}
		if ge.Y.E != e.X.E && ge.Y.E != e.Y.E {
			continue
		}
		return true
	}
	return false
}

func (g ArcsList[T]) AreAdjacent(a, b *Vertex[T]) bool {
	for i := range g.e {
		e := g.e[i]
		if e.X.E == a.E && e.Y.E == b.E {
			return true
		}
		if e.X.E == b.E && e.Y.E == a.E {
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
