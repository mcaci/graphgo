package testticket_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/test-ticket/internal"
	"github.com/mcaci/graphgo/visit"
)

func TestVisitTicketToRideUSA(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		g := graph.New[string](i)
		vs, es, err := graph.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
		if err != nil {
			t.Fatal(err)
		}
		graph.Fill(vs, es, g)
		tree := visit.Generic(g, &graph.Vertex[string]{E: "Chicago"})
		if tree.Size() != len(g.Vertices()) {
			t.Fatalf("could not compute correct tree, result is %v", tree)
		}
	}
}

func TestConnectionTicketToRideUSA(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		g := graph.New[string](i)
		vs, es, err := graph.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
		if err != nil {
			t.Fatal(err)
		}
		graph.Fill(vs, es, g)
		if !visit.Connected(g) {
			t.Log(len(g.Vertices()), visit.Generic(g, g.Vertices()[0]).Size())
			t.Fatalf("ticket to ride board should be connected but was not; graph: %v", g)
		}
	}
}
