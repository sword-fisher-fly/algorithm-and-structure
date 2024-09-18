package graph

import "testing"

func Test_dfs(t *testing.T) {
	type args struct {
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}
	tests := []struct {
		name string
		args args
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.visited[tt.args.x][tt.args.y] = true
			// t.Logf("visited path: %v", prettyPrintGrid(tt.args.visited))

			dfs(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)

			t.Logf("visited path: \n%v", prettyPrintGrid(tt.args.visited))
		})
	}
}
