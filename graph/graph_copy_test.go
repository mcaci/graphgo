package graph

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	a, b := &ArcsList[int]{}, &ArcsList[int]{}
	a.AddVertex(&Vertex[int]{E: 1})
	a.AddVertex(&Vertex[int]{E: 2})
	a.AddEdge(&Edge[int]{X: a.Vertices()[0], Y: a.Vertices()[1]})
	b.AddVertex(&Vertex[int]{E: 1})
	b.AddVertex(&Vertex[int]{E: 2})
	b.AddEdge(&Edge[int]{X: b.Vertices()[0], Y: b.Vertices()[1]})
	type args struct {
		g Graph[int]
	}
	tests := []struct {
		name string
		args args
		want Graph[int]
	}{
		{
			name: "Two vertices graph",
			args: args{g: a},
			want: b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
