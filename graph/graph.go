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
	Degree(n *Vertex[T]) int
	AdjacentNodes(n *Vertex[T]) []*Vertex[T]

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

func NewWithReader(graphType int, r io.Reader) Graph[string] {
	g := New[string](graphType)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		_, v, err := getVertex(g, &Vertex[string]{E: f[0]})
		if err != nil {
			v = &Vertex[string]{E: f[0]}
			g.AddVertex(v)
		}
		_, u, err := getVertex(g, &Vertex[string]{E: f[1]})
		if err != nil {
			u = &Vertex[string]{E: f[1]}
			g.AddVertex(u)
		}
		var w int
		if len(f) >= 3 {
			var err error
			w, err = strconv.Atoi(f[2])
			if err != nil {
				log.Panic(err)
			}
		}
		g.AddEdge(&Edge[string]{X: v, Y: u, P: EdgeProperty{W: w}})
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
