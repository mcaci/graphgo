package mst_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/mst"
	"github.com/mcaci/graphgo/mst/internal"
)

func testKruskalWith1Node(t *testing.T) {
	v := &graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v)
	og := mst.Kruskal[int](g)
	if !og.ContainsVertex(v) {
		t.Fatalf("Expecting graph %v to contain all vertexes but did not", og)
	}
}

func testKruskalWith3Nodes3Arcs(t *testing.T) {
	v1 := &graph.Vertex[int]{E: 1}
	v2 := &graph.Vertex[int]{E: 2}
	v3 := &graph.Vertex[int]{E: 3}
	e1 := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 10}}
	e2 := &graph.Edge[int]{X: v3, Y: v2, P: graph.EdgeProperty{W: 8}}
	e3 := &graph.Edge[int]{X: v1, Y: v3, P: graph.EdgeProperty{W: 7}}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)
	g.AddEdge(e1)
	g.AddEdge(e2)
	g.AddEdge(e3)
	og := mst.Kruskal[int](g)
	if og.ContainsEdge(e1) {
		t.Fatalf("Expecting graph %v to not contain edge %v but it did", og, e1)
	}
}

func TestTicketToRideUSA(t *testing.T) {
	g := graph.Create[string](graph.ArcsListType, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
	tree := mst.Kruskal(g)
	var cost int
	for _, e := range tree.Edges() {
		t.Log(e.P.W, e.X, e.Y)
		cost += e.P.W
	}
	t.Log(cost)
	if len(tree.Vertices()) != len(g.Vertices()) {
		t.Fatalf("could not compute correct tree, result is %v", tree.Edges())
	}
}
