package graph

import (
	"reflect"
	"testing"
)

// 1 -> 3
// 3 -> 5
// 1 -> 2
// 2 -> 4
// 4 -> 5
func TestNewGraphMatrix(t *testing.T) {
	g := NewGraphMatrix(5)
	g.AddEdge(1, 3)
	g.AddEdge(3, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)

	t.Logf("Pretty print graph matrix: \n%s", g.PrettyPrintGraphMatrix())
}

func TestGraphMatrix_DFS(t *testing.T) {
	type fields struct {
		n          int
		nodeMatrix [][]int64
	}
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]int
	}{
		// 1 -> 3
		// 3 -> 5
		// 1 -> 2
		// 2 -> 4
		// 4 -> 5
		{
			name: "graph matrix dfs",
			fields: fields{
				n: 5,
				nodeMatrix: [][]int64{
					{0, 0, 0, 0, 0, 0},
					{0, 0, 1, 1, 0, 0},
					{0, 0, 0, 0, 1, 0},
					{0, 0, 0, 0, 0, 1},
					{0, 0, 0, 0, 0, 1},
					{0, 0, 0, 0, 0, 0},
				},
			},
			args: args{
				x: 1,
				n: 5,
			},
			want: [][]int{
				{1, 2, 4, 5},
				{1, 3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphMatrix{
				n:          tt.fields.n + 1,
				nodeMatrix: tt.fields.nodeMatrix,
			}

			if got := g.DFS(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphMatrix.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
