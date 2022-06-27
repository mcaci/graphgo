package visit_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/visit"
	"github.com/mcaci/graphgo/visit/internal"
)

func TestGenericVisit(t *testing.T) {
	testdata := []struct {
		name  string
		setup func(int) (graph.Graph[int], *graph.Vertex[int])
	}{
		{name: "Empty graph", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int]) {
			v := &graph.Vertex[int]{E: 1}
			g := graph.New[int](graphType)
			return g, v
		}},
		{name: "One vertex graph", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int]) {
			v := &graph.Vertex[int]{E: 1}
			g := graph.New[int](graphType)
			g.AddVertex(v)
			return g, v
		}},
		{name: "Two vertices graph", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int]) {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := graph.Edge[int]{X: v1, Y: v2}
			g := graph.New[int](graphType)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddEdge(&e)
			return g, v1
		}},
		{name: "Five vertices graph", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int]) {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			v3 := &graph.Vertex[int]{E: 3}
			v4 := &graph.Vertex[int]{E: 4}
			v5 := &graph.Vertex[int]{E: 5}
			e1 := graph.Edge[int]{X: v1, Y: v2}
			e2 := graph.Edge[int]{X: v3, Y: v4}
			e3 := graph.Edge[int]{X: v1, Y: v5}
			e4 := graph.Edge[int]{X: v5, Y: v2}
			e5 := graph.Edge[int]{X: v4, Y: v2}
			e6 := graph.Edge[int]{X: v3, Y: v5}
			e7 := graph.Edge[int]{X: v3, Y: v1}
			e8 := graph.Edge[int]{X: v4, Y: v5}
			g := graph.New[int](graphType)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddVertex(v3)
			g.AddVertex(v4)
			g.AddVertex(v5)
			g.AddEdge(&e1)
			g.AddEdge(&e2)
			g.AddEdge(&e3)
			g.AddEdge(&e4)
			g.AddEdge(&e5)
			g.AddEdge(&e6)
			g.AddEdge(&e7)
			g.AddEdge(&e8)
			return g, v3
		}},
	}
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		for _, tc := range testdata {
			t.Run(tc.name+strconv.Itoa(i), func(t *testing.T) {
				g, v := tc.setup(i)
				tree := visit.Generic(g, v)
				if tree.Size() != len(g.Vertices()) {
					t.Fatalf("could not compute correct tree, result is %v", tree)
				}
			})
		}
	}
}

func TestTicketToRideUSA(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		g := graph.NewWithReader(i, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
		tree := visit.Generic(g, &graph.Vertex[string]{E: "Chicago"})
		if tree.Size() != len(g.Vertices()) {
			t.Fatalf("could not compute correct tree, result is %v", tree)
		}
		for _, v := range g.Vertices() {
			v.ResetVisited()
		}
		if !visit.Connected(g) {
			t.Log(len(g.Vertices()), visit.Generic(g, g.Vertices()[0]).Size())
			t.Fatalf("ticket to ride board should be connected but was not; graph: %v", g)
		}
	}
}
