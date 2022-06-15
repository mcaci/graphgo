package graph_test

import (
	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/graph/internal"
)

var testdata = []struct {
	name     string
	edges []string
	a, b     graph.Vertex[string]
	degree   int
	adjNodes []string
}{
	{name: "graph with 1 edge", edges: []string{
		"Vancouver,Seattle,1",
	},
		a: graph.Vertex[string]{E: "Vancouver"}, b: graph.Vertex[string]{E: "Seattle"},
		degree:   1,
		adjNodes: []string{"Seattle"},
	},
	{name: "graph with 2 edges", edges: []string{
		"Vancouver,Seattle,1",
		"Seattle,Helena,6",
	},
		a: graph.Vertex[string]{E: "Seattle"}, b: graph.Vertex[string]{E: "Helena"},
		degree:   2,
		adjNodes: []string{"Helena", "Vancouver"},
	},
	{name: "graph with 4 edges", edges: []string{
		"Vancouver,Seattle,1",
		"Seattle,Helena,6",
		"Helena,Salt Lake City,3",
		"Seattle,Portland,1",
	},
		a: graph.Vertex[string]{E: "Seattle"}, b: graph.Vertex[string]{E: "Portland"},
		degree:   3,
		adjNodes: []string{"Helena", "Vancouver", "Portland"},
	},
	{name: "ticket to ride graph", edges: internal.TicketToRideUSA,
		a: graph.Vertex[string]{E: "Chicago"}, b: graph.Vertex[string]{E: "Toronto"},
		degree:   5,
		adjNodes: []string{"Toronto", "Pittsburgh", "Duluth", "Omaha", "Saint Louis"},
	},
}
