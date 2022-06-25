package graph_test

import (
	"strconv"
	"testing"

	"github.com/mcaci/graphgo/graph"
)

func TestAddVertexOK(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			v := &graph.Vertex[int]{E: 5}
			g.AddVertex(v)
			if !g.ContainsVertex(v) {
				t.Fatalf("Expecting graph %v to contain vertex %v but did not", g, v)
			}
		})
	}
}

func TestAddVertexKO(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			v := &graph.Vertex[int]{E: 5}
			nov := &graph.Vertex[int]{E: 4}
			g.AddVertex(v)
			if g.ContainsVertex(nov) {
				t.Fatalf("Expecting graph %v not to contain vertex %v but it did", g, nov)
			}
		})
	}
}

func TestAddEdgeOK(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			if !g.ContainsEdge(e) {
				t.Fatalf("Expecting graph %v to contain edge %v but did not", g, e)
			}
		})
	}
}

func TestAddEdgeKO(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			noe := &graph.Edge[int]{X: x, Y: &graph.Vertex[int]{E: 4}}
			if g.ContainsEdge(noe) {
				t.Fatalf("Expecting graph %v not to contain edge %v but it did", g, noe)
			}
		})
	}
}

func TestRemoveVertexOK(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			v := &graph.Vertex[int]{E: 5}
			v1 := &graph.Vertex[int]{E: 6}
			g.AddVertex(v)
			g.AddVertex(v1)
			g.RemoveVertex(v)
			if g.ContainsVertex(v) {
				t.Fatalf("Expecting graph %v not to contain vertex %v but it did", g, v)
			}
		})
	}
}

func TestRemoveVertexKO(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			v := &graph.Vertex[int]{E: 5}
			g.AddVertex(v)
			g.RemoveVertex(&graph.Vertex[int]{E: 4})
			if !g.ContainsVertex(v) {
				t.Fatalf("Expecting graph %v to still contain vertex %v but it did not", g, v)
			}
		})
	}
}

func TestRemoveEdgeOK(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			g.RemoveEdge(e)
			if g.ContainsEdge(e) {
				t.Fatalf("Expecting graph %v to not contain edge %v but it did", g, e)
			}
		})
	}
}

func TestRemoveMultipleEdgeOK(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			x, y, z := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}, &graph.Vertex[int]{E: 9}
			g.AddVertex(x)
			g.AddVertex(y)
			g.AddVertex(z)
			e0 := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e0)
			e1 := &graph.Edge[int]{X: x, Y: z}
			g.AddEdge(e1)
			e2 := &graph.Edge[int]{X: y, Y: z}
			g.AddEdge(e2)
			g.RemoveEdge(e0)
			if g.ContainsEdge(e0) {
				t.Fatalf("Expecting graph %v to not contain edge %v but it did", g, e0)
			}
			g.RemoveEdge(e1)
			if g.ContainsEdge(e1) {
				t.Fatalf("Expecting graph %v to not contain edge %v but it did", g, e0)
			}
			g.RemoveEdge(e2)
			if g.ContainsEdge(e2) {
				t.Fatalf("Expecting graph %v to not contain edge %v but it did", g, e0)
			}
		})
	}
}

func TestRemomveEdgeKO(t *testing.T) {
	for i := 0; i <= graph.AdjacencyMatrixType; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := graph.New[int](i)
			x, y := &graph.Vertex[int]{E: 5}, &graph.Vertex[int]{E: 2}
			g.AddVertex(x)
			g.AddVertex(y)
			e := &graph.Edge[int]{X: x, Y: y}
			g.AddEdge(e)
			noe := &graph.Edge[int]{X: x, Y: &graph.Vertex[int]{E: 4}}
			g.RemoveEdge(noe)
			if !g.ContainsEdge(e) {
				t.Fatalf("Expecting graph %v not to still contain edge %v but it did not", g, e)
			}
		})
	}
}
