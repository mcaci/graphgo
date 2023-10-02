package graph_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
)

var testdata = []struct {
	name     string
	edges    []string
	a, b     graph.Vertex[string]
	degree   int
	adjNodes []string
}{
	{name: "graph with 1 edge", edges: []string{
		"Vancouver,Seattle",
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
	{name: "ticket to ride graph", edges: []string{
		"Vancouver,Seattle,1",
		"Seattle,Helena,6",
		"Helena,Salt Lake City,3",
		"Portland,Salt Lake City,6",
		"Seattle,Portland,1",
		"Seattle,Calgary,4",
		"Vancouver,Calgary,3",
		"Helena,Calgary,4",
		"Winnipeg,Calgary,6",
		"Winnipeg,Helena,4",
		"Duluth,Helena,6",
		"Duluth,Winnipeg,4",
		"Sault St Marie,Winnipeg,6",
		"Sault St Marie,Duluth,3",
		"Omaha,Helena,5",
		"Omaha,Duluth,2",
		"Omaha,Denver,4",
		"Kansas City,Denver,4",
		"Omaha,Kansas City,1",
		"Oklahoma City,Kansas City,1",
		"Oklahoma City,Denver,4",
		"Oklahoma City,Santa Fe,3",
		"Oklahoma City,El Paso,4",
		"Oklahoma City,Dallas,2",
		"Dallas,El Paso,4",
		"Houston,El Paso,6",
		"Dallas,Houston,1",
		"Portland,San Francisco,5",
		"Salt Lake City,San Francisco,5",
		"Salt Lake City,Las Vegas,3",
		"Los Angeles,Las Vegas,2",
		"Los Angeles,San Francisco,3",
		"Los Angeles,Phoenix,3",
		"Los Angeles,El Paso,6",
		"Phoenix,El Paso,3",
		"Phoenix,Denver,5",
		"Salt Lake City,Denver,3",
		"Helena,Denver,4",
		"Santa Fe,Denver,2",
		"Santa Fe,Phoenix,3",
		"Santa Fe,El Paso,2",
		"Sault St Marie,Montreal,5",
		"Sault St Marie,Toronto,2",
		"Duluth,Toronto,6",
		"Montreal,Toronto,3",
		"Toronto,Pittsburgh,2",
		"Duluth,Chicago,3",
		"Omaha,Chicago,4",
		"Chicago,Toronto,4",
		"Chicago,Pittsburgh,3",
		"Chicago,Saint Louis,2",
		"Kansas City,Saint Louis,2",
		"Pittsburgh,Saint Louis,5",
		"Pittsburgh,New York,2",
		"Boston,New York,2",
		"Boston,Montreal,2",
		"Montreal,New York,3",
		"Washington,New York,2",
		"Washington,Pittsburgh,2",
		"Washington,Raleigh,2",
		"Pittsburgh,Raleigh,2",
		"Pittsburgh,Nashville,4",
		"Raleigh,Nashville,3",
		"Saint Louis,Nashville,2",
		"Raleigh,Charleston,2",
		"Atlanta,Charleston,2",
		"Miami,Charleston,4",
		"Miami,Atlanta,5",
		"Miami,New Orleans,6",
		"Houston,New Orleans,2",
		"Little Rock,New Orleans,3",
		"Atlanta,New Orleans,4",
		"Atlanta,Nashville,1",
		"Little Rock,Nashville,3",
		"Atlanta,Raleigh,2",
		"Little Rock,Saint Louis,2",
		"Little Rock,Dallas,2",
		"Little Rock,Oklahoma City,2",
	},
		a: graph.Vertex[string]{E: "Chicago"}, b: graph.Vertex[string]{E: "Toronto"},
		degree:   5,
		adjNodes: []string{"Toronto", "Pittsburgh", "Duluth", "Omaha", "Saint Louis"},
	},
}

func TestGraphCreationOK(t *testing.T) {
	for _, tc := range testdata {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i <= graph.AdjacencyMatrixType; i++ {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					vs, es, err := graph.FromCSV(strings.NewReader(strings.Join(tc.edges, "\n")))
					if err != nil {
						t.Fatal(err)
					}
					g := graph.New[string](i)
					graph.Fill(vs, es, g)
					if !g.AreAdjacent(&tc.a, &tc.b) {
						t.Fatalf("expecting %v to contain (%v,%v)", g, tc.a, tc.b)
					}
				})
			}
		})
	}
}

func TestGraphDegreeOK(t *testing.T) {
	for _, tc := range testdata {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i <= graph.AdjacencyMatrixType; i++ {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					vs, es, err := graph.FromCSV(strings.NewReader(strings.Join(tc.edges, "\n")))
					if err != nil {
						t.Fatal(err)
					}
					g := graph.New[string](i)
					graph.Fill(vs, es, g)
					if d := g.Degree(&tc.a); d != tc.degree {
						t.Fatalf("expecting node %v from graph %v to have degree %v but was %v", tc.a, g, tc.degree, d)
					}
				})
			}
		})
	}
}

func TestAdjNodesOK(t *testing.T) {
	for _, tc := range testdata {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i <= graph.AdjacencyMatrixType; i++ {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					vs, es, err := graph.FromCSV(strings.NewReader(strings.Join(tc.edges, "\n")))
					if err != nil {
						t.Fatal(err)
					}
					g := graph.New[string](i)
					graph.Fill(vs, es, g)
					nodes := g.AdjacentNodes(&tc.a)
					for _, tcn := range tc.adjNodes {
						var found bool
						for _, n := range nodes {
							if n.E != tcn {
								continue
							}
							found = true
							break
						}
						if !found {
							t.Fatalf("expecting node %v from graph %v to be found in this list %v", tcn, g, nodes)
						}
					}
				})
			}
		})
	}
}
