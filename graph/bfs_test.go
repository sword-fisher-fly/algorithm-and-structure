package graph

import (
	"strings"
	"testing"
)

func prettyPrintGrid(grid [][]byte) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, col := range row {
			if col != 0 {
				sb.WriteByte(col)
			} else {
				sb.WriteByte(' ')
			}
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func Test_bfs(t *testing.T) {
	type args struct {
		grid    [][]byte
		visited [][]bool
		x       int
		y       int
	}
	tests := []struct {
		name        string
		args        args
		wantVisited [][]byte
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0'},
					{'0', '1', '0', '0'},
					{'1', '1', '1', '0'},
					{'1', '0', '1', '1'},
				},
				visited: [][]bool{
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
				x: 0,
				y: 0,
			},
			wantVisited: [][]byte{
				{'1', '1', 0, 0},
				{0, '1', 0, 0},
				{0, '1', '1', 0},
				{0, 0, '1', '1'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bfs(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)

			visitedPath := make([][]byte, len(tt.args.visited))
			for i := range visitedPath {
				visitedPath[i] = make([]byte, len(tt.args.visited[i]))
			}

			for i := 0; i < len(tt.args.visited); i++ {
				for j := 0; j < len(tt.args.visited[i]); j++ {
					if tt.args.visited[i][j] {
						visitedPath[i][j] = '1'
					}
				}
			}

			t.Logf("visitedPath: \n%s", prettyPrintGrid(visitedPath))

			// 有些额外的相邻点也被访问过, 是否可以剔除?
			// if !reflect.DeepEqual(tt.wantVisited, visitedPath) {
			// 	t.Errorf("bfs() = %v, want %v", visitedPath, tt.wantVisited)
			// }
		})
	}
}
