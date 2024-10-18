package graph

import "fmt"

type Vertex[T comparable] struct {
	E       T
	visited bool
}

func (v *Vertex[T]) Visit()        { v.visited = true }
func (v *Vertex[T]) Unvisit()        { v.visited = false }
func (v *Vertex[T]) Visited() bool { return v.visited }

func (v Vertex[T]) String() string {
	return fmt.Sprintf("(e:%v)", v.E)
}
