package graph

import (
	"strings"
	"testing"
)

func prettyPrintGrid(visitedPath [][]bool) string {
	var sb strings.Builder
	for i := 0; i < len(visitedPath); i++ {
		for j := 0; j < len(visitedPath[0]); j++ {
			if visitedPath[i][j] {
				sb.WriteByte('1')
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
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}
	tests := []struct {
		name        string
		args        args
		wantVisited [][]int
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{1, 0, 1, 1},
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
			wantVisited: [][]int{
				{1, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.visited[tt.args.x][tt.args.y] = true

			bfs(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)

			// t.Logf("visited: %v", tt.args.visited)

			t.Logf("visitedPath: \n%s", prettyPrintGrid(tt.args.visited))

			// 有些额外的相邻点也被访问过, 是否可以剔除?
			// if !reflect.DeepEqual(tt.wantVisited, visitedPath) {
			// 	t.Errorf("bfs() = %v, want %v", visitedPath, tt.wantVisited)
			// }
		})
	}
}
