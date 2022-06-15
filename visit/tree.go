package visit

type Tree[T comparable] struct {
	root     *Node[T]
	children []*Node[T]
}

type Node[T comparable] struct {
	element  T
	children []*Node[T]
}
