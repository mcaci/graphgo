package testticket_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/mst"
	"github.com/mcaci/graphgo/test-ticket/internal"
)

func TestMSTOnTicketToRideUSA(t *testing.T) {
	g := graph.NewFromCSV(graph.ArcsListType, strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")))
	tree := mst.Kruskal(g)
	var cost int
	for _, e := range tree.Edges() {
		cost += e.P.W
	}
	if len(tree.Vertices()) != len(g.Vertices()) {
		t.Fatalf("could not compute correct tree, result is %v", tree.Edges())
	}
}
