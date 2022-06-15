package graph

import "fmt"

type Vertex[T comparable] struct {
	E T
}

func (v Vertex[T]) String() string {
	return fmt.Sprintf("(e:%v)", v.E)
}
