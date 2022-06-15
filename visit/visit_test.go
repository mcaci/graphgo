package visit_test

import (
	"testing"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/visit"
)

func TestGenericVisit(t *testing.T) {
	v := graph.Vertex[int]{E: 1}
	g := &graph.ArcsList[int]{}
	tree := visit.Generic[int](g, v)
	t.Fatalf("tree %v", tree)
}
