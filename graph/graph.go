package graph

import (
	"bufio"
	"errors"
	"io"
	"log"
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

func NewFromCSV(graphType int, r io.Reader) Graph[string] {
	g := New[string](graphType)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		v1 := &Vertex[string]{E: f[0]}
		if !g.ContainsVertex(v1) {
			g.AddVertex(v1)
		}
		v2 := &Vertex[string]{E: f[1]}
		if !g.ContainsVertex(v2) {
			g.AddVertex(v2)
		}
		e := &Edge[string]{X: v1, Y: v2}
		if !g.ContainsEdge(e) {
			if len(f) < 3 {
				g.AddEdge(e)
				continue
			}
			w, err := strconv.Atoi(f[2])
			if err != nil {
				log.Panic(err)
			}
			e.P = EdgeProperty{W: w}
			g.AddEdge(e)
		}
	}
	return g
}

func getVertex[T comparable](g Graph[T], v *Vertex[T]) (int, *Vertex[T], error) {
	for i, u := range g.Vertices() {
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
