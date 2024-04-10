package graph_test

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/mcaci/graphgo/graph"
)

type EdgeWeight int

func FromCSV(r io.Reader) ([]*graph.Vertex[string], []*graph.Edge[string], error) {
	var vs []*graph.Vertex[string]
	var es []*graph.Edge[string]
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		f := strings.Split(l, ",")
		v := &graph.Vertex[string]{E: f[0]}
		u := &graph.Vertex[string]{E: f[1]}
		vs = append(vs, v, u)
		switch len(f) {
		case 3:
			w, err := strconv.Atoi(f[2])
			if err != nil {
				return nil, nil, err
			}
			es = append(es, &graph.Edge[string]{X: v, Y: u, P: EdgeWeight(w)})
		default:
			es = append(es, &graph.Edge[string]{X: v, Y: u})
		}
	}
	return vs, es, nil
}
