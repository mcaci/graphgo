package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/mcaci/graphgo/graph"
)

const csv = `Vancouver Seattle 1 grey
Seattle Helena 6 yellow
Helena SaltLakeCity 3 pink
Seattle Portland 1 grey`

func main() {
	g := graph.New[string](graph.AdjacencyListType)
	vs, es, err := FromSpaced(strings.NewReader(csv))
	if err != nil {
		log.Print(err)
	}
	graph.Fill(vs, es, g)
	fmt.Printf("g: %v\n", g)
}

type EdgeWeight int

func FromSpaced(r io.Reader) ([]*graph.Vertex[string], []*graph.Edge[string], error) {
	var vs []*graph.Vertex[string]
	var es []*graph.Edge[string]
	s := bufio.NewScanner(r)
	for s.Scan() {
		var vv, uu, cc string
		var ww int
		fmt.Sscanf(s.Text(), "%s %s %d %s", &vv, &uu, &ww, &cc)
		v := &graph.Vertex[string]{E: vv}
		u := &graph.Vertex[string]{E: uu}
		vs = append(vs, v, u)
		// replace C with cc
		e := &graph.Edge[string]{X: v, Y: u, P: EdgeWeight(ww)}
		es = append(es, e)
	}
	return vs, es, nil
}
