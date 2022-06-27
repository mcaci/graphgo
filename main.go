package main

import (
	"fmt"
	"strings"

	"github.com/mcaci/graphgo/graph"
)

const edges = `Vancouver Seattle 1 grey
Seattle Helena 6 yellow
Helena SaltLakeCity 3 pink
Seattle Portland 1 grey`

func main() {
	fmt.Printf("g: %v\n", graph.NewWithReader(graph.ArcsListType, strings.NewReader(edges)))
}
