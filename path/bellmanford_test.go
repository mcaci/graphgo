package path_test

import (
	"strconv"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/path"
)

type EdgeWeight int

func (e EdgeWeight) Weight() int { return int(e) }

func TestGenericDistance(t *testing.T) {
	testdata := []struct {
		name  string
		setup func(int) (graph.Graph[int], *graph.Vertex[int], *graph.Vertex[int])
		dist  int
	}{
		{name: "Distance with node to itself", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int], *graph.Vertex[int]) {
			v := &graph.Vertex[int]{E: 1}
			g := graph.New[int](graphType, false)
			g.AddVertex(v)
			return g, v, v
		}},
		{name: "Distance with two nodes", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int], *graph.Vertex[int]) {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			e := &graph.Edge[int]{X: v1, Y: v2, P: EdgeWeight(5)}
			g := graph.New[int](graphType, false)
			g.AddVertex(v1)
			g.AddVertex(v2)
			g.AddEdge(e)
			return g, v1, v2
		}, dist: 5,
		},
		{name: "Five vertices graph", setup: func(graphType int) (graph.Graph[int], *graph.Vertex[int], *graph.Vertex[int]) {
			v1 := &graph.Vertex[int]{E: 1}
			v2 := &graph.Vertex[int]{E: 2}
			v3 := &graph.Vertex[int]{E: 3}
			v4 := &graph.Vertex[int]{E: 4}
			v5 := &graph.Vertex[int]{E: 5}
			e1 := graph.Edge[int]{X: v1, Y: v2, P: EdgeWeight(3)}
			e2 := graph.Edge[int]{X: v3, Y: v4, P: EdgeWeight(2)}
			e3 := graph.Edge[int]{X: v1, Y: v5, P: EdgeWeight(8)}
			e4 := graph.Edge[int]{X: v5, Y: v2, P: EdgeWeight(4)}
			e5 := graph.Edge[int]{X: v4, Y: v2, P: EdgeWeight(8)}
			e6 := graph.Edge[int]{X: v3, Y: v5, P: EdgeWeight(6)}
			e7 := graph.Edge[int]{X: v3, Y: v1, P: EdgeWeight(11)}
			e8 := graph.Edge[int]{X: v4, Y: v5, P: EdgeWeight(9)}
			g := graph.New[int](graphType, false)
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
			return g, v4, v1
		}, dist: 11,
		},
	}
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		for _, tc := range testdata {
			t.Run(tc.name+strconv.Itoa(i), func(t *testing.T) {
				g, v, u := tc.setup(i)
				d := path.BellmanFordDist(g, v)
				if dist := d[u].Dist(); dist != tc.dist {
					t.Fatalf("Expecting a distance of %d but was %d; graph: %v", tc.dist, dist, g)
				}
			})
		}
	}
}
