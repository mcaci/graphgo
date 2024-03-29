package graph

import (
	"bufio"
	"errors"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"
)

type Graph[T comparable] interface {
	AreAdjacent(a, b *Vertex[T]) bool
	Degree(v *Vertex[T]) int
	AdjacentNodes(v *Vertex[T]) []*Vertex[T]

	Vertices() []*Vertex[T]
	Edges() []*Edge[T]
	AddVertex(v *Vertex[T])
	RemoveVertex(v *Vertex[T])
	ContainsVertex(v *Vertex[T]) bool
	AddEdge(e *Edge[T])
	RemoveEdge(e *Edge[T])
	ContainsEdge(e *Edge[T]) bool
}

const (
	ArcsListType = iota
	AdjacencyListType
	AdjacencyMatrixType
	// IncidenceListType
	// IncidenceMatrixType
)

func New[T comparable](graphType int) Graph[T] {
	switch graphType {
	case ArcsListType:
		return &ArcsList[T]{}
	case AdjacencyListType:
		return &AdjacencyLists[T]{}
	case AdjacencyMatrixType:
		return &AdjacencyMatrix[T]{}
	default:
		return &ArcsList[T]{}
	}
}

func Fill[T comparable](vs []*Vertex[T], es []*Edge[T], into Graph[T]) {
	for _, v := range vs {
		if into.ContainsVertex(v) {
			continue
		}
		into.AddVertex(v)
	}
	for _, e := range es {
		if into.ContainsEdge(e) {
			continue
		}
		_, v, _ := getVertexFromList(vs, e.X)
		_, u, _ := getVertexFromList(vs, e.Y)
		into.AddEdge(&Edge[T]{X: v, Y: u, P: e.P})
	}
}

func FromCSV(r io.Reader) ([]*Vertex[string], []*Edge[string], error) {
	var vs []*Vertex[string]
	var es []*Edge[string]
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		v := &Vertex[string]{E: f[0]}
		u := &Vertex[string]{E: f[1]}
		vs = append(vs, v, u)
		e := &Edge[string]{X: v, Y: u}
		if len(f) < 3 {
			es = append(es, e)
			continue
		}
		w, err := strconv.Atoi(f[2])
		if err != nil {
			return nil, nil, err
		}
		e.P.W = w
		es = append(es, e)
	}
	return vs, es, nil
}

func FromSpaced(r io.Reader) ([]*Vertex[string], []*Edge[string], error) {
	var vs []*Vertex[string]
	var es []*Edge[string]
	s := bufio.NewScanner(r)
	for s.Scan() {
		var vv, uu, cc string
		var ww int
		fmt.Sscanf(s.Text(), "%s %s %d %s", &vv, &uu, &ww, &cc)
		v := &Vertex[string]{E: vv}
		u := &Vertex[string]{E: uu}
		vs = append(vs, v, u)
		// replace C with cc
		e := &Edge[string]{X: v, Y: u, P: EdgeProperty{W: ww, C: color.Black}}
		es = append(es, e)
	}
	return vs, es, nil
}

func getVertex[T comparable](g interface{ Vertices() []*Vertex[T] }, v *Vertex[T]) (int, *Vertex[T], error) {
	return getVertexFromList(g.Vertices(), v)
}

func getVertexFromList[T comparable](vs []*Vertex[T], v *Vertex[T]) (int, *Vertex[T], error) {
	for i, u := range vs {
		if u.E != v.E {
			continue
		}
		return i, u, nil
	}
	return 0, nil, errors.New("Vertex not found")
}

func getEdge[T comparable](g Graph[T], e *Edge[T]) (int, *Edge[T], error) {
	for i, edge := range g.Edges() {
		switch {
		case edge.X.E == e.X.E && edge.Y.E == e.Y.E,
			edge.X.E == e.Y.E && edge.Y.E == e.X.E:
			return i, edge, nil
		}
	}
	return 0, nil, errors.New("Edge not found")
}
