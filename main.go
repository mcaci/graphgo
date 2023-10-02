package main

import (
	"fmt"
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
	vs, es, err := graph.FromSpaced(strings.NewReader(csv))
	if err != nil {
		log.Print(err)
	}
	graph.Fill(vs, es, g)
	fmt.Printf("g: %v\n", g)
}
