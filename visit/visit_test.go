package visit_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/visit"
	"github.com/mcaci/graphgo/visit/internal"
)

func TestGenericVisitWithSize0(t *testing.T) {
	v := &graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	tree := visit.Generic[int](g, v)
	if tree.Size() != 0 {
		t.Fatalf("could not compute correct tree, result is %v", tree)
	}
}

func TestGenericVisitWithSize1(t *testing.T) {
	v := &graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v)
	tree := visit.Generic[int](g, v)
	if tree.Size() != 1 {
		t.Fatalf("could not compute correct tree, result is %v", tree)
	}
}

func TestGenericVisitWithSize2(t *testing.T) {
	v1 := &graph.Vertex[int]{E: 1}
	v2 := &graph.Vertex[int]{E: 2}
	e := graph.Edge[int]{X: v1, Y: v2}
	g := &graph.AdjacencyLists[int]{}
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(&e)
	tree := visit.Generic[int](g, v1)
	if tree.Size() != 2 {
		t.Fatalf("could not compute correct tree, result is %v", tree)
	}
}

func TestGenericVisitWithSize5(t *testing.T) {
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
	g := &graph.AdjacencyMatrix[int]{}
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
	tree := visit.Generic[int](g, v3)
	if tree.Size() != 5 {
		t.Fatalf("could not compute correct tree, result is %v", tree)
	}
}

func TestTicketToRideUSA(t *testing.T) {
	g := graph.Create[string](graph.AdjacencyListType, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
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
