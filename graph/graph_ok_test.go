package graph_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mcaci/graphgo/graph"
)

func TestGraphCreationOK(t *testing.T) {
	for _, tc := range testdata {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i <= graph.AdjacencyMatrixType; i++ {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					g := graph.Create(i, strings.NewReader(strings.Join(tc.edges, "\n")))
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
					g := graph.Create(i, strings.NewReader(strings.Join(tc.edges, "\n")))
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
					g := graph.Create(i, strings.NewReader(strings.Join(tc.edges, "\n")))
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
