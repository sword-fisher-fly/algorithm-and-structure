package graph

import (
	"reflect"
	"testing"
)

func TestGraphList_BuildGraphList(t *testing.T) {
	type args struct {
		edges [][2]int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			// 1 -> 3
			// 3 -> 5
			// 1 -> 2
			// 2 -> 4
			// 4 -> 5
			name: "5 nodes and 4 edges",
			args: args{
				edges: [][2]int64{
					{1, 3},
					{3, 5},
					{1, 2},
					{2, 4},
					{4, 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraphList()
			g.BuildGraphList(tt.args.edges)

			t.Logf("All vertexs: %v", g.GetVertex())
			t.Logf("Pretty print graph: \n%v", g.PrettyPrintGraphList())
		})
	}
}

func TestGraphList_DFS(t *testing.T) {
	type args struct {
		edges [][2]int64
		x     int64
		n     int64
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{
			// 1 -> 3
			// 3 -> 5
			// 1 -> 2
			// 2 -> 4
			// 4 -> 5
			name: "5 nodes and 4 edges",
			args: args{
				edges: [][2]int64{
					{1, 3},
					{3, 5},
					{1, 2},
					{2, 4},
					{4, 5},
				},
				x: 1,
				n: 5,
			},
			want: [][]int64{
				{1, 3, 5},
				{1, 2, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraphList()
			g.BuildGraphList(tt.args.edges)

			if got := g.DFS(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphList.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
