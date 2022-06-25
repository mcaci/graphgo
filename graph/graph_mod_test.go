package graph_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mcaci/graphgo/graph"
)

func TestGenericGraphModificationOperation(t *testing.T) {
	testdata := []struct {
		name   string
		setup  func(int) graph.Graph[int]
		ok     func(graph.Graph[int]) bool
		errMsg func(graph.Graph[int]) string
	}{
		{name: "Check on empty graph", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v := &graph.Vertex[int]{E: 1}
			return !g.ContainsVertex(v)
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 1}
			return fmt.Sprintf("Expecting graph %v not to contain vertex %v but it did", g, v)
		}},
		{name: "Check vertex in one vertex graph", setup: func(graphType int) graph.Graph[int] {
			v := &graph.Vertex[int]{E: 1}
			g := graph.New[int](graphType)
			g.AddVertex(v)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v := &graph.Vertex[int]{E: 1}
			return g.ContainsVertex(v)
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 1}
			return fmt.Sprintf("Expecting graph %v to contain vertex %v but it did not", g, v)
		}},
		{name: "Check Edge in two vertices graph", setup: func(graphType int) graph.Graph[int] {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := graph.Edge[int]{X: v1, Y: v2}
			g := graph.New[int](graphType)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddEdge(&e)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := graph.Edge[int]{X: v1, Y: v2}
			return g.ContainsEdge(&e)
		}, errMsg: func(g graph.Graph[int]) string {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := graph.Edge[int]{X: v1, Y: v2}
			return fmt.Sprintf("Expecting graph %v to contain edge %v but it did not", g, e)
		}},
		{name: "Five vertices graph", setup: func(graphType int) graph.Graph[int] {
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
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v3 := &graph.Vertex[int]{E: 3}
			v5 := &graph.Vertex[int]{E: 5}
			e := graph.Edge[int]{X: v3, Y: v5}
			return g.ContainsEdge(&e)
		}, errMsg: func(g graph.Graph[int]) string {
			v3 := &graph.Vertex[int]{E: 3}
			v5 := &graph.Vertex[int]{E: 5}
			e := graph.Edge[int]{X: v3, Y: v5}
			return fmt.Sprintf("Expecting graph %v to contain edge %v but it did not", g, e)
		}},
		{name: "Check non existent vertex", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			v := &graph.Vertex[int]{E: 5}
			g.AddVertex(v)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			v := &graph.Vertex[int]{E: 4}
			return !g.ContainsVertex(v)
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 1}
			return fmt.Sprintf("Expecting graph %v not to contain vertex %v but it did", g, v)
		}},
		{name: "Check non existent edge", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			e := &graph.Edge[int]{X: &graph.Vertex[int]{E: 5}, Y: &graph.Vertex[int]{E: 4}}
			return !g.ContainsEdge(e)
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 1}
			return fmt.Sprintf("Expecting graph %v not to contain vertex %v but it did", g, v)
		}},
		{name: "Check vertex removal", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			v := &graph.Vertex[int]{E: 5}
			v1 := &graph.Vertex[int]{E: 6}
			g.AddVertex(v)
			g.AddVertex(v1)
			g.RemoveVertex(v)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			return !g.ContainsVertex(&graph.Vertex[int]{E: 5})
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 5}
			return fmt.Sprintf("Expecting graph %v not to contain vertex %v but it did", g, v)
		}},
		{name: "Check removal of non existent vertex", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			v := &graph.Vertex[int]{E: 5}
			g.AddVertex(v)
			g.RemoveVertex(&graph.Vertex[int]{E: 4})
			return g
		}, ok: func(g graph.Graph[int]) bool {
			return g.ContainsVertex(&graph.Vertex[int]{E: 5})
		}, errMsg: func(g graph.Graph[int]) string {
			v := &graph.Vertex[int]{E: 5}
			return fmt.Sprintf("Expecting graph %v to still contain vertex %v but it did not", g, v)
		}},
		{name: "Check edge removal", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			g.RemoveEdge(e)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: x, Y: y}
			return !g.ContainsEdge(e)
		}, errMsg: func(g graph.Graph[int]) string {
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: x, Y: y}
			return fmt.Sprintf("Expecting graph %v not to contain edge %v but it did", g, e)
		}},
		{name: "Check removal of non existent edge", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			g.RemoveEdge(e)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: x, Y: y}
			return !g.ContainsEdge(e)
		}, errMsg: func(g graph.Graph[int]) string {
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: x, Y: y}
			return fmt.Sprintf("Expecting graph %v not to still contain edge %v but it did not", g, e)
		}},
		{name: "Check removal of multiple edges", setup: func(graphType int) graph.Graph[int] {
			g := graph.New[int](graphType)
			x, y, z := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}, &graph.Vertex[int]{E: 9}
			g.AddVertex(x)
			g.AddVertex(y)
			g.AddVertex(z)
			e0 := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e0)
			e1 := &graph.Edge[int]{X: x, Y: z}
			g.AddEdge(e1)
			e2 := &graph.Edge[int]{X: y, Y: z}
			g.AddEdge(e2)
			g.RemoveEdge(e0)
			g.RemoveEdge(e1)
			g.RemoveEdge(e2)
			return g
		}, ok: func(g graph.Graph[int]) bool {
			x, y, z := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}, &graph.Vertex[int]{E: 9}
			e0 := &graph.Edge[int]{X: x, Y: y}
			e1 := &graph.Edge[int]{X: x, Y: z}
			e2 := &graph.Edge[int]{X: y, Y: z}
			return !g.ContainsEdge(e0) && !g.ContainsEdge(e1) && !g.ContainsEdge(e2)
		}, errMsg: func(g graph.Graph[int]) string {
			x, y, z := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}, &graph.Vertex[int]{E: 9}
			e0 := &graph.Edge[int]{X: x, Y: y}
			e1 := &graph.Edge[int]{X: x, Y: z}
			e2 := &graph.Edge[int]{X: y, Y: z}
			return fmt.Sprintf("Expecting graph %v not to contain edges %v, %v, %v but it did", g, e0, e1, e2)
		}},
	}
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		for _, tc := range testdata {
			t.Run(tc.name+strconv.Itoa(i), func(t *testing.T) {
				g := tc.setup(i)
				if !tc.ok(g) {
					t.Fatalf(tc.errMsg(g))
				}
			})
		}
	}
}