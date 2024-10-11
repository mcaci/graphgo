package visit

import (
	"fmt"
	"strings"
)

type Tree[T comparable] struct {
	element  *T
	children []*Tree[T]
}

func (t *Tree[T]) Size() int {
	switch {
	case t == nil, t.element == nil:
		return 0
	case t.children == nil:
		return 1
	default:
		size := 1
		for i := range t.children {
			size += t.children[i].Size()
		}
		return size
	}
}

func (t *Tree[T]) Find(e *T) *Tree[T] {
	switch {
	case t == nil, t.element == nil:
		return nil
	case *t.element == *e:
		return t
	case t.children == nil:
		return nil
	default:
		var tree *Tree[T]
		for i := range t.children {
			tree = t.children[i].Find(e)
			if tree == nil {
				continue
			}
			return tree
		}
		return nil
	}
}

func (t *Tree[T]) Contains(e *T) bool { return t.Find(e) != nil }

func (t *Tree[T]) String() string {
	switch {
	case t == nil, t.element == nil:
		return "()"
	case t.children == nil:
		return fmt.Sprintf("(element:%v)", *t.element)
	default:
		var children []string
		for i := range t.children {
			children = append(children, t.children[i].String())
		}
		return fmt.Sprintf("(element:%v,\nchildren:%v)", *t.element, strings.Join(children, ","))
	}
}
