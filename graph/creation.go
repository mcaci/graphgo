package graph

import "io"

const (
	ArcsListType = iota
	AdjacencyListType
	AdjacencyMatrixType
	// IncidenceListType
	// IncidenceMatrixType
)

func Create[T comparable](t int, r io.Reader) Type[string] {
	switch t {
	case ArcsListType:
		return NewArcsList(r)
	case AdjacencyListType:
		return NewAdjacencyList(r)
	case AdjacencyMatrixType:
		return NewAdjacencyMatrix(r)
	// case IncidenceListType:
	// 	return NewIncidenceList(r)
	// case IncidenceMatrixType:
	// 	return NewIncidenceMatrix(r)
	default:
		return NewArcsList(r)
	}
}
