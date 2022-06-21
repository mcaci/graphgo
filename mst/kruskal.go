package mst

import (
	"image/color"
	"sort"

	"github.com/mcaci/graphgo/graph"
)

type SortedEdges[T comparable] []*graph.Edge[T]

func (se *SortedEdges[T]) Len() int           { return len(*se) }
func (se *SortedEdges[T]) Less(i, j int) bool { return (*se)[i].P.W <= (*se)[j].P.W }
func (se *SortedEdges[T]) Swap(i, j int)      { (*se)[i], (*se)[j] = (*se)[j], (*se)[i] }

func Kruskal[T comparable](g graph.Graph[T]) graph.Graph[T] {
	var blueTrees []graph.Graph[T]
	for _, v := range g.Vertices() {
		bt := &graph.ArcsList[T]{}
		bt.AddVertex(v)
		blueTrees = append(blueTrees, bt)
	}
	se := SortedEdges[T](g.Edges())
	sort.Sort(&se)
	for _, e := range se {
		find := func(v *graph.Vertex[T]) int {
			for i, bt := range blueTrees {
				if !bt.ContainsVertex(v) {
					continue
				}
				return i
			}
			return -1
		}
		btx := find(e.X)
		bty := find(e.Y)
		switch btx == bty {
		case true:
			e.P.C = color.RGBA{R: 255}
		default:
			e.P.C = color.RGBA{B: 255}
			g := &graph.ArcsList[T]{}
			g.AddEdge(e)
			for _, v := range blueTrees[btx].Vertices() {
				g.AddVertex(v)
			}
			for _, e := range blueTrees[btx].Edges() {
				g.AddEdge(e)
			}
			for _, v := range blueTrees[bty].Vertices() {
				g.AddVertex(v)
			}
			for _, e := range blueTrees[bty].Edges() {
				g.AddEdge(e)
			}
			blueTrees = append(blueTrees, g)
			var b, x, y []graph.Graph[T]
			switch bty < btx {
			case true:
				b, x, y = blueTrees[:bty], blueTrees[bty+1:btx], blueTrees[btx+1:]
			case false:
				b, x, y = blueTrees[:btx], blueTrees[btx+1:bty], blueTrees[bty+1:]
			}
			blueTrees = append(b, append(x, y...)...)
		}
	}
	return blueTrees[0]
}
