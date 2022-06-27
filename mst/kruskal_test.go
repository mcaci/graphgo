package mst_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/mst"
	"github.com/mcaci/graphgo/mst/internal"
)

func TestGenericKruskalMST(t *testing.T) {
	testdata := []struct {
		name   string
		setup  func(int) graph.Graph[int]
		ok     func(graph.Graph[int]) bool
		errMsg func(graph.Graph[int]) string
	}{
		{name: "Distance with node to itself", setup: func(graphType int) graph.Graph[int] {
			v := &graph.Vertex[int]{E: 1}
			g := &graph.ArcsList[int]{}
			g.AddVertex(v)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v := &graph.Vertex[int]{E: 1}
			return g.ContainsVertex(v)
		}, errMsg: func(g graph.Graph[int]) string {
			return fmt.Sprintf("Expecting graph %v to contain all vertexes but did not", g)
		}},
		{name: "Distance with two nodes", setup: func(graphType int) graph.Graph[int] {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 5}}
			g := graph.New[int](graph.AdjacencyMatrixType)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddEdge(e)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 5}}
			return g.ContainsEdge(e)
		}, errMsg: func(g graph.Graph[int]) string {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 5}}
			return fmt.Sprintf("Expecting graph %v to contain edge %v but it did not", g, e)
		}},
		{name: "Distance with three nodes", setup: func(graphType int) graph.Graph[int] {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			v3 := &graph.Vertex[int]{E: 3}
			e1 := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 10}}
			e2 := &graph.Edge[int]{X: v3, Y: v2, P: graph.EdgeProperty{W: 8}}
			e3 := &graph.Edge[int]{X: v1, Y: v3, P: graph.EdgeProperty{W: 7}}
			g := graph.New[int](graph.AdjacencyMatrixType)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddVertex(v3)
			g.AddEdge(e1)
			g.AddEdge(e2)
			g.AddEdge(e3)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 10}}
			return !g.ContainsEdge(e)
		}, errMsg: func(g graph.Graph[int]) string {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 10}}
			return fmt.Sprintf("Expecting graph %v to not contain edge %v but it did", g, e)
		}},
		{name: "Five vertices graph", setup: func(graphType int) graph.Graph[int] {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			v3 := &graph.Vertex[int]{E: 3}
			v4 := &graph.Vertex[int]{E: 4}
			v5 := &graph.Vertex[int]{E: 5}
			e1 := graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 3}}
			e2 := graph.Edge[int]{X: v3, Y: v4, P: graph.EdgeProperty{W: 2}}
			e3 := graph.Edge[int]{X: v1, Y: v5, P: graph.EdgeProperty{W: 8}}
			e4 := graph.Edge[int]{X: v5, Y: v2, P: graph.EdgeProperty{W: 4}}
			e5 := graph.Edge[int]{X: v4, Y: v2, P: graph.EdgeProperty{W: 8}}
			e6 := graph.Edge[int]{X: v3, Y: v5, P: graph.EdgeProperty{W: 6}}
			e7 := graph.Edge[int]{X: v3, Y: v1, P: graph.EdgeProperty{W: 11}}
			e8 := graph.Edge[int]{X: v4, Y: v5, P: graph.EdgeProperty{W: 9}}
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
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v2 := &graph.Vertex[int]{E: 2}
			v5 := &graph.Vertex[int]{E: 5}
			e4 := &graph.Edge[int]{X: v5, Y: v2, P: graph.EdgeProperty{W: 4}}
			return g.ContainsEdge(e4)
		}, errMsg: func(g graph.Graph[int]) string {
			v2 := &graph.Vertex[int]{E: 2}
			v5 := &graph.Vertex[int]{E: 5}
			e4 := graph.Edge[int]{X: v5, Y: v2, P: graph.EdgeProperty{W: 4}}
			return fmt.Sprintf("Expecting graph %v to contain edge %v but it did not", g, e4)
		}},
	}
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		for _, tc := range testdata {
			t.Run(tc.name+strconv.Itoa(i), func(t *testing.T) {
				g := tc.setup(i)
				msTree := mst.Kruskal(g)
				if !tc.ok(msTree) {
					t.Fatalf(tc.errMsg(msTree))
				}
			})
		}
	}
}

func TestKruskalWith1Node(t *testing.T) {
	v := &graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v)
	og := mst.Kruskal[int](g)
	if !og.ContainsVertex(v) {
		t.Fatalf("Expecting graph %v to contain all vertexes but did not", og)
	}
}

func TestKruskalWith3Nodes3Arcs(t *testing.T) {
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
	g := graph.NewWithReader(graph.ArcsListType, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
	tree := mst.Kruskal(g)
	var cost int
	for _, e := range tree.Edges() {
		cost += e.P.W
	}
	if len(tree.Vertices()) != len(g.Vertices()) {
		t.Fatalf("could not compute correct tree, result is %v", tree.Edges())
	}
}
