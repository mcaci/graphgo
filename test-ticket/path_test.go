package testticket_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/path"
	"github.com/mcaci/graphgo/test-ticket/internal"
)

func TestDistanceInTicketToRide(t *testing.T) {
	g := graph.New[string](graph.AdjacencyMatrixType)
	vs, es, err := internal.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")), func(w int) graph.EdgeProperty { return path.EdgeWeight(w) })
	if err != nil {
		t.Fatal(err)
	}
	graph.Fill(vs, es, g)
	vs = g.Vertices()
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
	if dist := d[v2].Dist(); dist != 16 {
		t.Fatalf("Expecting a distance of 16 but was %d; graph: %v", dist, g)
	}
}

func TestPathInTicketToRide(t *testing.T) {
	g := graph.New[string](graph.AdjacencyMatrixType)
	vs, es, err := internal.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")), func(w int) graph.EdgeProperty { return path.EdgeWeight(w) })
	if err != nil {
		t.Fatal(err)
	}
	graph.Fill(vs, es, g)
	vs = g.Vertices()
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
