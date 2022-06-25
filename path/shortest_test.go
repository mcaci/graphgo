package path_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/path"
	"github.com/mcaci/graphgo/path/internal"
)

func TestPathOfNodeWithItself(t *testing.T) {
	v := &graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v)
	d := path.BellmanFordDist[int](g, v)
	p := path.Shortest[int](g, d, v, v)
	if len(p) != 0 {
		t.Fatalf("Expecting a path of len 0 but was %v; graph: %v", p, g)
	}
}

func TestPathOf2Nodes(t *testing.T) {
	v1 := &graph.Vertex[int]{E: 1}
	v2 := &graph.Vertex[int]{E: 2}
	e := &graph.Edge[int]{X: v1, Y: v2, P: graph.EdgeProperty{W: 5}}
	g := &graph.ArcsList[int]{}
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddEdge(e)
	d := path.BellmanFordDist[int](g, v1)
	p := path.Shortest[int](g, d, v1, v2)
	if len(p) != 2 {
		t.Fatalf("Expecting a path of len 2 but was %v; graph: %v", p, g)
	}
}

func TestPathInTicketToRide(t *testing.T) {
	g := graph.Create(graph.ArcsListType, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
	vs := g.Vertices()
	var v1, v2 *graph.Vertex[string]
	for i := range vs {
		switch vs[i].E {
		case "Chicago":
			v1 = vs[i]
		case "Vancouver":
			v2 = vs[i]
		}
	}
	d := path.BellmanFordDist(g, v1)
	p := path.Shortest(g, d, v1, v2)
	if len(p) != 5 {
		t.Fatalf("Expecting a path of len 5 but was %v; graph: %v", p, g)
	}
}
