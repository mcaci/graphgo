package graph_test

import (
	"testing"

	"github.com/mcaci/graphgo/graph"
)

func TestIsVisited(t *testing.T) {
	v := graph.Vertex[int]{E: 1}
	if v.IsVisited() {
		t.Fatalf("Expecting %v to be not visited but it was", v)
	}
}

func TestSetVisited(t *testing.T) {
	v := graph.Vertex[int]{E: 1}
	v.SetVisited()
	if !v.IsVisited() {
		t.Fatalf("Expecting %v to be visited but it was not", v)
	}
}
