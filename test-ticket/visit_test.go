package testticket_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/test-ticket/internal"
	"github.com/mcaci/graphgo/visit"
)

func TestVisitTicketToRideUSA(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[string](i, false)
			vs, es, err := internal.FromCSV(strings.NewReader(strings.Join(internal.MiniTicket, "\n")), nil)
			if err != nil {
				t.Fatal(err)
			}
			graph.Fill(vs, es, g)
			vs = g.Vertices()
			var s *graph.Vertex[string]
		found:
			for i := range vs {
				switch vs[i].E {
				case "Chicago":
					s = vs[i]
					break found
				}
			}
			tree := visit.Generic(g, s)
			if tree.Size() != len(g.Vertices()) {
				t.Log(tree.Size(), len(g.Vertices()))
				t.Fatalf("could not compute correct tree, result is %v", tree)
			}
		})
	}
}

func TestConnectionTicketToRideUSA(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[string](i, false)
			vs, es, err := internal.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")), nil)
			if err != nil {
				t.Fatal(err)
			}
			graph.Fill(vs, es, g)
			if !visit.Connected(g) {
				t.Log(len(g.Vertices()), visit.Generic(g, g.Vertices()[0]).Size())
				t.Fatalf("ticket to ride board should be connected but was not; graph: %v", g)
			}
		})
	}
}
