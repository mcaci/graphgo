package testticket_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/mst"
	"github.com/mcaci/graphgo/test-ticket/internal"
)

func TestMSTOnTicketToRideUSA(t *testing.T) {
	g := graph.New[string](graph.ArcsListType, false)
	vs, es, err := internal.FromCSV(strings.NewReader(strings.Join(internal.TicketToRideUSA, "\n")), func(w int) graph.EdgeProperty { return mst.EdgeWeightAndColor{W: w} })
	if err != nil {
		t.Fatal(err)
	}
	graph.Fill(vs, es, g)
	tree := mst.Kruskal(g)
	var cost int
	for _, e := range tree.Edges() {
		cost += int(e.P.(mst.EdgeWeightAndColor).W)
	}
	if len(tree.Vertices()) != len(g.Vertices()) {
		t.Fatalf("could not compute correct tree, result is %v", tree.Edges())
	}
}
