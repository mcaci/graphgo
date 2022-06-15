package graph

import (
	"fmt"
	"image/color"
)

type Edge[T comparable] struct {
	X, Y *Vertex[T]
	P    EdgeProperty
}

func (e Edge[T]) String() string {
	return fmt.Sprintf("(x:%v,y:%v,p:%v)", e.X, e.Y, e.P)
}

type EdgeProperty struct {
	W int
	C color.Color
}

func (e EdgeProperty) String() string {
	return fmt.Sprintf("(w:%d,c:%v)", e.W, e.C)
}
