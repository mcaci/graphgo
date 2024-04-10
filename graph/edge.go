package graph

import (
	"fmt"
)

type Edge[T comparable] struct {
	X, Y *Vertex[T]
	P    EdgeProperty
}

func (e Edge[T]) String() string {
	return fmt.Sprintf("(x:%v,y:%v,p:%v)", e.X, e.Y, e.P)
}

type EdgeProperty any
