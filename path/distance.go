package path

import (
	"fmt"

	"github.com/mcaci/graphgo/graph"
)

type Distance[T comparable] struct {
	v, u *graph.Vertex[T]
	d    int
}

func (d Distance[T]) Dist() int             { return d.d }
func (d *Distance[T]) setDistance(dist int) { d.d = dist }

func (d Distance[T]) String() string {
	return fmt.Sprintf("(v: %v, u:%v, d:%d)", d.v.E, d.u.E, d.d)
}
