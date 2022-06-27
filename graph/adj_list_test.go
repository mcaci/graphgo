package graph_test

import (
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
)

func TestAdjacencyEdgesCreationOK(t *testing.T) {
	for _, tc := range testdata {
		t.Run(tc.name, func(t *testing.T) {
			g := graph.NewWithReader(graph.AdjacencyListType, strings.NewReader(strings.Join(tc.edges, "\n")))
			if edges := g.Edges(); len(edges) == 0 {
				t.Fatalf("expecting %v to be filled but was not; graph: %v", edges, g)
			}
		})
	}
}
