package graph

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	type args struct {
		n     int64
		edges [][2]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "case test",
			args: args{
				n:     5,
				edges: [][2]int64{{0, 1}, {0, 2}, {1, 3}, {2, 4}},
			},
			want: []int64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopologicalSort(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopologicalSortFromGraphList(t *testing.T) {
	type args struct {
		graph *GraphList
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "case 1",
			args: args{
				graph: graphList5Vertex4Edge,
			},
			want: []int64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopologicalSortFromGraphList(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopologicalSortFromGraphList() = %v, want %v", got, tt.want)
			}
		})
	}
}
